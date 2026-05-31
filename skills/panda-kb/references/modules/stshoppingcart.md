---
module_name: stshoppingcart
display_name: Shopping cart mod
version: 1.5.8
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: helpers
hooks: [displayHeader, displayMobileBar, displaySideBar, displayTop]
hooks_count: 4
is_configurable: true
files_count: 31
loc: 2168
author: SUNNYTOO.COM
source_path: dependencies/modules/stshoppingcart
tags: [panda, module, helpers]
---

# Shopping cart mod

> [!info] Panda module
> **Category:** helpers · **Version:** 1.5.8 · **Hooks:** 4 · **LOC:** 2,168 across 31 files.

## Purpose

Adds a block containing the customer&#039;s shopping cart.

## Hooks registered

Total: **4** hooks.

- `displayHeader`
- `displayMobileBar`
- `displaySideBar`
- `displayTop`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stshoppingcart.php`.

## Options (from translations)

Translations folder present at `stshoppingcart/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stshoppingcart.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stshoppingcart/config.xml`
- `stshoppingcart/stshoppingcart.php`

> [!note] Full source path
> `/dependencies/modules/stshoppingcart/` in the Panda theme archive.
