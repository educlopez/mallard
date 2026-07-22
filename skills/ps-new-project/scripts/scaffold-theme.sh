#!/usr/bin/env bash
# scaffold-theme.sh — create a PrestaShop child-theme skeleton for a new project.
# Two tracks: panda (parent: panda, legacy) and elementflow (parent: classic +
# builder layouts, emerging default). Creates dirs + theme.yml + config .htaccess.
# Does NOT touch git, the server, or the CSS build (ps-css-build owns that).
#
# Usage: scaffold-theme.sh <panda|elementflow> <name> [display_name] [repo-root]
set -euo pipefail

TRACK="${1:-}"
NAME="${2:-}"
DISPLAY="${3:-$NAME}"
ROOT="${4:-.}"

[ -z "$TRACK" ] || [ -z "$NAME" ] && { echo "usage: scaffold-theme.sh <panda|elementflow> <name> [display] [root]" >&2; exit 2; }
case "$TRACK" in panda|elementflow) ;; *) echo "track must be panda|elementflow" >&2; exit 2 ;; esac

cd "$ROOT" || { echo "cannot cd to $ROOT" >&2; exit 1; }
[ -d themes ] || { echo "no themes/ dir here — is this a PrestaShop repo root?" >&2; exit 1; }

DEST="themes/$NAME"
if [ -e "$DEST" ]; then
  echo "⚠ $DEST already exists — refusing to clobber. Remove it or pick another --name." >&2
  exit 1
fi

# Dir skeleton.
mkdir -p "$DEST"/config \
         "$DEST"/templates/{_partials,catalog,checkout,customer,errors} \
         "$DEST"/modules \
         "$DEST"/assets/{css,js}
touch "$DEST"/templates/_partials/.gitkeep \
      "$DEST"/modules/.gitkeep \
      "$DEST"/assets/css/.gitkeep \
      "$DEST"/assets/js/.gitkeep

# config/.htaccess — deny direct access (Apache 2.2 + 2.4 guarded).
cat > "$DEST/config/.htaccess" <<'HT'
<IfModule mod_authz_core.c>
    Require all denied
</IfModule>
<IfModule !mod_authz_core.c>
    Order deny,allow
    Deny from all
</IfModule>
HT

# theme.yml per track.
if [ "$TRACK" = "panda" ]; then
  cat > "$DEST/config/theme.yml" <<YML
parent: panda
name: $NAME
display_name: $DISPLAY
version: 1.0
assets:
  use_parent_assets: true
YML
else
  cat > "$DEST/config/theme.yml" <<YML
parent: classic
name: $NAME
display_name: $DISPLAY
version: 1.0.0
author:
  name: "Cinetic"
  email: ""
  url: ""
meta:
  compatibility:
    from: 1.7.1.0
    to: '~'
  available_layouts:
    layout-full-width:
      name: Full Width
      description: No side columns, ideal for distraction-free pages such as product pages.
    layout-both-columns:
      name: Three Columns
      description: One large central column and 2 side columns.
    layout-left-column:
      name: Two Columns, small left column
      description: Two columns with a small left column
    layout-right-column:
      name: Two Columns, small right column
      description: Two columns with a small right column
assets:
  use_parent_assets: true
global_settings:
  configuration:
    PS_IMAGE_QUALITY: png
  image_types:
    cart_default:
      width: 125
      height: 125
      scope: [ products ]
    small_default:
      width: 98
      height: 98
      scope: [ products, categories, manufacturers, suppliers ]
    medium_default:
      width: 452
      height: 452
      scope: [ products, manufacturers, suppliers ]
    home_default:
      width: 250
      height: 250
      scope: [ products ]
    large_default:
      width: 800
      height: 800
      scope: [ products, manufacturers, suppliers ]
    category_default:
      width: 141
      height: 180
      scope: [ categories ]
YML
fi

echo "✓ scaffolded $TRACK child theme at $DEST"
echo "  next: ps-css-build (wire the CSS build) → ps-project-doc (CLAUDE.md) → activate in BO"
