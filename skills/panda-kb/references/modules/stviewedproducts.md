---
module_name: stviewedproducts
display_name: Viewed products block
version: 1.2.4
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: product
hooks: [actionProductDelete, displayProductAdditionalInfo, displaySideBar]
hooks_count: 3
is_configurable: true
files_count: 22
loc: 957
author: SUNNYTOO.COM
source_path: dependencies/modules/stviewedproducts
tags: [panda, module, product]
---

# Viewed products block

> [!info] Panda module
> **Category:** product · **Version:** 1.2.4 · **Hooks:** 3 · **LOC:** 957 across 22 files.

## Purpose

Adds a block displaying recently viewed products.

## Hooks registered

Total: **3** hooks.

- `actionProductDelete`
- `displayProductAdditionalInfo`
- `displaySideBar`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stviewedproducts.php`.

## Options (from translations)

Translations folder present at `stviewedproducts/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stviewedproducts.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stviewedproducts/config.xml`
- `stviewedproducts/stviewedproducts.php`

> [!note] Full source path
> `/dependencies/modules/stviewedproducts/` in the Panda theme archive.
