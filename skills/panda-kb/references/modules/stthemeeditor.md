---
module_name: stthemeeditor
display_name: Theme editor
version: 2.9.2
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: helpers
hooks: [actionProductSearchAfter, actionShopDataDuplication, actionStAssemble, displayHeader, displayRightColumnProduct]
hooks_count: 5
is_configurable: true
files_count: 153
loc: 25271
author: SUNNYTOO.COM
source_path: dependencies/modules/stthemeeditor
tags: [panda, module, helpers]
---

# Theme editor

> [!info] Panda module
> **Category:** helpers · **Version:** 2.9.2 · **Hooks:** 5 · **LOC:** 25,271 across 153 files.

## Purpose

Allows to change theme design

## Hooks registered

Total: **5** hooks.

- `actionProductSearchAfter`
- `actionShopDataDuplication`
- `actionStAssemble`
- `displayHeader`
- `displayRightColumnProduct`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stthemeeditor.php`.

## Options (from translations)

Translations folder present at `stthemeeditor/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

- Central theme configuration hub. Many other Panda modules read their options through this editor.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stthemeeditor.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- Very large module (25,271 LOC) — likely bundles vendor assets (jQuery plugins, WYSIWYG, etc.). Heavy on the asset pipeline.

## Source references

- `stthemeeditor/config.xml`
- `stthemeeditor/stthemeeditor.php`

> [!tip] Deep-dive
> Esta fiche es shallow. Para el mapa completo de 38 tabs + 529 Configuration keys + hook pipeline + lookup matrix, ver [[stthemeeditor-deep-dive]].

> [!note] Full source path
> `/dependencies/modules/stthemeeditor/` in the Panda theme archive.
