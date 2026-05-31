---
module_name: stproductlinknav
display_name: Next and previous products
version: 1.0
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: product
hooks: [displayFooterProduct, displayProductExtraContent]
hooks_count: 2
is_configurable: false
files_count: 13
loc: 501
author: SUNNYTOO.COM
source_path: dependencies/modules/stproductlinknav
tags: [panda, module, product]
---

# Next and previous products

> [!info] Panda module
> **Category:** product · **Version:** 1.0 · **Hooks:** 2 · **LOC:** 501 across 13 files.

## Purpose

This module adds next and previous links on the product page.

## Hooks registered

Total: **2** hooks.

- `displayFooterProduct`
- `displayProductExtraContent`

## Admin configuration

No BO config screen detected. Module runs with defaults from hooks, or exposes config through the theme editor (`stthemeeditor`).

## Options (from translations)

Translations folder present at `stproductlinknav/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stproductlinknav.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stproductlinknav/config.xml`
- `stproductlinknav/stproductlinknav.php`

> [!note] Full source path
> `/dependencies/modules/stproductlinknav/` in the Panda theme archive.
