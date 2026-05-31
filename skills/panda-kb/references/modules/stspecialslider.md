---
module_name: stspecialslider
display_name: Special Products Slider
version: 1.1.0
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: product
hooks: [displayAdminProductPriceFormFooter, displayFooter]
hooks_count: 2
is_configurable: true
files_count: 20
loc: 955
author: SUNNYTOO.COM
source_path: dependencies/modules/stspecialslider
tags: [panda, module, product]
---

# Special Products Slider

> [!info] Panda module
> **Category:** product · **Version:** 1.1.0 · **Hooks:** 2 · **LOC:** 955 across 20 files.

## Purpose

Display special products slider on hompage.

## Hooks registered

Total: **2** hooks.

- `displayAdminProductPriceFormFooter`
- `displayFooter`

## Admin configuration

Admin template(s) present: `stspecialslider.tpl`.

See source at `stspecialslider/views/templates/admin/` for the full form layout.

## Options (from translations)

Translations folder present at `stspecialslider/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stspecialslider.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stspecialslider/config.xml`
- `stspecialslider/stspecialslider.php`
- `stspecialslider/views/templates/admin/` (1 template)

> [!note] Full source path
> `/dependencies/modules/stspecialslider/` in the Panda theme archive.
