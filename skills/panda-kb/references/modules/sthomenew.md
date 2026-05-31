---
module_name: sthomenew
display_name: New Products Slider
version: 1.7.9
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: product
hooks: [displayHomeBottom]
hooks_count: 1
is_configurable: true
files_count: 17
loc: 613
author: SUNNYTOO.COM
source_path: dependencies/modules/sthomenew
tags: [panda, module, product]
---

# New Products Slider

> [!info] Panda module
> **Category:** product · **Version:** 1.7.9 · **Hooks:** 1 · **LOC:** 613 across 17 files.

## Purpose

Display new products on homepage.

## Hooks registered

Total: **1** hook.

- `displayHomeBottom`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `sthomenew.php`.

## Options (from translations)

Translations folder present at `sthomenew/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `sthomenew.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `sthomenew/config.xml`
- `sthomenew/sthomenew.php`

> [!note] Full source path
> `/dependencies/modules/sthomenew/` in the Panda theme archive.
