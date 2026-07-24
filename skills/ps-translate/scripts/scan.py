#!/usr/bin/env python3
"""
scan.py — Detecta strings sin traducir en una instalación PrestaShop.

Salida: JSON con los strings agrupados por dominio PS.

Uso:
  python scan.py --base /ruta/ps [--theme milagros] [--lang es-CO]
                 [--domain ShopThemePanda] [--output /tmp/missing.json]
                 [--include-admin]
"""

import argparse
import glob
import json
import os
import re
import sys
from collections import defaultdict

try:
    from detect import detect_theme  # type: ignore
except ImportError:
    sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
    from detect import detect_theme

# ── Heurística "ya traducido" — genérica por idioma ────────────────────────────
# Marcadores tipográficos/diacríticos por prefijo ISO. Sirve para saltar strings
# que un dev escribió directamente en el idioma destino dentro del template.
# El diff contra los XLF existentes hace el grueso del trabajo; esto solo evita
# re-traducir strings ya nativos. Idioma desconocido → sin heurística.
LANG_HINTS = {
    'es': 'áéíóúÁÉÍÓÚñÑüÜ¡¿',
    'fr': 'àâçéèêëîïôûùüÿœæÀÂÇÉÈÊËÎÏÔÛÙÜ',
    'de': 'äöüßÄÖÜ',
    'it': 'àèéìíîòóùúÀÈÉÌÒÙ',
    'pt': 'ãõáâàéêíóôúçÃÕÁÂÀÉÊÍÓÔÚÇ',
    'ca': 'àèéíïòóúçÀÈÉÍÏÒÓÚÇ·',
    'nl': 'ëïéèáä',
    'pl': 'ąćęłńóśźżĄĆĘŁŃÓŚŹŻ',
}
ADMIN_DOMAINS = re.compile(
    r'^Admin\.|^psgdpr|^ps_themecusto|^steasybuilder|'
    r'^ps_facebook|^ps_accounts|^allinone|^AdminTheme'
)

L_PATTERN = re.compile(r"\{l\s+s='((?:[^'\\]|\\.)*)'\s[^}]*d='([^']+)'")
SOURCE_PAT = re.compile(r'<source>(.*?)</source>', re.DOTALL)


def looks_translated(s: str, lang: str) -> bool:
    """True si el string parece ya escrito en el idioma destino (diacríticos)."""
    hint = LANG_HINTS.get(lang.split('-')[0].lower())
    if not hint:
        return False
    return any(ch in hint for ch in s)


def domain_to_key(domain: str) -> str:
    """Shop.Theme.Checkout → ShopThemeCheckout"""
    return domain.replace('.', '')


def scan_templates(base: str) -> dict:
    used = defaultdict(set)
    paths = (
        glob.glob(f'{base}/themes/**/*.tpl', recursive=True) +
        glob.glob(f'{base}/modules/**/*.tpl', recursive=True)
    )
    for path in paths:
        try:
            content = open(path, encoding='utf-8', errors='ignore').read()
            for m in L_PATTERN.finditer(content):
                s, d = m.group(1), m.group(2)
                if len(s.strip()) > 1:
                    used[d].add(s)
        except Exception:
            pass
    return {k: v for k, v in used.items()}


def load_translations(base: str, theme: str, lang: str, parent: str | None = None) -> dict:
    translated = defaultdict(set)
    dirs = [
        f'{base}/themes/{theme}/translations/{lang}',
        f'{base}/translations/{lang}',
    ]
    if parent:
        dirs.append(f'{base}/themes/{parent}/translations/{lang}')
    for d in dirs:
        if not os.path.isdir(d):
            continue
        for xlf in glob.glob(f'{d}/*.xlf'):
            name = os.path.basename(xlf)
            key = re.sub(r'\.(es-\w+|[a-z]{2}-[A-Z]{2})\.xlf$', '', name)
            try:
                content = open(xlf, encoding='utf-8').read()
                translated[key].update(SOURCE_PAT.findall(content))
            except Exception:
                pass
    return dict(translated)


def main():
    p = argparse.ArgumentParser(description='Detecta strings PS sin traducir')
    p.add_argument('--base', default='.', help='Raíz instalación PS')
    p.add_argument('--theme', default=None, help='Child theme (auto-detecta si se omite)')
    p.add_argument('--lang', default=None, help='Locale objetivo, ej. es-ES (auto si hay uno solo)')
    p.add_argument('--domain', default=None, help='Filtrar por dominio')
    p.add_argument('--output', default=None, help='Guardar JSON en fichero')
    p.add_argument('--include-admin', action='store_true')
    p.add_argument('--theme-only', action='store_true',
                   help='Solo dominios de storefront (Shop.*), ignora módulos terceros')
    args = p.parse_args()

    base = os.path.abspath(args.base)

    # ── Auto-detección de theme / locale ──────────────────────────────────────
    child, parent = detect_theme(base)
    theme = args.theme or child
    if not theme:
        print("ERROR: no se detectó child theme. Pasa --theme <nombre>.", file=sys.stderr)
        sys.exit(1)

    lang = args.lang
    if not lang:
        from detect import detect_locales  # type: ignore
        locales, _ = detect_locales(base)
        targets = [l for l in locales if l != 'en-US']
        if len(targets) == 1:
            lang = targets[0]
            print(f"ℹ️  Locale auto-detectado: {lang}", file=sys.stderr)
        else:
            print(f"ERROR: pasa --lang. Locales instalados: {', '.join(locales) or '(ninguno)'}",
                  file=sys.stderr)
            sys.exit(1)

    used = scan_templates(base)
    translated = load_translations(base, theme, lang, parent)

    missing = {}
    for domain, strings in sorted(used.items()):
        if args.domain and args.domain.lower() not in domain.lower():
            continue
        if args.theme_only and not domain.startswith('Shop.'):
            continue
        if not args.include_admin and ADMIN_DOMAINS.match(domain):
            continue
        key = domain_to_key(domain)
        existing = translated.get(key, set())
        need = sorted([
            s for s in strings
            if s not in existing and not looks_translated(s, lang) and len(s.strip()) > 1
        ])
        if need:
            missing[domain] = need

    total = sum(len(v) for v in missing.values())
    result = {
        'locale': lang,
        'theme': theme,
        'parent': parent,
        'base': base,
        'total': total,
        'domains': missing,
    }

    output = json.dumps(result, ensure_ascii=False, indent=2)

    if args.output:
        open(args.output, 'w', encoding='utf-8').write(output)
        print(f"✅ {total} strings faltantes en {len(missing)} dominios → {args.output}",
              file=sys.stderr)
    else:
        print(output)


if __name__ == '__main__':
    main()
