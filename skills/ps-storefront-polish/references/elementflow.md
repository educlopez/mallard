# ElementFlow (stsitebuilder) notes

Frictions found applying this skill to a store on the `elementflow` theme (parent `classic`, builder
module `stsitebuilder`). Class prefixes like `gb-` below are the project's own; use yours.

Theme CSS source lives in `_dev/css/*.css`, built with `pnpm run css` (lightningcss) to
`assets/css/custom.css`. Head assets are added in `templates/_partials/head.tpl` (`hook_extra`
block), not in `theme.yml`.

## Selectors that differ from Hummingbird

| Thing | Hummingbird | ElementFlow |
|---|---|---|
| Fixed header | `#header` | `#stsb-header` (made fixed by the project on scroll). Wrapper `#header` is only a min-height spacer |
| Product tile | `.js-product-miniature` | same class, inside `.stsb_grid_wrapper`; image is `.stsb_pm_image` |
| Product page photo | `.product__carousel .carousel-item.active img` | `.stsb_pro_gallery_item` (two exist, one is `display:none`, so one visible name only) |
| Product list container | `#js-product-list` | none. Wrap `.stsb_grid_wrapper .js-product-miniature` |
| Header icons | | `.stsb_bar_item`, `.stsb_cart_widget_link`, `.stsb_customer_link` |
| Search input | | `input[type=text].stsb_search_form_input` |

## Facets and view transitions

stsitebuilder swaps the grid synchronously inside its `updateProductList` handler, then emits
`stsbDomUpdate`. So the `prestashop.emit` wrapper from `ps-view-transitions` layer 2 works. Verified
on a category page: 1 `startViewTransition` call, `ready` and `finished` resolved, the list class
removed. Clear the inline `view-transition-name` on tiles when the transition finishes, or each tile
keeps a stacking context.

Name `#stsb-header` only, never both `#header` and `#stsb-header` with the same name: duplicate
names abort the whole transition.

## Traps

- `input[type="text"]{font-size:14px}` in the base sheet is (0,1,1), so `.stsb_form_control{font-size:16px}`
  loses. Use `input[type=text].stsb_search_form_input`.
- A promo rotator can be wider than the viewport on purpose (clipped). Do not report it as overflow.
- Full-page screenshots show the off-canvas cart inline and logo walls empty (lazy images). Judge
  layout from viewport screenshots and from `scrollWidth`.
- Lando: if `lando start` fails with `network ... not found`, run `lando rebuild -y` for that app.
  The DB lives in the named volume `<app>_data_database`, so it survives.
- Console error seen on the product page: `productCommentMandatoryMessage is not defined`
  (the productcomments module var is defined in `post-comment-modal.tpl`, which is not rendered).

## Surfaces, hover and tokens

- `.elementor a{box-shadow:none}` is (0,1,1): any shadow on a tile wrapper (an `<a>`) needs
  `.elementor a.<class>` or it never shows.
- Hover lift/zoom must sit in `@media (hover:hover) and (pointer:fine)`. The theme had none, so tiles
  stayed lifted after a tap.
- `stsb_pm_image_wrapper` needs `background:#fff` when photos are white-matte JPG/PNG: on a tinted
  section they look like white squares.
- Active filters: `#js-active-search-filters` (ps_facetedsearch) can be styled as pill chips; the
  clear-all link is `.js-search-filters-clear-all`.
- Keep elevation tokens, press feedback (`scale:.97`) and the global reduced-motion net in one
  partial imported last.

## Typography, forms, touch

- Builder markup has no accessible names: header search/menu/cart (`.e-a-search`, `.e-menu`,
  `.e-b-cart` svg classes), `.wishlist-button-add` (text is the ligature `favorite_border`), share
  links, price slider handles. Label them from a small JS file per `<html lang>`. `textContent` is
  never a real name here (count bubble, ligature), so always set `aria-label`.
- Off-canvas panels (`.stsb_side_panel_content`) and `.swiper-slide` stay tabbable while parked
  off-screen. Toggle `inert` with an IntersectionObserver. TRAP: a panel parked exactly on the
  viewport edge is `isIntersecting:true, ratio:0`, so threshold 0 never flips it. Use
  `rootMargin:'-1px'` and `isIntersecting`.
- Builder headings set `line-height` equal to `font-size` (1.0). Force about 1.15 with
  `text-wrap:balance`.
- Muted text colors tend to fail contrast (a `#8a8a8a` was 3.45:1 on white). The tax label is 11px and
  its builder rule needs `.elementor .stsb_pm_price_item.stsb_pm_tax_label` to win.
- Registration/login fields ship without `autocomplete`: add tokens by field name from JS.
- Check `font-synthesis:none` only when every used weight is actually loaded.
- Leave third-party z-indexes alone (header 999, Bootstrap 1000/1040, module toasts 9999) unless a
  clash is seen.
- Audit script checks: size <12px, contrast, heading wrap, measure, tabular price, input
  labels/type/autocomplete, unnamed controls, tabbable off-screen, image dims/alt, small taps.

## Builder content, translations and rebuilds

- Builder text lives in `ps_st_site_builder_postmeta` (`meta_key='_elementor_data'`, JSON per doc).
  Strings are per language with an ISO suffix: `title_es`, `title_en`, `title_fr`, `title_ca`. A doc
  seeded with English demo text has the English in `*_es`. To translate: set `_es` to the shop
  language and add the others. Back the row up first (`SELECT meta_value ... > file`) and write with
  `UNHEX('<hex>')` to avoid quoting bugs.
- Labels that come from widget PHP (`EbTrans::ll('Subject')`, contact form, cart "Remove") are NOT in
  the JSON: add them to the theme XLF `ModulesSitebuilderShop.<locale>.xlf` (domain
  `Modules.Sitebuilder.Shop`). ps-translate `write_xlf.py` does it. The module concatenates
  `number + ' ' + 'days ago'`, so "ago" strings must be suffix-style, never prefix-style.
- Compiled templates `modules/stsitebuilder/views/templates/front/template/*.tpl` (gitignored): page
  docs are `{lang}-{shop}-{doc}-{type}.tpl` and regenerate on request; MINIATURES
  (`{lang}-{doc}-{type}.tpl`, e.g. cart item, product tile) do NOT regenerate on request. NEVER
  `rm *.tpl`: it removes miniatures and the cart/tiles break. If it happens: restore from a copy, or
  run the module's own rebuild.
- Own rebuild without the back office (works): a throwaway PHP in the project root that sets
  `$_SERVER['HTTP_HOST']`, requires `config/config.inc.php`, sets
  `Context->controller = new FrontController()`, shop/language/currency/cookie/customer/cart, then
  `Module::getInstanceByName('stsitebuilder')->elementPlugin(); ->ClearFileAndCSSCache();`. It
  deletes page tpl + all `post-*.css`, then recompiles miniature tpl + CSS. Page CSS regenerates on
  request. Without a `FrontController` it fatals (`Elementor\Core\Breakpoints\Manager not found`).
- Compile-time `ll()` runs with the CLI context, which does not see theme XLF: after a CLI rebuild
  some miniature labels (cart "Remove") are English again. Patch the compiled tpl text per language
  or save the doc from the back office.
- An empty "new products" slider (`product_slider_post_type: new`, `PS_NB_DAYS_NEW_PRODUCT`) renders a
  blank band. Hide with
  `.e-con.e-parent:has(> .e-con-inner > .elementor-widget-product_slider):not(:has(.js-product-miniature)){display:none}`.
- `.swiper.stsb-swiper-container` is `overflow:hidden`: card ring/shadow is clipped. Give it padding
  and an equal negative margin.
- Regressions to avoid: a pseudo-element hit area (`::before{inset:-12px -19px}`) on the price slider
  handle made the category sidebar scroll sideways; `padding-block` on facet links pushed the checkbox
  out of line (centre the `.facet-label` row and zero the link margin with `!important`, classic sets
  `margin-top:7px`).
- Demo promo claims left in English on purpose ("Free Shipping On All Orders", "Receive 15% off your
  first order"): they need the shop's real policy before translating.
