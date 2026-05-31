---
module_name: stblog
display_name: Blog module
version: 1.7.1
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: blog
hooks: [actionShopDataDuplication, displayBackOfficeHeader, gSitemapAppendUrls, moduleRoutes, overrideLayoutTemplate]
hooks_count: 5
is_configurable: true
files_count: 118
loc: 12796
author: SUNNYTOO.COM
source_path: dependencies/modules/stblog
tags: [panda, module, blog]
---

# Blog module

> [!info] Panda module
> **Category:** blog · **Version:** 1.7.1 · **Hooks:** 5 · **LOC:** 12,796 across 118 files.

## Purpose

Blog module for Prestashop

## Hooks registered

Total: **5** hooks.

- `actionShopDataDuplication`
- `displayBackOfficeHeader`
- `gSitemapAppendUrls`
- `moduleRoutes`
- `overrideLayoutTemplate`

## Admin configuration

Admin controller(s): `index`, `AdminStBlogController`, `AdminStBlogCategoryController`.

Admin template(s) present: `content.tpl`.

See source at `stblog/views/templates/admin/` for the full form layout.

## Options (from translations)

Translations folder present at `stblog/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stblog.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- Very large module (12,796 LOC) — likely bundles vendor assets (jQuery plugins, WYSIWYG, etc.). Heavy on the asset pipeline.

## Source references

- `stblog/config.xml`
- `stblog/stblog.php`
- `stblog/controllers/admin/index.php`
- `stblog/controllers/admin/AdminStBlogController.php`
- `stblog/controllers/admin/AdminStBlogCategoryController.php`
- `stblog/views/templates/admin/` (1 template)

> [!note] Full source path
> `/dependencies/modules/stblog/` in the Panda theme archive.
