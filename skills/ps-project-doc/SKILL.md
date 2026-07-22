---
name: ps-project-doc
description: >
  Generate a standardized, private CLAUDE.md for a PrestaShop repo by
  introspecting it (PS version, theme + parent, theme stack Panda/ElementFlow,
  Lando URL/port, CSS build, git remote, modules). Auto-fills most fields,
  asks the user only for gotchas/contacts/module purposes, and writes a lean
  ~70-90 line file. The CLAUDE.md is PRIVATE — always gitignored, never
  committed. Reusable non-interactively by the new-project bootstrap.
  Use for the /project-doc command, when a PrestaShop repo has no CLAUDE.md
  or an inconsistent one, or when bootstrapping a new project's docs.
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# ps-project-doc — standardized private CLAUDE.md generator

Writes a consistent CLAUDE.md so Claude stops re-inferring the stack each session. **CLAUDE.md is private working context — never committed.**

## Hard rules

- **CLAUDE.md is gitignored, never committed.** Before finishing, ensure `.gitignore` contains a `CLAUDE.md` line (append it if missing). Never `git add`/commit CLAUDE.md. Do not touch git history for repos where it's already tracked unless the user explicitly asks (then `git rm --cached CLAUDE.md`).
- Because it's private, Gotchas can be candid (blunt notes-to-self are fine).
- English-only content. No company names in the skill itself (the generated CLAUDE.md may name the client — that file is private and untracked, so it's fine there).

## Flow

### 1. Verify repo

Confirm cwd is a PrestaShop repo (has `composer.json` mentioning prestashop, or both `themes/` and `modules/`). If not, stop with a clear message.

### 2. Detect facts

```bash
~/.claude/skills/ps-project-doc/scripts/detect.sh
```

Returns JSON: `ps_version, php_version, themes[] (dir/name/parent), theme_stack, lando_name, lando_url, mysql_port, css_build, git_remote, modules[], has_claudemd, claudemd_gitignored`.

- If Lando is down, `lando_url`/`mysql_port` come from `.lando.yml` parsing — acceptable.
- For anything empty/ambiguous, read the source file directly to fill it (e.g. open `themes/*/config/theme.yml`, `composer.json`).
- Child theme = the theme dir with a `parent` set (e.g. `parent: panda`/`classic`); the parent is that value.

### 3. Prompt the user (short batch)

Ask once, briefly, allowing skips:
- **Gotchas** — what trips a new dev/agent here? (legacy quirks, "don't touch X", broken tooling)
- **Contacts** — client + Cinetic point of contact (optional)
- **Module purposes** — for the custom-looking modules in `modules[]`, a one-liner each (or leave `TODO`)

### 4. Fill the template

Use the section template below. Auto-fill from step 2/3. Inject the **fixed** Deploy + override-priority text verbatim.

### 5. Write

- If `CLAUDE.md` exists and `--refresh` was NOT passed → stop, tell the user (offer `--print` or `--refresh`).
- With `--refresh` → copy existing to `CLAUDE.md.bak`, then overwrite.
- With `--print` → output to chat only, write nothing.
- Ensure `.gitignore` has `CLAUDE.md` (append if missing). Never stage/commit it.
- Print a summary: fields auto-filled vs left `TODO`.

## Template

```markdown
# CLAUDE.md

## Project Overview
<what the store sells> · Live: <url> · PrestaShop <ps_version> · Theme stack: <Panda | ElementFlow | custom> · Lang: <langs>

## Development Environment (Lando)
- URL: <lando_url>   MySQL: 127.0.0.1:<mysql_port> (user/pass/db per project)
- PHP <php_version>. Commands: `lando start` · `lando stop` · `lando mysql` · `lando ssh`. Mail: MailHog. xdebug: <on/off>

## Build & Assets
<css_build: e.g. "lightningcss — `pnpm run css` in _dev/"  |  "no build — edit assets/css/custom.css">
<if theme_stack transitioning: note Panda→ElementFlow status>

## Architecture
- Override: `override/` — most-touched: <list or "none">
- Theme: parent **<parent>** (do not edit directly) → child **<child dir>** (work here)
- Custom modules: <name — purpose> (one per line; TODO where unknown)
- Template override priority: child module → child theme → parent module → parent theme → module default

## Deploy
Manual `git pull` over SSH on the server — never FTP.
- `test` branch → **staging**
- `master`/`main` branch → **production**
- SSH: <host / path>  ·  Access: <who>

## Git Workflow
Remote: <git_remote> · Main branch: <master|main> · Feature branches: `<intervals-task-id>` (see /task). Commits: Conventional Commits, English.

## Key Conventions
<code style, PHPStan/CS-Fixer if present, comment/commit language, naming patterns>

## Gotchas
<candid notes — this file is private. TODO if none given>

## Contacts
<client + Cinetic contact — optional / TODO>
```

## Non-interactive reuse (P1 bootstrap)

When called by `ps-new-project`, accept pre-supplied field values and skip the step-3 prompts (bootstrap already knows them for a fresh scaffold). Same template, same gitignore rule.
