---
module_name: stblogeditor
display_name: Blog editor
version: 1.0.0
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: blog
hooks: [displayStBlogArticleFooter]
hooks_count: 1
is_configurable: true
files_count: 33
loc: 2083
author: SUNNYTOO.COM
source_path: dependencies/modules/stblogeditor
tags: [panda, module, blog]
---

# Blog editor

> [!info] Panda module
> **Category:** blog · **Version:** 1.0.0 · **Hooks:** 1 · **LOC:** 2,083 across 33 files.

## Purpose

Change blog settings.

## Hooks registered

Total: **1** hook.

- `displayStBlogArticleFooter`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stblogeditor.php`.

## Options (from translations)

Translations folder present at `stblogeditor/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

- **Requires `stblog`** — this is a blog sub-module that depends on the core blog engine.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stblogeditor.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stblogeditor/config.xml`
- `stblogeditor/stblogeditor.php`

> [!note] Full source path
> `/dependencies/modules/stblogeditor/` in the Panda theme archive.
