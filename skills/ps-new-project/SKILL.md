---
name: ps-new-project
description: >
  Orchestrate the day-1 setup that remains after cloning a PrestaShop repo and
  running lando-setup: scaffold the child theme (panda or elementflow track),
  wire the CSS build, and write CLAUDE.md. Composes existing skills
  (ps-css-build, ps-project-doc, lando-setup). Does NOT create the GitLab repo,
  the server, or install PrestaShop core. Use when starting a new client
  project locally and you need the standard scaffold fast.
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# ps-new-project — new client day-1 scaffold

Runs the repetitive setup that's still manual after clone + `lando-setup`: **child theme → CSS build → CLAUDE.md**. Pure orchestration over skills that already exist. Server/repo/PS-core stay out of scope (manual / PM-owned).

## Universal rule: always a child theme
Every project gets a child theme, on **both** tracks — you never edit the parent (`panda`/`classic`) directly. The child is the safe override layer; parent files stay untouched so parent updates don't wipe your work. Creating the child is not optional.

## Tracks
- `panda` — legacy majority: child of the Panda theme (`parent: panda`). Storefront work (templates + CSS) happens in the child.
- `elementflow` — emerging default: child of `classic`, relies on the `stsitebuilder` module. The child is still **mandatory** (same rule). What differs is *where the visual work goes*: in ElementFlow the site layout/content is built in the **builder** (stsitebuilder), so the child's hand-work concentrates on the **checkout and my-account pages** (building those with the builder causes compatibility issues). Don't build the whole site in the child — but you still always have one.

**EF child — prefer the official download over generating.** The canonical EF child theme is downloaded from the store itself so it **preserves existing image types** (installing a theme on PS 1.7/8 erases image types; the BO download re-adds them):
- PS 1.7 / 8: **BO → Element Flow → Settings → Child theme** (downloads with the store's image types baked in).
- PS 9: the `elementflow.zip` (classic) / `elementflow_hb.zip` (Hummingbird) — PS 9 keeps image types on install.

This skill's generated EF skeleton (below) is a **fallback/starting point** only. When the store already has the `stsitebuilder` module, tell the user to grab the official child instead, and warn: on PS 1.7/8 a hand-generated child will not carry the store's image types — re-add them or use the BO download. Consult the `elementflow-kb` skill for builder specifics (navigation items, terms modal `terms-and-conditions-modal.tpl`, one-column mobile checkout).

The CSS build for both tracks is the **same** lightningcss pipeline `ps-css-build` ships (bundle+minify `_dev/css/custom.css` → `assets/css/custom.css`, chokidar watch, zero-dep pre-commit fallback), so one build system serves both.

## Flow

### 0. Preconditions
Confirm cwd is a PS repo (`themes/` + `modules/`). Detect installed base: `themes/panda` → Panda available; `modules/stsitebuilder` → ElementFlow available. If `--track` mismatches what's installed, warn before proceeding.

### 1. Child theme scaffold
```bash
~/.claude/skills/ps-new-project/scripts/scaffold-theme.sh <track> <name> "<display_name>" .
```
- Default `<name>` = `<repo-slug>-child` (derive slug from the repo dir name; override with `--name`).
- The script creates `themes/<name>/` with `config/theme.yml` (per track), `config/.htaccess`, `templates/{_partials,catalog,checkout,customer,errors}/`, `modules/`, `assets/{css,js}/`.
- For `elementflow`, verify `modules/stsitebuilder` exists; warn if missing. **Also recommend the official EF child download** (BO → Element Flow → Settings → Child theme on PS 1.7/8, or the zip on PS 9) since it preserves image types — the generated skeleton is a fallback. On PS 1.7/8 warn that a generated child won't carry the store's image types.
- It refuses to clobber an existing `themes/<name>/`.

### 2. CSS build — reuse `ps-css-build`
Run the `ps-css-build` skill targeting `themes/<name>/` (lightningcss + pnpm + pre-commit). Same for both tracks. Skip with `--no-css`.

### 3. CLAUDE.md — reuse `ps-project-doc` (non-interactive)
Call `ps-project-doc` with known values (track, theme name, PS version, git remote) so it writes CLAUDE.md without re-prompting. Stays gitignored, never committed. Skip with `--no-doc`.

### 4. Lando — reuse `lando-setup` (optional)
If Lando isn't configured and `--no-lando` not passed, hand off to `lando-setup`. Usually already done → skip.

### 5. Report
List created/skipped, the child theme path, and remaining manual steps: activate the theme in the Back Office, and (if from scratch) create the repo/server — which this skill deliberately does not touch.

## Hard rules
- Never create the GitLab repo or server, never install PS core, never deploy.
- Never clobber an existing child theme without explicit confirmation.
- English-only content; no company names in the skill. (Generated CLAUDE.md may name the client — it's private/untracked.)
