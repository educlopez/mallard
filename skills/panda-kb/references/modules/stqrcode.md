---
module_name: stqrcode
display_name: QR code
version: 1.0
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: content
hooks: [displayHeader, displayProductCartRight, displaySideBar]
hooks_count: 3
is_configurable: true
files_count: 14
loc: 817
author: SUNNYTOO.COM
source_path: dependencies/modules/stqrcode
tags: [panda, module, content]
---

# QR code

> [!info] Panda module
> **Category:** content · **Version:** 1.0 · **Hooks:** 3 · **LOC:** 817 across 14 files.

## Purpose

Add QR code to your site.

## Hooks registered

Total: **3** hooks.

- `displayHeader`
- `displayProductCartRight`
- `displaySideBar`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stqrcode.php`.

## Options (from translations)

Translations folder present at `stqrcode/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stqrcode.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stqrcode/config.xml`
- `stqrcode/stqrcode.php`

> [!note] Full source path
> `/dependencies/modules/stqrcode/` in the Panda theme archive.
