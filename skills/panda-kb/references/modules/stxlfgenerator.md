---
module_name: stxlfgenerator
display_name: Translation files generator
version: 1.0.0
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: helpers
hooks: []
hooks_count: 0
is_configurable: true
files_count: 11
loc: 919
author: SUNNYTOO.COM
source_path: dependencies/modules/stxlfgenerator
tags: [panda, module, helpers]
---

# Translation files generator

> [!info] Panda module
> **Category:** helpers · **Version:** 1.0.0 · **Hooks:** 0 · **LOC:** 919 across 11 files.

## Purpose

Generate xlif files in order to translate the theme.

## Hooks registered

_No hooks declared in `$this->hooks` or `registerHook()` calls — this module likely acts as a library, override, or admin-only tool._

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stxlfgenerator.php`.

## Options (from translations)

Translations folder present at `stxlfgenerator/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stxlfgenerator.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- Zero hooks declared — verify whether this module is a library/helper (e.g. translation generator, override loader) rather than a user-facing module.
- Dev/build helper — not intended for end users.

## Source references

- `stxlfgenerator/config.xml`
- `stxlfgenerator/stxlfgenerator.php`

> [!note] Full source path
> `/dependencies/modules/stxlfgenerator/` in the Panda theme archive.
