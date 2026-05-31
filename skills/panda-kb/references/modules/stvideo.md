---
module_name: stvideo
display_name: Product video and size chart
version: 1.1.1
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: content
hooks: [actionObjectCategoryDeleteAfter, actionObjectManufacturerDeleteAfter, actionShopDataDuplication, displayAdminProductsExtra, displayHeader, displayProductExtraContent]
hooks_count: 6
is_configurable: true
files_count: 22
loc: 2254
author: SUNNYTOO.COM
source_path: dependencies/modules/stvideo
tags: [panda, module, content]
---

# Product video and size chart

> [!info] Panda module
> **Category:** content · **Version:** 1.1.1 · **Hooks:** 6 · **LOC:** 2,254 across 22 files.

## Purpose

Show youtube, vimeo and size chart on your product page.

## Hooks registered

Total: **6** hooks.

- `actionObjectCategoryDeleteAfter`
- `actionObjectManufacturerDeleteAfter`
- `actionShopDataDuplication`
- `displayAdminProductsExtra`
- `displayHeader`
- `displayProductExtraContent`

## Admin configuration

Admin template(s) present: `stvideo.tpl`.

See source at `stvideo/views/templates/admin/` for the full form layout.

## Options (from translations)

Translations folder present at `stvideo/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stvideo.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stvideo/config.xml`
- `stvideo/stvideo.php`
- `stvideo/views/templates/admin/` (1 template)

> [!note] Full source path
> `/dependencies/modules/stvideo/` in the Panda theme archive.
