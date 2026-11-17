#!/usr/bin/env bash
# detect-context.sh — collect facts a project's AGENTS.md usually lacks, from
# the repo itself. Stack-aware (prestashop | magento | laravel | generic).
# Best-effort and read-only: empty values mean "not detected", never "absent".
#
# Usage: detect-context.sh [repo-root]   (defaults to cwd)
set -uo pipefail
cd "${1:-.}" 2>/dev/null || { echo '{"error":"cannot cd to repo root"}'; exit 1; }

j()    { python3 -c 'import json,sys; print(json.dumps(sys.argv[1]))' "${1:-}" 2>/dev/null || printf '"%s"' "${1:-}"; }
jarr() { python3 -c 'import sys,json; print(json.dumps([l.rstrip("\n") for l in sys.stdin if l.strip()]))' 2>/dev/null || echo '[]'; }

# --- Stack ---
STACK="generic"
if [ -f bin/magento ] || grep -qs '"magento/' composer.json; then STACK="magento"
elif [ -d themes ] && [ -d modules ] && { [ -d classes ] || [ -d src ] || grep -qs 'prestashop' composer.json; }; then STACK="prestashop"
elif [ -f artisan ] && grep -qs 'laravel/framework' composer.json; then STACK="laravel"
fi

# --- Who to ask: top 2 committers (name only, last 18 months, no merges, no bots) ---
TOP=$(git log --no-merges --since="18 months ago" --format='%aN' 2>/dev/null \
  | grep -viE 'bot|dependabot|renovate|gitlab|github' | sort | uniq -c | sort -rn | head -2 \
  | sed -E 's/^ *([0-9]+) (.*)$/\2 (\1 commits)/' | jarr)

# --- Overview hint: description/name from the usual manifests ---
DESC=""
[ -f composer.json ] && command -v jq >/dev/null 2>&1 && DESC=$(jq -r '.description // empty' composer.json 2>/dev/null)
[ -z "$DESC" ] && [ -f package.json ] && command -v jq >/dev/null 2>&1 && DESC=$(jq -r '.description // empty' package.json 2>/dev/null)

# --- Languages / stores ---
LANGS="[]"
case "$STACK" in
  prestashop) LANGS=$(ls -d translations/*/ 2>/dev/null | xargs -n1 basename 2>/dev/null | grep -E '^[a-z]{2}-[A-Z]{2}$' | jarr) ;;
  magento)    LANGS=$(ls -d app/i18n/*/* 2>/dev/null | jarr) ;;
  laravel)    LANGS=$( (ls -d lang/*/ resources/lang/*/ 2>/dev/null | xargs -n1 basename 2>/dev/null) | jarr) ;;
esac

# --- Custom code (team-owned, not vendor) ---
CUSTOM="[]"
case "$STACK" in
  prestashop) # tracked module dirs that are not standard core/partner modules
    CUSTOM=$(git ls-files modules 2>/dev/null | cut -d/ -f2 | sort -u \
      | grep -E '^[A-Za-z0-9_-]+$' | grep -vE '^(ps_|block|dash|graph|grid|pages|product|contact|gsitemap|psgdpr|stats|index)' | head -30 | jarr) ;;
  magento)    CUSTOM=$(ls -d app/code/*/* 2>/dev/null | head -30 | jarr) ;;
  laravel)    CUSTOM=$(ls -d app/Services/* app/Modules/* 2>/dev/null | head -30 | jarr) ;;
esac

# --- Patches to review on upgrade ---
PATCHES="[]"
{ ls patches/*.patch patches/*/*.patch 2>/dev/null; ls -d override/*/ 2>/dev/null; grep -ls 'cweagans/composer-patches' composer.json 2>/dev/null; } | head -20 | jarr > /tmp/.agmd_patches.$$ 2>/dev/null && PATCHES=$(cat /tmp/.agmd_patches.$$); rm -f /tmp/.agmd_patches.$$

# --- Integrations (keywords in manifests and module names; a hint, never a fact) ---
INTEG=$( { cat composer.json package.json 2>/dev/null; ls modules app/code/* 2>/dev/null; } \
  | grep -oiE 'erp|sap|odoo|holded|navision|business.?central|stripe|paypal|redsys|mollie|adyen|klarna|gls|correos|dhl|ups|seur|mondial|akeneo|pimcore|salsify|mailchimp|klaviyo|brevo|sendinblue|algolia|elasticsearch|meilisearch|api.?client' \
  | tr 'A-Z' 'a-z' | sort -u | jarr)

# --- Automated processes ---
CRON="[]"
{ ls crontab* cron/ cronjobs/ 2>/dev/null; find . -maxdepth 4 \( -name 'cron_groups.xml' -o -name 'crontab.xml' -o -name '*cron*.php' \) -not -path './vendor/*' -not -path './node_modules/*' 2>/dev/null; grep -ls 'schedule->' app/Console/Kernel.php routes/console.php 2>/dev/null; } | head -15 | jarr > /tmp/.agmd_cron.$$ 2>/dev/null && CRON=$(cat /tmp/.agmd_cron.$$); rm -f /tmp/.agmd_cron.$$

# --- Tests ---
TESTS=$( { ls phpunit.xml* playwright.config.* cypress.config.* codeception.yml 2>/dev/null; ls -d tests/ test/ 2>/dev/null; } | jarr)

# --- Environments (names only; never URLs with credentials) ---
ENVS=$( { ls .lando.yml docker-compose*.yml .ddev/config.yaml 2>/dev/null; grep -hoE '^[[:space:]]*environment:[[:space:]]*\n?[[:space:]]*name:[[:space:]]*[a-z0-9_-]+' .gitlab-ci.yml 2>/dev/null; } | jarr)
CI="none"; [ -d .github/workflows ] && CI="github"; [ -f .gitlab-ci.yml ] && CI="gitlab"

# --- What broke before: recent fix/revert/hotfix subjects ---
BROKE=$(git log --no-merges -i -E --grep='^(fix|hotfix|revert)|regression|broke|rollback' -n 40 --format='%h %s' 2>/dev/null | grep -vE '^[0-9a-f]+ (feat|refactor|chore|docs|style)' | head -12 | jarr)

# --- Branching ---
DEFB=$(git symbolic-ref --short refs/remotes/origin/HEAD 2>/dev/null | sed 's#^origin/##')
BRANCHES=$(git for-each-ref --format='%(refname:short)' refs/remotes/origin 2>/dev/null | sed 's#^origin/##' | grep -vE '^(HEAD|flow/|feature/|fix/|docs/)' | head -8 | jarr)

# --- Existing AGENTS.md: which template sections are already present ---
SECT="[]"
[ -f AGENTS.md ] && SECT=$(grep -E '^## ' AGENTS.md | sed 's/^## //' | jarr)
CHARS=0; [ -f AGENTS.md ] && CHARS=$(wc -c < AGENTS.md | tr -d ' ')

# --- Agent files: ignored? tracked? (an ignore rule hides a missing file from git add)
AGIGN=false; git check-ignore -q AGENTS.md 2>/dev/null && AGIGN=true
CLIGN=false; git check-ignore -q CLAUDE.md 2>/dev/null && CLIGN=true
CLTRACK=false; [ -n "$(git ls-files CLAUDE.md 2>/dev/null)" ] && CLTRACK=true
CLLINES=0; [ -f CLAUDE.md ] && CLLINES=$(wc -l < CLAUDE.md | tr -d ' ')

# --- Working tree: dirty or behind the remote default branch means use a clean worktree
DIRTY=$(git status --porcelain 2>/dev/null | wc -l | tr -d ' ')
BEHIND=0
if [ -n "$DEFB" ]; then BEHIND=$(git rev-list --count "HEAD..origin/$DEFB" 2>/dev/null || echo 0); fi

cat <<JSON
{
  "stack": $(j "$STACK"),
  "description_hint": $(j "$DESC"),
  "top_committers": $TOP,
  "languages": $LANGS,
  "custom_code": $CUSTOM,
  "patches_and_overrides": $PATCHES,
  "integration_hints": $INTEG,
  "automated_processes": $CRON,
  "tests": $TESTS,
  "environments": $ENVS,
  "ci": $(j "$CI"),
  "default_branch": $(j "$DEFB"),
  "long_lived_branches": $BRANCHES,
  "recent_fixes": $BROKE,
  "agents_md_sections": $SECT,
  "agents_md_chars": $CHARS,
  "agents_md_ignored": $AGIGN,
  "claude_md_tracked": $CLTRACK,
  "claude_md_ignored": $CLIGN,
  "claude_md_lines": $CLLINES,
  "dirty_entries": $DIRTY,
  "behind_origin": $BEHIND
}
JSON
