---
module_name: stblogcomments
display_name: Blog Module - Comments
version: 1.1
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: blog
hooks: [actionDeleteGDPRCustomer, actionExportGDPRData, displayCustomerAccount, displayHeader, displayStBlogArticleSecondary, displayStBlogLeftColumn, displayStBlogRightColumn, registerGDPRConsent]
hooks_count: 8
is_configurable: true
files_count: 30
loc: 2450
author: SUNNYTOO.COM
source_path: dependencies/modules/stblogcomments
tags: [panda, module, blog]
---

# Blog Module - Comments

> [!info] Panda module
> **Category:** blog · **Version:** 1.1 · **Hooks:** 8 · **LOC:** 2,450 across 30 files.

## Purpose

Allows users to post comments.

## Hooks registered

Total: **8** hooks.

- `actionDeleteGDPRCustomer`
- `actionExportGDPRData`
- `displayCustomerAccount`
- `displayHeader`
- `displayStBlogArticleSecondary`
- `displayStBlogLeftColumn`
- `displayStBlogRightColumn`
- `registerGDPRConsent`

## Admin configuration

`<is_configurable>1</is_configurable>` in config.xml — configuration is likely handled via the module's `getContent()` method. Check `stblogcomments.php`.

## Options (from translations)

Translations folder present at `stblogcomments/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

- **Requires `stblog`** — this is a blog sub-module that depends on the core blog engine.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stblogcomments.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stblogcomments/config.xml`
- `stblogcomments/stblogcomments.php`

> [!note] Full source path
> `/dependencies/modules/stblogcomments/` in the Panda theme archive.
