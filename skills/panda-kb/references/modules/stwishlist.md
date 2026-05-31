---
module_name: stwishlist
display_name: Wishlist block
version: 1.0.1
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: product
hooks: [actionProductSearchAfter, actionStAssemble, displayBeforeBodyClosingTag, displayCustomerAccount, displayHeader, displayProductCenterColumn, displaySideBar]
hooks_count: 7
is_configurable: true
files_count: 42
loc: 3464
author: SUNNYTOO.COM
source_path: dependencies/modules/stwishlist
tags: [panda, module, product]
---

# Wishlist block

> [!info] Panda module
> **Category:** product · **Version:** 1.0.1 · **Hooks:** 7 · **LOC:** 3,464 across 42 files.

## Purpose

Dispaly wishlist buttons on your store.

## Hooks registered

Total: **7** hooks.

- `actionProductSearchAfter`
- `actionStAssemble`
- `displayBeforeBodyClosingTag`
- `displayCustomerAccount`
- `displayHeader`
- `displayProductCenterColumn`
- `displaySideBar`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stwishlist.php`.

## Options (from translations)

Translations folder present at `stwishlist/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stwishlist.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stwishlist/config.xml`
- `stwishlist/stwishlist.php`
- Public docs: [[docs/module-love-buttons]]

> [!note] Full source path
> `/dependencies/modules/stwishlist/` in the Panda theme archive.
