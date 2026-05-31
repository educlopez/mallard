---
module_name: stblogtags
display_name: Blog Module - Tags block
version: 1.0
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: blog
hooks: [displayStBlogLeftColumn, displayStBlogRightColumn]
hooks_count: 2
is_configurable: true
files_count: 10
loc: 433
author: SUNNYTOO.COM
source_path: dependencies/modules/stblogtags
tags: [panda, module, blog]
---

# Blog Module - Tags block

> [!info] Panda module
> **Category:** blog · **Version:** 1.0 · **Hooks:** 2 · **LOC:** 433 across 10 files.

## Purpose

Adds a block containing blog tags.

## Hooks registered

Total: **2** hooks.

- `displayStBlogLeftColumn`
- `displayStBlogRightColumn`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stblogtags.php`.

## Options (from translations)

Translations folder present at `stblogtags/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

- **Requires `stblog`** — this is a blog sub-module that depends on the core blog engine.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stblogtags.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stblogtags/config.xml`
- `stblogtags/stblogtags.php`

> [!note] Full source path
> `/dependencies/modules/stblogtags/` in the Panda theme archive.
