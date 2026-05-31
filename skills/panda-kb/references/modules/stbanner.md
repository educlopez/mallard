---
module_name: stbanner
display_name: Advanced banner
version: 1.9.0
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: content
hooks: [actionObjectCategoryDeleteAfter, actionObjectManufacturerDeleteAfter, actionOutputHTMLBefore, actionShopDataDuplication, displayCategoryHeader, displayCMSExtra, displayHeader, vcBeforeInit]
hooks_count: 8
is_configurable: true
files_count: 74
loc: 6821
author: SUNNYTOO.COM
source_path: dependencies/modules/stbanner
tags: [panda, module, content]
---

# Advanced banner

> [!info] Panda module
> **Category:** content · **Version:** 1.9.0 · **Hooks:** 8 · **LOC:** 6,821 across 74 files.

## Purpose

This module is used to add banners in your shop.

## Hooks registered

Total: **8** hooks.

- `actionObjectCategoryDeleteAfter`
- `actionObjectManufacturerDeleteAfter`
- `actionOutputHTMLBefore`
- `actionShopDataDuplication`
- `displayCategoryHeader`
- `displayCMSExtra`
- `displayHeader`
- `vcBeforeInit`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stbanner.php`.

## Options (from translations)

Translations folder present at `stbanner/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stbanner.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stbanner/config.xml`
- `stbanner/stbanner.php`
- Public docs: [[docs/module-advanced-banner]]

> [!note] Full source path
> `/dependencies/modules/stbanner/` in the Panda theme archive.
