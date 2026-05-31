---
module_name: stinstagram
display_name: Advanced instagram feed
version: 1.1.3
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: content
hooks: [actionOutputHTMLBefore, displayHeader]
hooks_count: 2
is_configurable: true
files_count: 76
loc: 6486
author: SUNNYTOO.COM
source_path: dependencies/modules/stinstagram
tags: [panda, module, content]
---

# Advanced instagram feed

> [!info] Panda module
> **Category:** content · **Version:** 1.1.3 · **Hooks:** 2 · **LOC:** 6,486 across 76 files.

## Purpose

Display your instagram media to your web page.

## Hooks registered

Total: **2** hooks.

- `actionOutputHTMLBefore`
- `displayHeader`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stinstagram.php`.

## Options (from translations)

Translations folder present at `stinstagram/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stinstagram.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stinstagram/config.xml`
- `stinstagram/stinstagram.php`
- Public docs: [[docs/module-advanced-social-feed]]

> [!note] Full source path
> `/dependencies/modules/stinstagram/` in the Panda theme archive.
