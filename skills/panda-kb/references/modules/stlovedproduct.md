---
module_name: stlovedproduct
display_name: Love buttons
version: 1.0.3
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: product
hooks: [actionAuthentication, actionProductDelete, actionProductSearchAfter, actionStAssemble, displayBeforeBodyClosingTag, displayCustomerAccount, displayHeader, displayProductCenterColumn, displaySideBar, displayStBlogArticleInfo]
hooks_count: 10
is_configurable: true
files_count: 31
loc: 2821
author: SUNNYTOO.COM
source_path: dependencies/modules/stlovedproduct
tags: [panda, module, product]
---

# Love buttons

> [!info] Panda module
> **Category:** product · **Version:** 1.0.3 · **Hooks:** 10 · **LOC:** 2,821 across 31 files.

## Purpose

Display love buttons on your store.

## Hooks registered

Total: **10** hooks.

- `actionAuthentication`
- `actionProductDelete`
- `actionProductSearchAfter`
- `actionStAssemble`
- `displayBeforeBodyClosingTag`
- `displayCustomerAccount`
- `displayHeader`
- `displayProductCenterColumn`
- `displaySideBar`
- `displayStBlogArticleInfo`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stlovedproduct.php`.

## Options (from translations)

Translations folder present at `stlovedproduct/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stlovedproduct.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stlovedproduct/config.xml`
- `stlovedproduct/stlovedproduct.php`
- Public docs: [[docs/module-love-buttons]]

> [!note] Full source path
> `/dependencies/modules/stlovedproduct/` in the Panda theme archive.
