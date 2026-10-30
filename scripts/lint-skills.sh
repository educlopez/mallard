#!/usr/bin/env bash
# Validate frontmatter, local links and README coverage for skills, commands
# and agents. No dependencies beyond bash, awk, grep and sed.
set -u
cd "$(dirname "$0")/.."

fail=0
err() { echo "ERROR $1: $2"; fail=1; }

fm()    { awk 'NR==1&&$0!="---"{exit} NR>1&&$0=="---"{exit} NR>1{print}' "$1"; }
field() { fm "$1" | grep -E "^$2:" | head -1 | sed -E "s/^$2:[[:space:]]*//; s/^[\"']//; s/[\"']\$//"; }

# check_meta <file> <expected-name> <name-required:0|1>
check_meta() {
  local f=$1 want=$2 need_name=$3 name desc ver
  [ "$(head -1 "$f")" = "---" ] || { err "$f" "missing frontmatter"; return; }
  name=$(field "$f" name)
  if [ -n "$name" ]; then
    [ "$name" = "$want" ] || err "$f" "name '$name' != '$want'"
  elif [ "$need_name" = 1 ]; then
    err "$f" "missing name"
  fi
  desc=$(field "$f" description)
  if [ -z "$desc" ]; then
    err "$f" "missing description"
  elif [[ $desc =~ ^[\>\|]-?$ ]] && ! fm "$f" | grep -A1 -E '^description:' | tail -1 | grep -qE '^[[:space:]]+[^[:space:]]'; then
    err "$f" "empty description block"
  fi
  ver=$(field "$f" version)
  [[ $ver =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]] || err "$f" "version '$ver' is not semver (X.Y.Z)"
}

for d in skills/*/; do
  s=${d%/}; s=${s#skills/}
  [ -f "${d%/}/SKILL.md" ] || { err "$d" "no SKILL.md"; continue; }
  check_meta "${d%/}/SKILL.md" "$s" 1
  grep -qF "\`$s\`" README.md || err README.md "skill '$s' not listed"
  # local references/ paths must exist (scripts/ in skills often name target-project files)
  for r in $(grep -oE '(^|[^A-Za-z0-9_/.-])references/[A-Za-z0-9._/-]+' "${d%/}/SKILL.md" | sed -E 's/^[^r]//' | sort -u); do
    [[ $r == */ ]] && continue
    [ -e "${d%/}/$r" ] || err "${d%/}/SKILL.md" "broken path $r"
  done
done

for f in claude/commands/*.md; do
  c=$(basename "$f" .md)
  check_meta "$f" "$c" 0
  grep -qF "\`/$c\`" README.md || err README.md "command '/$c' not listed"
done

for f in claude/agents/*.md; do
  a=$(basename "$f" .md)
  check_meta "$f" "$a" 1
  grep -qF "\`$a\`" README.md || err README.md "agent '$a' not listed"
done

[ $fail = 0 ] && echo "skills lint: ok"
exit $fail
