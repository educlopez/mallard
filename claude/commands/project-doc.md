---
name: project-doc
description: Generate a lean, team-shared AGENTS.md for the current PrestaShop repo (only what the code does not show), with a one-line CLAUDE.md bridge and a gitignored CLAUDE.local.md for private notes. Enforces a 40-line target and 80-line cap.
version: "0.2.0"
metadata:
  author: Eduardo Calvo
---

# project-doc — lean shared AGENTS.md

Run inside a PrestaShop repo. Writes the short instruction file every agent reads on each task. **AGENTS.md is committed and shared; private notes go to `CLAUDE.local.md` (gitignored).**

## Arguments

`$ARGUMENTS` (optional):
- (none) — create `AGENTS.md` if missing; audit it if it exists.
- `--refresh` — regenerate even if one exists (backs up to `AGENTS.md.bak` first).
- `--print` — preview only, write nothing.

## Instructions

Load and follow the **`ps-project-doc`** skill (`~/.claude/skills/ps-project-doc/SKILL.md`):

1. Verify the cwd is a PrestaShop repo.
2. Run `scripts/detect.sh` for PS version, theme, build, overrides and the state of AGENTS.md / CLAUDE.md.
3. Ask one short batch: production PHP, why each override exists, core/third-party patches, gotchas.
4. Fill the template. Stay under 80 lines. No directory tables, module lists, overviews or snippets.
5. Ensure `CLAUDE.md` is the single line `@AGENTS.md`, `CLAUDE.local.md` is gitignored, and AGENTS.md is not ignored.

Never commit. Never delete or replace an existing CLAUDE.md without showing it and getting an OK.
