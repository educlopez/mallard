#!/usr/bin/env bash
# gmail_search.sh — opt-in Gmail enricher for the task-context skill.
# Searches the user's mailbox via the `gws` CLI and prints COMPACT JSON
# (one object per message: from / date / subject / snippet) so only a small
# summary enters the model context — never raw email bodies.
#
# Usage: gmail_search.sh "<client>" "<pm-owner>" "<keywords>" [max]
#   client   free-text client/company name (from the Intervals task `client`)
#   pm-owner PM display name (from the task `owners`) — used as a soft from: hint
#   keywords 2-4 salient words from the task title
#   max      max messages to fetch (default 5)
#
# This script is ONLY run when the user passes --enrich. It is never called
# automatically.
set -euo pipefail

CLIENT="${1:-}"
PM="${2:-}"
KEYWORDS="${3:-}"
MAX="${4:-5}"

if ! command -v gws >/dev/null 2>&1; then
  echo '{"error":"gws CLI not found on PATH"}'
  exit 1
fi

# Build the Gmail search query. Keep recent (180d) to avoid ancient noise.
# PM name is a soft hint; drop it if it yields nothing (caller can retry).
QUERY="${KEYWORDS} ${CLIENT} newer_than:180d"
QUERY="$(echo "$QUERY" | sed 's/^ *//;s/ *$//')"

# Pick a JSON processor.
if command -v jq >/dev/null 2>&1; then
  PARSE=jq
elif command -v python3 >/dev/null 2>&1; then
  PARSE=python3
else
  echo '{"error":"need jq or python3 to parse gws output"}'
  exit 1
fi

# 1) List matching message ids.
LIST_PARAMS=$(printf '{"userId":"me","q":"%s","maxResults":%s}' "$QUERY" "$MAX")
LIST_JSON="$(gws gmail users messages list --params "$LIST_PARAMS" --format json 2>/dev/null || echo '{}')"

if [ "$PARSE" = jq ]; then
  IDS="$(echo "$LIST_JSON" | jq -r '.messages[]?.id // empty')"
else
  IDS="$(echo "$LIST_JSON" | python3 -c 'import sys,json;d=json.load(sys.stdin);[print(m["id"]) for m in d.get("messages",[])]' 2>/dev/null || true)"
fi

if [ -z "${IDS}" ]; then
  printf '{"query":"%s","results":[],"note":"no matches — try different keywords or add the PM email"}\n' "$QUERY"
  exit 0
fi

# 2) For each id, fetch metadata headers + snippet only (cheap).
echo "{\"query\":\"$QUERY\",\"results\":["
FIRST=1
for id in $IDS; do
  GET_PARAMS=$(printf '{"userId":"me","id":"%s","format":"metadata","metadataHeaders":["From","Subject","Date"]}' "$id")
  MSG="$(gws gmail users messages get --params "$GET_PARAMS" --format json 2>/dev/null || echo '{}')"
  if [ "$PARSE" = jq ]; then
    ROW="$(echo "$MSG" | jq -c '{from:(.payload.headers[]?|select(.name=="From")|.value), date:(.payload.headers[]?|select(.name=="Date")|.value), subject:(.payload.headers[]?|select(.name=="Subject")|.value), snippet:.snippet}')"
  else
    ROW="$(echo "$MSG" | python3 -c '
import sys,json
d=json.load(sys.stdin)
h={x["name"]:x["value"] for x in d.get("payload",{}).get("headers",[])}
print(json.dumps({"from":h.get("From"),"date":h.get("Date"),"subject":h.get("Subject"),"snippet":d.get("snippet")}))' 2>/dev/null || echo '{}')"
  fi
  if [ "$FIRST" -eq 1 ]; then FIRST=0; else printf ','; fi
  printf '%s' "$ROW"
done
echo "]}"
