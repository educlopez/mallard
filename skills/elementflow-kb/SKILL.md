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
version: "0.2.0"
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

## Gaps to be honest about

- Most of these 50 docs are **short FAQ-style tips** ("here's how to fix X specific problem"), not exhaustive settings references. For the real, complete list of controls/options a widget exposes, the source of truth is the vendored Elementor widget PHP in the actual project (`modules/stsitebuilder/libs/elementor/includes/widgets/*.php`), not this KB.
- The doc site (`elementflow.io`) has no crawlable sitemap — the full page list only appears in the client-side-rendered sidebar nav, not in each page's static HTML. If ElementFlow ships new doc pages later, re-discover them via a real browser render (not a plain HTTP fetch) before assuming this index is complete.
- Screenshots referenced in the original posts were not captured — where a doc says "see the screenshot below," that visual context is missing here.

## Cross-skill pointer

For the PrestaShop platform mechanics ElementFlow sits on top of (hooks lifecycle, theme.yml/parent-child cascade, Symfony BO, Smarty syntax, migration 8→9), see the sibling skill `prestashop-kb`. For GitLab project hygiene on an ElementFlow-based client repo, see `gitlab-project-bootstrap`.
