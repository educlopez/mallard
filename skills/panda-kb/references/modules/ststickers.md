---
module_name: ststickers
display_name: Stickers block
version: 1.1.3
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: content
hooks: [actionObjectCategoryDeleteAfter, actionObjectManufacturerDeleteAfter, actionProductSearchAfter, actionShopDataDuplication, actionStAssemble, displayAdminProductsExtra, displayHeader, displayProductExtraContent]
hooks_count: 8
is_configurable: true
files_count: 26
loc: 2953
author: SUNNYTOO.COM
source_path: dependencies/modules/ststickers
tags: [panda, module, content]
---

# Stickers block

> [!info] Panda module
> **Category:** content · **Version:** 1.1.3 · **Hooks:** 8 · **LOC:** 2,953 across 26 files.

## Purpose

Show sticker icon or label on product images.

## Hooks registered

Total: **8** hooks.

- `actionObjectCategoryDeleteAfter`
- `actionObjectManufacturerDeleteAfter`
- `actionProductSearchAfter`
- `actionShopDataDuplication`
- `actionStAssemble`
- `displayAdminProductsExtra`
- `displayHeader`
- `displayProductExtraContent`

## Admin configuration

Admin template(s) present: `ststickers.tpl`.

See source at `ststickers/views/templates/admin/` for the full form layout.

## Options (from translations)

Translations folder present at `ststickers/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `ststickers.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `ststickers/config.xml`
- `ststickers/ststickers.php`
- `ststickers/views/templates/admin/` (1 template)
- Public docs: [[docs/module-stickers-block]]

> [!note] Full source path
> `/dependencies/modules/ststickers/` in the Panda theme archive.
