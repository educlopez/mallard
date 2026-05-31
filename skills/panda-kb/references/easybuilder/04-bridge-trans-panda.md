---
kind: module-fiche
module_name: steasy_trans_panda
display_name: Easy builder plugin for Transformer theme and Panda theme
version: 1.1.9
config_version: 1.1.8
vendor: SunnyToo (sunnytoo.com)
role: bridge module — Easy Builder ↔ Panda theme
pricing: paid optional addon (ships with Easy Builder purchase)
ps_version: [1.7, 8, 9]
ps_versions_compliancy:
  min: "1.7"
  max: "_PS_VERSION_"
tab: front_office_features
hooks_registered:
  - displayHeader
  - actionEasyBuilderInit
elementor_namespace: ElementorTransPanda
widgets_registered: 20
depends_on:
  hard: [steasybuilder, panda]
  soft: [stblog, stthemeeditor, revsliderprestashop]
  incompatible: [jscomposer, creativeelements, crazyelements]
source_path: <source-archive>/
tags: [easybuilder, panda, bridge, addon, paid, elementor, ps-aware-widgets]
---

# Easy builder plugin for Transformer theme and Panda theme (`steasy_trans_panda`)

> [!info] What this fiche covers
> This is the **bridge module** between [[00-overview|Easy Builder (steasybuilder)]] and the Panda theme. If `steasybuilder` is the generic Elementor engine ported to PrestaShop, `steasy_trans_panda` is the pack of **PrestaShop-aware widgets** that let the editor produce real catalog pages (product sliders, category sliders, brand sliders, Instagram, newsletter, etc.). Without this module, Easy Builder can only assemble generic marketing pages with no PS data.

## Role

`steasy_trans_panda` is a **thin plugin module** (~10 files + a widget library) that piggybacks on `steasybuilder`'s Elementor fork and injects **20 PrestaShop-aware widgets** into the editor. It registers itself on the custom `actionEasyBuilderInit` hook emitted by the main Easy Builder module, at which point it boots its own `ElementorTransPanda\Plugin` singleton (see `elementor/plugin.php`). On `displayHeader` it enqueues its own `frontend.js` and `frontend.css` with priority 60 (after Easy Builder's own frontend at priority 50 — a comment in the source notes in Chinese "qu bao zai easybuilder zhi hou", "must load after easybuilder").

It does **not** register any visible BO screen. The bundled `AdminStEasy_Trans_Panda` tab is marked `active: 0` and exists only to host an `ajaxProcessElementorAjax` endpoint that proxies Elementor's `wp_ajax_elementor_ajax` action into PrestaShop's admin dispatcher.

> [!important] Hard dependency on two modules
> The module checks on install that `jscomposer`, `creativeelements` and `crazyelements` are **not** installed (all three are competing page builders for PrestaShop) and refuses to install otherwise. It also requires `revsliderprestashop >= 5.4.7.5` if that module is present. If `steasybuilder` itself is missing, the configuration screen shows: *"The Easybuilder is composed by two modules, the one you've installed is 'steasy_trans_panda', you've not installed the steasybuilder module yet."*

## PrestaShop-aware widgets

The widget registry is a static array `$build_widgets_filename` in `elementor/plugin.php::init_widgets()`. It registers **20 widgets** under the `ElementorTransPanda\Widget_*` namespace (all extend `\Elementor\Widget_Base` from [[00-overview|steasybuilder]]). Each `.php` file in `elementor/widgets/` that is **not** a base class (`slider-base`, `product-slider-base`, `footer-base`, `dropdown_base`) defines one widget. Two files are currently commented out in the registry: `header-dropdown-content` and `comparison-table` — they ship in the tree but are not wired.

See also [[01-widgets-catalog]] for the 31 generic Elementor widgets bundled in `steasybuilder` itself (headings, buttons, images, tabs, etc.). **The table below is the PrestaShop-aware overlay — catalog widgets that the generic set does not provide.**

| Widget (name) | Display title | Category | PS data source | Key controls |
|---|---|---|---|---|
| `product_slider` | Product Slider | prestashop | `ProductSearchQuery` → products (featured / new / bestseller / category / manual IDs) | source mode, count, autoplay, columns per breakpoint, show price / name / rating, cart-button |
| `product_slider_tab` | Tabbed Product Slider | prestashop | Multiple `ProductSearchQuery` tabs | tab list, per-tab source, active tab, slider config, icon per tab |
| `categoryslider` | Category slider | prestashop | `Category` root + children (tree) | root category, depth, show image / count, columns, autoplay |
| `brandslider` | Brand slider | prestashop | `Manufacturer::getManufacturers()` | brand IDs or "all", show name, columns, autoplay, image size |
| `stblog_slider` | Blog Slider | prestashop | `stblog` module (`StBlogClass`) — **soft dep** | source (latest / category / manual), count, excerpt length, show author/date |
| `psmodules` | PrestaShop Modules | prestashop | Any installed PS module rendered into a hook | target hook, module selector, wrapper markup |
| `stswiper` | Swiper | prestashop | Static slides (image + caption + link), repeater | slides repeater, autoplay, nav arrows, pagination dots, effect (slide/fade) |
| `stbanner` | Banner | prestashop | Static image + link + typography overlay | image, link, title/subtitle, text position, hover effect |
| `sttextslider` | Content slider | prestashop | Repeater of text blocks | slides, animation, interval, alignment |
| `stinstagram` | Instagram | prestashop | Instagram feed (static embed or API) | account, count, columns, spacing |
| `stnewsletter` | Newsletter | prestashop | PrestaShop `emailsubscription` module | title, placeholder, submit label, success message |
| `lookbook` | Lookbook | prestashop | Background image + product hotspots (`Product::getProductProperties`) | bg image, hotspots repeater (x/y + product id), tooltip style |
| `search` | Search | prestashop | Front search controller | placeholder, show categories dropdown, style |
| `stlogo` | Site Logo | header | Shop logo (`Configuration::get('PS_LOGO')`) | image source, size, link target |
| `header_modules` | Header modules | header | PS modules rendered in header (cart, user, wishlist, currency, language) | module checklist, order, icon style |
| `header-dropdown-list` | Dropdown | header | Any flat list (currency / language / custom) | items source, trigger label, icon, dropdown animation |
| `footer-column` | Footer link | footer | Static links or CMS pages (`CMS::listCMS`) | list of links, columns, title, icon |
| `footer-title-text` | Footer Text | footer | Static text | title, text, typography |
| `custom_template` | Custom template | prestashop | A saved Elementor template from the library | template selector, render mode |
| `revsliderprestashop-manager` | Revolution slider | prestashop | [Revolution Slider](https://revolution.themepunch.com/) via `revsliderprestashop` — **soft dep** | slider selector (dropdown of existing rev sliders), custom CSS class |

> [!note] Sliders everywhere
> All sliders share a common backbone — `slider-base.php` (Swiper.js wrapper with autoplay / loop / breakpoints / navigation / pagination / effects / RTL support). Product-driven sliders additionally extend `product-slider-base.php`, which centralizes the `ProductSearchQuery` bridge (source = featured / new / bestseller / category-id / tag / manual IDs). `dropdown_base.php` and `footer-base.php` play the same role for their families.

## Custom templates

**Empty directory.** The module ships a `custom_templates/` folder but it contains no files in the installer — it is an empty hook point. No Panda `.tpl` overrides are distributed here.

The hypothesis that `custom_templates/` would hold Panda theme overrides is **refuted** by the source. Template integration happens a different way:

- The **`custom_template` widget** (see table above) renders a saved Elementor template **by ID** at runtime. Templates are authored in the BO via the [[03-admin-controllers|AdminStEasyBuilderManager controller]] and stored in `ps_st_easy_builder` by the parent `steasybuilder` module. This widget is the only way Panda pages embed Easy-Builder-authored blocks.
- Rendering on the front is handled by the **`render` front controller** (`controllers/front/render.php`), which loads the PS context, includes `steasybuilder/elementor/define.php`, and returns the rendered Elementor tree.

| Template | Overrides | Purpose |
|---|---|---|
| *(none shipped)* | — | `custom_templates/` is an empty installer hook; no `.tpl` files distributed. |

## Hooks registered

From `install()` in `steasy_trans_panda.php`:

- `displayHeader` → `hookDisplayHeader()` — registers `modules/steasy_trans_panda/elementor/assets/js/frontend.js` (priority 60) and `frontend.css` (+ `-rtl.css` variant when RTL).
- `actionEasyBuilderInit` → `hookActionEasyBuilderInit()` — one-liner that instantiates `\ElementorTransPanda\Plugin::instance()`. This is the **single extension point** used by the bridge — it latches onto the custom hook emitted by the parent module when Easy Builder boots the editor.

No other PS hooks (no `displayHome`, no `actionFrontControllerAfterInit`, etc.). The module is deliberately narrow: load CSS/JS, then register widgets into the Elementor engine and let the parent module do everything else.

## Admin controllers

| Controller | Purpose |
|---|---|
| `AdminStEasy_Trans_PandaController` | **Invisible** (`active = 0` in `installTab()`). Sole method: `ajaxProcessElementorAjax()` → forwards to `do_action('wp_ajax_elementor_ajax')`, bridging Elementor's WP-style AJAX router into PrestaShop's admin dispatch. No BO screen rendered. |
| `controllers/front/render.php` | Front controller, route `render`. Bootstraps `steasybuilder/elementor/define.php` and proxies rendering for widgets that need live server rendering (e.g. product sliders with ProductSearchQuery). |

No additional BO menu entries — the real configuration lives entirely inside the [[03-admin-controllers|four `AdminStEasyBuilder*` controllers]] of the parent module.

## Integration with Easy Builder

- **Extension point used**: `actionEasyBuilderInit` (one of the 7 custom hooks registered by [[00-overview#Hook emission on the front|steasybuilder]]). The bridge does **not** use `actionOtherAddEasyBuilderData`, `actionOtherEasyBuilderType`, `actionEasyBuilderPreviewUrl`, or the `Other*` family — those are reserved for heavier extensions that inject new builder types or persist extra data.
- **Elementor category**: all widgets are registered under the existing `prestashop`, `header`, or `footer` Elementor categories. No new categories created.
- **Namespace**: `ElementorTransPanda\Plugin` (PSR-0-style autoload via `spl_autoload_register` inside `plugin.php`). Widget classes: `ElementorTransPanda\Widget_*`.
- **Runtime entry**: on `elementor/init` action (Elementor internal), `on_elementor_init()` loops over `$build_widgets_filename`, `include()`s each widget file, then calls `\Elementor\Plugin::instance()->widgets_manager->register_widget_type( new $class_name() )`.

## Integration with Panda

- **Theme hooks targeted**: none directly. The bridge does **not** emit or consume any Panda-specific hooks (`displayFullWidthTop2`, `displayFullWidthCategoryHeader`, `displayProductDescRightColumn`, etc.). That plumbing stays in the parent [[00-overview|steasybuilder]] module.
- **Panda module behavior**: the widgets read Panda-bundled data providers when available — `stblog` for the blog slider, `steasycontent` and `stthemeeditor` indirectly via `steasybuilder`, `revsliderprestashop` for the Rev-Slider widget.
- **Shared assets with Panda**: none. CSS/JS live in `modules/steasy_trans_panda/elementor/assets/` and are registered as `trans_panda_frontend` (not the Panda theme bundle). The bridge assumes Panda's front Swiper/Bootstrap stack is already loaded by the theme; it does not ship its own jQuery or bootstrap-compat shims.

## Dependencies

- **Hard**:
  - `steasybuilder` — module refuses to function without it; config screen shows an explicit error message.
  - Panda (or Transformer) theme — the widgets expect Panda's CSS grid, Swiper styling, and front JS helpers. Usable with Transformer theme (shared heritage), as the display name suggests ("plugin for Transformer theme and Panda theme").
- **Soft** (widget subset unavailable without them):
  - `stblog` — gates the `stblog_slider` widget.
  - `revsliderprestashop` (≥ 5.4.7.5) — gates the `revsliderprestashop-manager` widget.
  - `stthemeeditor` — affects widget rendering in theme-editor contexts.
- **Incompatible** (install blocked): `jscomposer` (Visual Composer port), `creativeelements`, `crazyelements`. All three are competing PS page-builders; coexistence is not supported.

## Known limits / caveats

- **Chinese comments** in source (consistent with `steasybuilder` origin): e.g. `// default 50, qu bao zai easybuilder zhi hou` ("priority 50, must load after easybuilder"). Not a bug, just a vendor signal.
- **Empty `custom_templates/` directory** — looks like a leftover hook from an abandoned template-override feature. Do not expect Panda-theme `.tpl` overrides here.
- **Two widgets shipped but not registered**: `comparison-table` and `header-dropdown-content` sit in `elementor/widgets/` but are commented out in `init_widgets()`. They can be re-enabled by editing the array, but they are **unsupported** in stock builds.
- **Version skew**: `config.xml` declares version `1.1.8` while the runtime class declares `$this->version = '1.1.9'`. On install, PS uses the class-level version — clients may see `1.1.9` in the BO module list even though the config manifest lags.
- **Widget-to-module coupling is rigid**: the registry is a hard-coded PHP array, not a filter. Third-party modules cannot inject widgets into Easy Builder without forking this bridge or writing a new bridge module that also hooks `actionEasyBuilderInit`.
- **BO tab is hidden by design** (`active: 0`). If the client sees `AdminStEasy_Trans_Panda` in the URL bar, it's only because the Elementor editor is POST-ing to its ajax endpoint.

## Source references

- `config.xml` — manifest (name, display name 1.1.8, author `sunnytoo.com`).
- `steasy_trans_panda.php` — main class `StEasy_Trans_Panda` (install / uninstall / `hookDisplayHeader` / `hookActionEasyBuilderInit`).
- `elementor/plugin.php` — `ElementorTransPanda\Plugin` singleton, `init_widgets()` loop, autoloader.
- `elementor/widgets/` — 25 widget files (20 registered, 4 base classes, 1 commented out).
- `elementor/assets/js/frontend.js`, `elementor/assets/css/frontend.css` — front bundle.
- `classes/EasyBuilderProductBase.php` — shared product-query helper (wraps `ProductSearchQuery`).
- `controllers/admin/AdminStEasy_Trans_PandaController.php` — hidden ajax endpoint.
- `controllers/front/render.php` — front render proxy.

## Related KB

- [[00-overview|Easy Builder overview]] — the parent module this bridge attaches to.
- [[01-widgets-catalog|31 generic Elementor widgets]] — stock Elementor set; this bridge is the PS-aware overlay on top.
- [[02-presets-catalog|Easy Builder presets]] — bootstrap data (some presets reference bridge widgets).
- [[03-admin-controllers|BO controllers]] — the 4 `AdminStEasyBuilder*` screens (none of them live in this bridge).
- [[README|Easy Builder KB index]].
- [[../00-overview|Panda KB root]].
