---
module_name: stbrandsslider
display_name: Brands Slider
version: 1.3.7
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: content
hooks: [actionObjectManufacturerDeleteAfter, actionObjectManufacturerUpdateAfter, displayHomeBottom, vcBeforeInit]
hooks_count: 4
is_configurable: true
files_count: 25
loc: 1841
author: SUNNYTOO.COM
source_path: dependencies/modules/stbrandsslider
tags: [panda, module, content]
---

# Brands Slider

> [!info] Panda module
> **Category:** content · **Version:** 1.3.7 · **Hooks:** 4 · **LOC:** 1,841 across 25 files.

## Purpose

Brands slider on your home page.

## Hooks registered

Total: **4** hooks.

- `actionObjectManufacturerDeleteAfter`
- `actionObjectManufacturerUpdateAfter`
- `displayHomeBottom`
- `vcBeforeInit`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stbrandsslider.php`.

## Options (from translations)

Translations folder present at `stbrandsslider/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stbrandsslider.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stbrandsslider/config.xml`
- `stbrandsslider/stbrandsslider.php`

> [!note] Full source path
> `/dependencies/modules/stbrandsslider/` in the Panda theme archive.
