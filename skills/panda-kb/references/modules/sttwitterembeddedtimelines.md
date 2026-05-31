---
module_name: sttwitterembeddedtimelines
display_name: Twitter Embedded Timelines
version: 1.0
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: content
hooks: [displayFooter, vcBeforeInit]
hooks_count: 2
is_configurable: true
files_count: 11
loc: 1004
author: SUNNYTOO.COM
source_path: dependencies/modules/sttwitterembeddedtimelines
tags: [panda, module, content]
---

# Twitter Embedded Timelines

> [!info] Panda module
> **Category:** content · **Version:** 1.0 · **Hooks:** 2 · **LOC:** 1,004 across 11 files.

## Purpose

Display the recent tweets of a twitter user

## Hooks registered

Total: **2** hooks.

- `displayFooter`
- `vcBeforeInit`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `sttwitterembeddedtimelines.php`.

## Options (from translations)

Translations folder present at `sttwitterembeddedtimelines/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `sttwitterembeddedtimelines.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `sttwitterembeddedtimelines/config.xml`
- `sttwitterembeddedtimelines/sttwitterembeddedtimelines.php`

> [!note] Full source path
> `/dependencies/modules/sttwitterembeddedtimelines/` in the Panda theme archive.
