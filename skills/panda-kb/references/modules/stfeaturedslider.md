---
module_name: stfeaturedslider
display_name: Featured products slider
version: 1.1.8
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: product
hooks: [displayFullWidthTop]
hooks_count: 1
is_configurable: true
files_count: 17
loc: 674
author: SUNNYTOO.COM
source_path: dependencies/modules/stfeaturedslider
tags: [panda, module, product]
---

# Featured products slider

> [!info] Panda module
> **Category:** product · **Version:** 1.1.8 · **Hooks:** 1 · **LOC:** 674 across 17 files.

## Purpose

Displays featured products on homepage.

## Hooks registered

Total: **1** hook.

- `displayFullWidthTop`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stfeaturedslider.php`.

## Options (from translations)

Translations folder present at `stfeaturedslider/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stfeaturedslider.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stfeaturedslider/config.xml`
- `stfeaturedslider/stfeaturedslider.php`

> [!note] Full source path
> `/dependencies/modules/stfeaturedslider/` in the Panda theme archive.
