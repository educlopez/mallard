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
# Twig mail templates (PS_MAIL_THEME=modern/classic under mails/themes/**/*.twig) use a
# completely different translation call than Smarty {l}: 'string'|trans({}, 'Domain', locale).
# Same underlying Symfony translation domain/XLF mechanism as everything else — just a second
# syntax the scanner needs to recognize. Emails.Body is the domain used by mail bodies; email
# SUBJECT lines are set in core PHP (Emails.Subject) and aren't scannable from templates at all —
# see SKILL.md's email-translation notes.
TWIG_STRING_TRANS_START = re.compile(r"'((?:[^'\\]|\\.)*)'\s*\|\s*trans\(")
TWIG_DOMAIN_AFTER_PARAMS = re.compile(r"^\s*,\s*'([^']+)'")
SOURCE_PAT = re.compile(r'<source>(.*?)</source>', re.DOTALL)


def xml_unescape(s: str) -> str:
    return (s.replace('&lt;', '<').replace('&gt;', '>')
             .replace('&quot;', '"').replace('&amp;', '&'))


def find_twig_trans_calls(content: str):
    """Yields (source_string, domain) for every 'string'|trans({...}, 'Domain', locale) call.

    Some layouts nest a SECOND |trans() call inside the first argument's params dict, e.g.
    'Follow your order...'|trans({'%label%': 'Order history'|trans({}, 'Shop.Theme.X', locale)},
    'Emails.Body', locale) — a flat regex for the params dict (`\\{[^}]*\\}`) stops at the first
    `}` it sees, which is the INNER call's closing brace, then misreads the inner domain
    ('Shop.Theme.X') as if it were the outer one. Manually balancing braces from the opening `{`
    of the params dict correctly skips past the nested call to find the real outer domain.
    """
    for m in TWIG_STRING_TRANS_START.finditer(content):
        source = m.group(1)
        pos = m.end()  # just after 'trans('
        if pos >= len(content) or content[pos] != '{':
            continue
        depth = 0
        i = pos
        while i < len(content):
            if content[i] == '{':
                depth += 1
            elif content[i] == '}':
                depth -= 1
                if depth == 0:
                    i += 1
                    break
            i += 1
        else:
            continue  # unbalanced — skip rather than misparse
        dm = TWIG_DOMAIN_AFTER_PARAMS.match(content[i:i + 200])
        if dm:
            yield source, dm.group(1)


def looks_translated(s: str, lang: str) -> bool:
    """True si el string parece ya escrito en el idioma destino (diacríticos)."""
    hint = LANG_HINTS.get(lang.split('-')[0].lower())
    if not hint:
        return False
    return any(ch in hint for ch in s)


def domain_to_key(domain: str) -> str:
    """Shop.Theme.Checkout → ShopThemeCheckout"""
    return domain.replace('.', '')


def scan_templates(base: str, include_emails: bool = False) -> dict:
    used = defaultdict(set)
    # PDF invoices/delivery-slips/credit-slips live at the CORE root /pdf/**/*.tpl — outside
    # both themes/ and modules/, so they were silently never scanned before. Child-theme PDF
    # overrides (themes/<theme>/pdf/**) are already covered by the themes/**/*.tpl glob.
    paths = (
        glob.glob(f'{base}/themes/**/*.tpl', recursive=True) +
        glob.glob(f'{base}/modules/**/*.tpl', recursive=True) +
        glob.glob(f'{base}/pdf/**/*.tpl', recursive=True)
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

    if include_emails:
        # mails/themes/**/*.twig — the modern/classic Twig mail theme (PS_MAIL_THEME setting).
        # These templates are LANGUAGE-AGNOSTIC (one file for all locales; `locale` is a Twig
        # variable) — translation happens purely through the Emails.Body domain, same XLF flow
        # as everything else. The legacy mails/<lang>/*.html+.txt per-language file duplicates
        # are a SEPARATE, older mechanism this scanner does not touch — check which mail theme
        # is actually active (ps_configuration.PS_MAIL_THEME) before assuming either is live.
        for path in glob.glob(f'{base}/mails/themes/**/*.twig', recursive=True):
            try:
                content = open(path, encoding='utf-8', errors='ignore').read()
                for s, d in find_twig_trans_calls(content):
                    if len(s.strip()) > 1:
                        used[d].add(s)
            except Exception:
                pass

    return {k: v for k, v in used.items()}


def list_mail_domains(base: str) -> set:
    """Every domain referenced anywhere in mails/themes/**/*.twig (both the outer 'string'|trans
    call AND any domain used by a nested |trans() call inside its params dict, e.g. a translated
    %label% built from a DIFFERENT domain like Shop.Theme.Customeraccount).

    Why this matters: PS's mail-generation pipeline (prestashop:mail:generate) resolves ALL of
    these domains through a core-only translator that does NOT read theme-scoped translation
    dirs — confirmed empirically (see SKILL.md). So every domain this returns needs a CORE-level
    XLF copy (write_xlf.py --core-domains "<comma-separated list>"), not just Emails.Body, even
    if that domain is ALSO a normal theme-scoped domain used correctly elsewhere on the storefront.
    """
    domains = set()
    for path in glob.glob(f'{base}/mails/themes/**/*.twig', recursive=True):
        try:
            content = open(path, encoding='utf-8', errors='ignore').read()
            for _, d in find_twig_trans_calls(content):
                domains.add(d)
        except Exception:
            pass
    return domains


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
                # <source> content is XML-escaped on disk but `used` (from find_twig_trans_calls
                # / L_PATTERN) holds raw, unescaped strings — comparing escaped vs unescaped meant
                # ANY string containing <, >, ", or & (all the HTML-tag-bearing email strings)
                # never matched and kept reporting as "missing" even after being translated.
                translated[key].update(xml_unescape(s) for s in SOURCE_PAT.findall(content))
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
    p.add_argument('--include-emails', action='store_true',
                   help='Incluye plantillas de email Twig (mails/themes/**/*.twig, dominio Emails.Body)')
    p.add_argument('--list-mail-domains', action='store_true',
                   help='Solo imprime, separados por comas, TODOS los dominios referenciados en '
                        'mails/themes/**/*.twig (Emails.Body + cualquier otro dominio usado por un '
                        '|trans() anidado, ej. Shop.Theme.Customeraccount) — pásalo tal cual a '
                        'write_xlf.py --core-domains. No requiere --lang ni --theme.')
    args = p.parse_args()

    base = os.path.abspath(args.base)

    if args.list_mail_domains:
        print(','.join(sorted(list_mail_domains(base))))
        return

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

    used = scan_templates(base, include_emails=args.include_emails)
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
