---
module_name: strelatedproducts
display_name: Related products
version: 1.1.0
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: product
hooks: [displayProductRightColumn]
hooks_count: 1
is_configurable: true
files_count: 16
loc: 814
author: SUNNYTOO.COM
source_path: dependencies/modules/strelatedproducts
tags: [panda, module, product]
---

# Related products

> [!info] Panda module
> **Category:** product · **Version:** 1.1.0 · **Hooks:** 1 · **LOC:** 814 across 16 files.

## Purpose

Add related products on product pages.

## Hooks registered

Total: **1** hook.

- `displayProductRightColumn`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `strelatedproducts.php`.

## Options (from translations)

Translations folder present at `strelatedproducts/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `strelatedproducts.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `strelatedproducts/config.xml`
- `strelatedproducts/strelatedproducts.php`

> [!note] Full source path
> `/dependencies/modules/strelatedproducts/` in the Panda theme archive.
