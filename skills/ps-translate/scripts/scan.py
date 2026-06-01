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

# ── Detección de strings ya en español ────────────────────────────────────────

SPANISH_CHARS = re.compile(r'[áéíóúÁÉÍÓÚñÑüÜ¡¿]')
SPANISH_STARTERS = re.compile(
    r'^(El |La |Los |Las |Un |Una |De |Del |En |Con |Por |Para |Sin |Sobre |'
    r'Este |Esta |No |Sí |Hola |Bienvenid|Añadir|Ver |Editar|Crear|'
    r'Cancelar|Guardar|Cerrar|Enviar|Buscar|Continuar|Confirmar|'
    r'Seleccionar|Mostrar|Ocultar|Ordenar|Filtrar|Pedido|Carrito|'
    r'Dirección|Pago|Envío|Precio|Producto|Cantidad|Total)',
    re.IGNORECASE
)
ADMIN_DOMAINS = re.compile(
    r'^Admin\.|^psgdpr|^ps_themecusto|^steasybuilder|'
    r'^ps_facebook|^ps_accounts|^allinone|^AdminTheme'
)

L_PATTERN = re.compile(r"\{l\s+s='((?:[^'\\]|\\.)*)'\s[^}]*d='([^']+)'")
SOURCE_PAT = re.compile(r'<source>(.*?)</source>', re.DOTALL)


def is_spanish(s: str) -> bool:
    return bool(SPANISH_CHARS.search(s)) or bool(SPANISH_STARTERS.match(s))


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


def load_translations(base: str, theme: str, lang: str) -> dict:
    translated = defaultdict(set)
    dirs = [
        f'{base}/themes/{theme}/translations/{lang}',
        f'{base}/translations/{lang}',
        f'{base}/themes/panda/translations/{lang}',
        f'{base}/themes/panda/translations/es-ES',
    ]
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
    p.add_argument('--theme', default='milagros', help='Child theme')
    p.add_argument('--lang', default='es-CO', help='Locale objetivo')
    p.add_argument('--domain', default=None, help='Filtrar por dominio')
    p.add_argument('--output', default=None, help='Guardar JSON en fichero')
    p.add_argument('--include-admin', action='store_true')
    args = p.parse_args()

    base = os.path.abspath(args.base)

    used = scan_templates(base)
    translated = load_translations(base, args.theme, args.lang)

    missing = {}
    for domain, strings in sorted(used.items()):
        if args.domain and args.domain.lower() not in domain.lower():
            continue
        if not args.include_admin and ADMIN_DOMAINS.match(domain):
            continue
        key = domain_to_key(domain)
        existing = translated.get(key, set())
        need = sorted([
            s for s in strings
            if s not in existing and not is_spanish(s) and len(s.strip()) > 1
        ])
        if need:
            missing[domain] = need

    total = sum(len(v) for v in missing.values())
    result = {
        'locale': args.lang,
        'theme': args.theme,
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
