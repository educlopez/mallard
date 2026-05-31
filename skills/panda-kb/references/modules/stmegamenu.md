---
module_name: stmegamenu
display_name: Megamenu
version: 2.0.8
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: navigation
hooks: [actionCategoryAdd, actionCategoryDelete, actionCategoryUpdate, actionObjectCmsDeleteAfter, actionObjectCmsUpdateAfter, actionObjectManufacturerDeleteAfter, actionObjectManufacturerUpdateAfter, actionObjectSupplierDeleteAfter, actionObjectSupplierUpdateAfter, actionProductAdd, actionProductDelete, actionProductUpdate, actionShopDataDuplication, displayHeader, displayLeftColumn, displayLeftColumnProduct, displayMainMenu, displayMobileBarLeft, displaySideBar]
hooks_count: 19
is_configurable: true
files_count: 45
loc: 6879
author: SUNNYTOO.COM
source_path: dependencies/modules/stmegamenu
tags: [panda, module, navigation]
---

# Megamenu

> [!info] Panda module
> **Category:** navigation · **Version:** 2.0.8 · **Hooks:** 19 · **LOC:** 6,879 across 45 files.

## Purpose

Add a menu on top of your shop.

## Hooks registered

Total: **19** hooks.

- `actionCategoryAdd`
- `actionCategoryDelete`
- `actionCategoryUpdate`
- `actionObjectCmsDeleteAfter`
- `actionObjectCmsUpdateAfter`
- `actionObjectManufacturerDeleteAfter`
- `actionObjectManufacturerUpdateAfter`
- `actionObjectSupplierDeleteAfter`
- `actionObjectSupplierUpdateAfter`
- `actionProductAdd`
- `actionProductDelete`
- `actionProductUpdate`
- `actionShopDataDuplication`
- `displayHeader`
- `displayLeftColumn`
- `displayLeftColumnProduct`
- `displayMainMenu`
- `displayMobileBarLeft`
- `displaySideBar`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stmegamenu.php`.

## Options (from translations)

Translations folder present at `stmegamenu/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stmegamenu.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stmegamenu/config.xml`
- `stmegamenu/stmegamenu.php`
- Public docs: [[docs/module-megamenu]]

> [!note] Full source path
> `/dependencies/modules/stmegamenu/` in the Panda theme archive.
