---
module_name: stoverride
display_name: Overriding default behaviors
version: 1.0.0
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: helpers
hooks: []
hooks_count: 0
is_configurable: false
files_count: 10
loc: 286
author: SUNNYTOO.COM
source_path: dependencies/modules/stoverride
tags: [panda, module, helpers]
---

# Overriding default behaviors

> [!info] Panda module
> **Category:** helpers · **Version:** 1.0.0 · **Hooks:** 0 · **LOC:** 286 across 10 files.

## Purpose

Overriding PrestaShops classes and controllers

## Hooks registered

_No hooks declared in `$this->hooks` or `registerHook()` calls — this module likely acts as a library, override, or admin-only tool._

## Admin configuration

No BO config screen detected. Module runs with defaults from hooks, or exposes config through the theme editor (`stthemeeditor`).

## Options (from translations)

Translations folder present at `stoverride/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

- Panda core override loader. Depended on by most `st*` modules.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stoverride.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- Zero hooks declared — verify whether this module is a library/helper (e.g. translation generator, override loader) rather than a user-facing module.

## Source references

- `stoverride/config.xml`
- `stoverride/stoverride.php`

> [!note] Full source path
> `/dependencies/modules/stoverride/` in the Panda theme archive.
