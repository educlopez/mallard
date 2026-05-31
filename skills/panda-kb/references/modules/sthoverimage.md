---
module_name: sthoverimage
display_name: Hover image
version: 1.0
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: product
hooks: [actionStAssemble, displayAdminProductsExtra, displayHeader]
hooks_count: 3
is_configurable: true
files_count: 25
loc: 1095
author: SUNNYTOO.COM
source_path: dependencies/modules/sthoverimage
tags: [panda, module, product]
---

# Hover image

> [!info] Panda module
> **Category:** product · **Version:** 1.0 · **Hooks:** 3 · **LOC:** 1,095 across 25 files.

## Purpose

Products thumb image change on mouse hover

## Hooks registered

Total: **3** hooks.

- `actionStAssemble`
- `displayAdminProductsExtra`
- `displayHeader`

## Admin configuration

Admin template(s) present: `sthoverimage.tpl`, `view.tpl`.

See source at `sthoverimage/views/templates/admin/` for the full form layout.

## Options (from translations)

Translations folder present at `sthoverimage/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `sthoverimage.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `sthoverimage/config.xml`
- `sthoverimage/sthoverimage.php`
- `sthoverimage/views/templates/admin/` (2 templates)

> [!note] Full source path
> `/dependencies/modules/sthoverimage/` in the Panda theme archive.
