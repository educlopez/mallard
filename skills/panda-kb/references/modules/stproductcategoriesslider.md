---
module_name: stproductcategoriesslider
display_name: Product Slider
version: 1.7.3
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: product
hooks: [actionCategoryDelete, actionObjectCategoryDeleteAfter, actionProductDelete, actionProductUpdate, addproduct, deleteproduct, updateproduct, vcBeforeInit]
hooks_count: 8
is_configurable: true
files_count: 25
loc: 3263
author: SUNNYTOO.COM
source_path: dependencies/modules/stproductcategoriesslider
tags: [panda, module, product]
---

# Product Slider

> [!info] Panda module
> **Category:** product · **Version:** 1.7.3 · **Hooks:** 8 · **LOC:** 3,263 across 25 files.

## Purpose

Display products from different categories on the homepage.

## Hooks registered

Total: **8** hooks.

- `actionCategoryDelete`
- `actionObjectCategoryDeleteAfter`
- `actionProductDelete`
- `actionProductUpdate`
- `addproduct`
- `deleteproduct`
- `updateproduct`
- `vcBeforeInit`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stproductcategoriesslider.php`.

## Options (from translations)

Translations folder present at `stproductcategoriesslider/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stproductcategoriesslider.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stproductcategoriesslider/config.xml`
- `stproductcategoriesslider/stproductcategoriesslider.php`

> [!note] Full source path
> `/dependencies/modules/stproductcategoriesslider/` in the Panda theme archive.
