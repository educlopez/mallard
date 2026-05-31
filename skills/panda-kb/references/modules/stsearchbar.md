---
module_name: stsearchbar
display_name: Search bar mod
version: 1.6.7
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: navigation
hooks: [displayFullWidthBottom, displayheader, displayHeaderCenter, displayMobileNav, displaySideBar, displayStBlogFullWidthBottom, filterProductSearch]
hooks_count: 7
is_configurable: true
files_count: 35
loc: 2747
author: SUNNYTOO.COM
source_path: dependencies/modules/stsearchbar
tags: [panda, module, navigation]
---

# Search bar mod

> [!info] Panda module
> **Category:** navigation · **Version:** 1.6.7 · **Hooks:** 7 · **LOC:** 2,747 across 35 files.

## Purpose

Adds a quick search field to your website.

## Hooks registered

Total: **7** hooks.

- `displayFullWidthBottom`
- `displayheader`
- `displayHeaderCenter`
- `displayMobileNav`
- `displaySideBar`
- `displayStBlogFullWidthBottom`
- `filterProductSearch`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stsearchbar.php`.

## Options (from translations)

Translations folder present at `stsearchbar/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stsearchbar.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stsearchbar/config.xml`
- `stsearchbar/stsearchbar.php`

> [!note] Full source path
> `/dependencies/modules/stsearchbar/` in the Panda theme archive.
