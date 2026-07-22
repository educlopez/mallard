---
name: project-doc
description: Generate a standardized private CLAUDE.md for the current PrestaShop repo by introspecting it. Auto-fills stack/theme/lando/build/deploy, asks briefly for gotchas/contacts, and keeps the file gitignored (never committed).
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# project-doc — standardized private CLAUDE.md

Run inside a PrestaShop repo. Writes a lean, consistent `CLAUDE.md` so Claude stops guessing the stack each session. **The file is private — always gitignored, never committed.**

## Arguments

`$ARGUMENTS` (optional):
- (none) — create `CLAUDE.md` if missing.
- `--refresh` — regenerate even if one exists (backs up to `CLAUDE.md.bak` first).
- `--print` — preview only, write nothing.

## Instructions

Load and follow the **`ps-project-doc`** skill (`~/.claude/skills/ps-project-doc/SKILL.md`):

1. Verify the cwd is a PrestaShop repo.
2. Run `scripts/detect.sh` to auto-detect PS version, theme + parent, theme stack (Panda/ElementFlow), Lando URL/port, CSS build, git remote, modules. Fill gaps by reading source files.
3. Ask a short batch for gotchas / contacts / module purposes (skippable).
4. Fill the template; inject the fixed Deploy (git pull over SSH; test=staging, master=prod) and override-priority text.
5. Respect `--refresh` / `--print`. Ensure `.gitignore` contains `CLAUDE.md`. Never stage or commit it.

Never commit CLAUDE.md. Never overwrite an existing one without `--refresh`.
