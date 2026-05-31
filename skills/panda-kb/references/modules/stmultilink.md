---
module_name: stmultilink
display_name: Custom Links
version: 1.9.4
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: navigation
hooks: [actionObjectCategoryDeleteAfter, actionObjectCategoryUpdateAfter, actionObjectCmsDeleteAfter, actionObjectCmsUpdateAfter, actionObjectManufacturerDeleteAfter, actionObjectManufacturerUpdateAfter, actionObjectSupplierDeleteAfter, actionObjectSupplierUpdateAfter, actionShopDataDuplication, displayHeader, displayMobileNav, vcBeforeInit]
hooks_count: 12
is_configurable: true
files_count: 31
loc: 3354
author: SUNNYTOO.COM
source_path: dependencies/modules/stmultilink
tags: [panda, module, navigation]
---

# Custom Links

> [!info] Panda module
> **Category:** navigation · **Version:** 1.9.4 · **Hooks:** 12 · **LOC:** 3,354 across 31 files.

## Purpose

This module is used to manage and display collections of links to your shop.

## Hooks registered

Total: **12** hooks.

- `actionObjectCategoryDeleteAfter`
- `actionObjectCategoryUpdateAfter`
- `actionObjectCmsDeleteAfter`
- `actionObjectCmsUpdateAfter`
- `actionObjectManufacturerDeleteAfter`
- `actionObjectManufacturerUpdateAfter`
- `actionObjectSupplierDeleteAfter`
- `actionObjectSupplierUpdateAfter`
- `actionShopDataDuplication`
- `displayHeader`
- `displayMobileNav`
- `vcBeforeInit`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stmultilink.php`.

## Options (from translations)

Translations folder present at `stmultilink/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stmultilink.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stmultilink/config.xml`
- `stmultilink/stmultilink.php`

> [!note] Full source path
> `/dependencies/modules/stmultilink/` in the Panda theme archive.
