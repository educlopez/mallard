---
module_name: stcustomersignin
display_name: Customer &quot;Sign in&quot; link mod
version: 0.3.5
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: helpers
hooks: [displayCheckoutHeader, displayCheckoutMobileNav, displayHeader, displayMobileNav, displayNav2]
hooks_count: 5
is_configurable: true
files_count: 18
loc: 1050
author: SUNNYTOO.COM
source_path: dependencies/modules/stcustomersignin
tags: [panda, module, helpers]
---

# Customer &quot;Sign in&quot; link mod

> [!info] Panda module
> **Category:** helpers · **Version:** 0.3.5 · **Hooks:** 5 · **LOC:** 1,050 across 18 files.

## Purpose

Adds a block that displays information about the customer.

## Hooks registered

Total: **5** hooks.

- `displayCheckoutHeader`
- `displayCheckoutMobileNav`
- `displayHeader`
- `displayMobileNav`
- `displayNav2`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stcustomersignin.php`.

## Options (from translations)

Translations folder present at `stcustomersignin/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stcustomersignin.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stcustomersignin/config.xml`
- `stcustomersignin/stcustomersignin.php`
- Public docs: [[docs/module-customer-sign-in]]

> [!note] Full source path
> `/dependencies/modules/stcustomersignin/` in the Panda theme archive.
