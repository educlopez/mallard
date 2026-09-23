---
name: elementflow-kb
description: >
  Curated knowledge base for ElementFlow — an Elementor-based page-builder
  theme/module for PrestaShop (backed by the `stsitebuilder` module: header,
  footer, home, category, product-page, and CMS-page builders, plus a widget
  library). Use when answering questions about ElementFlow-specific widgets
  (Newsletter, Container, Grid, PrestaShop module, Image, Registration/Login/
  Sign-in forms, Custom template, Tabs & megamenu, Divider, Product name/
  comment/gallery, Search, Text editor, Slider, Accordion), the page builders
  (product/category miniature, product/category page, shopping cart, header,
  mobile menu, CMS page), or ElementFlow-specific features (sticky header,
  wishlist, shortcodes, display conditions, drag & drop editor, sidebar/popup,
  side cart, SVG icons, checkout/my-account child theme, blog, JS events).
  Also trigger for install/upgrade/demo-import/dev-to-production questions
  about ElementFlow, or when the user mentions `stsitebuilder`, `st_site_builder`,
  or a project using the ElementFlow child theme. Trigger PROACTIVELY — even
  without an explicit ElementFlow question — whenever the working project is a
  PrestaShop store using ElementFlow/`stsitebuilder` (e.g. a `themes/*` child
  theme with `parent: classic` alongside a `modules/stsitebuilder/` directory)
  and the task touches storefront UI, layout, header/footer/home/category/
  product-page structure, or theme CSS — check this KB before guessing at
  builder behavior or writing CSS that duplicates something the builder
  already handles natively. For PrestaShop platform-level questions unrelated
  to the builder (Symfony BO, Twig, Smarty core, hooks,
  theme.yml mechanics, migration 8→9) use the `prestashop-kb` skill instead.
  For the Panda theme (`st*` modules, SunnyToo, Easy Builder) use `panda-kb`.
version: "0.3.0"
metadata:
  author: Eduardo Calvo
---

# ElementFlow Knowledge Base

This skill exposes a curated ElementFlow knowledge base: 50 docs scraped from ElementFlow's own demo/docs site, plus empirically-learned gotchas from working on a real ElementFlow project (`ps9-gaudibarcelonashop`) that aren't documented anywhere official.

## When this skill is relevant

Load and read from this KB when the user mentions:

- **ElementFlow** as a theme/builder, or the **`stsitebuilder`** PrestaShop module (its Elementor-based engine — vendored under `modules/stsitebuilder/libs/elementor/` in a project that uses it).
- Any of its **widgets** by name (Newsletter, Container/Flex, Grid, PrestaShop module, Image, Registration form, Login form, Sign in, Custom template, Tabs & megamenu, Divider, Product name/comment/gallery, Search, Text editor, Slider, Accordion).
- Its **page builders** (product/category miniature, product page, category page, shopping cart page, header, mobile menu, CMS page).
- Its **features** (sticky header, wishlist, shortcodes `{SSBC id=X}`, display conditions, drag & drop editor behavior, sidebar/popup, side cart, SVG handling, checkout/my-account child theme, built-in blog, JS events).
- Install/upgrade/demo-import/dev-to-production questions about a project running ElementFlow.

Do NOT load for: PrestaShop core mechanics unrelated to the builder (use `prestashop-kb`), the Panda theme (use `panda-kb`), or non-PrestaShop e-commerce.

## KB layout

```
references/
├── README.md              # conventions + index
└── docs/
    ├── _index.md           # table of all 50 files by section, with source URLs
    ├── welcome.md, faqs.md, system-requirements.md, installation.md,
    │   upgrade.md, import-demo-data.md, setup-service.md,
    │   dev-to-production.md, changelog.md          # Getting started (9)
    ├── product-miniature.md, category-miniature.md,
    │   product-page-builder.md, category-page-builder.md,
    │   shopping-cart-page.md, header-builder.md,
    │   mobile-menu-builder.md, cms-page-builder.md  # Page builders (8)
    ├── widget-*.md                                  # Widgets (18)
    └── feature-*.md, howto-*.md                     # Features (15)
```

## How to use the KB

1. **Start with the index**: `docs/_index.md` for the full table, grouped by section, with the original source URL for each.
2. **Open the specific file** for the widget/builder/feature in question.
3. **"How do I get ElementFlow running?"** → `docs/system-requirements.md`, `docs/installation.md`, `docs/import-demo-data.md`, `docs/dev-to-production.md` in that order.
4. **Changelog / "what changed in vX"** → `docs/changelog.md`.

## Key facts to surface to agents

- **Server requirements**: 512M+ `memory_limit`, 16MB `post_max_size` and `upload_max_filesize`.
- **Install**: via Back Office → Modules → Upload a module (`sitebuilder.zip`), not composer/FTP-only. If widgets don't render after upload, clear cache via FTP as a fallback.
- **Upgrade**: either the "Live Upgrade" one-click path or a manual package swap — either way, clear cache afterward.
- **Demo import**: 4 distinct ways (Import feature, template Library, copy-from-live-frontend, copy-from-editor) — not just one.
- **Dev → production**: domain change, full server transfer, and data-only transfer are three different documented workflows; moving from a subfolder install to root needs an explicit path fix.
- **Shortcodes**: `{SSBC id=X}` reuses a saved widget/section across pages/templates — this is ElementFlow's answer to "reusable blocks."
- **Display conditions**: widgets can be conditionally shown by customer group, cart state, product/category state, and conditions can nest.
- **Side cart**: two different implementation paths exist (a single "Sidebar Cart" widget, vs. composing it from separate product/summary/button/voucher widgets) — check which one a given project actually uses before assuming.
- **PrestaShop module integration**: ElementFlow ships specific compatibility patches for a number of popular third-party modules (SEO Audit, Gift card, Amazzing/Easy filter, WebP, PrestaShop Checkout, loyalty rewards, min/max quantity, product comparison, notify-me) — `docs/feature-ps-module-integration.md` lists exactly which and what the patch does.

## Gotchas learned empirically (not in the official docs)

These came from working on a real ElementFlow project, not from the scraped docs — trust these over guessing from the docs alone when they conflict:

- **PrestaShop's CCC (combine/compress/cache)** will happily bundle a *stale* copy of a theme's `custom.css` if `PS_CSS_THEME_CACHE=1` and the cache wasn't cleared after a CSS rebuild. If a CSS fix "isn't showing up" after a rebuild, check `themes/<theme>/assets/cache/theme-*.css` for stale bundles before assuming the fix didn't apply.
- **Per-widget/global colors**: ElementFlow's `stsitebuilder` DOES support real Elementor-style Global Colors (`__globals__.text_color: "globals/colors?id=..."` in a widget's stored settings), resolved via a separately auto-generated stylesheet (`libs/elementor/js/elementor/css/elementor/css/post-global-setting-css-{shop_id}.css`, regenerated automatically if missing) — NOT via a full "Kit" document/post the way vanilla Elementor works. Don't assume a `__globals__` color reference is broken just because there's no dedicated kit post in the DB; check whether that global-setting stylesheet actually defines the referenced `--e-global-color-*` variable before concluding it's orphaned.
- **`html,body{overflow-x:hidden}`** (a common "kill the mobile horizontal scroll" fix) combined with a `height:100%` reset on `html`/`body` (which Classic/child themes often carry for sticky-footer layouts) triggers a legacy CSS spec quirk: when one axis is non-`visible` and the other isn't set, the browser forces the other axis to `auto` — turning `<body>` into its own internally-scrolling box and silently breaking anything that listens to `window`/`document` scroll events (sticky headers, scroll-triggered JS). Use `overflow-x: clip` instead of `hidden` to sidestep this.
- **A widget's `_css_classes` / `css_classes` field** (set via the builder's Advanced tab — `_css_classes` on widgets, `css_classes` on containers) is the correct way to give theme CSS a stable hook — prefer it over Elementor's auto-generated per-element IDs/classes (`.elementor-element-xxxxxxx`), which are long, unstable across re-saves, and hard to read.
- **Header/footer/home content is genuinely split** between builder-native config (colors, typography, per-widget custom CSS — set via the admin UI, stored in DB, wiped on a DB refresh from server) and theme CSS (structural/responsive/JS-driven behavior the builder UI can't express). When deciding where a fix belongs, ask which category it falls into — see the `gitlab-project-bootstrap` skill's sibling notes on this for the DB-refresh risk of builder-native settings.
- **A container's `content_width` setting (`boxed` vs `full`) changes the actual DOM shape**, not just the visual width: with `full`, widgets render as direct children of the container; with the default/boxed setting, Elementor auto-wraps them in an extra `.e-con-inner` div first. A CSS selector written against one shape (e.g. `.my-col > .elementor-widget`) can silently match nothing on a sibling page/container that uses the other setting, even though both were built with "the same" container. Cover both shapes (`.my-col .e-con-inner, .my-col > .elementor-widget { ... }`) rather than assuming one.
- **Elementor's own auto-generated per-widget CSS (`post-{id}.css`) loads AFTER theme `custom.css`** in `<head>`, and can carry the exact same selector specificity as a theme override (e.g. both are two classes deep). At equal specificity, source order wins — so the widget's own generated rule (commonly `max-width:100%` from a width/size setting) silently beats a theme CSS override that "should" apply. If a CSS property you set on a widget isn't taking effect despite the selector clearly matching, suspect this cascade order first and add `!important` rather than hunting for a specificity bug that isn't there.
- **A Grid-type container (`container_type: grid`) does not default its column-gap to zero** — Elementor's own grid gap default leaves a visible strip of page background between columns even when both columns' own padding is zeroed. Set `gap: 0` explicitly on the grid container if you want the columns to touch.
- **`content:url(...)` on a CSS pseudo-element (`::before`/`::after`) is not a real replaced element for sizing purposes**: it paints the image at its native pixel resolution regardless of any `width`/`height` set on the pseudo-element, and does NOT auto-preserve the image's intrinsic aspect ratio the way a real `<img>` does. This bites specifically when faking a builder-style layout on a page ElementFlow has no document for (e.g. PrestaShop-core `password-recovery`/`forgot-password`, which isn't a page you can open in the builder) and you reach for a pseudo-element to inject a logo or hero image via CSS alone. Use `background-image` + `background-size: contain` (or `cover`) on an empty (`content:""`) pseudo-element instead — that correctly rescales, same as `object-fit` does for a real `<img>`.
- **Faking a builder-style split/grid layout on a non-builder PrestaShop-core page** (no ElementFlow document exists for it) means you're fighting Bootstrap, not Elementor: Bootstrap's `.container` sets an explicit `width` per breakpoint (capped well below the viewport even at its largest breakpoint), not just `max-width` — overriding `max-width: none` alone leaves it boxed with dead margins on wide screens; you also need `width: 100%`. Likewise `.form-check`/`.form-check-label` (Bootstrap checkboxes) carry legacy absolute-position offsets/padding-left that will double up with a flex `gap` if you re-lay them out — zero the old padding explicitly rather than just adding a gap on top of it.

## `_elementor_data` is the source; `.tpl` and `post-N.css` are build artifacts

The single most expensive lesson: editing `bs_st_site_builder_postmeta._elementor_data` directly is only safe for **pure content**, never for structure or anything that emits CSS. Two separate generated artifacts go stale:

- **Almost every document type compiles to a per-language Smarty template** under `modules/stsitebuilder/views/templates/front/template/`. Two different naming schemes coexist: **miniatures** are `{id_lang}-{doc}-{type}.tpl` (e.g. `1-8-60.tpl`), while **header/footer/mobile-menu/page** docs are `{id_shop}-{id_lang}-{doc}-{type}.tpl` (e.g. `1-1-2-30.tpl`, `1-2-2-30.tpl`, `1-3-4-120.tpl`). Only an **editor publish** regenerates them (`page.php::render_template`), and the front reads the `.tpl`, not the JSON — `SiteBuilderBase.php:~1520` even prints "Template … not found." when one is missing. Consequence: after a raw DB edit the default language may look right while the others serve stale content. **Always verify every language on the front, not just the default one.** (When the files are absent entirely the front still renders live from data, so absence is not fatal — staleness is.)
- **`post-{doc}.css` is also only rebuilt on publish.** Deleting a setting from `_elementor_data` does NOT remove the CSS rule it generated. Symptom: you remove e.g. `_flex_size: grow` in the DB and the layout is unchanged, which reads like a specificity bug but is a stale artifact.
- **An open Elementor editor autosaves and will silently clobber DB edits.** Close the editor tab before touching `_elementor_data`, and re-read the row afterwards to confirm your write survived.
- Workable pattern when you must bulk-edit: close the editor → write the DB → open the editor fresh (it loads your data) → **Publish** (regenerates every artifact for every language).

## Flex layout inside a header container

- **Every child of a header row typically carries `content_width: full`**, i.e. each asks for `width: 100%` of the row with `flex: 0 1 auto`. With two children they shrink to fit and it looks intentional; **add a third and all three collapse to equal thirds**. Worse, a middle child that demands 100% and refuses to shrink (`_flex_size: grow`, or shrink 0 without a content-based width) starves its siblings to `width: 0` — a logo silently disappears and icon groups overlap.
- The fix for a content-sized item between two flexible ones (e.g. a nav menu between logo and cart icons): on that container set **Width = `max-content`** (custom unit) **plus** `_flex_size: 'custom'` and `_flex_shrink: 0`. `fit-content` does NOT work — it still shrinks.
- To place an element between existing siblings without dragging, use the builder's **Order** control (`_flex_order`: `start` | `end` | `custom`) on a sibling — e.g. setting the icons container to `end` turns logo|icons|menu into logo|menu|icons.

## Automating the editor: what works and what doesn't

- **Repeater controls are hostile to scripted editing.** Setting `tabs` / `list` item fields via injected JS repeatedly froze the renderer (CDP `Runtime.evaluate` timeouts), and a row duplicate that appeared in the DOM did not survive Publish. Non-repeater controls script fine. For bulk item edits prefer the DB route (then publish) or the clipboard route below.
- **Drag & drop cannot be automated.** In the Structure navigator, dropping onto a container row always nests INSIDE it — never as a sibling — and dropping on a row boundary or on the parent row does not produce a sibling either. A canvas drag from the container grip does nothing. A human dropping on the boundary between two containers works fine.
- **Cross-environment element transfer without the system clipboard** (the reliable alternative to both): ElementFlow's own "Paste from other site" needs a *real* `cmd+V` and fails on a synthetic one with *"Make sure that both sites are updated to last version of Elementor…"*. Instead, note that **Copy writes the payload to `localStorage['elementor'].clipboard`** as `{type:'elementor', siteurl:'<origin>/', elements:[…]}`. Copy in the source, rewrite `siteurl` to the target origin, then in the target tab merge it in — `const st=JSON.parse(localStorage.getItem('elementor')||'{}'); st.clipboard=payload; localStorage.setItem('elementor',JSON.stringify(st));` — and the normal **Paste** becomes enabled immediately, no reload needed. Gzip+base64 the payload and decode in-page with `DecompressionStream('gzip')` to keep it small. **Paste always inserts INTO the right-clicked container as its last child**, so paste the smallest piece you need and fix order with `_flex_order`.
- **Safety net when you have no DB access to the target** (e.g. a staging server): the Theme builder row dropdown has **Duplicar**. Duplicate the document first — the copy lands active but *unbound*, so it renders nowhere — and if the edit goes wrong repoint the "Bind to pages" entry at the duplicate.

## Menu widgets: two different link-type conventions

Both the header menu and the mobile menu can build links natively instead of hardcoding URLs, but **they don't share the same schema** — check which widget you're editing:

- **`nested-tabs`** ("Tabs & Megamenu", the desktop navbar): per-item `type` uses **word values** — `''` (default/custom) | `category` | `information` | `my_account` | `cms` | `manufacturer` (from `SiteBuilderBase::getMenuTypes()`), with conditional `information_link` / `cms_link` / `category_link` fields, plus `tab_title_{iso}` and `tab_url_{iso}`. The title is **always** taken from `tab_title_*`, whatever the type.
- **`mobilemenu_list`** (mobile sidebar; controls come from `menu_base.php::_reg_list_source`): the repeater is called `list` and its per-item `link_type` uses **numeric strings** — `0` Custom | `1` Cms | `2` Manufacturer | `3` Category | `4` My account | `5` Information | `6` Product. Critically, its `text_{iso}` ("Custom title") and `link_{iso}` ("Custom link") are `'condition' => ['link_type' => '0']`, so **for any native type the widget resolves BOTH the URL and the label from the linked entity and ignores your custom title**. If you need a label that differs from the entity's own name (e.g. "Novedades" rather than PrestaShop's "Nuevos productos"), you must use `link_type: '0'` and supply the per-language URLs yourself.
- `information` link options are `prices-drop`, `new-products`, `best-sales`, `stores`, `contact`, `sitemap` (`classes/data/element/Linker.php`).

## Desktop and mobile menus are separate documents

The mobile sidebar is its own builder document (type `120`, typically id 4, "Mobile menu"), completely independent of the header — editing the navbar does nothing to it. On a fresh demo import it is full of demo data (nested accordions titled "Men" containing "Underwear / Undershirts / Pajamas / Collections"), which is what a client sees as "men men" in the burger menu. For a flat first-level menu, replace the demo `nested-accordion` with a single `mobilemenu_list`; keep the accordion only when you actually need expandable sub-levels. Also check the small label widget above it — on the demo it says "Categorías" with EN/CA left as "Add Your Text Here".

**The sidebar background is dark (`#1A1A1A`) but widget text defaults to black**, so a freshly added `mobilemenu_list` is invisible (measured 1:1 contrast). The demo widgets carried their colour in `__globals__` rather than in a plain colour field, so it's easy to drop when hand-building settings. Set `__globals__.mobile_menu_color = 'globals/colors?id=secondary'` — in this theme `secondary` resolves to pure white (17.4:1 on the sidebar) — and the chevron SVG follows automatically because its `fill` inherits `currentColor`. The demo also set the colour only on the `_mobile_extra` responsive variant (`mobile_menu_color_mobile_extra`), which leaves the base breakpoint black; set the base key instead.

## Driving the editor through Elementor's own JS API

Far more reliable than clicking the panel, and it goes through the model so it saves and regenerates CSS normally:

```js
// find a widget by type, anywhere in the document
let target=null;
const iter=c=>c&&c.each&&c.each(m=>{ if(m.get('widgetType')==='mobilemenu_list') target=m; const e=m.get('elements'); if(e) iter(e); });
iter(elementor.documents.getCurrent().container.model.get('elements'));

const c = elementor.getContainer(target.get('id'));
$e.run('document/elements/select',   {container:c});
$e.run('document/elements/settings', {container:c, settings:{ text_color:'#FFFFFF99' }});
await $e.run('document/save/publish');            // publish without hunting for the button
elementor.documents.getCurrent().editor.isChanged // false once saved
```

`elementor.getContainer(id)` + `$e.run('document/elements/settings', …)` also sidesteps the repeater-freeze problem for non-repeater fields, and `document/save/publish` avoids coordinate-clicking a Publish button whose position shifts.

**But `document/elements/create` cannot populate a NESTED widget's panels.** `nested-tabs` owns its child containers (one per `tabs` repeater row) and creates them itself when rows are added through the UI. Passing a pre-populated `elements` array in the model, or calling `create` with the nested-tabs as the target container, both silently lose the children on save — in one case 6 panels came back as 1, and the `layout` setting was dropped too. What DOES survive: writing the children straight into `_elementor_data` (with the editor closed), then opening the editor fresh and publishing — verified with both a 7th navbar tab and 6 megamenu panels. So: **settings via the JS API, nested children via the DB.**

## Building a megamenu panel that matches a two-column design

The desktop navbar is a `nested-tabs` with `layout: menu`; each top item's dropdown is just an empty container you fill. For the common "links on the left, image that swaps on hover on the right" design, **nest a second `nested-tabs` inside that container** — no custom JS needed:

- `layout: 'tabs'`, `open_on: 'hover'`
- `tabs_direction: 'inline-start'` → titles column to the LEFT of the content (`block-start`/`block-end`/`inline-end`/`inline-start`; the default puts titles above)
- `tabs_heading_direction: 'column'` → titles stack vertically
- `tabs_width: {unit:'px', size:280}` → left column width (emits `--n-tabs-heading-width`)
- link look: `title_typography_*` group (plus `title_typography_hover` / `_active`), `tabs_title_padding`, `tabs_title_align_items`, `tabs_title_space_between`
- one child container per tab = the right-hand panel: background image `cover`/`center`, `min_height` = panel height, `flex_justify_content: 'flex-end'` to sit content at the bottom, and a gradient overlay (`background_overlay_background:'gradient'`, `background_overlay_color:'rgba(0,0,0,0.6)'`, `background_overlay_color_b:'rgba(0,0,0,0)'`, `background_overlay_gradient_angle: 0deg`) for the caption legibility.

Reference relative image URLs (`/img/c/10-category_cover.jpg`) rather than absolute ones so the same config works in local and staging.

**Watch the dropdown-width trap**: the dropdown can never be wider than the menu widget's own container. If you gave that container `width: max-content` to stop it starving its siblings, you have also capped the dropdown. `submenu_width: '1'` ("stretch to full width") is documented as working *conditionally* and does nothing in that case. The fix is to invert which children are content-sized: give the logo and icon containers `width: max-content` + `_flex_size:'custom'` + `_flex_shrink: 0`, and let the menu container stay full-width with `flex_justify_content: 'center'`. Nothing collapses, and the dropdown gets the full menu width.

**Titles vs linked entities differ between the two menu widgets** (see the section above): in `nested-tabs`, `EBB::textAndLinkLang('tab_title', …)` wins and the entity name is only a fallback, so you can label a tab "Muebles" while it links to a category actually named "Mobiliario". In `mobilemenu_list` you cannot.

## Gaps to be honest about

- Most of these 50 docs are **short FAQ-style tips** ("here's how to fix X specific problem"), not exhaustive settings references. For the real, complete list of controls/options a widget exposes, the source of truth is the vendored Elementor widget PHP in the actual project (`modules/stsitebuilder/libs/elementor/includes/widgets/*.php`), not this KB.
- The doc site (`elementflow.io`) has no crawlable sitemap — the full page list only appears in the client-side-rendered sidebar nav, not in each page's static HTML. If ElementFlow ships new doc pages later, re-discover them via a real browser render (not a plain HTTP fetch) before assuming this index is complete.
- Screenshots referenced in the original posts were not captured — where a doc says "see the screenshot below," that visual context is missing here.

## Cross-skill pointer

For the PrestaShop platform mechanics ElementFlow sits on top of (hooks lifecycle, theme.yml/parent-child cascade, Symfony BO, Smarty syntax, migration 8→9), see the sibling skill `prestashop-kb`. For GitLab project hygiene on an ElementFlow-based client repo, see `gitlab-project-bootstrap`.

### Megamenú con `nested-tabs`: el desplegable no puede ser más ancho que su columna (y `submenu_width` no lo arregla)

En `layout: menu`, el panel del dropdown es `.e-n-tabs-content` con `position: absolute`, y su **bloque contenedor es el container que envuelve al widget** (Elementor pone `position: relative` en todos los `.e-con`). Si el widget vive en una columna de 985px, el panel se queda en 985px hagas lo que hagas: `submenu_width` no tiene efecto en este layout, y `_position: absolute` sobre el container hijo del panel **tampoco** (Elementor no emite CSS de posicionado para containers con ese control; queda `position: relative`).

La única salida es mover el bloque contenedor hacia arriba con CSS: `position: static` en la columna del menú → el bloque contenedor pasa a ser la fila del header, y ahí ya puedes centrar el panel con `width: min(calc(100% - 24px), Npx); left: 50%; transform: translateX(-50%)`. Comprueba `offsetParent` del `.e-n-tabs-content` para saber contra qué estás midiendo antes de escribir números: al cambiar el `position` de un ancestro, `top`/`left` se recalculan contra otro elemento y el panel puede acabar solapando el header (necesitarás `top: 100%`).

### Fondo de un container: sin `background_size: cover` la foto sale a tamaño natural

Al poner una imagen de fondo en un container (por ejemplo el panel de un tab que hace de "foto grande"), el `background-size` por defecto deja la imagen a su tamaño natural centrada, con el fondo del container asomando a los lados. Parece un bug de anchura y no lo es: `background_size: 'cover'` + `background_position: 'center center'` es nativo y es lo que equivale al `object-fit: cover` de un `<img>`.

### `post-{doc}.css` gana a `custom.css` por orden de fuente

Ya está anotado más arriba para widgets; aplica igual a las variables/paddings que emiten los `nested-tabs`. Si pones un `padding` en un elemento interno del widget desde el tema y el computed sigue mostrando el valor del builder, no busques un fallo de especificidad: son iguales y el generado carga después. `!important` es la respuesta correcta aquí.

### Publicar por `$e.run('document/save/publish')` a veces no resuelve

El `await` puede colgarse (timeouts de `Runtime.evaluate` a los 45s) aunque los settings sí se hayan aplicado al modelo. No asumas que falló: vuelve a leer `elementor.documents.getCurrent().editor.isChanged` en una llamada aparte; si sigue `true`, relanza el publish. Y ojo al navegar fuera del editor: un `beforeunload` bloquea la navegación si quedan cambios sin guardar.

### `tabs_justify_vertical` es el control de alineación vertical de los títulos

Con `tabs_heading_direction: column`, el control que alinea la lista de títulos arriba/centro/abajo es `tabs_justify_vertical` (`start|center|end|stretch`), NO `tabs_justify_content` (ése está condicionado a direction `row`). Su diccionario de selectores reescribe también `--n-tabs-title-flex-basis`, así que al cambiarlo se pierde la separación que dieran los paddings: recupérala con `tabs_title_space_between`.

### Las miniaturas de categoría/producto de PrestaShop pueden traer relleno blanco — no las uses como imagen de fondo

Si la tienda tiene la generación de miniaturas en modo "fit" con relleno (lo habitual en las importaciones de demo), los ficheros `img/c/{id}-category_cover.jpg`, `-category_default.jpg`, etc. contienen la foto original **centrada sobre un lienzo blanco** del tamaño nominal. Al usar uno de esos ficheros como `background_image` de un container con `background-size: cover`, el relleno se escala con la foto: se ve la imagen pequeña en el centro con franjas a los lados, y el síntoma se confunde con "el contenedor es más ancho de lo que debería" o "el cover no se aplica".

Comprueba siempre el fichero en disco (`sips -g pixelWidth -g pixelHeight`) y compáralo con lo que se ve renderizado; si el ratio del contenido visible no coincide con el ratio del fichero, hay relleno. La única sin relleno es la **original** (`img/c/{id}.jpg`).

### El mismo widget puede salir `position: relative` en un entorno y no en otro

Un `nested-tabs` idéntico (mismos settings, replicado por clipboard) generó `position: relative` sobre el propio widget en un entorno y no en otro. Si tu CSS depende de cuál es el bloque contenedor de un descendiente absoluto (por ejemplo un dropdown de megamenú), eso significa que el layout funciona en local y se rompe en el servidor sin que haya diferencia de datos aparente.

No confíes en heredar el `position` que "toca": fuerza explícitamente a `static` **todos** los ancestros que quieras saltar (`!important`, porque el relative lo pone el CSS generado del widget, que carga después del tema), y verifica con `element.offsetParent` en cada entorno en lugar de asumir la cadena que viste en local.
