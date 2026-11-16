---
name: ps-project-doc
description: >
  Generate a lean, team-shared AGENTS.md for a PrestaShop repo (committed to the
  default branch) containing only what an agent cannot discover from the code:
  production PHP/PS version, why each override/ file exists, patches to core or
  third-party modules, build quirks, hard rules and gotchas. Replaces the old
  private CLAUDE.md generator. Candid or personal notes go to the gitignored
  CLAUDE.local.md instead. Auto-detects what it can, asks only for what code
  cannot say, and enforces a size budget. Use for the /project-doc command, when
  a PrestaShop repo has no AGENTS.md, has a bloated one, or is bootstrapping
  its docs for agentic development.
version: "0.2.0"
metadata:
  author: Eduardo Calvo
---

# ps-project-doc — lean AGENTS.md generator

Writes the file an agent (Claude Code, Codex, Cursor, Copilot) reads on every task in this repo. It is read on **every** run, so every line costs context and attention. Evidence (ETH Zurich / LogicStar, arXiv 2602.11988): context files that restate what the repo already shows lower task success and raise cost by 20%+; LLM-generated overviews do not help agents find files. Only facts the agent cannot discover earn a line.

## Hard rules

- **AGENTS.md is shared and committed** to the default branch. Never put personal preferences, tooling of one developer (memory tools, personal skills, `glab`), credentials, or private client contacts in it.
- **Budget: target 40 lines, hard cap 80** (Claude docs advise under 200; the Codex cap is 32 KiB total and silently truncates). If detect + answers exceed 80 lines, cut, do not compress prose.
- **Never write what the code shows**: no directory tables, no module lists, no architecture overview, no style rules a linter enforces (PHPStan / CS-Fixer config already says it), no code snippets (they go stale), no copies of README text. Point to a file instead (`see docs/deploy.md`).
- **Every line must pass the test**: "if this line were missing, would the agent make a mistake it cannot catch by reading the code?" If not, drop it.
- English-only content. No company names in this skill.
- Never commit on the user's behalf. Write the file, show the diff, stop.

## Where things live

| File | Committed | Content |
|---|---|---|
| `AGENTS.md` | yes | the lean shared facts below |
| `CLAUDE.md` | yes, 1 line | `@AGENTS.md` only (bridge for Claude Code versions or setups that skip AGENTS.md when a CLAUDE.md exists; the import is never read twice) |
| `CLAUDE.local.md` | **no** (gitignored) | candid gotchas, contacts, personal notes |

Claude Code reads `AGENTS.md` natively only when no `CLAUDE.md` exists, and ignores it otherwise unless `/config` → Project instructions is set to both. The one-line `CLAUDE.md` bridge removes that ambiguity. Check with `/context all`.

## Flow

### 1. Verify repo

Confirm cwd is a PrestaShop repo (`composer.json` mentions prestashop, or both `themes/` and `modules/`). Otherwise stop with a clear message.

### 2. Detect facts

```bash
~/.claude/skills/ps-project-doc/scripts/detect.sh
```

JSON: `ps_version, php_version (local Lando, NOT production), themes[], theme_stack, lando_name, lando_url, css_build, git_remote, modules[], overrides[], has_claudemd, claudemd_tracked, claudemd_gitignored, has_agentsmd, agentsmd_gitignored, agentsmd_lines, has_claude_local`.

Read source files for anything empty or ambiguous (`themes/*/config/theme.yml`, `composer.json`). `modules[]` is only a hint for step 3; it never goes into the file as a list.

### 3. Ask (one short batch, skippable)

- **Production PHP version** (and PS version if it differs from local). Not detectable from the repo; check the server if the user has a registry, else ask.
- **Why each `overrides[]` file exists.** One line each. List only files that exist.
- **Patches to core or third-party modules** (files edited in `classes/`, `controllers/`, `src/` or inside `modules/<third-party>/`) and why. Check `git log` on those paths for hints before asking.
- **Gotchas**: what burns a new dev or agent here that the code does not show. Skippable.
- Anything candid or personal → goes to `CLAUDE.local.md`, not AGENTS.md.

### 4. Write AGENTS.md from the template

Omit any section with nothing to say. Do not leave `TODO` lines in a committed file: ask, or drop the line.

```markdown
# AGENTS.md

PrestaShop <ps_version> · PHP <prod_php> in production · theme: <child> (parent <parent>; never edit the parent) · <builds with `<cmd>` | no build step>

## Commands
- Local env: `lando start` (see `.lando.yml`)
- <CSS build / PHPStan / CS-Fixer: exact command, only if present>

## Overrides
- `override/<path>` — <why it exists>

## Core / third-party patches
- `<path>` — <what changed and why; what breaks if the vendor updates it>

## Rules
- Branch: `flow/<task-number>-<slug>`; open a merge request, never push to `<default_branch>`.
- Conventional Commits, English.
- Not in git: `app/config/parameters.php`, `img/`, `upload/`, `download/`, `var/cache/`, `var/logs/`, `vendor/`, `node_modules/`.
- Modules installed or updated from the Back Office are committed the same day.
- Deploy is manual and done by a person: agents never deploy or touch servers.

## Gotchas
- <non-obvious trap, one line each>
```

### 5. Bridge, gitignore, write

- `CLAUDE.md` missing → create it containing exactly `@AGENTS.md`.
- `CLAUDE.md` exists with real content (`claudemd_tracked` or not) → **do not delete it silently.** Show the user what is in it, propose: move the shared, non-discoverable parts into AGENTS.md, move candid parts into `CLAUDE.local.md`, replace `CLAUDE.md` with `@AGENTS.md`. Apply only after OK.
- Ensure `.gitignore` contains `CLAUDE.local.md`. Remove a stale `CLAUDE.md` ignore line only if the bridge is being committed (tell the user; ignoring the bridge defeats it).
- `AGENTS.md` ignored (`agentsmd_gitignored`) → remove that ignore line and tell the user.
- Existing `AGENTS.md` without `--refresh` → stop, offer `--print` or `--refresh`. `--refresh` backs up to `AGENTS.md.bak` first; never stage the `.bak`.
- `--print` writes nothing.

### 6. Report

Print: line count vs the 40/80 budget, what was auto-filled, what was asked, what was cut. If over 40 lines, say which lines are the weakest candidates to drop.

## Audit mode (existing files)

When `agentsmd_lines > 80`, or the file contains a directory table, a module list, an overview paragraph or snippets: list those blocks and propose removals. Typical legacy `CLAUDE.md` files in this org are 100-340 lines; most of that is discoverable and should go.

## Non-interactive reuse

When called by `ps-new-project`, accept pre-supplied values (track, theme name, PS version, default branch) and skip step 3 questions that the scaffold already answers. Production PHP and overrides stay unknown on a fresh project: omit those sections.
