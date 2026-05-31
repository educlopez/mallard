---
module_name: stcurrencyselector
display_name: Currency selector block mod
version: 0.3.3
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: helpers
hooks: [actionAdminCurrenciesControllerSaveAfter, displayMobileNav, displayNav1]
hooks_count: 3
is_configurable: true
files_count: 15
loc: 789
author: SUNNYTOO.COM
source_path: dependencies/modules/stcurrencyselector
tags: [panda, module, helpers]
---

# Currency selector block mod

> [!info] Panda module
> **Category:** helpers · **Version:** 0.3.3 · **Hooks:** 3 · **LOC:** 789 across 15 files.

## Purpose

Adds a block allowing customers to choose their preferred shopping currency.

## Hooks registered

Total: **3** hooks.

- `actionAdminCurrenciesControllerSaveAfter`
- `displayMobileNav`
- `displayNav1`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stcurrencyselector.php`.

## Options (from translations)

Translations folder present at `stcurrencyselector/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stcurrencyselector.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stcurrencyselector/config.xml`
- `stcurrencyselector/stcurrencyselector.php`

> [!note] Full source path
> `/dependencies/modules/stcurrencyselector/` in the Panda theme archive.
