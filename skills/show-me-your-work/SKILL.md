---
name: show-me-your-work
description: >
  Keep a reviewable decision trail (one TSV row per decision) for long or
  unattended work. Use when the user says "deja rastro de las decisiones",
  "show me your work", "registro de decisiones", or for long multi-phase jobs a
  human reviews later: a knowledge-base rewrite, a theme migration, a bulk
  translation, a multi-store rollout.
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# show-me-your-work

Adapted from pstack's `show-me-your-work`. Only for long tasks. Skip it for anything that fits in one sitting.

## Log

One file, one row per decision, append-only, cells on a single line. Evidence is a pointer, never prose.

```
ts	phase	decision	why	evidence	result
```

- `ts`: ISO 8601. `phase`: workstream.
- `decision`: what was chosen or done, one line.
- `why`: the reason in plain words.
- `evidence`: commit SHA, `file:line`, screenshot or artifact path.
- `result`: `tests green`, `reverted`, `open`, `INCONCLUSIVE`.

Create the header on first use. If a cell starts with `=`, `+`, `-` or `@`, prefix it with `'` so spreadsheets do not read it as a formula.

## What to log

Forks chosen, units finished with their verification result, pivots and reverts with the trigger, blockers, fixed gates. Skip trivial actions. A wrong call gets a new row that supersedes it; never edit history.

## Where it lives

Default `decisions.tsv` in the work directory, kept out of git (`.git/info/exclude`). Commit it only when a reviewer needs the trail to trust the result, and only if the user agrees.

## Before handing back

Check each row maps to a real action and its evidence resolves and shows what the row claims. Fix or supersede any row that does not.
