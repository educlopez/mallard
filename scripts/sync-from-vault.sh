#!/usr/bin/env bash
# sync-from-vault.sh
# Sync the curated Panda + PrestaShop knowledge base from Eduardo's Obsidian
# vault into the plugin's skills/*/references/ folders, and scrub any
# author-local absolute paths so the resulting KB is self-contained.
#
# Run from the repo root: ./scripts/sync-from-vault.sh
#
# Idempotent: deletes destination references/ and re-copies fresh.

set -euo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
VAULT_KB="${PANDA_KB_SOURCE:-$HOME/Library/Mobile Documents/iCloud~md~obsidian/Documents/Second Brain/Resources/Panda Knowledge Base}"

if [[ ! -d "$VAULT_KB" ]]; then
  echo "ERROR: vault KB not found at: $VAULT_KB" >&2
  echo "Override with PANDA_KB_SOURCE=/path/to/Panda\\ Knowledge\\ Base" >&2
  exit 1
fi

PANDA_DST="$REPO_ROOT/skills/panda-kb/references"
PS_DST="$REPO_ROOT/skills/prestashop-kb/references"

# Stage to tmp dirs; only swap into place if all copies succeed.
# This protects against partial wipes when iCloud TCC blocks a copy mid-script.
STAGE_DIR="$(mktemp -d -t prestashop-experts-sync.XXXXXX)"
PANDA_STAGE="$STAGE_DIR/panda-kb"
PS_STAGE="$STAGE_DIR/prestashop-kb"
mkdir -p "$PANDA_STAGE" "$PS_STAGE"

cleanup() { rm -rf "$STAGE_DIR"; }
trap cleanup EXIT

echo "==> Syncing from: $VAULT_KB"
echo "==> Panda dst:    $PANDA_DST"
echo "==> PS dst:       $PS_DST"
echo "==> Staging at:   $STAGE_DIR"

# Copy file-by-file via `cat` to avoid macOS TCC issues with cp/rsync/tar/ditto
# on iCloud-protected paths. No metadata is preserved (we don't need any).
copy_tree() {
  local src="$1" dst="$2"
  local errors=0
  # find + read loop: relative path, mkdir parent, cat the file.
  # Suppress per-file failures to keep going on iCloud TCC errors; count them.
  while IFS= read -r -d '' f; do
    local rel="${f#$src/}"
    local target="$dst/$rel"
    mkdir -p "$(dirname "$target")" 2>/dev/null
    cat "$f" > "$target" 2>/dev/null || errors=$((errors + 1))
  done < <(find "$src" \( -name '.DS_Store' -prune \) -o -type f -print0 2>/dev/null)
  if [[ "$errors" -gt 0 ]]; then
    echo "  WARN: $errors file(s) failed to copy from $src" >&2
    return 1
  fi
  return 0
}

# --- panda-expert KB (to stage) ---
# theme Panda + Easy Builder + scraped docs + module fiches + demos
for sub in docs modules demos easybuilder forum; do
  if [[ -d "$VAULT_KB/$sub" ]]; then
    mkdir -p "$PANDA_STAGE/$sub"
    copy_tree "$VAULT_KB/$sub" "$PANDA_STAGE/$sub" || echo "  WARN: partial copy on $sub" >&2
  fi
done
# top-level orientation files
for f in README.md 00-Research-Report.md; do
  if [[ -f "$VAULT_KB/$f" ]]; then
    cat "$VAULT_KB/$f" > "$PANDA_STAGE/$f" 2>/dev/null || echo "  WARN: skipped $f" >&2
  fi
done

# --- prestashop-expert KB (to stage) ---
# platform-level docs (Symfony, smarty, themes, hooks, modules, install, version-migration)
if [[ -d "$VAULT_KB/prestashop" ]]; then
  copy_tree "$VAULT_KB/prestashop" "$PS_STAGE" || echo "  WARN: partial copy on prestashop/" >&2
fi

# Validate stage is non-empty before swapping
panda_stage_count=$(find "$PANDA_STAGE" -type f -name '*.md' | wc -l | tr -d ' ')
ps_stage_count=$(find "$PS_STAGE" -type f -name '*.md' | wc -l | tr -d ' ')

if [[ "$panda_stage_count" -lt 50 ]] || [[ "$ps_stage_count" -lt 30 ]]; then
  echo "ERROR: stage too small (panda=$panda_stage_count, ps=$ps_stage_count). Aborting without touching plugin tree." >&2
  echo "Current plugin KB left intact." >&2
  exit 2
fi

# Atomic swap: only now do we wipe and replace
rm -rf "$PANDA_DST" "$PS_DST"
mv "$PANDA_STAGE" "$PANDA_DST"
mv "$PS_STAGE" "$PS_DST"

# --- scrub absolute paths so KB is self-contained ---
echo "==> Scrubbing absolute paths..."
scrub() {
  local target="$1"
  # Author's local prestashop setup mirror — KB shouldn't depend on it.
  find "$target" -type f -name '*.md' -print0 | xargs -0 sed -i '' -E \
    -e 's|/Users/eduardocalvolopez/Documents/prestashop/setup/[^ )"`]*|<source-archive>/|g' \
    -e 's|/Users/eduardocalvolopez/Library/Mobile Documents/iCloud~md~obsidian/Documents/Second Brain/Resources/Panda Knowledge Base/|<KB>/|g' \
    -e 's|/Users/eduardocalvolopez/|<HOME>/|g'
}
scrub "$PANDA_DST"
scrub "$PS_DST"

# --- report ---
panda_files=$(find "$PANDA_DST" -type f -name '*.md' | wc -l | tr -d ' ')
ps_files=$(find "$PS_DST" -type f -name '*.md' | wc -l | tr -d ' ')
remaining=$(grep -rl '/Users/eduardocalvolopez' "$PANDA_DST" "$PS_DST" 2>/dev/null | wc -l | tr -d ' ')

echo ""
echo "Done."
echo "  panda-kb:       $panda_files markdown files"
echo "  prestashop-kb:  $ps_files markdown files"
echo "  remaining absolute /Users/eduardocalvolopez references: $remaining"
if [[ "$remaining" != "0" ]]; then
  echo "  (run: grep -rl /Users/eduardocalvolopez skills/  to inspect)"
fi
