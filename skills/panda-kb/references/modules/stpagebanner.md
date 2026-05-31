---
module_name: stpagebanner
display_name: Breadcrumbs
version: 1.7.0
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: product
hooks: [actionObjectCategoryDeleteAfter, actionObjectCmsDeleteAfter, actionObjectManufacturerDeleteAfter, actionObjectSupplierDeleteAfter, actionShopDataDuplication, displayBreadcrumb, displayHeader]
hooks_count: 7
is_configurable: true
files_count: 26
loc: 2938
author: SUNNYTOO.COM
source_path: dependencies/modules/stpagebanner
tags: [panda, module, product]
---

# Breadcrumbs

> [!info] Panda module
> **Category:** product · **Version:** 1.7.0 · **Hooks:** 7 · **LOC:** 2,938 across 26 files.

## Purpose

This module allows you put different backgrounds to breadcrumbs on different pages.

## Hooks registered

Total: **7** hooks.

- `actionObjectCategoryDeleteAfter`
- `actionObjectCmsDeleteAfter`
- `actionObjectManufacturerDeleteAfter`
- `actionObjectSupplierDeleteAfter`
- `actionShopDataDuplication`
- `displayBreadcrumb`
- `displayHeader`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stpagebanner.php`.

## Options (from translations)

Translations folder present at `stpagebanner/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stpagebanner.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stpagebanner/config.xml`
- `stpagebanner/stpagebanner.php`

> [!note] Full source path
> `/dependencies/modules/stpagebanner/` in the Panda theme archive.
