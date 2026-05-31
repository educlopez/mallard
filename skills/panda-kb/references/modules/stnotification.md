---
module_name: stnotification
display_name: Notification
version: 1.1
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: content
hooks: [displayBanner, displayBeforeBodyClosingTag, displayHeader]
hooks_count: 3
is_configurable: true
files_count: 29
loc: 2140
author: SUNNYTOO.COM
source_path: dependencies/modules/stnotification
tags: [panda, module, content]
---

# Notification

> [!info] Panda module
> **Category:** content · **Version:** 1.1 · **Hooks:** 3 · **LOC:** 2,140 across 29 files.

## Purpose

Display a responsive and custom warning notification on your site.

## Hooks registered

Total: **3** hooks.

- `displayBanner`
- `displayBeforeBodyClosingTag`
- `displayHeader`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stnotification.php`.

## Options (from translations)

Translations folder present at `stnotification/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stnotification.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stnotification/config.xml`
- `stnotification/stnotification.php`
- Public docs: [[docs/module-notification]]

> [!note] Full source path
> `/dependencies/modules/stnotification/` in the Panda theme archive.
