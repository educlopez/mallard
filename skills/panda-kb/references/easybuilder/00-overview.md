---
product_name: Easy Builder (steasybuilder)
vendor: SunnyToo (sunnytoo.com)
pricing: paid optional addon
version: 1.2.3
config_version: 1.2.1
ps_version: [1.7, 8, 9]
ps_versions_compliancy:
  min: "1.7"
  max: "_PS_VERSION_"
tab: front_office_features
custom_hooks_registered:
  - actionEasyBuilderInit
  - filterSupplierContent
  - actionEasyBuilderPreviewUrl
  - actionOtherEasyBuilderType
  - actionOtherAddEasyBuilderData
  - actionOtherDuplicateEasyBuilderData
  - actionOtherDeleteEasyBuilderData
ps_hooks_registered:
  - displayHeader
  - displayHome
  - actionMaintenancePageSave
  - actionMaintenancePageForm
  - displayFullWidthTop2
  - displayNavFullWidth
  - displayFooterBefore
  - actionAdminStEasyBuilderListingResultsModifier
  - displayBackOfficeHeader
  - actionAdminImagesControllerSaveAfter
  - actionAdminImagesControllerDeleteAfter
  - actionObjectLanguageDeleteAfter
  - actionFrontControllerAfterInit
  - overrideLayoutTemplate
  - dashboardZoneOne
dependencies:
  required: []
  optional: [stblog, stthemeeditor]
  incompatible: [jscomposer, creativeelements, crazyelements]
source_path: <source-archive>/
tags: [easybuilder, panda, addon, paid, elementor, page-builder]
---

# Easy Builder (steasybuilder) — Overview

> [!info] Separate track from Panda KB
> Easy Builder is a **paid optional addon** sold by SunnyToo on top of the Panda theme. It is **not always installed** — only clients who pay for it have it enabled. The main Panda KB (`../`) does not depend on this folder; references are made only via wikilinks when relevant. For the bridge module that feeds PrestaShop data into Easy Builder widgets, see [[../easybuilder-bridge/]] (task #8, not yet processed).

## What it is

Paid optional page builder for PrestaShop, technically a **fork/port of the WordPress Elementor plugin** adapted to run inside PrestaShop. Lets clients edit layouts visually (drag-and-drop) without touching code or the stock PrestaShop backoffice forms. The editor authors content as Elementor JSON (`_elementor_data` post meta) and renders on the front via Elementor's frontend engine.

Evidence of the WordPress origin:

- `elementor/wp-includes/` — ported WordPress globals and functions.
- `elementor/includes/` mirrors the Elementor plugin layout (`api.php`, `plugin.php`, `widgets/`, `elements/`, `controls/`, `managers/`, `template-library/`).
- `elementor/core/` mirrors Elementor core (`app/`, `base/`, `editor/`, `frontend/`, `kits/`, `schemes/`, `documents-manager.php`, `modules-manager.php`).
- Widget icons (`eicon-*`) and keyword set (e.g. `tabs`, `toggle`, `accordion`, `heading`, `image-carousel`) are Elementor-native.
- Post-meta keys use WordPress conventions: `_elementor_data`, `_elementor_version`, `_elementor_edit_mode`, `_elementor_template_type`, `_elementor_page_settings`.

> [!warning] Legal / technical caveat
> The code carries Elementor copyrights inside `elementor/`. Any upstream work should stay aware that this is derivative code, not a SunnyToo invention.

## When it is installed

Client-dependent. Some clients pay the extra license, some don't. **Panda runs fine without Easy Builder** — the stock `steasycontent` module (bundled with Panda) already covers basic block/HTML editing. Easy Builder is the upgrade for clients who want true visual editing.

Incompatible with other PS builders; install blocks if any of these are active: `jscomposer`, `creativeelements`, `crazyelements`.

## Architecture overview

```
steasybuilder/
├── steasybuilder.php              # main module (95 KB) — hooks, builder types, install
├── config.xml                     # manifest
├── classes/
│   ├── EasyBuilderBase.php        # EBB / core helper (referenced as EBB::getInstance())
│   ├── EasyBuilderTranslation.php # EbTrans::l() translator
│   ├── StModuleLicense.php        # license / purchase-code validation
│   ├── DB/StEasyBuilderClass.php  # ObjectModel for st_easy_builder table
│   └── data/
│       ├── provider/              # search/query providers (product, cms, brand, category)
│       └── translation/
├── controllers/
│   ├── admin/                     # 4 BO controllers (see 03-admin-controllers)
│   └── front/
│       ├── editor.php             # editor bootstrap
│       └── maintenance.php        # maintenance page renderer
├── elementor/                     # THE PORTED ELEMENTOR ENGINE
│   ├── core/                      # 18 subsystems (documents-manager, editor, kits, schemes, etc.)
│   ├── includes/
│   │   ├── widgets/               # 31 widget PHP files (Elementor-standard)
│   │   ├── elements/              # column.php, section.php, repeater.php
│   │   ├── controls/              # control types
│   │   ├── managers/              # widgets/elements/controls managers
│   │   └── template-library/
│   ├── modules/                   # 8 Elementor modules (wp-cli, library, history, page-templates, usage, gutenberg, dynamic-tags)
│   ├── wp-includes/               # ported WP functions
│   ├── assets/  js/  data/  bin/
├── views/                         # .tpl templates (block, header, bo_header, bo_header1760)
├── translations/                  # incl. ModulesEasybuilderShop.xlf (PS 1.7+ Symfony)
├── import_data/                   # bootstrap data (see 02-presets-catalog)
├── upgrade/                       # install-1.0.4.php, install-1.0.5.php, install-1.1.0.php
└── other-setting/
```

### Data model

Core table: `ps_st_easy_builder` (declared in `classes/DB/StEasyBuilderClass.php`). Key columns: `id_st_easy_builder` (PK), `type`, `location`, `hide_on_mobile`, `active`, `position`, `id_category`, `id_product` (csv), `id_manufacturer`, `id_cms`, `id_st_blog`, `full_screen`, `display_on`, `contr`, `post_status`, `post_parent`, `post_type`, `post_title`, `post_content`, `last_edited`, `post_name`, `default_lang`. Multi-language via companion table `ps_st_easy_builder_yuyan` (note the romanized-Chinese suffix `yuyan` = "language"). Multi-shop via standard `_shop` companion.

The `post_*` field names and `_yuyan` suffix are direct WordPress legacy — confirms the Elementor fork origin.

## Integration with Panda

Easy Builder plugs into the theme via a large matrix of PS hooks exposed by Panda. Each "location type" maps to one theme hook:

- **Theme-wide slots (Panda)**: `FullWidthTop2`, `FullWidthTop`, `NavFullWidth`, `FooterBefore`, `TopColumn`, `HeaderLeft`, `HeaderCenter`, `Top` (header right), homepage quarter slots (`HomeFirstQuarter`..`HomeFourthQuarter`), `HomeLeft`, `HomeRight`, `HomeSecondaryLeft/Right`.
- **Product page (Panda product hooks)**: `ProductRightColumn`, `ProductLeftColumn`, `ProductCenterColumn`, `ProductExtraContent`, `ProductDescRightColumn`, `FooterProduct`, `MiddleProduct`.
- **Category / Manufacturer / CMS**: `CategoryHeader`, `CategoryFooter`, `FullWidthCategoryHeader`, `FullWidthCategoryFooter`, `ManufacturerHeader`, `ManufacturerFooter`, `CMSExtra`.
- **Blog (optional)**: `StBlogArticleTop`, `StBlogContent`, `StBlogHome`, `StBlogHomeTop`, `StBlogHomeBottom`, `StBlogFullWidthTop`, `StBlogFullWidthBottom`, `StBlogLeftColumn`, `StBlogRightColumn` — only available if `stblog` module is installed and active.
- **Utility**: `NotFound` (404), `Maintenance`, `ShoppingCartFooter`, `BelowContactForm`.

These hooks are declared by Panda's theme + child modules (stmegamenu, stproductpage, stcustompages...). Easy Builder emits the rendered layout into them via custom module hook handlers (`hookDisplayFullWidthTop2`, `hookDisplayNavFullWidth`, `hookDisplayFooterBefore`, and the catch-all listener on `actionFrontControllerAfterInit` / `overrideLayoutTemplate`).

### Relation to `steasycontent` (bundled in Panda)

`steasycontent` is the **free** content block manager bundled with Panda. Easy Builder is a **superset** — same general idea (inject layouts into hooks) but with a visual Elementor editor instead of TinyMCE / stock PS forms. Both can coexist; most clients use `steasycontent` for text blocks and Easy Builder only for richer page layouts.

### Relation to `steasy_trans_panda` (bridge, task #8)

`steasy_trans_panda` is a **separate bridge module** (not processed in this task) that exposes PS catalog objects (products, categories, manufacturers) as Elementor "dynamic tags". Without it, Easy Builder widgets only render static content; with it, widgets can pull live product data. Deferred to task #8.

### Relation to `stthemeeditor`

If `stthemeeditor/classes/easy.php` exists, Easy Builder auto-skips the License tab in the BO menu and bypasses the purchase-code activation check — the theme-editor package acts as a bundled license. See `installTab()` and `AdminStEasyBuilderEditorController::initContent()`.

## Hook emission on the front

On install, Easy Builder registers **7 custom hooks** (listed in frontmatter, prefixed `actionEasyBuilder*` / `actionOtherEasyBuilder*`). These let other modules (e.g. `steasy_trans_panda`) inject dynamic tag providers, preview URLs, or extra builder types at runtime. The hook `actionOtherAddEasyBuilderData` (added in upgrade 1.1.0) is the extension point third-party modules use to push data into builder pages.

## Bundled presets

Ships with three bootstrap files under `import_data/`:

- `base_easy_builder.sql` (953 B) — schema/seed SQL, loaded on install via `installDefaultData()`.
- `easy_builder_postmeta.json` (9.5 KB) — initial post-meta records (Elementor data for demo pages).
- `settings.json` (77 KB) — initial `styleGlobals` (system colors: Primary, Secondary, Text, Accent; system typography: Vollkorn; kit schemes).

Catalog in [[02-presets-catalog]].

## Licensing / upgrade path

- **Licensing**: Purchase-code activation managed by `classes/StModuleLicense.php`. Validated domain-locked; users activate via the License BO screen by submitting their code, which calls back to the SunnyToo server. Invalid/expired licenses disable editor features with a "not activated" banner (see `AdminStEasyBuilderEditorController::initContent()` line 64-65). License check is skipped automatically if `stthemeeditor` is present.
- **Upgrade scripts** (`upgrade/`):
  - `install-1.0.4.php`
  - `install-1.0.5.php`
  - `install-1.1.0.php` — registers `actionFrontControllerAfterInit` + `actionOtherAddEasyBuilderData` (bridge extension point).
- **Auto-update check**: `AdminStEasyBuilderLinceseController::postProcess()` pings remote for new versions ("This module checks update every day automatically.").

## Related KB

- [[01-widgets-catalog]] — 31 Elementor widgets table.
- [[02-presets-catalog]] — bundled import_data contents.
- [[03-admin-controllers]] — 4 BO controllers.
- [[README]] — subfolder index.
- [[../00-overview|Panda KB root]] — main Panda Knowledge Base (Easy Builder is a separate track).
