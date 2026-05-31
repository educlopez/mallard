---
module_name: stbloglinknav
display_name: Blog Module - Next and previous links
version: 1.0
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: blog
hooks: [displayStBlogArticleFooter]
hooks_count: 1
is_configurable: false
files_count: 11
loc: 345
author: SUNNYTOO.COM
source_path: dependencies/modules/stbloglinknav
tags: [panda, module, blog]
---

# Blog Module - Next and previous links

> [!info] Panda module
> **Category:** blog · **Version:** 1.0 · **Hooks:** 1 · **LOC:** 345 across 11 files.

## Purpose

This module adds Next and Previous links on the blog page.

## Hooks registered

Total: **1** hook.

- `displayStBlogArticleFooter`

## Admin configuration

No BO config screen detected. Module runs with defaults from hooks, or exposes config through the theme editor (`stthemeeditor`).

## Options (from translations)

Translations folder present at `stbloglinknav/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

- **Requires `stblog`** — this is a blog sub-module that depends on the core blog engine.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stbloglinknav.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stbloglinknav/config.xml`
- `stbloglinknav/stbloglinknav.php`

> [!note] Full source path
> `/dependencies/modules/stbloglinknav/` in the Panda theme archive.
