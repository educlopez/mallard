---
name: gitlab-wiki-bootstrap
description: >
  Bootstrap or extend a team-facing GitLab Wiki for a project — explores the
  actual codebase (not generic boilerplate) to populate Home plus
  type-specific pages (PrestaShop: Modules/Theme/Local Development/Gotchas;
  Laravel: Integrations/Jobs & Queues/Deployment/Testing/Gotchas). Wiki
  upkeep afterwards is a PERSONAL habit (see Step 5) — never written into
  the project's own CLAUDE.md/AGENTS.md, since that would force every
  teammate's agent to depend on `glab` CLI just from touching the repo.
  Works both for brand-new projects with no wiki AND for established
  projects with months/years of history — for the latter, audits what
  already exists (wiki pages, README, docs/, CLAUDE.md) first and
  extends/corrects rather than starting from a blank slate. Use when the user
  says "monta la wiki", "crea la wiki para X", "documenta este proyecto en
  GitLab", "actualiza/completa la wiki de X", or wants team documentation set
  up or brought up to date for any project regardless of its age. Distinct
  from a personal notes/second-brain tool — this is for anything teammates
  need, not the user alone.
version: "0.3.0"
metadata:
  author: Eduardo Calvo
---

# GitLab Wiki Bootstrap

Sets up (or extends) a real, populated GitLab Wiki — not an empty shell. The
value is in exploring the actual codebase for concrete detail (real module
names, real integration behavior, real gotchas), the same way you'd explain
the project to a new teammate. Generic boilerplate pages are worse than no
wiki at all — don't ship those.

Works for two starting points:
- **Brand-new project, no wiki** — Steps 0 → 2 → 3 → 4, straightforward.
- **Established project, possibly with history/existing docs/existing wiki
  pages** — add Step 1 (audit) before writing anything. Don't treat this as
  a from-scratch bootstrap: mine what's already there first.

## Step 0 — check prerequisites

```bash
glab auth status
```

Get the project's numeric ID (needed for all wiki API calls):
```bash
glab api "projects/<group%2Fsubgroup%2Fproject>" | jq .id
```

Check the Wiki feature is enabled (it usually is by default):
```bash
glab api "projects/<id>" | jq .wiki_access_level
```
If `disabled`, enable it:
```bash
glab api -X PUT projects/<id> -f "wiki_access_level=enabled"
```

## Step 1 — audit existing state (skip only for a truly brand-new project)

An "existing project" here means: any repo with real commit history, even if
it has never had a wiki. Don't assume empty — check:

1. **Existing wiki pages**:
   ```bash
   glab api "projects/<id>/wikis" | jq -r '.[].slug'
   ```
   If any exist, `glab api "projects/<id>/wikis/<slug>"` each one and read the
   content. Note what's already covered, what's stale (check against the
   actual current code — theme changes, removed integrations, renamed
   modules), and what's missing entirely.

2. **Existing human-written docs**: `README.md`, `docs/*`, `CLAUDE.md`,
   `CONTRIBUTING.md`. These often contain real institutional knowledge
   (design decisions, environment quirks, historical context) that code
   exploration alone won't surface — mine them for content, but verify
   against the actual code before trusting anything version-specific or
   procedural (a stale `docs/DEVELOPMENT.md` describing a setup the project
   no longer uses is a real failure mode — flag it, don't propagate it into
   the wiki as if current).

3. **Git history for context, sparingly**: `git log --oneline -50` and
   skimming a few merge/feature commit messages can surface *why* something
   is built a certain way (a past incident, a client requirement) that isn't
   visible in the code itself. Don't do a deep archaeology dig — a quick
   skim is enough; the codebase exploration in Step 3 is the primary source.

**Decide your mode before writing anything**:
- No wiki, no existing docs worth mining → full bootstrap (Steps 2-4 as if new)
- Existing wiki, mostly accurate → **extend**: add missing pages, patch stale
  sections in place (edit existing pages via PUT, don't replace wholesale
  unless the whole page is wrong)
- Existing docs (README/docs/CLAUDE.md) but no wiki → **migrate + verify**:
  pull genuinely accurate content into wiki pages, correct anything stale
  against the real code, then still explore (Step 3) for what those docs
  never covered (most human docs undersell integration quirks/gotchas)

## Step 2 — detect project type

- **Laravel**: `composer.json` requires `laravel/framework`
- **PrestaShop**: has `modules/` + a `themes/<name>/theme.yml` with `parent: classic` (or similar)
- **Other**: fall back to a minimal generic structure (Home + Architecture + Gotchas)

## Step 3 — explore the codebase (delegate, don't guess)

Spawn an Explore-type agent per page group to gather **factual, specific**
detail — not generic framework docs. The agent should read actual controllers/
modules/config, not just the README. See the prompts used for Cinetic's
first two rollouts (apply-animalmax, ps9-gaudibarcelonashop) as a model:

**For PrestaShop**, explore and report on:
1. Installed modules (`modules/` dir) — name, purpose, customized vs vendored,
   skip pure `ps_*` core modules unless overridden
2. Theme structure — `_dev/` vs built `assets/`, CSS build pipeline
   (check `package.json` scripts), custom JS, template overrides, naming
   conventions for custom classes
3. Any custom PHP modules (not vendored builder engines/core)
4. Local dev config (Lando/Docker), URLs, PHP/DB versions
5. Anything stale/misleading in existing docs (README, `docs/`) — flag it,
   don't propagate it

**For Laravel**, explore and report on:
1. Each external integration (CRM, ERP, payment, etc) — what it actually
   does, OAuth/auth flow, sync direction, webhooks (inbound or just
   outbound-listing), error/retry handling, where to debug it (logs, DB
   tables, dev panels)
2. Jobs/queues — what's actually queued vs synchronous, queue driver,
   scheduled tasks (`routes/console.php` or `app/Console/Kernel.php`), how
   they run in dev vs prod
3. Emails/notifications — what triggers them, any hardcoded
   recipients/BCCs, dev preview routes
4. Testing — what's ACTUALLY covered vs scaffolding-only; call this out
   explicitly if integrations have zero tests, don't soften it
5. Deployment — the actual deploy script/CI, not assumptions
6. Gotchas — commented-out logging, duplicated logic across controllers,
   unthrottled public endpoints hitting external services, anything fragile

Push back on inventing behavior you can't confirm from code — factual only.
If Step 1 found existing wiki pages, pass their current content into the
explore agent's prompt so it can flag what's stale/missing rather than
re-deriving everything blind.

## Step 4 — write pages

Structure (adapt names to what Step 3 found):

**PrestaShop**: `home`, `modules`, `theme-<name>`, `local-development`, `gotchas`
**Laravel**: `home`, `integrations`, `jobs-and-queues`, `deployment`, `testing`, `gotchas`
**Generic**: `home`, `architecture`, `gotchas`

Cross-link pages with `[[slug]]` or `[[slug|Display Text]]` wiki-link syntax.
`home` should list all pages with a one-line description of what's on each.
If extending an existing wiki, keep its existing slugs/structure unless
they're genuinely wrong — don't rename pages just to match this skill's
naming exactly.

Create each page (new):
```bash
glab api -X POST "projects/<id>/wikis" \
  -f "title=<slug>" \
  -f "format=markdown" \
  -f "content=$(cat page.md)"
```

Update later:
```bash
glab api -X PUT "projects/<id>/wikis/<slug>" \
  -f "content=$(cat page.md)" \
  -f "format=markdown"
```

**Watch out for shell variable naming collisions** when scripting multiple
pages in a loop — reusing the same variable name across `cat`/upload calls for
different pages can silently push empty or wrong content to a page. Verify
each page's content after upload if scripting more than 2-3 at once:
```bash
glab api "projects/<id>/wikis/<slug>" | jq -r .content | head -5
```

## Step 5 — wiki upkeep is PERSONAL, not project config

**Do NOT write a "keep the wiki updated" section into the project's
CLAUDE.md/AGENTS.md.** That was this skill's approach in earlier versions
(v0.2.0 and before) and it was a real mistake: that file is read by every
agent any teammate runs against the repo, so an instruction to use
`glab api` there forces `glab` CLI (GitLab CLI) on people who don't have it
installed — they'd get install prompts just from touching the repo. Several
Cinetic projects had this section retroactively removed for that reason
(see `git log` on those repos' `CLAUDE.md` for the revert commits if you
want the paper trail).

Instead, wiki upkeep is a **personal habit** that lives in the *user's own*
global Claude Code config (`~/.claude/CLAUDE.md`), scoped explicitly to
Cinetic/work projects — never in the repo itself, never assumed for anyone
else. If the user doesn't already have this in their global config, offer to
add it (to `~/.claude/CLAUDE.md`, not the project):

```markdown
## GitLab Wiki maintenance (Cinetic projects ONLY)
- Applies ONLY to Cinetic client/work projects hosted on GitLab (cineticd
  namespace — a `gl-cinetic`-style SSH remote alias, if the user has one, is
  just a local machine convenience, not something to rely on or assume
  elsewhere). Do NOT apply to personal projects or anything not hosted on
  GitLab.
- This is a PERSONAL preference, not a project convention — do NOT write
  "keep the wiki updated" instructions into any project's
  CLAUDE.md/AGENTS.md, even for Cinetic projects. Teammates without `glab`
  CLI installed would get forced install prompts just from working on the
  repo.
- When asked to update a Cinetic project's GitLab wiki, use `glab api`
  directly (read current page, edit incrementally, don't overwrite blindly)
  — keep the instruction here in personal config, never baked into the repo
  itself.
```

There's no reliable way to auto-detect "wiki-worthy" changes mechanically,
so the update habit has to live somewhere a human re-triggers it — that's
this personal config entry, picked up only in the user's own sessions, not
a project-level mechanism every contributor inherits.

## Checklist

- [ ] Wiki feature confirmed enabled
- [ ] Existing wiki pages / README / docs / CLAUDE.md audited (skip only for
      a genuinely brand-new project) and mode decided (bootstrap / extend /
      migrate+verify)
- [ ] Project type detected (Laravel / PrestaShop / generic)
- [ ] Codebase explored for real detail — no generic boilerplate content
- [ ] Home page created or updated, links to all other pages
- [ ] Type-specific pages created/extended and cross-linked
- [ ] Each uploaded page spot-checked (`glab api .../wikis/<slug> | jq -r .content`)
- [ ] `CLAUDE.md` maintenance section added or updated (create file if missing,
      don't duplicate the section if one already exists)
- [ ] Changes NOT auto-committed — left for the user to review/commit
