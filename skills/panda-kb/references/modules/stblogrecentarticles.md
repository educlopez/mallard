---
module_name: stblogrecentarticles
display_name: Blog Module - Recent articles
version: 1.0.2
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: blog
hooks: [displayStBlogHome]
hooks_count: 1
is_configurable: true
files_count: 16
loc: 547
author: SUNNYTOO.COM
source_path: dependencies/modules/stblogrecentarticles
tags: [panda, module, blog]
---

# Blog Module - Recent articles

> [!info] Panda module
> **Category:** blog · **Version:** 1.0.2 · **Hooks:** 1 · **LOC:** 547 across 16 files.

## Purpose

Display rencent articles on your store.

## Hooks registered

Total: **1** hook.

- `displayStBlogHome`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stblogrecentarticles.php`.

## Options (from translations)

Translations folder present at `stblogrecentarticles/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

- **Requires `stblog`** — this is a blog sub-module that depends on the core blog engine.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stblogrecentarticles.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stblogrecentarticles/config.xml`
- `stblogrecentarticles/stblogrecentarticles.php`

> [!note] Full source path
> `/dependencies/modules/stblogrecentarticles/` in the Panda theme archive.
