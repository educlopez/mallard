---
name: task
description: Turn an Intervals task into ready-to-work context — fetch it, build a brief, classify it, create the branch, and route it to the right tool. Intervals is the sole source of truth; Google Workspace is an opt-in enricher only.
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# task — Intervals task cockpit

Compresses "understand → execute" for a single Intervals task. Run it from **inside the client repo** you'll work in.

## Arguments

`$ARGUMENTS` — one of:

- `<localid>` — the numeric Intervals task ID (as shown in the web UI, e.g. `41953`). Fetch + brief + branch + route.
- `<localid> --enrich` — also pull Gmail/Drive context (opt-in only; never automatic).
- `<localid> --no-branch` — skip git branch creation.
- `<localid> --type <type>` — override the auto-classified type.
- `done <localid>` — finish flow: log time to Intervals (time only, no git).

## Instructions

Load and follow the **`task-context`** skill (`~/.claude/skills/task-context/SKILL.md`). It holds the full flow:

1. Fetch the task (`get_task`) + notes (`get_task_notes`), decode HTML entities.
2. Build the one-screen brief; print it AND save to engram (`task/<localid>/brief`).
3. If the task looks thin, only **label** it — never search anything unless `--enrich` is passed.
4. On `--enrich`: run the gws Gmail/Drive helper, fold snippets in.
5. Classify the type and name the routing skill/agent.
6. Create the git branch `<localid>` (or `<localid>-<slug>` on collision) unless `--no-branch`.
7. On `done`: run the time model (suggest a 15-min-band value, take the real one from the user, ask billable, then `add_time_entry`). Never commit or push.

Do NOT auto-enrich. Do NOT auto-log time. Keep everything Intervals-first.
