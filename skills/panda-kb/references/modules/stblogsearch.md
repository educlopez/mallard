---
module_name: stblogsearch
display_name: Blog Module - Search
version: 1.0
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: blog
hooks: [displayStBlogLeftColumn, displayStBlogRightColumn, moduleRoutes]
hooks_count: 3
is_configurable: true
files_count: 19
loc: 1179
author: SUNNYTOO.COM
source_path: dependencies/modules/stblogsearch
tags: [panda, module, blog]
---

# Blog Module - Search

> [!info] Panda module
> **Category:** blog · **Version:** 1.0 · **Hooks:** 3 · **LOC:** 1,179 across 19 files.

## Purpose

Blog search.

## Hooks registered

Total: **3** hooks.

- `displayStBlogLeftColumn`
- `displayStBlogRightColumn`
- `moduleRoutes`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stblogsearch.php`.

## Options (from translations)

Translations folder present at `stblogsearch/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

- **Requires `stblog`** — this is a blog sub-module that depends on the core blog engine.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stblogsearch.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stblogsearch/config.xml`
- `stblogsearch/stblogsearch.php`

> [!note] Full source path
> `/dependencies/modules/stblogsearch/` in the Panda theme archive.
