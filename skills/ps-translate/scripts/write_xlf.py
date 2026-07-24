#!/usr/bin/env python3
"""
write_xlf.py — Escribe traducciones al XLF del child theme PS.

Acepta un JSON con traducciones agrupadas por dominio y las añade a los
archivos XLF correspondientes en themes/<theme>/translations/<lang>/.

Formato del JSON de entrada:
{
  "ShopThemePanda": {"Filter": "Filtrar", "Sort by": "Ordenar por"},
  "ShopThemeActions": {"Buy now": "Comprar ahora"}
}

Uso:
  python write_xlf.py --base /ruta/ps --theme milagros --lang es-CO \
                       --translations /tmp/traducciones.json

  # O desde stdin:
  cat traducciones.json | python write_xlf.py --base /ruta/ps --stdin
"""

import argparse
import hashlib
import json
import os
import re
import sys

try:
    from detect import detect_theme  # type: ignore
except ImportError:
    sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
    from detect import detect_theme

XLF_TEMPLATE = '''<?xml version="1.0" encoding="utf-8"?>
<xliff xmlns="urn:oasis:names:tc:xliff:document:1.2" version="1.2">
  <file source-language="en-US" target-language="{lang}" datatype="plaintext" original="file.ext">
    <header>
      <tool tool-id="symfony" tool-name="Symfony"/>
    </header>
    <body>
    </body>
  </file>
</xliff>'''


def xlf_id(source: str) -> str:
    return 'eco_' + hashlib.md5(source.encode()).hexdigest()[:7]


def xml_escape(s: str) -> str:
    return (s.replace('&', '&amp;')
             .replace('<', '&lt;')
             .replace('>', '&gt;')
             .replace('"', '&quot;'))


def write_domain(base: str, theme: str, lang: str,
                  domain_key: str, translations: dict, core: bool = False) -> tuple[int, int]:
    """
    Escribe traducciones al XLF del dominio.

    `core=True` escribe en la raíz `translations/<lang>/` (con puntos preservados en
    el nombre de fichero) en vez de `themes/<theme>/translations/<lang>/`. Necesario
    para dominios como `Emails.Body`/`Emails.Subject` (email transaccional generado
    por `prestashop:mail:generate`) y cualquier otro resuelto por el traductor Symfony
    "core" — confirmado empíricamente que este traductor NO lee directorios de
    traducción por-theme en absoluto (probado con y sin puntos en el nombre de
    fichero, en ambos casos "missing" hasta ponerlo en la raíz `translations/`).
    Esto es una EXCEPCIÓN deliberada a la norma general de "nunca tocar core, todo
    en el child theme" — solo para los dominios que genuinamente lo requieren.
    Returns (added, skipped).
    """
    if core:
        xlf_dir = os.path.join(base, 'translations', lang)
        filename_key = domain_key  # puntos preservados: Emails.Body.fr-FR.xlf
    else:
        xlf_dir = os.path.join(base, 'themes', theme, 'translations', lang)
        filename_key = domain_key.replace('.', '')  # convención theme: EmailsBody.fr-FR.xlf
    os.makedirs(xlf_dir, exist_ok=True)
    xlf_path = os.path.join(xlf_dir, f'{filename_key}.{lang}.xlf')

    if os.path.exists(xlf_path):
        content = open(xlf_path, encoding='utf-8').read()
    else:
        content = XLF_TEMPLATE.format(lang=lang)

    added = skipped = 0
    for source, target in translations.items():
        if not source.strip() or not target.strip():
            skipped += 1
            continue
        src_esc = xml_escape(source)
        if f'resname="{src_esc}"' in content or f'<source>{src_esc}</source>' in content:
            skipped += 1
            continue
        tgt_esc = xml_escape(target)
        entry = (
            f'      <trans-unit id="{xlf_id(source)}" resname="{src_esc}">\n'
            f'        <source>{src_esc}</source>\n'
            f'        <target>{tgt_esc}</target>\n'
            f'      </trans-unit>'
        )
        content = content.replace('    </body>', entry + '\n    </body>')
        added += 1

    if added:
        open(xlf_path, 'w', encoding='utf-8').write(content)

    return added, skipped


def main():
    p = argparse.ArgumentParser(description='Escribe traducciones PS a XLF')
    p.add_argument('--base', default='.', help='Raíz instalación PS')
    p.add_argument('--theme', default=None, help='Child theme (auto-detecta si se omite)')
    p.add_argument('--lang', default=None, help='Locale objetivo, ej. es-ES (requerido)')
    p.add_argument('--translations', default=None,
                   help='Ruta al JSON de traducciones')
    p.add_argument('--stdin', action='store_true',
                   help='Leer JSON desde stdin')
    p.add_argument('--dry-run', action='store_true',
                   help='Preview sin escribir')
    p.add_argument('--core-domains', default='',
                   help='Lista separada por comas de dominios (con puntos, ej. '
                        '"Emails.Body,Emails.Subject") que deben escribirse en la '
                        'raíz translations/<lang>/ en vez de themes/<theme>/translations/ '
                        '— necesario para dominios resueltos por el traductor Symfony '
                        '"core" (emails transaccionales). Ver docstring de write_domain().')
    args = p.parse_args()

    if args.stdin:
        data = json.load(sys.stdin)
    elif args.translations:
        data = json.load(open(args.translations, encoding='utf-8'))
    else:
        print("ERROR: Necesitas --translations <ruta> o --stdin", file=sys.stderr)
        sys.exit(1)

    base = os.path.abspath(args.base)

    # ── Auto-detección de theme / validación de locale ────────────────────────
    theme = args.theme or detect_theme(base)[0]
    if not theme:
        print("ERROR: no se detectó child theme. Pasa --theme <nombre>.", file=sys.stderr)
        sys.exit(1)
    lang = args.lang or data.get('locale') if isinstance(data, dict) else args.lang
    if not lang:
        print("ERROR: pasa --lang <locale> (ej. es-ES).", file=sys.stderr)
        sys.exit(1)

    # El JSON puede venir en dos formatos:
    # 1. Plano: {"DomainKey": {"source": "target", ...}, ...}
    # 2. Anidado del scan: {"domains": {"Domain.With.Dots": [...], ...}}
    #    En este caso las claves del dict de traducciones son los dominios con puntos

    core_domains = {d.strip() for d in args.core_domains.split(',') if d.strip()}

    total_added = total_skipped = 0
    wrote_core = False

    for domain_key, translations in data.items():
        if domain_key in ('locale', 'theme', 'parent', 'base', 'total', 'domains'):
            continue
        if not isinstance(translations, dict):
            continue

        is_core = domain_key in core_domains
        # Normalizar: quitar puntos si los tiene (Shop.Theme.X → ShopThemeX) — salvo
        # que este dominio vaya a core, donde el traductor Symfony espera el nombre
        # de fichero con puntos preservados (ver write_domain).
        normalized_key = domain_key if is_core else domain_key.replace('.', '')

        if args.dry_run:
            dest = 'CORE translations/' if is_core else f'themes/{theme}/translations/'
            print(f"  [DRY] {normalized_key} → {dest}: {len(translations)} strings")
            continue

        added, skipped = write_domain(
            base, theme, lang, normalized_key, translations, core=is_core
        )
        total_added += added
        total_skipped += skipped
        wrote_core = wrote_core or (is_core and added)
        if added:
            dest_label = ' (CORE)' if is_core else ''
            print(f"  ✅ {normalized_key}{dest_label}: +{added} añadidos ({skipped} ya existían)")

    if not args.dry_run:
        print(f"\n🎉 Total: {total_added} traducciones escritas, {total_skipped} ya existían")
        print(f"\nPróximos pasos:")
        print(f"  php bin/console cache:clear   # rm -rf var/cache/* NO es suficiente, hace falta el comando")
        if wrote_core:
            print(f"  # Dominios core escritos — si son de email, regenera las plantillas:")
            print(f"  php bin/console prestashop:mail:generate <mail-theme> {lang} --overwrite")
        print(f"  git add themes/{theme}/translations/" + (" translations/" if wrote_core else ""))
        print(f"  git commit -m 'feat(i18n): traducciones automáticas ps-translate'")


if __name__ == '__main__':
    main()
