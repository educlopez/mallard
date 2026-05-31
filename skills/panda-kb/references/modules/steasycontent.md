---
module_name: steasycontent
display_name: Advanced custom content
version: 1.7.0
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: content
hooks: [actionProductSave, actionShopDataDuplication, displayAdminProductsExtra, displayCMSExtra, displayHeader, displayModuleCustomContent, vcBeforeInit]
hooks_count: 7
is_configurable: true
files_count: 182
loc: 16048
author: SUNNYTOO.COM
source_path: dependencies/modules/steasycontent
tags: [panda, module, content]
---

# Advanced custom content

> [!info] Panda module
> **Category:** content · **Version:** 1.7.0 · **Hooks:** 7 · **LOC:** 16,048 across 182 files.

## Purpose

This module is used to create tabs, accordions, FAQ, parallax block, products sliders, lists, textbox sliders, testimonals and more.

## Hooks registered

Total: **7** hooks.

- `actionProductSave`
- `actionShopDataDuplication`
- `displayAdminProductsExtra`
- `displayCMSExtra`
- `displayHeader`
- `displayModuleCustomContent`
- `vcBeforeInit`

## Admin configuration

Admin template(s) present: `steasycontent.tpl`.

See source at `steasycontent/views/templates/admin/` for the full form layout.

## Options (from translations)

Translations folder present at `steasycontent/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `steasycontent.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- Very large module (16,048 LOC) — likely bundles vendor assets (jQuery plugins, WYSIWYG, etc.). Heavy on the asset pipeline.

## Source references

- `steasycontent/config.xml`
- `steasycontent/steasycontent.php`
- `steasycontent/views/templates/admin/` (1 template)
- Public docs: [[docs/module-advanced-custom-content]]

> [!note] Full source path
> `/dependencies/modules/steasycontent/` in the Panda theme archive.
