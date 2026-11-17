---
name: agents-md-completeness
description: >
  Verification pass for an existing AGENTS.md: checks it against a standard
  six-section outline (what it is, how it is built, what is special, how we
  work here, watch out, who to ask), auto-detects what the repo itself can
  answer (top committers, languages, custom code, patches, integrations,
  cron jobs, tests, environments, recent fixes) and proposes only the missing
  lines. Stack-aware for PrestaShop, Magento, Laravel and custom projects;
  every check is optional. Use after ps-project-doc or on any repo whose
  AGENTS.md exists but feels thin, when asked to "complete", "verify" or
  "audit" AGENTS.md, or before a tool with a character cap reads it.
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# agents-md-completeness — fill the gaps, from the repo

Second pass over an `AGENTS.md` that already exists. It does not generate the base file (use `ps-project-doc` for PrestaShop, or write it by hand). It finds which standard sections are missing, fills what the repository can answer by itself, and asks only for the rest.

## Hard rules

- **Optional by design.** Every check is a proposal. Skip any section the project does not need; never write a placeholder or `TODO` line.
- **Only what the code does not show.** A line must pass: "if it were missing, would an agent or newcomer make a mistake they cannot catch by reading the code?"
- **Hints are not facts.** Detected values are a starting point. Confirm each one with the user before it enters the file.
- **Never secrets.** No passwords, tokens, hostnames with credentials, access details or private contacts. Those belong in a secret store.
- **Size.** Keep the file under the character cap of whatever reads it. Default 6000 characters; check `agents_md_chars` from the detect script. Cut before compressing prose.
- **Do not quote CI guardrail patterns.** If the repo's CI scans added lines for debug calls (for example `console.log`, `dd(`, `dump(`, `debugger`), the AGENTS.md itself trips it when it spells them out. Say "leftover debug calls" and point to the CI file instead. Check `.gitlab-ci.yml` for `forbidden-patterns` before writing.
- **English only.** Write every added line in English, whatever the language of the conversation or of the team outline; it costs fewer tokens on every read. Keep the six section headings in English too.
- **No organisation names.** Neither in this skill nor in the generated lines unless the project itself uses them.
- **Never commit.** Write the changes, show the diff, stop.

## Outline checked

| Section | What it holds | Detected from |
|---|---|---|
| What it is | One or two sentences: what the project sells or does, for whom, what is particular (B2B, several countries, ERP) | manifest description, README, user |
| How it is built | Stores, languages, domains if more than one; integrations and where data enters; what is custom vs third-party in prices, stock, checkout | `languages`, `integration_hints`, user |
| What is special | Custom modules, one line each; things NOT done the usual way, and why; patches or overrides to review at every upgrade | `custom_code`, `patches_and_overrides`, user |
| How we work here | Branches and how publishing works; what is and is not in the repo; tests and how to run them, or say there are none; environments by name, never access | `default_branch`, `long_lived_branches`, `tests`, `environments`, `ci` |
| Watch out | What broke before; what the client is sensitive about; automated processes that must not be interrupted | `recent_fixes`, `automated_processes`, user |
| Who to ask | Project manager and the people who know the code best | `top_committers`, user |

Match sections by meaning, not by exact heading. An existing `Rules` section covers part of "How we work here"; do not duplicate it, extend it.

## Flow

### 0. Check before touching anything

Someone may have worked on this repo already. Before writing a line:

- `git fetch`, then look at the default branch on the remote: does `AGENTS.md` exist there, who wrote it and when (`git log -3 -- AGENTS.md`), and are there open merge requests or branches active in the last two weeks?
- An `AGENTS.md` written by someone else is theirs: do not rewrite, translate or restructure it. Verify its facts against the repo, report what is wrong, and propose additions only. If the owner wants it left alone, stop and record that.
- If the working tree is dirty or behind `origin` (see `dirty_entries` and `behind_origin` below), work in a clean worktree from the remote default branch, for example `git worktree add <tmp-dir> -b docs/agents-md origin/<default-branch>`. Never touch uncommitted work.
- Stage only your own files. `git status` in a worktree can list files that are not yours.

### 1. Detect

```bash
~/.claude/skills/agents-md-completeness/scripts/detect-context.sh
```

JSON fields: `stack` (prestashop | magento | laravel | generic), `description_hint`, `top_committers`, `languages`, `custom_code`, `patches_and_overrides`, `integration_hints`, `automated_processes`, `tests`, `environments`, `ci`, `default_branch`, `long_lived_branches`, `recent_fixes`, `agents_md_sections`, `agents_md_chars`, `agents_md_ignored`, `claude_md_tracked`, `claude_md_ignored`, `claude_md_lines`, `dirty_entries`, `behind_origin`.

No `AGENTS.md` yet: stop and say so; this skill needs a base file.

### 2. Gap report

Compare `agents_md_sections` with the outline. Print a table: section, present / partial / missing, and what the repo detected for it. Include the character budget left (`cap - agents_md_chars`).

### 3. Stack checks (all optional)

Run only the block for the detected `stack`; for `generic`, run none and rely on the outline.

**PrestaShop**
- Modules installed or updated from the back office are committed the same day.
- `override/` files each have a one-line reason; patches to core or third-party modules are listed with what breaks on update.
- Which shops, languages and currencies are active; if the repo only shows translation folders, ask which are really used.
- Production matches the repository: no files edited only on the server.

**Magento**
- Core changes exist only as Composer patches, never by editing `vendor/`. List the patch files.
- Custom modules under `app/code` (one line each) and any `di.xml` preferences that replace core behaviour.
- Store views, websites and locales if more than one; indexers and cron groups that must not be interrupted.

**Laravel**
- Every database change is a migration, never manual. Say whether seeders are safe to run in production.
- Scheduled commands (`schedule->`) and queue workers that must keep running.
- Packages published or forked inside the repo.

**Custom / generic**
- How it is built and run (one command), where configuration lives, what is generated and must not be edited.
- Anything the framework would normally do that this project does differently, and why.

### 4. Fill, then ask

1. Draft lines for every section the detect output can support, marking each as a hint.
2. Ask one short batch for what the repo cannot say: project manager, what is sensitive for the client, past incidents worth a warning, planned integrations. Skippable; unanswered means the line is omitted.
3. `top_committers`: propose at most the two people with the most commits as "who knows the code". Names only, no emails. Confirm with the user first; commit counts show activity, not ownership.
4. `recent_fixes`: read the subjects and propose at most three durable lessons ("X broke when Y"), not a changelog. Drop one-off fixes.

### 5. Agent files: ignores and the CLAUDE.md bridge

- `agents_md_ignored` or `claude_md_ignored` true: remove that line from `.gitignore` and say so. Add `CLAUDE.local.md` instead. Note that `git update-index --add` and `git add -f` bypass an ignore rule silently, so check with `git check-ignore -v` before staging.
- `claude_md_tracked` with real content: read it all before replacing it with the one-line bridge `@AGENTS.md`. Sort each block: facts the code does not show go to `AGENTS.md`; commands and notes still true go to `docs/<topic>.md`, linked from `AGENTS.md`; stale claims (check each against the repo: active theme, branches, overrides) and personal preferences are dropped. Tell the user what moved and what was dropped.
- Local environment tool (Lando, DDEV, Docker Compose): say which one is used in `AGENTS.md`. Commit its config only if the team agrees; do not decide that here.
- For branch protection in "How we work here", check the member list as well: protection is hollow when most members are Maintainer (see `gitlab-project-bootstrap`, section 7 and 8).

### 6. Write and report

- Add or extend sections in place; keep the existing wording of anything already written.
- Show the diff. Report: sections added, characters before and after against the cap, hints the user rejected, questions left unanswered.
- Never stage or commit.
