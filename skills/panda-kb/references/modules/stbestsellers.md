---
module_name: stbestsellers
display_name: Top sellers Slider
version: 1.1.1
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: product
hooks: [actionOrderStatusPostUpdate, displayLeftColumn]
hooks_count: 2
is_configurable: true
files_count: 17
loc: 613
author: SUNNYTOO.COM
source_path: dependencies/modules/stbestsellers
tags: [panda, module, product]
---

# Top sellers Slider

> [!info] Panda module
> **Category:** product · **Version:** 1.1.1 · **Hooks:** 2 · **LOC:** 613 across 17 files.

## Purpose

Display top-selling products on homepage.

## Hooks registered

Total: **2** hooks.

- `actionOrderStatusPostUpdate`
- `displayLeftColumn`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stbestsellers.php`.

## Options (from translations)

Translations folder present at `stbestsellers/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stbestsellers.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stbestsellers/config.xml`
- `stbestsellers/stbestsellers.php`

> [!note] Full source path
> `/dependencies/modules/stbestsellers/` in the Panda theme archive.
