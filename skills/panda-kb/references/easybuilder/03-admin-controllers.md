---
parent: "[[00-overview]]"
controller_count: 4
source_dir: <source-archive>/
tags: [easybuilder, admin, backoffice, controllers]
---

# Easy Builder — Admin Controllers

> [!info] BO menu shape
> Easy Builder registers 2 parent tabs + 4 child tabs under PS "IMPROVE" menu. The parents are shell `AdminParent*` classes that just render a container; real work happens in the children. The menu tree (from `$_bo_menu` in `steasybuilder.php`):
> 
> ```
> IMPROVE
> └── AdminParentStEasyBuilder               "Easy builder"              (parent, icon: explicit)
>     ├── AdminStEasyBuilder                 "Easy builder manager"      (visible)
>     ├── AdminStEasyBuilderEditor           "Easy page editor"          (hidden: active=0)
>     └── AdminParentStEasyBuilderGeneralSetting   "Easy builder settings"   (parent)
>         ├── AdminStEasyBuilderGeneralSetting    "Settings"
>         └── AdminStEasyBuilderLincese           "License"               (auto-hidden if stthemeeditor installed)
> ```

## Controllers

| File | Size | Class | BO screen | Main CRUD / actions |
| --- | --- | --- | --- | --- |
| `AdminStEasyBuilderController.php` | 48 KB | `AdminStEasyBuilderController extends ModuleAdminController` | **Easy builder manager** — list of all Easy Builder pages (ObjectModel-backed, table `ps_st_easy_builder`). | list / add / edit / **delete** / **duplicate** / **easybuilderpreview** (custom row action). Columns shown in `fields_list`: ID, Type, Name / Show on, Status, Position, Edit with. Supports PS multi-shop (`Shop::addTableAssociation`). Provides helpers for per-object editing (product, category, manufacturer, cms, brand) with direct links into the editor. |
| `AdminStEasyBuilderEditorController.php` | 23 KB | `AdminStEasyBuilderEditorController extends ModuleAdminController` | **Easy page editor** — loads the Elementor editor IFRAME/SPA for a specific page. Gated by license: if not activated and `stthemeeditor` not present, shows "not activated" panel. | `initContent()`, plus AJAX endpoints: `processElementorJsLog`, `ajaxProcessHeartbeat`, `ajaxProcessStyleGlobals`, `processElementorLibraryDirectActions`, `ajaxProcessElementorAjax` (catch-all Elementor ajax), `ajaxProcessPsAc` (PrestaShop autocomplete for product/cms/brand/category pickers), `ajaxProcessSwitcherAc`, `ajaxProcessPreviewAc`. This is the bridge between Elementor's JS client and the PS backend. |
| `AdminStEasyBuilderGeneralSettingController.php` | 34 KB | `AdminStEasyBuilderGeneralSettingController extends ModuleAdminController` | **Settings** — global module configuration form. | `fields_options` with entries prefixed `ST_EASY_B_`: `PS_PRODUCT_BTN`, `PS_CATEGORY_BTN`, `PS_CMS_BTN`, `PS_MANUFACTURER_BTN`, `PS_ST_BLOG_BTN` (toggles for the "Edit with Easy Builder" button on each PS object type); `PS_MODULE`, `PS_HOOKS`, `LOAD_SWIPER`, `USE_LEGACY_TRANS`, `CUSTOM_HEADER_MODULES`. Also conditionally exposes `tp_tongbu_settings` + `tp_tongbu_settings_search_bar` sections (likely theme-editor sync settings, `tongbu` = Chinese for "sync"). |
| `AdminStEasyBuilderLinceseController.php` | 9 KB | `AdminStEasyBuilderLinceseController extends ModuleAdminController` | **License** — purchase-code activation screen. Auto-hidden when `stthemeeditor/classes/easy.php` is present. | `postProcess()` handles: `act=check_update` (calls `license->checkUpdate()` against remote), `act=dereg_purchase_code` (deactivate), plus registration via `license->registerLicense()` (implicit). Renders domain/store-scoped activation status (active / invalid / not-activated) with masked purchase code. Built on `classes/StModuleLicense.php` (constants: `IN_TEST_MODE`, `ERROR_IN_ALL_STORE`). Note the filename typo: "Lincese" (not "License") is intentional and used everywhere — DO NOT fix casually. |

## Front controllers (for context)

| File | Role |
| --- | --- |
| `controllers/front/editor.php` | Trivial bootstrap (`ModuleFrontController`) for front-side editor invocation — most editor work is BO-side via `AdminStEasyBuilderEditorController`. |
| `controllers/front/maintenance.php` | Renders the maintenance page when `type=30` (Maintenance) Easy Builder pages are active; hooks into PS maintenance mode via `actionMaintenancePageSave` / `actionMaintenancePageForm` in the main module. |

## AJAX protocol notes

`AdminStEasyBuilderEditorController` is the operational hot-spot for any editor bug. Key handlers:

- `ajaxProcessElementorAjax` — catch-all dispatcher; routes to Elementor's internal ajax manager. Most `/editor` requests go through here.
- `ajaxProcessHeartbeat` — keep-alive + autosave (mirrors WP heartbeat).
- `ajaxProcessStyleGlobals` — read/write global style kit (tied to `import_data/settings.json`).
- `processElementorLibraryDirectActions` — template library downloads/imports.
- `ajaxProcessPsAc` — autocomplete for PS objects (products, categories, cms, manufacturers) used when picking "Show on" targets. Implemented via `EBB::getInstance()->getProvider('search')->query(...)->product()/cms()/brand()/etc.` chain (see `classes/data/provider/`).

## Database interaction surface

- `AdminStEasyBuilderController` → `StEasyBuilderClass` (ObjectModel, table `ps_st_easy_builder` + `_yuyan` for i18n + `_shop` for multi-shop).
- `AdminStEasyBuilderGeneralSettingController` → PS `Configuration` table (key/value with `ST_EASY_B_*` prefix).
- `AdminStEasyBuilderLinceseController` → PS `Configuration` (license values) + remote SunnyToo license server.
- `AdminStEasyBuilderEditorController` → proxies to Elementor's internal managers which ultimately hit `ps_easy_builder_postmeta` (ported `wp_postmeta`).

## Cross-links

- [[00-overview#Data model]] — table schema backing this list.
- [[01-widgets-catalog]] — what the editor loads in `initContent`.
- [[02-presets-catalog]] — what `ajaxProcessStyleGlobals` reads on first boot.
- [[README]] — subfolder index.
