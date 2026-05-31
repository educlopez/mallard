---
module_name: stsidebar
display_name: Sidebar and mobile header bar
version: 1.1.0
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: navigation
hooks: [displayHeader, displayMobileBar, displayMobileBarBottom, displayMobileBarCenter, displayMobileBarLeft, displayRightBar, displaySideBar]
hooks_count: 7
is_configurable: true
files_count: 23
loc: 2619
author: SUNNYTOO.COM
source_path: dependencies/modules/stsidebar
tags: [panda, module, navigation]
---

# Sidebar and mobile header bar

> [!info] Panda module
> **Category:** navigation · **Version:** 1.1.0 · **Hooks:** 7 · **LOC:** 2,619 across 23 files.

## Purpose

Manage items in the sidebar.

## Hooks registered

Total: **7** hooks.

- `displayHeader`
- `displayMobileBar`
- `displayMobileBarBottom`
- `displayMobileBarCenter`
- `displayMobileBarLeft`
- `displayRightBar`
- `displaySideBar`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stsidebar.php`.

## Options (from translations)

Translations folder present at `stsidebar/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stsidebar.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stsidebar/config.xml`
- `stsidebar/stsidebar.php`

> [!note] Full source path
> `/dependencies/modules/stsidebar/` in the Panda theme archive.
