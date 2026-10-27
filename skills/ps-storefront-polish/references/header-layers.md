# Header, mega menu and overlay layering

Reference: `assets/css/custom.css` header block (`.pies-header`, `.pies-nav`,
`.pies-mega`, `.pies-mm`), the `--pies-z-*` tokens, and the WhatsApp rules next
to the drawer block.

## The bug: mega menu opens under the cart summary

Hummingbird makes `.cart-grid__aside-wrapper` (and the checkout aside)
`position:sticky; z-index:1020`. `#header` is static with z-index 1021 but is not
a stacking context, so a nav with a lower z-index inside it paints under the
sticky panel. Fix: make the header rows own a stacking context above 1020.

```css
:root{--pies-z-nav:1025;--pies-z-sticky-bar:1030;--pies-z-mobile-menu:1080;--pies-z-cookie:8005}
.pies-header{position:relative;z-index:var(--pies-z-nav);background:#fff}
.pies-nav{position:relative;z-index:var(--pies-z-nav)}
.pies-mega{position:absolute;left:0;right:0;top:100%;z-index:30}   /* z inside the nav context */
.pies-mm{position:fixed;inset:0;z-index:var(--pies-z-mobile-menu)}
.pies-sticky-atc{position:fixed;bottom:0;z-index:var(--pies-z-sticky-bar)}
#idxrcookies>div{z-index:var(--pies-z-cookie)}
```

Scale (keep it in one place, README "Design tokens"):

| Layer | z-index |
|---|---|
| Sticky cart/checkout aside (Hummingbird) | 1020 |
| Header, nav, mega menu | 1025 |
| Mobile sticky add-to-cart bar | 1030 |
| Drawer veil / drawer | 1040 / 1045 |
| Mobile menu | 1080 |
| Product zoom lightbox | 1090 |
| Cookie card | 8005 |
| WhatsApp widget (module) | 10000, so hide it, see below |

Do not give `#header` itself a `view-transition-name` (it would become a
stacking context under the hero). Name the rows instead (`ps-view-transitions`).

How to verify: open the mega menu with a sticky aside visible, then
`document.elementFromPoint(x,y)` at a point where both overlap and check that
`.closest('.pies-mega')` matches. Repeat on the cart page and the checkout.

## WhatsApp (or any floating widget) vs overlays

The module's button is `position:fixed; z-index:10000`, so it floats over
veils, the mobile menu and the sticky add-to-cart bar. Hide it under overlays and
lift it above the bar:

```css
html.pies-drawer-lock .whatsapp,body.pies-mm-open .whatsapp,body:has(.offcanvas.show) .whatsapp{display:none}
@media (max-width:991.98px){
  body:has(.pies-sticky-atc.is-visible) .whatsapp{bottom:calc(74px + 12px + env(safe-area-inset-bottom))!important}
}
```

## Main nav versus search box (no overlap)

Row: list of items (flex, nowrap) plus a fixed-basis search. With 9 items the
list needs about 870px at `.6rem` padding, so the search basis must shrink with
the container:

```css
.pies-nav__item>a{padding:1rem .6rem;font-size:.8125rem}
.pies-nav__search{flex:0 0 320px;max-width:320px;margin-left:auto}
@media (min-width:1200px) and (max-width:1399.98px){.pies-nav__search{flex-basis:240px;max-width:240px}}
@media (min-width:992px) and (max-width:1199.98px){
  .pies-nav__row{gap:1rem}
  .pies-nav__item>a{padding:1rem .4rem;font-size:.75rem;letter-spacing:0}
  .pies-nav__search{flex-basis:170px;max-width:170px}
}
```

Below 992px the desktop nav is replaced by the burger menu, the search sits
under the header. Retest whenever an item is added: measure the last item's
`right` against the search's `left` at 992, 1024, 1100, 1199, 1200, 1280, 1399,
1400, 1440, 1920 (same-origin iframes of each width in one script work well).

## Mega menu notes

- Full-width panel under the nav, content inside a `.container`, opens on
  `:hover` and `:focus-within`, `visibility`+`opacity` so it stays keyboard reachable.
- Elevation: `--pies-elevation-floating`, no top border.
- Touch devices have no hover: a tap on the parent link navigates to the
  category. Acceptable; first-tap-opens needs JS.
