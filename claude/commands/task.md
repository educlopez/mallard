---
name: task
description: Turn a tracker task into ready-to-work context — fetch it, build a brief, classify it, create the branch, and route it to the right tool. The configured tracker is the sole source of truth; Google Workspace is an opt-in enricher only.
version: "0.2.0"
metadata:
  author: Eduardo Calvo
---

# task — tracker task cockpit

Compresses "understand → execute" for a single tracker task. Run it from **inside the client repo** you'll work in.

## Arguments

`$ARGUMENTS` — one of:

- `<number>` — the numeric tracker task ID (as shown in the web UI, e.g. `41953`). Fetch + brief + branch + route.
- `<number> --enrich` — also pull Gmail/Drive context (opt-in only; never automatic).
- `<number> --no-branch` — skip git branch creation.
- `<number> --type <type>` — override the auto-classified type.
- `done <number>` — finish flow: log time to the tracker (time only, no git).

## Instructions

Load and follow the **`task-context`** skill (`~/.claude/skills/task-context/SKILL.md`). It holds the full flow:

0. Read the tracker config (`~/.config/mallard/tracker.md`); if missing, offer to create it from the skill's `tracker.template.md`.
1. Fetch the task with the configured `get_task` operation; open any attachment it mentions.
2. Build the one-screen brief; print it AND save to engram (`task/<number>/brief`).
3. If the task looks thin, only **label** it — never search anything unless `--enrich` is passed.
4. On `--enrich`: run the gws Gmail/Drive helper, fold snippets in.
5. Classify the type and name the routing skill/agent.
6. Create the git branch from the configured `branch.pattern` unless `--no-branch`.
7. On `done`: run the time model (suggest a value rounded up as configured, take the real one from the user in the tracker's time format, confirm, then the configured `log_time` with a short note of the work done, always, unless the user says to omit it). Never commit or push.

Do NOT auto-enrich. Do NOT auto-log time. Keep everything tracker-first.
