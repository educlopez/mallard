---
module_name: stswiper
display_name: Swiper slider
version: 2.2.2
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: navigation
hooks: [actionObjectCategoryDeleteAfter, actionObjectManufacturerDeleteAfter, actionOutputHTMLBefore, actionShopDataDuplication, DisplayCategoryHeader, displayCMSExtra, displayHeader, displayManufacturerHeader, vcBeforeInit]
hooks_count: 9
is_configurable: true
files_count: 47
loc: 5318
author: SUNNYTOO.COM
source_path: dependencies/modules/stswiper
tags: [panda, module, navigation]
---

# Swiper slider

> [!info] Panda module
> **Category:** navigation · **Version:** 2.2.2 · **Hooks:** 9 · **LOC:** 5,318 across 47 files.

## Purpose

Most modern mobile touch slider.

## Hooks registered

Total: **9** hooks.

- `actionObjectCategoryDeleteAfter`
- `actionObjectManufacturerDeleteAfter`
- `actionOutputHTMLBefore`
- `actionShopDataDuplication`
- `DisplayCategoryHeader`
- `displayCMSExtra`
- `displayHeader`
- `displayManufacturerHeader`
- `vcBeforeInit`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stswiper.php`.

## Options (from translations)

Translations folder present at `stswiper/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

- Provides slider/carousel functionality for other Panda modules (Swiper.js wrapper).

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stswiper.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stswiper/config.xml`
- `stswiper/stswiper.php`

> [!note] Full source path
> `/dependencies/modules/stswiper/` in the Panda theme archive.
