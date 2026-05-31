---
module_name: stupgrader
display_name: stupgrader
version: unknown
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: helpers
hooks: []
hooks_count: 0
is_configurable: false
files_count: 0
loc: 0
author: SunnyToo.com
source_path: dependencies/modules/stupgrader
tags: [panda, module, helpers]
---

# stupgrader

> [!info] Panda module
> **Category:** helpers · **Version:** ? · **Hooks:** 0 · **LOC:** 0 across 0 files.

## Purpose

No description in config.xml. See source for details.

## Hooks registered

_No hooks declared in `$this->hooks` or `registerHook()` calls — this module likely acts as a library, override, or admin-only tool._

## Admin configuration

No BO config screen detected. Module runs with defaults from hooks, or exposes config through the theme editor (`stthemeeditor`).

## Options (from translations)

No translations folder — module has no user-facing UI strings, or strings live in shared theme translations.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stupgrader.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- Zero hooks declared — verify whether this module is a library/helper (e.g. translation generator, override loader) rather than a user-facing module.
- No translations folder — the module may rely on parent/shared Panda strings.
- Upgrade tool — destructive operations possible. Back up before running in production.

## Source references

- `stupgrader/config.xml`
- `stupgrader/stupgrader.php`

> [!note] Full source path
> `/dependencies/modules/stupgrader/` in the Panda theme archive.
