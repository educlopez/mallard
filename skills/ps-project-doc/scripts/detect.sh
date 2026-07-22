#!/usr/bin/env bash
# detect.sh — introspect a PrestaShop repo and emit JSON of facts for the
# ps-project-doc skill. Best-effort: missing values come back empty/null and
# the skill fills gaps by reading files directly. Never writes anything.
#
# Usage: detect.sh [repo-root]   (defaults to cwd)
set -uo pipefail
ROOT="${1:-.}"
cd "$ROOT" 2>/dev/null || { echo '{"error":"cannot cd to repo root"}'; exit 1; }

# JSON-encode a scalar string safely.
j() { python3 -c 'import json,sys; print(json.dumps(sys.argv[1]))' "${1:-}" 2>/dev/null || printf '"%s"' "${1:-}"; }
# Read a simple "key: value" scalar from a YAML file (first match).
yval() { grep -m1 -E "^[[:space:]]*$2[[:space:]]*:" "$1" 2>/dev/null \
  | sed -E "s/^[^:]*:[[:space:]]*//; s/[\"']//g; s/[[:space:]]+#.*$//; s/[[:space:]]*$//"; }

# --- PrestaShop version ---
PSVER=""
if [ -f composer.json ] && command -v jq >/dev/null 2>&1; then
  PSVER=$(jq -r '(.require["prestashop/prestashop"] // .require["prestashop/core"] // empty)' composer.json 2>/dev/null)
fi
[ -z "$PSVER" ] && [ -f composer.json ] && PSVER=$(grep -m1 -E "prestashop/(prestashop|core)" composer.json | grep -oE "[0-9]+\.[0-9]+(\.[0-9]+)?" | head -1)
[ -z "$PSVER" ] && [ -f app/AppKernel.php ] && PSVER=$(grep -m1 -iE "VERSION" app/AppKernel.php | grep -oE "[0-9]+\.[0-9]+\.[0-9.]+" | head -1)
[ -z "$PSVER" ] && [ -f config/settings.inc.php ] && PSVER=$(grep -m1 _PS_VERSION_ config/settings.inc.php | grep -oE "[0-9]+\.[0-9]+\.[0-9.]+" | head -1)
[ -z "$PSVER" ] && PSVER=$(grep -rhoE "define\('_PS_VERSION_', *'[0-9.]+'" config 2>/dev/null | grep -oE "[0-9]+\.[0-9]+\.[0-9.]+" | head -1)
# Fallback: infer major from the repo directory name (ps8-*, ps9-*, ps17-*, *-v9, *-v8).
if [ -z "$PSVER" ]; then
  BN=$(basename "$(pwd)")
  case "$BN" in
    ps17*|*-17*)  PSVER="1.7.x" ;;
    ps9*|*-v9*|*v9.*) PSVER="9.x" ;;
    ps8*|*-v8*|*v8.*) PSVER="8.x" ;;
  esac
fi

# --- Themes (name + parent per child) ---
THEMES_JSON="[]"
if compgen -G "themes/*/config/theme.yml" >/dev/null 2>&1; then
  rows=""
  for ty in themes/*/config/theme.yml; do
    dir=$(basename "$(dirname "$(dirname "$ty")")")
    rows="$rows{\"dir\":$(j "$dir"),\"name\":$(j "$(yval "$ty" name)"),\"parent\":$(j "$(yval "$ty" parent)")},"
  done
  THEMES_JSON="[${rows%,}]"
fi

# --- Theme stack ---
STACK="other/custom"
if [ -d modules/stsitebuilder ]; then
  STACK="elementflow"
elif ls -d themes/*/ 2>/dev/null | grep -qiE "panda"; then
  STACK="panda"
elif echo "$THEMES_JSON" | grep -qiE '"parent":[[:space:]]*"panda"'; then
  STACK="panda"
fi

# --- Lando ---
LNAME=""; MYSQLPORT=""; PHPVER=""
if [ -f .lando.yml ]; then
  LNAME=$(yval .lando.yml name)
  MYSQLPORT=$(grep -A8 -iE "database:" .lando.yml | grep -m1 -iE "portforward" | grep -oE "[0-9]+" | head -1)
  PHPVER=$(grep -m1 -iE "^[[:space:]]*php:" .lando.yml | grep -oE "[0-9]+\.[0-9]+" | head -1)
fi
LURL=""; [ -n "$LNAME" ] && LURL="https://$LNAME.lndo.site/"

# --- CSS build tooling ---
# Search package.json in root, _dev/, and any theme's _dev/ (globs, depth-limited).
CSSBUILD="none"
PKGS=$(ls package.json _dev/package.json _dev/*/package.json themes/*/package.json themes/*/_dev/package.json 2>/dev/null)
for pj in $PKGS; do
  [ -f "$pj" ] || continue
  if   grep -qiE "lightningcss|@parcel/css" "$pj"; then CSSBUILD="lightningcss"; break
  elif grep -qiE "webpack"                  "$pj"; then CSSBUILD="webpack"; break
  elif grep -qiE "\"vite\"|esbuild"         "$pj"; then CSSBUILD="vite/esbuild"; break
  elif grep -qiE "gulp"                     "$pj"; then CSSBUILD="gulp"; break
  fi
done
# Fallback: a _dev/ source dir exists but no known tool matched → still a build project.
if [ "$CSSBUILD" = "none" ]; then
  if ls -d _dev themes/*/_dev 2>/dev/null | grep -q .; then
    CSSBUILD="custom (_dev present — check package.json)"
  fi
fi

# --- Git remote ---
REMOTE=$(git remote get-url origin 2>/dev/null || echo "")

# --- Modules (dir names, capped) ---
MODS="[]"
if [ -d modules ]; then
  m=$(ls -d modules/*/ 2>/dev/null | xargs -n1 basename 2>/dev/null | head -60 \
      | python3 -c 'import sys,json; print(json.dumps([l.strip() for l in sys.stdin if l.strip()]))' 2>/dev/null)
  [ -n "$m" ] && MODS="$m"
fi

# --- CLAUDE.md status ---
HASCMD=false; [ -f CLAUDE.md ] && HASCMD=true
CMDIGN=false; grep -qxE "[[:space:]]*/?CLAUDE.md[[:space:]]*" .gitignore 2>/dev/null && CMDIGN=true

cat <<EOF
{
  "ps_version": $(j "$PSVER"),
  "php_version": $(j "$PHPVER"),
  "themes": $THEMES_JSON,
  "theme_stack": $(j "$STACK"),
  "lando_name": $(j "$LNAME"),
  "lando_url": $(j "$LURL"),
  "mysql_port": $(j "$MYSQLPORT"),
  "css_build": $(j "$CSSBUILD"),
  "git_remote": $(j "$REMOTE"),
  "modules": $MODS,
  "has_claudemd": $HASCMD,
  "claudemd_gitignored": $CMDIGN
}
EOF
