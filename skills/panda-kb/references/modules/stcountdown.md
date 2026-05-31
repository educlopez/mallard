---
module_name: stcountdown
display_name: Countdown module
version: 1.0
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: product
hooks: [actionProductSearchAfter, displayHeader]
hooks_count: 2
is_configurable: true
files_count: 17
loc: 1750
author: SUNNYTOO.COM
source_path: dependencies/modules/stcountdown
tags: [panda, module, product]
---

# Countdown module

> [!info] Panda module
> **Category:** product · **Version:** 1.0 · **Hooks:** 2 · **LOC:** 1,750 across 17 files.

## Purpose

Display countdown special products.

## Hooks registered

Total: **2** hooks.

- `actionProductSearchAfter`
- `displayHeader`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stcountdown.php`.

## Options (from translations)

Translations folder present at `stcountdown/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stcountdown.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stcountdown/config.xml`
- `stcountdown/stcountdown.php`

> [!note] Full source path
> `/dependencies/modules/stcountdown/` in the Panda theme archive.
