---
title: themes/ folder layout in PrestaShop 9
source: direct inspection of clean PS 9.1.0 source clone, 2026-05-20
ps_version: [9]
tags: [themes, ps9, folder-structure, _core, _libraries, hummingbird]
last-updated: 2026-05-20
---

# `themes/` folder layout — PrestaShop 9

> [!info] PS 9 ships with shared theme resources at the `themes/` root, not nested per theme. New compared to PS 8.

## Canonical layout (clean PS 9 source clone with Panda + child + EB)

```
themes/
├── _core/                    # shared JS used by all themes (PS 9)
│   └── js/
├── _libraries/               # shared libraries used by all themes (PS 9)
│   └── font-awesome/
│       ├── css/
│       └── fonts/
├── classic/                  # PS bundled default (legacy fallback)
├── hummingbird/              # PS 9 default theme (replaces classic as primary)
├── panda/                    # premium theme (SunnyToo) — present if installed
├── panda-v2_child/           # child of panda — your custom code lives here
├── debug.tpl                 # global debug bar template
├── javascript.tpl            # global JS loader template
├── index.php                 # PrestaShop standard guard file
├── package.json              # npm config for compiling theme assets (mainly Hummingbird)
├── package-lock.json
├── preview-fallback.png      # fallback theme preview image
├── webpack.config.js         # asset compilation entrypoint
├── .eslintrc.js              # JS linting for theme source
├── .eslintignore
└── .htaccess
```

## Items by role

### Real themes (folders with `config/theme.yml`)

- `classic/` — legacy default, always present in PS source.
- `hummingbird/` — **PS 9 default theme**. Requires `npm ci && npm run build` from `themes/` root if you want to activate it (assets must be compiled).
- `panda/` — premium, version-locked at `2.9.2` in the current ecosystem. Parent theme.
- `panda-v2_child/` (or `<name>_child/`) — child theme with `parent: panda` in its `config/theme.yml`.

### Shared resources (NOT themes — don't confuse them)

- `_core/js/` — shared JS modules used across themes. Treat as read-only; PS 9 expects them at this path.
- `_libraries/font-awesome/` — font-awesome icons used by multiple themes.

### Build & dev files

- `package.json`, `package-lock.json`, `webpack.config.js` — used by Hummingbird's asset pipeline (and any future SunnyToo theme that opts in). Run `npm ci && npm run build` from `themes/` to compile.
- `.eslintrc.js`, `.eslintignore` — JS linting for theme source.

### Shared templates

- `debug.tpl` — rendered when `_PS_MODE_DEV_ = true`. Bottom debug bar.
- `javascript.tpl` — used by `Tools::renderHTMLContent()` to inject inline scripts safely.
- `index.php` — `header("Location: ../"); exit;` guard. Standard PS pattern across many folders.

## Discovery checklist for agents

When entering a PS 9 project, distinguish themes from non-themes:

```bash
# REAL themes only (have config/theme.yml)
for d in themes/*/; do
  [[ -f "$d/config/theme.yml" ]] && echo "theme: $d"
done

# Shared resources to ignore as theme candidates
ls themes/_core themes/_libraries 2>/dev/null
```

## Common gotchas

1. **Hummingbird present but broken visually**: it requires asset compilation. Run `npm ci && npm run build` from `themes/` root. Without that, switching to Hummingbird shows raw markup with no styling.

2. **Child theme name does NOT need to follow `<parent>_child`** convention. `panda-v2_child` works. The relationship is in `config/theme.yml` via `parent:` field, not in folder naming.

3. **`_core/` and `_libraries/` are NOT themes**. Some agents may try to read their `config/theme.yml` — it doesn't exist. Check `[[ -f "$d/config/theme.yml" ]]` before treating as theme.

4. **Active theme is in DB**, not derived from folder presence. Query:
   ```sql
   SELECT id_shop, theme_name FROM ps_shop;
   ```

5. **package.json at `themes/` root, not at theme level** — affects `npm install` location. Run from `themes/`, not from `themes/hummingbird/`.

## Compare with PS 8 layout

| Item | PS 8 | PS 9 |
|---|---|---|
| Shared `_core/js/` | absent (per-theme JS) | present, shared |
| Shared `_libraries/font-awesome/` | absent (per-theme) | present, shared |
| `hummingbird/` | absent | present (default) |
| `classic/` | present (default) | present (legacy) |
| `debug.tpl`, `javascript.tpl` at root | absent | present |
| `package.json` at root | absent | present |

## Related KB

- [[parent-child-feature]] — child theme mechanism (`parent:` field).
- [[hummingbird-v2]] — Hummingbird v2 decision matrix vs other themes.
- [[theme-yml]] — theme.yml fields canonical reference.
- [[recovery-parent-edits]] — recovery protocol when client edits parent.
