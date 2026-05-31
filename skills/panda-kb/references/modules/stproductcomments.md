---
module_name: stproductcomments
display_name: Product Comments
version: 1.0.5
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: product
hooks: [actionDeleteGDPRCustomer, actionExportGDPRData, actionProductDelete, actionStAssemble, dashboardZoneOne, displayAdminOrder, displayCustomerAccount, displayHeader, displayProductExtraContent, displayUnderProductName, registerGDPRConsent, vcBeforeInit]
hooks_count: 12
is_configurable: true
files_count: 103
loc: 11737
author: SUNNYTOO.COM
source_path: dependencies/modules/stproductcomments
tags: [panda, module, product]
---

# Product Comments

> [!info] Panda module
> **Category:** product · **Version:** 1.0.5 · **Hooks:** 12 · **LOC:** 11,737 across 103 files.

## Purpose

Allows users to post comments for products after placed orders.

## Hooks registered

Total: **12** hooks.

- `actionDeleteGDPRCustomer`
- `actionExportGDPRData`
- `actionProductDelete`
- `actionStAssemble`
- `dashboardZoneOne`
- `displayAdminOrder`
- `displayCustomerAccount`
- `displayHeader`
- `displayProductExtraContent`
- `displayUnderProductName`
- `registerGDPRConsent`
- `vcBeforeInit`

## Admin configuration

Admin template(s) present: `list_action_noabuse.tpl`, `list_action_reply.tpl`, `list_action_approve.tpl`.

See source at `stproductcomments/views/templates/admin/` for the full form layout.

## Options (from translations)

Translations folder present at `stproductcomments/translations/`. Admin labels are translated via `$this->l('...')`. See the module's `.tpl` files and PHP for specific fields.

## Dependencies

No declared dependencies on other `st*` modules in `config.xml`. May still use shared Panda JS/CSS assets loaded via `stthemeeditor`.

## PS 8 vs PS 9 notes

`ps_versions_compliancy`: **1.7.1.0** to **current (_PS_VERSION_)**. No version-specific branching observed in the header of `stproductcomments.php`. Compatible with PS 8 and PS 9 per theme's `theme.yml`.

## Known limits / caveats

- Very large module (11,737 LOC) — likely bundles vendor assets (jQuery plugins, WYSIWYG, etc.). Heavy on the asset pipeline.

## Source references

- `stproductcomments/config.xml`
- `stproductcomments/stproductcomments.php`
- `stproductcomments/views/templates/admin/` (3 templates)

> [!note] Full source path
> `/dependencies/modules/stproductcomments/` in the Panda theme archive.
