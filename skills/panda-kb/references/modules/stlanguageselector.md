---
module_name: stlanguageselector
display_name: Language selector block mod
version: 1.3.4
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: helpers
hooks: [displayMobileNav, displayNav1]
hooks_count: 2
is_configurable: true
files_count: 15
loc: 751
author: SUNNYTOO.COM
source_path: dependencies/modules/stlanguageselector
tags: [panda, module, helpers]
---

# Language selector block mod

> [!info] Panda module
> **Category:** helpers · **Version:** 1.3.4 · **Hooks:** 2 · **LOC:** 751 across 15 files.

## Purpose

Adds a language switcher to your site.

## Hooks registered

Total: **2** hooks.

- `displayMobileNav`
- `displayNav1`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stlanguageselector.php`.

## Options (from translations)

Translations folder present at `stlanguageselector/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stlanguageselector.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stlanguageselector/config.xml`
- `stlanguageselector/stlanguageselector.php`

> [!note] Full source path
> `/dependencies/modules/stlanguageselector/` in the Panda theme archive.
