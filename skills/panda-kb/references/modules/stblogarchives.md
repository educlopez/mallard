---
module_name: stblogarchives
display_name: Blog Module - Archives
version: 1.1.3
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: blog
hooks: [displayHeader, displayStBlogLeftColumn, displayStBlogRightColumn, gSitemapAppendUrls, moduleRoutes]
hooks_count: 5
is_configurable: false
files_count: 18
loc: 892
author: SUNNYTOO.COM
source_path: dependencies/modules/stblogarchives
tags: [panda, module, blog]
---

# Blog Module - Archives

> [!info] Panda module
> **Category:** blog · **Version:** 1.1.3 · **Hooks:** 5 · **LOC:** 892 across 18 files.

## Purpose

The archives module allows you to display a tree list of the months and past months.

## Hooks registered

Total: **5** hooks.

- `displayHeader`
- `displayStBlogLeftColumn`
- `displayStBlogRightColumn`
- `gSitemapAppendUrls`
- `moduleRoutes`

## Admin configuration

No BO config screen detected. Module runs with defaults from hooks, or exposes config through the theme editor (`stthemeeditor`).

## Options (from translations)

Translations folder present at `stblogarchives/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

- **Requires `stblog`** — this is a blog sub-module that depends on the core blog engine.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stblogarchives.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stblogarchives/config.xml`
- `stblogarchives/stblogarchives.php`

> [!note] Full source path
> `/dependencies/modules/stblogarchives/` in the Panda theme archive.
