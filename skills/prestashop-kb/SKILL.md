---
name: prestashop-kb
description: Curated knowledge base for the PrestaShop platform — versions 8 and 9. Covers theme mechanics (theme.yml, parent-child, template inheritance, assets), Smarty templating + PS extensions ({l}, {hook}, {widget}), hooks (overview + listings), modules architecture (config.xml, controllers, lifecycle), install vs distribution (PS 9 source clone), and version migration PS 8 → 9 (Symfony BO, Twig, upgrade process). Use for platform-level PrestaShop questions independent of any specific theme. For Panda-theme-specific questions (`st*` modules, SunnyToo, Easy Builder) use the `panda-kb` skill instead.
version: "0.1.3"
metadata:
  author: Eduardo Calvo
---

# PrestaShop Knowledge Base

This skill exposes the curated PrestaShop platform knowledge base bundled with the `prestashop-experts` plugin. The `prestashop-expert` agent consults this KB before answering, but any agent can read these files directly.

## When this skill is relevant

Load and read from this KB when the user mentions:

- **PrestaShop** mechanics independent of theme — hooks, modules, Smarty, themes parent-child, theme.yml, install vs distribution, migration.
- **PS 8 vs PS 9** differences.
- **Symfony BO** (the new Symfony 6 + Twig back office in PS 9).
- **Smarty** templating in PS context — `{l}`, `{hook}`, `{widget}`, `{capture}`.
- **PS source clone** vs distribution zip (`install-dev/` vs `install/`, `admin-dev/` vs `admin/`).
- **Hummingbird v2** theme — the default theme of PS 9 — and "which theme should I use" decision matrix.
- **Composer/vendor**, `parameters.php`, `.env` in a PrestaShop project.

Do NOT load for: Panda-specific modules (use `panda-kb`), non-PrestaShop e-commerce, generic PHP/Symfony questions outside PrestaShop.

## KB layout

```
references/
├── README.md                    # convenciones y índice de la KB
├── themes/                      # 13 docs
│   ├── parent-child-feature.md  # child theme mechanism (file path does NOT contain "child")
│   ├── theme-yml.md
│   ├── template-inheritance.md
│   ├── hummingbird-v2.md        # PS 9 default theme + decision matrix vs Panda
│   └── ... (assets, layouts, etc.)
├── modules/                     # 11 docs: arquitectura, config.xml, controllers, services.yml, lifecycle, hooks-in-modules
├── hooks/                       # 3 docs: overview + listing of available hooks
├── smarty/                      # 40 docs: Smarty templating + PS extensions
├── install/                     # 1 doc: source-vs-distribution.md (PS 9 clone vs zip)
└── version-migration/           # 8 docs: PS 8 → 9, Symfony/Twig BO, upgrade process
```

## How to use the KB

1. **Identify the topic** of the question — themes / hooks / modules / smarty / install / migration — and jump to the matching subfolder.
2. **`themes/parent-child-feature.md`** is the canonical doc on child theme mechanics — note the file does NOT have "child" in its name (discoverability gotcha).
3. **For Symfony BO (PS 9) work**: `version-migration/` has the most context. Source code in CWD lives in `./src/PrestaShopBundle/`.
4. **For legacy controllers (PS 8 + some PS 9 screens)**: still in `./controllers/admin/` and `./classes/`. Not deeply covered in the KB — read source.
5. **Smarty syntax** → `smarty/` (40 files) covers Smarty base + PS extensions. Read before improvising.
6. **install/source-vs-distribution.md**: critical when the CWD has `install-dev/`, `admin-dev/`, or a missing `vendor/` folder — those indicate a source clone, not a distribution zip.

## Key facts to surface to agents

- **PS 9 default theme**: Hummingbird v2 (accessibility-focused, EAA-aligned). Panda is still a valid premium alternative.
- **PHP 8.1+** required in PS 9. PS 8 accepts PHP 7.4+.
- **Symfony BO**: PS 9 moves most BO to Symfony 6 + Twig. Legacy `controllers/admin/AdminXController.php` still exist for non-migrated screens.
- **Active theme is in DB**: `ps_shop.theme_name`. Files in `./themes/` are not authoritative on which theme is active.
- **Hooks**:
  - `displayX` hooks → inject HTML on the front.
  - `actionX` hooks → execute logic on events.
  - Definitive read in CWD: `./classes/Hook.php`, `./src/Core/Module/...`.
  - Installed hooks: DB table `ps_hook`.
- **Theme.yml** is authoritative on layouts, PS compatibility, assets, and child-parent relationship (`parent:` declares the parent — **not** `parent_theme:`).

## Project shape heuristics (PS 9)

When entering a project, check these to identify what you're dealing with:

| Signal | What it means |
|---|---|
| `./install-dev/` (not `./install/`) | Source clone from GitHub, not distribution zip. |
| `./admin-dev/` (not `./admin/`) | Same — `/admin/` URL won't respond until rename. |
| Missing `./vendor/` | `composer install` pending. PS won't boot. |
| `./themes/hummingbird/` present but no compiled assets | Needs `npm ci && npm run build` inside the theme to render. |
| `./.env` / `./app/config/parameters.php` | Env config — DB credentials, secret, etc. |

## Gaps to be honest about

- **Coverage is uneven**: `smarty/` (40 docs) is well covered, `install/` (1 doc) and `hooks/` (3 docs) are minimal. For gaps, read source in CWD or consult `devdocs.prestashop-project.org` externally.
- **PS 8-specific quirks** are less covered than PS 9. If the question is PS 8 and the KB doesn't have it, say so.
- **Performance, caching, profiling** are not deeply covered. Recognize the gap and read source.

## Cross-skill pointer

For Panda-theme-specific questions (`st*` modules, SunnyToo widgets, Easy Builder, Panda demos, "which Panda module for X"), use the sibling skill `panda-kb` and the agent `panda-expert`.
