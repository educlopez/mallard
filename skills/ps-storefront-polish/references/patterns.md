# Page-level patterns (copy the idea, then the files)

Reference repo: `ps9-espaciopiessanos`, theme `themes/pies-sanos-child`. The
commit is the quickest way to see a pattern whole: `git show <hash> --stat`.

| Pattern | Where in the theme | Commit |
|---|---|---|
| Footprints page loader | `templates/components/page-loader.tpl`, CSS "Page loader" | 16fe418b |
| 404 page | `templates/errors/404.tpl`, CSS "404" | a14d223a |
| Cookie card (idxcookies) | CSS "Cookie banner", module template untouched | 6aa51028 |
| Shared drawer component | CSS "Pies drawer", JS "Pies drawer" | bb0edfbd |
| Filters drawer (facets) | CSS "Faceted search", override of facets template | 699326a9, f027cbdb |
| Mini-cart drawer | `modules/ps_shoppingcart/ps_shoppingcart.tpl`, footer partial shell, JS "Mini cart drawer" | 55a500d4 |
| Cart list as one card | CSS "Cart", override `cart-detailed-product-line.tpl` | 29d3eb17 |
| One-page checkout styling | CSS "One-page checkout" (see `ps-one-page-checkout`) | 4f7c6d6f |
| Category compact description | CSS "Category description", JS drawer fill | 34ced4be |
| Product page rebuild, sticky ATC | CSS "Product page", JS "Mobile sticky add-to-cart bar" | bb0edfbd |
| Product tiles with variant dots | CSS "Product tile" | e8299e70 |
| Login and sign-up pages | CSS "Login page", `templates/customer/...` | fdaed924 |
| View transitions | see skill `ps-view-transitions` | cc7517d4 |
| Size guides module | `modules/...` plus theme integration | 2246fc5f |

## Page loader (footprints)

Override `templates/components/page-loader.tpl`. Keep the `page-loader` and
`js-page-loader` classes: core JS toggles `d-none` on them. Four inline SVG
footprints with staggered `opacity` keyframes (2s, 0.35s delay each), alternate
feet mirrored with `scaleX(-1)`; reduced motion shows the first two statically.
Swap the glyph for the client's mark when reusing.

## 404

`templates/errors/404.tpl`: an SVG trail of prints that fade in one by one and
stop at a dashed circle with a question mark, big "404", short friendly title
and two actions (back to home, contact us). 8s loop, reduced motion shows the final state.
Copy from the brand, not from a generic template.

## Cookie card (idxcookies)

Theme-only override of the module markup: `#idxrcookies>div` becomes a floating
card, bottom-right on desktop and full width on mobile. Accept and Reject are
the same size and weight (equal choice). Preferences open a fancybox modal restyled
with the same tokens. Needs `!important` because the module injects inline styles.
Check the module for removed core static methods on PS 9 before blaming CSS: the
HTTP 500 on Accept came from `Tools::jsonDecode/jsonEncode`, replace with `json_decode/json_encode`.

## Mini-cart drawer

- Header cart icon opens the drawer instead of navigating (on the cart and
  checkout pages it keeps navigating).
- The module widget ships the drawer content inside
  `<template class="js-pies-cart-tpl">`. The persistent shell `#pies-cart-drawer`
  lives in the footer partial. A `MutationObserver` on `.ps-shoppingcart`
  copies fresh content into the shell after every `updateCart`.
- Quantity and remove use the URLs the cart presenter already provides
  (`data-pies-cart-url`, `ajax=1&action=update`), then `prestashop.emit('updateCart', ...)`.
- On the cart page, refresh after closing the drawer if something changed.
- Stock-limit errors come back as `json.hasError`: show them in the row.

## Cart as one card

Row layout: image, info, then total, stepper and an icon-only trash button.
The stepper height uses `--pies-btn-h`. Hide the unit-price repetition when the
catalogue data makes it misleading. Checked at 390 and 320px.

## Category description

BO text is pasted rich text with inline fonts. Normalise with
`:is(.category__description,.category__drawer-content) [style]{font-family:inherit!important;...}`.
Long text shows a 6.25rem clamped preview with a masked fade and a "Leer mas →"
button that opens the shared drawer, whose body is a copy filled on first open
(`data-pies-fill-from`). The category header is re-rendered by facets AJAX, so
bind with delegated events.

## Mobile sticky add-to-cart bar

Fixed bottom, white, 74px, soft shadow, only below 992px. Shown while the real
button is above the viewport (IntersectionObserver, no scroll listener). It
mirrors price and disabled state, re-reading after PrestaShop's `updatedProduct`
because the real nodes are replaced, and clicks the REAL button so the stock AJAX
and modal run. `inert` and `aria-hidden` while hidden. Lift the WhatsApp button
above it (`header-layers.md`).

## Product tiles

Order: image, price, brand, name, colour dots. Dots are the variants; wishlist
heart overlay at top-right with `z-index:3`, flags top-left `z-index:2`. Module
wishlist rules are `!important` in `@layer ps-modules`: override from the same layer.

## Login and sign-up

Centred single column, social buttons as one card ("Accede mas rapido"), fields
on the shared tokens. The registration route redirects logged-in customers to My
Account (core `RegistrationController::checkAccess`), so test it logged out.
