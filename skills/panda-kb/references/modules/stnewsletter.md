---
module_name: stnewsletter
display_name: Newsletter popup
version: 1.4.6
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: content
hooks: [actionCustomerAccountAdd, actionCustomerAccountUpdate, actionDeleteGDPRCustomer, actionExportGDPRData, additionalCustomerFormFields, displayAdminCustomersForm, displayHeader, displaySideBar, registerGDPRConsent, vcBeforeInit]
hooks_count: 10
is_configurable: true
files_count: 57
loc: 4602
author: SUNNYTOO.COM
source_path: dependencies/modules/stnewsletter
tags: [panda, module, content]
---

# Newsletter popup

> [!info] Panda module
> **Category:** content · **Version:** 1.4.6 · **Hooks:** 10 · **LOC:** 4,602 across 57 files.

## Purpose

Adds a block for newsletter subscription.

## Hooks registered

Total: **10** hooks.

- `actionCustomerAccountAdd`
- `actionCustomerAccountUpdate`
- `actionDeleteGDPRCustomer`
- `actionExportGDPRData`
- `additionalCustomerFormFields`
- `displayAdminCustomersForm`
- `displayHeader`
- `displaySideBar`
- `registerGDPRConsent`
- `vcBeforeInit`

## Admin configuration

Admin template(s) present: `newsletter_subscribe.tpl`.

See source at `stnewsletter/views/templates/admin/` for the full form layout.

## Options (from translations)

Translations folder present at `stnewsletter/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stnewsletter.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- No obvious issues in static inspection. Validate in a live PS 8 and PS 9 install before production use.

## Source references

- `stnewsletter/config.xml`
- `stnewsletter/stnewsletter.php`
- `stnewsletter/views/templates/admin/` (1 template)

> [!note] Full source path
> `/dependencies/modules/stnewsletter/` in the Panda theme archive.
