---
module_name: stblogblockcategory
display_name: Blog Module - Category block
version: 1.2.0
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: blog
hooks: [displayStBlogLeftColumn]
hooks_count: 1
is_configurable: true
files_count: 11
loc: 510
author: SUNNYTOO.COM
source_path: dependencies/modules/stblogblockcategory
tags: [panda, module, blog]
---

# Blog Module - Category block

> [!info] Panda module
> **Category:** blog · **Version:** 1.2.0 · **Hooks:** 1 · **LOC:** 510 across 11 files.

## Purpose

Adds a block featuring blog categories.

## Hooks registered

Total: **1** hook.

- `displayStBlogLeftColumn`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stblogblockcategory.php`.

## Options (from translations)

Translations folder present at `stblogblockcategory/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

- **Requires `stblog`** — this is a blog sub-module that depends on the core blog engine.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stblogblockcategory.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stblogblockcategory/config.xml`
- `stblogblockcategory/stblogblockcategory.php`

> [!note] Full source path
> `/dependencies/modules/stblogblockcategory/` in the Panda theme archive.
