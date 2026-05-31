---
module_name: stfeaturedcategories
display_name: Featured categories
version: 1.1.8
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: product
hooks: [actionCategoryAdd, actionCategoryDelete, actionCategoryUpdate, displayHomeBottom, vcBeforeInit]
hooks_count: 5
is_configurable: true
files_count: 21
loc: 2015
author: SUNNYTOO.COM
source_path: dependencies/modules/stfeaturedcategories
tags: [panda, module, product]
---

# Featured categories

> [!info] Panda module
> **Category:** product · **Version:** 1.1.8 · **Hooks:** 5 · **LOC:** 2,015 across 21 files.

## Purpose

Display featured categories on your homepage.

## Hooks registered

Total: **5** hooks.

- `actionCategoryAdd`
- `actionCategoryDelete`
- `actionCategoryUpdate`
- `displayHomeBottom`
- `vcBeforeInit`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stfeaturedcategories.php`.

## Options (from translations)

Translations folder present at `stfeaturedcategories/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stfeaturedcategories.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stfeaturedcategories/config.xml`
- `stfeaturedcategories/stfeaturedcategories.php`

> [!note] Full source path
> `/dependencies/modules/stfeaturedcategories/` in the Panda theme archive.
