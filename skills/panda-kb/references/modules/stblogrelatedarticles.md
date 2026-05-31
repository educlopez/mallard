---
module_name: stblogrelatedarticles
display_name: Blog Module - Related articles
version: 1.0
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: blog
hooks: [actionAdminStBlogFormModifier, actionObjectStBlogClassAddAfter, actionObjectStBlogClassDeleteAfter, actionObjectStBlogClassUpdateAfter, displayStBlogArticleFooter]
hooks_count: 5
is_configurable: true
files_count: 21
loc: 1527
author: SUNNYTOO.COM
source_path: dependencies/modules/stblogrelatedarticles
tags: [panda, module, blog]
---

# Blog Module - Related articles

> [!info] Panda module
> **Category:** blog · **Version:** 1.0 · **Hooks:** 5 · **LOC:** 1,527 across 21 files.

## Purpose

Add related articles on blog artice pages.

## Hooks registered

Total: **5** hooks.

- `actionAdminStBlogFormModifier`
- `actionObjectStBlogClassAddAfter`
- `actionObjectStBlogClassDeleteAfter`
- `actionObjectStBlogClassUpdateAfter`
- `displayStBlogArticleFooter`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stblogrelatedarticles.php`.

## Options (from translations)

Translations folder present at `stblogrelatedarticles/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

- **Requires `stblog`** — this is a blog sub-module that depends on the core blog engine.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stblogrelatedarticles.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stblogrelatedarticles/config.xml`
- `stblogrelatedarticles/stblogrelatedarticles.php`

> [!note] Full source path
> `/dependencies/modules/stblogrelatedarticles/` in the Panda theme archive.
