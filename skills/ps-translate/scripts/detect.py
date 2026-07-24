#!/usr/bin/env python3
"""
detect.py — Detecta el child theme y los idiomas instalados en una instalación PrestaShop.

Estrategia de idiomas: DB primero (tabla ps_lang vía parameters.php), fallback a
escaneo de filesystem (dirs translations/<locale>/). Robusto: cualquier fallo en DB
cae limpio a filesystem.

Uso:
  python detect.py --base /ruta/ps [--json]

Salida JSON:
  {
    "base": "/ruta/ps",
    "theme": "milagros",
    "parent": "panda",
    "locales": ["es-ES", "fr-FR"],
    "source": "db" | "filesystem"
  }
"""

import argparse
import glob
import json
import os
import re
import sys

LOCALE_RE = re.compile(r'^[a-z]{2}-[A-Z]{2}$')
PARAM_KEYS = (
    'database_host', 'database_port', 'database_name',
    'database_user', 'database_password', 'database_prefix',
)


# ── Theme ──────────────────────────────────────────────────────────────────────

def read_theme_yml(path: str) -> dict:
    """Parser mínimo de theme.yml — solo claves top-level tipo `name:` / `parent:`."""
    out = {}
    try:
        for line in open(path, encoding='utf-8', errors='ignore'):
            m = re.match(r'^([a-zA-Z_]+):\s*(.+?)\s*$', line)
            if m:
                out[m.group(1)] = m.group(2).strip().strip('"\'')
    except Exception:
        pass
    return out


def detect_theme(base: str) -> tuple[str | None, str | None]:
    """
    Devuelve (child_theme, parent_theme).
    Child = el theme cuyo theme.yml declara `parent:`. Si hay varios, prefiere el
    que no sea classic/panda/hummingbird (los padres típicos).
    """
    parents = {'classic', 'panda', 'hummingbird'}
    with_parent = []
    standalone = []  # theme.yml sin `parent:` (temas vendor usados directos)

    for yml in glob.glob(f'{base}/themes/*/config/theme.yml'):
        theme_dir = os.path.basename(os.path.dirname(os.path.dirname(yml)))
        data = read_theme_yml(yml)
        parent = data.get('parent')
        if parent:
            with_parent.append((theme_dir, parent))
        elif theme_dir.lower() not in parents:
            standalone.append(theme_dir)

    # 1º: child theme que declara parent (caso normal parent-child)
    for child, parent in with_parent:
        if child.lower() not in parents:
            return child, parent
    if with_parent:
        return with_parent[0]

    # 2º: tema vendor standalone usado directo (ej. at_movic sin child extraído)
    if standalone:
        return standalone[0], None

    return None, None


# ── Locales: DB ──────────────────────────────────────────────────────────────

def read_parameters(base: str) -> dict | None:
    """Lee parameters.php (PS 1.7/8) buscando credenciales DB."""
    for rel in ('app/config/parameters.php', 'config/parameters.php'):
        path = os.path.join(base, rel)
        if not os.path.isfile(path):
            continue
        try:
            content = open(path, encoding='utf-8', errors='ignore').read()
        except Exception:
            continue
        params = {}
        for key in PARAM_KEYS:
            m = re.search(rf"'{key}'\s*=>\s*'([^']*)'", content)
            if m is None:
                m = re.search(rf"'{key}'\s*=>\s*(\d+)", content)
            if m:
                params[key] = m.group(1)
        if params.get('database_name'):
            return params
    return None


def detect_locales_db(base: str) -> list[str] | None:
    """Consulta ps_lang.locale de idiomas activos. None si no se puede."""
    params = read_parameters(base)
    if not params:
        return None
    try:
        import pymysql  # type: ignore
    except ImportError:
        return None

    prefix = params.get('database_prefix', 'ps_')
    host = params.get('database_host', '127.0.0.1')
    port = int(params.get('database_port') or 3306)
    # host:port embebido en database_host (formato PS común)
    if ':' in host:
        host, _, p = host.partition(':')
        if p.isdigit():
            port = int(p)

    try:
        conn = pymysql.connect(
            host=host, port=port,
            user=params.get('database_user', ''),
            password=params.get('database_password', ''),
            database=params['database_name'],
            connect_timeout=4, read_timeout=4,
        )
        with conn.cursor() as cur:
            cur.execute(
                f"SELECT locale FROM `{prefix}lang` WHERE active = 1 AND locale <> ''"
            )
            rows = [r[0] for r in cur.fetchall() if r and r[0]]
        conn.close()
        return sorted(set(rows)) or None
    except Exception:
        return None


# ── Locales: filesystem ──────────────────────────────────────────────────────

def detect_locales_fs(base: str) -> list[str]:
    """Escanea dirs translations/<locale>/ en core + themes."""
    found = set()
    dirs = (
        glob.glob(f'{base}/translations/*') +
        glob.glob(f'{base}/themes/*/translations/*')
    )
    for d in dirs:
        if os.path.isdir(d):
            name = os.path.basename(d)
            if LOCALE_RE.match(name):
                found.add(name)
    return sorted(found)


def detect_locales(base: str) -> tuple[list[str], str]:
    """Devuelve (locales, source). DB primero, fallback filesystem."""
    db = detect_locales_db(base)
    if db:
        return db, 'db'
    return detect_locales_fs(base), 'filesystem'


# ── CLI ─────────────────────────────────────────────────────────────────────

def detect_all(base: str) -> dict:
    base = os.path.abspath(base)
    child, parent = detect_theme(base)
    locales, source = detect_locales(base)
    return {
        'base': base,
        'theme': child,
        'parent': parent,
        'locales': locales,
        'source': source,
    }


def main():
    p = argparse.ArgumentParser(description='Detecta theme e idiomas de PrestaShop')
    p.add_argument('--base', default='.', help='Raíz instalación PS')
    p.add_argument('--json', action='store_true', help='Salida JSON (default)')
    args = p.parse_args()

    result = detect_all(args.base)

    if args.json or True:
        print(json.dumps(result, ensure_ascii=False, indent=2))

    if not result['locales']:
        print("⚠️  No se detectaron locales instalados.", file=sys.stderr)
    if not result['theme']:
        print("⚠️  No se detectó child theme (theme.yml con parent).", file=sys.stderr)


if __name__ == '__main__':
    main()
