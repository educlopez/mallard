---
module_name: stsocial
display_name: Social share buttons
version: 1.0.3
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: content
hooks: [actionProductSearchAfter, displayHeader, displayProductCartRight, displaySideBar]
hooks_count: 4
is_configurable: true
files_count: 30
loc: 2636
author: SUNNYTOO.COM
source_path: dependencies/modules/stsocial
tags: [panda, module, content]
---

# Social share buttons

> [!info] Panda module
> **Category:** content · **Version:** 1.0.3 · **Hooks:** 4 · **LOC:** 2,636 across 30 files.

## Purpose

Dispaly Social sharing buttons on your store.

## Hooks registered

Total: **4** hooks.

- `actionProductSearchAfter`
- `displayHeader`
- `displayProductCartRight`
- `displaySideBar`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stsocial.php`.

## Options (from translations)

Translations folder present at `stsocial/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stsocial.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stsocial/config.xml`
- `stsocial/stsocial.php`

> [!note] Full source path
> `/dependencies/modules/stsocial/` in the Panda theme archive.
