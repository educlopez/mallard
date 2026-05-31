---
module_name: stblogfeaturedarticles
display_name: Blog Module - Article Slider
version: 1.0.2
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: blog
hooks: [vcBeforeInit]
hooks_count: 1
is_configurable: true
files_count: 21
loc: 2085
author: SUNNYTOO.COM
source_path: dependencies/modules/stblogfeaturedarticles
tags: [panda, module, blog]
---

# Blog Module - Article Slider

> [!info] Panda module
> **Category:** blog · **Version:** 1.0.2 · **Hooks:** 1 · **LOC:** 2,085 across 21 files.

## Purpose

Display articles from different categories.

## Hooks registered

Total: **1** hook.

- `vcBeforeInit`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stblogfeaturedarticles.php`.

## Options (from translations)

Translations folder present at `stblogfeaturedarticles/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

- **Requires `stblog`** — this is a blog sub-module that depends on the core blog engine.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stblogfeaturedarticles.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stblogfeaturedarticles/config.xml`
- `stblogfeaturedarticles/stblogfeaturedarticles.php`

> [!note] Full source path
> `/dependencies/modules/stblogfeaturedarticles/` in the Panda theme archive.
