---
module_name: sttags
display_name: Tags block
version: 1.2.2
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: content
hooks: [displayFooterAfter, vcBeforeInit]
hooks_count: 2
is_configurable: true
files_count: 13
loc: 896
author: SUNNYTOO.COM
source_path: dependencies/modules/sttags
tags: [panda, module, content]
---

# Tags block

> [!info] Panda module
> **Category:** content · **Version:** 1.2.2 · **Hooks:** 2 · **LOC:** 896 across 13 files.

## Purpose

Adds a block containing your product tags.

## Hooks registered

Total: **2** hooks.

- `displayFooterAfter`
- `vcBeforeInit`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `sttags.php`.

## Options (from translations)

Translations folder present at `sttags/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `sttags.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `sttags/config.xml`
- `sttags/sttags.php`

> [!note] Full source path
> `/dependencies/modules/sttags/` in the Panda theme archive.
