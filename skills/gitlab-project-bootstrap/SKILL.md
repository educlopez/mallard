---
name: gitlab-project-bootstrap
description: >
  Set up or audit GitLab project metadata hygiene (avatar, description, topics,
  badges, README) for Cinetic Digital's private client repos on gitlab.com. Use
  this whenever the user is starting a brand-new client project, asks to
  "configure" or "bootstrap" a GitLab project, wants project
  descriptions/tags/topics/badges set, mentions a project's README is
  missing/stock/hidden/needs improving, or asks what's missing / what should be
  set up on a GitLab repo. Also trigger when the user mentions a specific client
  project by name (e.g. a PrestaShop or Laravel repo) and asks about its GitLab
  page, its description, or wants it to "look proper." Works both for brand-new
  projects (apply everything from minute 1) and for auditing/fixing an existing
  one.
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# GitLab project bootstrap (Cinetic Digital client repos)

Checklist + concrete commands for making a private client repo on gitlab.com (group `cineticd`) look and behave properly from day one — description, topics, badges, README, and a couple of org-hygiene checks worth flagging. Built from real work across PrestaShop projects (`ps9-gaudibarcelonashop`, `vives-8`, `milagros-colombia-b2b-v9.1`) and a Laravel one (`apply-animalmax`).

Everything here is done via `glab` (GitLab CLI) — either `glab api` for project-metadata-only changes (no repo files, no PR needed), or a normal clone → branch → MR → merge flow when a repo file (README) needs to change.

## Before you start: find the project and its SSH remote alias

If you don't have the exact `namespace/project` path, search for it:

```bash
glab api "projects?membership=true&per_page=100&search=<name>" | python3 -c "
import json,sys
for p in json.load(sys.stdin):
    print(p['path_with_namespace'], '|', p.get('description'))
"
```

If there are multiple near-matches (e.g. `vives-8`, `vives-17-nuevo`, `ps17-vivescortadaimport`), confirm the right one with the user rather than guessing — don't touch the wrong client's repo.

Check `~/.ssh/config` for the SSH host alias used to reach gitlab.com — this org uses a `Host gl-cinetic` alias (not `git@gitlab.com:` directly) with its own IdentityFile. Cloning with the wrong host will fail with "Permission denied (publickey)" even though your GitLab account has access. Always clone as `git@gl-cinetic:<path>.git`, and run `glab mr create`/`glab mr merge` from *inside* the cloned repo (not with `--repo <path>` from elsewhere) — `glab` otherwise tries to add its own remote using the default host and fails the same way.

## 1. Avatar

GitLab defaults to a colored initial-letter avatar if none is set — looks unfinished on a client project page. This one's manual: **Settings → General → avatar upload**, a square logo image (≥128×128px). No API shortcut worth using here; just point the user at the setting if they haven't done it.

## 2. Description

One line, via API, no repo file touched:

```bash
enc=$(python3 -c "import urllib.parse; print(urllib.parse.quote('<namespace>/<project>', safe=''))")
glab api -X PUT "projects/$enc" --field "description=<one-liner>"
```

Name the stack, and for PrestaShop projects, the **parent theme** too (this is the detail that makes descriptions actually useful, not just decorative):

- `Gaudí & Barcelona Shop — PrestaShop 9.1.1 storefront (ElementFlow / stsitebuilder page builder)`
- `Milagros Colombia — B2B PrestaShop 9.1 storefront, Panda child theme`
- `Vives — PrestaShop 8 storefront, Panda child theme (viveschild)`
- `Animalmax — job application portal (Laravel + Inertia + React)`

**How to find the real active theme — don't guess from folder names.** A PS install usually ships several theme folders (`classic/`, `panda/`, `hummingbird/`, plus the client's custom one) but only one is active. Read the client's own theme's `config/theme.yml` for its `parent:` field to know what it's built on, then confirm it's the one actually in use — cross-check against `CLAUDE.md` if the repo has one, but don't trust it blindly: it can be stale. A quick sanity check that worked well: whichever theme directory has the most recent commit/mtime is almost always the one being actively developed, i.e. the active one.

```bash
cat themes/<client-theme>/config/theme.yml | grep parent
git log -1 --format=%ci -- themes/<theme-dir>   # per theme dir, compare dates
```

## 3. Topics (tags)

Also via API, same no-file-touched pattern:

```bash
glab api -X PUT "projects/$enc" --field "topics=prestashop,php,ecommerce"
```

Pick 3-5 real stack keywords, not generic filler. Examples that worked:
- PrestaShop storefront: `prestashop,php,ecommerce` (+ `b2b` if it's a B2B store)
- Laravel + Inertia + React app: `laravel,php,react,typescript,inertia`

## 4. Badges

Also API-only (`POST projects/<id>/badges` with `name`, `link_url`, `image_url` fields) — no PR needed.

**First, check if the repo actually has CI** (`ls .gitlab-ci.yml` in a clone, or `glab api "projects/$enc/repository/files/.gitlab-ci.yml/raw?ref=<default_branch>"`). This decides what kind of badge is honest to add:

- **No CI**: static shields.io badges only — stack version, PHP version, theme. Nothing else to report, so don't add pipeline/coverage badges (they'd have nothing behind them).
  ```bash
  glab api -X POST "projects/$enc/badges" \
    --field "name=PrestaShop" \
    --field "link_url=https://www.prestashop-project.org/" \
    --field "image_url=https://img.shields.io/badge/PrestaShop-9.1.1-df0067?style=flat-square"
  ```

- **Has CI**: add a real pipeline badge using GitLab's placeholder URL syntax (it resolves `%{project_path}`/`%{default_branch}` server-side, so it keeps working across renames/branch changes):
  ```bash
  glab api -X POST "projects/$enc/badges" \
    --field "name=Pipeline" \
    --field "link_url=https://gitlab.com/%{project_path}/-/commits/%{default_branch}" \
    --field "image_url=https://gitlab.com/%{project_path}/badges/%{default_branch}/pipeline.svg"
  ```
  **Only add a Coverage badge if the CI config actually produces a coverage report.** Read the `.gitlab-ci.yml` jobs first — a security-scan-only pipeline (e.g. just a Trivy dependency scan) has no coverage number, so a coverage badge would sit at "unknown" forever, which reads as broken/misleading rather than honest. When in doubt, skip it and say why.

## 5. README

### The hidden-README gotcha

GitLab only renders a project's overview from a non-dotfile named `README`/`README.md`/etc. A repo can have perfectly good docs sitting in a **dot-prefixed** file like `.README.md` that silently never displays on the project page. Always check for this before writing a new README from scratch:

```bash
ls -la  # look for .README.md, .readme, etc. alongside (or instead of) README.md
```

If found, `git mv .README.md README.md` and layer a short overview section on top — don't duplicate the existing content, it's usually solid dev documentation, just invisible.

### Structure that's worked well

```markdown
Project Name
============

[![Stack](badge)](link)
[![Stack2](badge)](link)

One-line description of what this project is.

<p align="center">
  <img src="docs/screenshot-home.webp" alt="..." width="900"/>
</p>

Stack
-----

- **Bullet list**, bold the key technology names
- Anything architecturally notable (fork-tracking strategy, override system, custom modules, page-builder module, etc.)

Environments
------------

- **Production/Pre:** <real URL> (note explicitly if it's down/in maintenance rather than silently skipping)

Local setup / Theme build
--------------------------

Only include this section if there's a real local dev flow worth documenting (e.g. Lando + a CSS build pipeline). Don't invent one.

Conventions
-----------

Pointer to CLAUDE.md if the repo has one — don't duplicate its content here.
```

See `references/readme-template.md` for a copy-pasteable skeleton.

### Screenshot rule

If there's a reachable instance (local dev, staging, or production — in that order of preference for a NEW project; for an EXISTING project with a real deployed URL, use that URL, not a local one, see below), take a homepage screenshot and compress it before committing — a raw PNG from a browser screenshot tool is typically 1-1.5MB, which is heavy for a repo. Convert to WebP:

```bash
cwebp -q 82 screenshot.png -o docs/screenshot-home.webp   # ~125KB from a 1.5MB PNG, negligible quality loss
rm screenshot.png
```

If the reachable instance is down/in maintenance, say so in the README rather than silently omitting the screenshot — that's more honest than pretending there's nothing to show.

### Never link local dev URLs in a README that has a real environment

If the project has a real staging/production URL, use *that* in the README — never a `*.lndo.site` or `localhost` link. Local dev links are only appropriate when the project genuinely has no deployed environment yet (e.g. it hasn't shipped anywhere).

### Respect existing ownership

If a specific team member actively relies on and edits a project's README (ask if unsure — "does anyone use this doc day to day?"), don't touch the file even if it looks improvable. In that case, still do steps 2-4 (description/topics/badges) since those are GitLab metadata, not the README file, and don't step on anyone's workflow.

## 6. Applying a README change (existing project, no local clone yet)

```bash
git clone git@gl-cinetic:<namespace>/<project>.git /path/to/scratch
cd /path/to/scratch
git checkout -b docs/readme-<whatever>
# edit README.md
git add README.md [docs/screenshot-home.webp]
git commit -m "docs: <conventional-commit-message-in-english>"   # no Co-Authored-By line
git push -u origin docs/readme-<whatever>
glab mr create --source-branch docs/readme-<whatever> --target-branch <default_branch> --title "..." --description "..."
```

Then decide on merge timing:
- **No CI, or the repo owner has said "no CI, merge directly"**: `glab mr merge <id> --yes` right away.
- **Has real CI**: it's still fine to merge a docs-only change immediately (nothing it touches affects the pipeline), but ask first if unsure — don't assume every repo works like the no-CI ones.

## 7. Org-hygiene things worth flagging (don't auto-fix)

These involve people/access, not just files — check and report, then ask before changing anything:

- **Branch protection can be hollow.** Check the default branch's protection AND the member list together:
  ```bash
  glab api "projects/$enc/protected_branches"
  glab api "projects/$enc/members/all" | python3 -c "
  import json,sys
  for m in json.load(sys.stdin): print(m['username'], m['access_level'])
  "
  ```
  Access levels: 10=Guest, 20=Reporter, 30=Developer, 40=Maintainer, 50=Owner. If a branch is "protected to Maintainers" but most/all members are ALSO Maintainer, the protection does nothing — anyone can push straight to the default branch or self-approve. Recommend most members sit at Developer, with only 1-2 people at Maintainer/Owner.
- **Required MR approval rules** — often unset by default (anyone can merge their own MR with zero review).
- **CODEOWNERS** for sensitive paths (payment modules, config/credentials-adjacent files, overrides).
- **Secret push protection** — GitLab has a built-in toggle that blocks commits containing likely secrets before they land.

Report findings plainly (e.g. "14/14 members are Maintainer, branch protection is effectively decorative") and let the user decide the fix — this is their team's access model, not a call to make unilaterally.
