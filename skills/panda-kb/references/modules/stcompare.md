---
module_name: stcompare
display_name: Product Comparison
version: 1.0.1
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: product
hooks: [actionProductSearchAfter, actionStAssemble, displayHeader, displayNav2, displayProductCenterColumn]
hooks_count: 5
is_configurable: true
files_count: 20
loc: 1537
author: SUNNYTOO.COM
source_path: dependencies/modules/stcompare
tags: [panda, module, product]
---

# Product Comparison

> [!info] Panda module
> **Category:** product · **Version:** 1.0.1 · **Hooks:** 5 · **LOC:** 1,537 across 20 files.

## Purpose

Adds a product comparison feature to your PrestaShop 1.7 site.

## Hooks registered

Total: **5** hooks.

- `actionProductSearchAfter`
- `actionStAssemble`
- `displayHeader`
- `displayNav2`
- `displayProductCenterColumn`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stcompare.php`.

## Options (from translations)

Translations folder present at `stcompare/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stcompare.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stcompare/config.xml`
- `stcompare/stcompare.php`

> [!note] Full source path
> `/dependencies/modules/stcompare/` in the Panda theme archive.
