---
module_name: stowlcarousel
display_name: Owl Carousel slider
version: 2.2.2
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: helpers
hooks: [actionObjectCategoryDeleteAfter, actionObjectManufacturerDeleteAfter, actionShopDataDuplication, displayCategoryHeader, displayCMSExtra, displayHeader, displayManufacturerHeader, vcBeforeInit]
hooks_count: 8
is_configurable: true
files_count: 44
loc: 4727
author: SUNNYTOO.COM
source_path: dependencies/modules/stowlcarousel
tags: [panda, module, helpers]
---

# Owl Carousel slider

> [!info] Panda module
> **Category:** helpers · **Version:** 2.2.2 · **Hooks:** 8 · **LOC:** 4,727 across 44 files.

## Purpose

Touch enabled jquery plugin that lets you create beautiful responsive carousel slider.

## Hooks registered

Total: **8** hooks.

- `actionObjectCategoryDeleteAfter`
- `actionObjectManufacturerDeleteAfter`
- `actionShopDataDuplication`
- `displayCategoryHeader`
- `displayCMSExtra`
- `displayHeader`
- `displayManufacturerHeader`
- `vcBeforeInit`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stowlcarousel.php`.

## Options (from translations)

Translations folder present at `stowlcarousel/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

- Provides slider/carousel functionality for other Panda modules (Owl Carousel wrapper).

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stowlcarousel.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stowlcarousel/config.xml`
- `stowlcarousel/stowlcarousel.php`

> [!note] Full source path
> `/dependencies/modules/stowlcarousel/` in the Panda theme archive.
