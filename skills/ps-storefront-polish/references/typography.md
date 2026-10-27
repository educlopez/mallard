# Typography

Generic rules (scale, leading, tracking, measure, wrap modes, tabular figures,
font synthesis, formats, floors): see `emil-typography`. This file holds the
PrestaShop and Hummingbird specifics and what the Pies Sanos pass changed.

## Font files

- One variable woff2 beats three static TTFs. Pies Sanos went from
  `montserrat-300/400/700.ttf` (528 KB, 176 KB each) to a single
  `montserrat-latin-wght.woff2` (38 KB, weights 100 to 900).
- Source: the `@fontsource-variable/montserrat` package (OFL), file
  `files/montserrat-latin-wght-normal.woff2` from jsDelivr. It is already
  subset to latin, so there is no tooling to run.
- `@font-face` for it:
  ```css
  @font-face{font-family:Montserrat;font-style:normal;font-weight:100 900;font-display:swap;src:url(../fonts/montserrat-latin-wght.woff2) format('woff2')}
  ```
- Why it mattered: the design used weights 500 and 600 (65 rules) that were not
  loaded, so 600 painted as 700 and 500 as 400. Check with
  `grep -o "font-weight:[0-9]*" custom.css | sort | uniq -c` against the weights in `@font-face`.
  After the change 500 and 600 are real (canvas width of the same string:
  168, 170, 172, 175 px at 400, 500, 600, 700).
- The latin subset has no arrows (U+2190 to U+2193 except two). `content:"\2192"`
  falls back to the system font and still looks fine in the chips. Add the
  `latin-ext`/symbols file only if the arrow glyph must be identical everywhere.
- Preload: Hummingbird's `_partials/preload.tpl` includes `assets/preload.html`
  from the ACTIVE theme dir and rewrites `href="../` to the theme assets URL.
  A child theme only needs this file, no template override:
  ```html
  <link rel="preload" href="../fonts/montserrat-latin-wght.woff2" as="font" type="font/woff2" crossorigin>
  ```
  `crossorigin` is mandatory for fonts. Verify: one request, initiator `link`, no console warning.

## Scale as role tokens

28 different `font-size` values collapsed to 10 role tokens in `:root`
(roots at 16px). Every `font-size` in the stylesheet reads one, except form
fields (kept literal 16px for iOS) and the 404 display sizes.

```css
--pies-fs-caption:.75rem;   /* 12 */
--pies-fs-small:.8125rem;   /* 13 */
--pies-fs-ui:.875rem;       /* 14 */
--pies-fs-body-sm:.9375rem; /* 15 */
--pies-fs-body:1rem;        /* 16 */
--pies-fs-lead:1.125rem;    /* 18 */
--pies-fs-title-sm:1.25rem; /* 20 */
--pies-fs-title:1.5rem;     /* 24 */
--pies-fs-title-lg:1.75rem; /* 28 */
--pies-fs-title-xl:2rem;    /* 32 */
```

Migration method (scriptable, zero visual change): regex every
`font-size:<value>` whose value equals a step (px and rem forms), snap the few
outliers (12.8, 15.2, 17, 30px) to the nearest step, skip lines about
`.form-control`, `.form-select` and the search input. Name tokens by role, not by
heading level: a site's `h2` is often 24px, so `--pies-fs-h2` would lie.
No step below 12px (the floor); an 11px legal note went to 12.

## Base block

```css
html{-webkit-font-smoothing:antialiased;-moz-osx-font-smoothing:grayscale;font-synthesis:none}
:where(h1,h2,h3,h4,.section-title,.pies-cats__title,.pies-404__title){text-wrap:balance}
:where(.category__description,.pies-sequra__text,.pies-404__text){text-wrap:pretty}
:where(.price,.product__price,.current-price,.product-line__price,.cart-summary-line .value,.header-block__badge){font-variant-numeric:tabular-nums}
.category__description{max-width:min(62rem,75ch)}
```

- `font-synthesis:none` removes the faux italic too: `<em>` from the back office
  renders upright. Acceptable when the face has no italic; say so to the client.
- Use `:where()` so the base block never wins over a component rule.
- Prices and totals change by AJAX (cart, mini cart, variant switch): tabular
  figures stop the layout jitter. List every price class the theme has, not only `.price`.
- BO-pasted descriptions are the only long reading text: cap the measure there.

## Verification method that caught nothing by eye

Snapshot computed `fontSize|fontWeight|fontStyle|lineHeight|letterSpacing` of
every text node per page into `localStorage` BEFORE the edit (key by tag, class,
and occurrence index), then diff after. Pies Sanos: 0 differences on category,
product and cart, 8 on home (all the intended 17px to 16px). Pages: home,
category, product, cart. Do it at 1280, then spot-check 390.

## Gotcha: stale compiled CSS under the same file name

The compiled `assets/cache/theme-<hash>.css` name can stay the same across
content changes and differs per page type (home, product and cart each had
their own hash). The browser keeps the old copy even with "reload ignoring
cache". Force it:

```js
await Promise.all([...document.querySelectorAll('link[rel=stylesheet]')].map(l => l.href).filter(h => /theme-/.test(h)).map(h => fetch(h, {cache: 'reload'})))
```
then reload. Verify a new token with `getComputedStyle(document.documentElement).getPropertyValue('--new-token')`
before trusting any measurement. On a real server visitors may hold the old
file until the hash changes: clear the cache from the back office (Advanced
Parameters > Performance, "Clear cache") so PrestaShop issues a new name.
