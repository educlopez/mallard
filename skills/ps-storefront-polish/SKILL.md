---
name: ps-storefront-polish
description: >
  Polish method and reusable details for the storefront of ANY PrestaShop 8/9
  theme (Hummingbird, Classic, Panda, ElementFlow or a custom child). A repeatable
  pass over design tokens, colour, surfaces, forms and focus, typography, motion
  and overlay layering (z-index), each verified in the browser, plus a library of
  tested recipes (shared drawer, mini cart, page loader, 404, cookie card, sticky
  add-to-cart, mega menu fit, Lottie wishlist like, quiet scrollbars and scroll
  fades, Sonner-style toasts, module modals) taken from the Pies Sanos
  Hummingbird build. Use
  when polishing or starting a storefront, when the user says "pulir diseno",
  "design tokens", "z-index" ("el menu se abre debajo del carrito"), "el input
  hace zoom en iOS", "focus", fonts and type scale, "animacion de like / corazon
  de wishlist / lottie", "scrollbars", "scroll area", "toasts", "fluid hover" (a highlight that glides between items), "modales de
  wishlist", "replica esto de otra web / como en <web>" (reverse engineer a
  reference with rea), or wants the same details as Pies Sanos on another project. For page and filter transitions use
  ps-view-transitions; for the native one-page checkout use ps-one-page-checkout.
version: "0.5.0"
---

# PrestaShop storefront polish

Two layers:

1. **Method and principles (any theme).** Pick the passes, run the generic skill
   for each, fix, verify in the browser, record what you learned. The rules,
   numbers and checks in this file and in `tokens-surfaces.md`, `forms.md`,
   `typography.md`, `motion.md` and the z-index part of `header-layers.md` apply
   whatever builds the pages.
2. **Recipes (adapt per theme).** Code that was written and measured on one
   Hummingbird child theme (`ps9-espaciopiessanos`, theme `pies-sanos-child`):
   shared drawer, mini cart, page loader, 404, cookie card, sticky bar, nav fit.
   Copy the idea and rewrite the selectors for the theme in front of you.

## Step 0: map the theme before touching anything

| Question | Where it changes the work |
|---|---|
| Which theme and parent (`config/theme.yml`)? Hummingbird, Classic, Panda (SunnyToo `st*` modules), ElementFlow (`stsitebuilder`), custom | Selectors, which pieces already exist (`panda-kb`, `elementflow-kb`, `prestashop-kb`) |
| Where does CSS live and how is it built? (`assets/css/custom.css`, `_dev/css` plus a build, a builder's Custom CSS) | Never edit generated output. See `ps-css-build` |
| Who owns each area: template, module, or a visual builder? | In a builder, change the setting or the class in the editor, not generated ids |
| Does the project have its own rules file (`CLAUDE.md`)? | It wins over this skill |
| What does the theme already do (tokens, focus ring, drawers)? | Reuse it, do not stack a second system |

Then record the answers in a new `references/<theme>.md` while you work.

Companion skills for related topics: `ps-view-transitions`, `ps-one-page-checkout`,
`ps-css-build`, `prestashop-kb`, `panda-kb`, `elementflow-kb`.

## Companion skills (generic principles live there, not here)

This skill keeps only what is specific to PrestaShop and to the Pies Sanos
decisions. Run the generic skill for the principle, then use this one for the
Hummingbird details. If a companion is not installed on the machine, skip it.

| Topic | Generic skill | What this skill adds |
|---|---|---|
| Colour, contrast, oklch | `emil-color` | the token values and the AA fixes (`tokens-surfaces.md`) |
| Shadows, borders, radii, gradients | `emil-surfaces` | the elevation tokens, hairline, scrim and mask snippets |
| Inputs, labels, focus, autofill | `emil-forms-and-inputs` | Hummingbird overrides, checkout autocomplete, birthday mask (`forms.md`) |
| What to animate and how | `find-animation-opportunities`, `improve-animations`, `emil-animations` | the snippets that shipped and the rejected list (`motion.md`) |
| Type scale, font files, wrap, figures | `emil-typography` | role tokens, variable woff2, preload, base block (`typography.md`) |
| Review of the whole storefront | `emil-ui-review` | the audits, cascade and overflow traps, slim checkout footer (`ui-review.md`) |
| Fine details (z-index, focus, scrollbars) | `emil-ui-polish` | the z-index scale and overlay fixes (`header-layers.md`) |
| Lottie like on the wishlist heart, recolour, lazy load | `emil-animations`, `emil-performance` | the whole recipe, licence rules and the Vue `characterData` trap (`lottie-like.md`) |
| Scrollbars, scroll areas, edge fades | `emil-ui-polish` | the three-tier recipe after fluidfunctionalism (`scrollbars.md`) |
| Hover highlight that glides between items | `emil-animations` | the nearest-item component, axes y/x/xy, traps (`fluid-hover.md`) |
| Toasts, module modals (blockwishlist) | `emil-animations`, `emil-surfaces` | Sonner port, mirroring module toasts, one dialog language (`toasts-modals.md`) |
| Copying a detail from another website | `reverse-engineer-anything` | the rea method for storefront details, licence rules, the x.com 403 and logged-in Chrome traps (`replicate-from-web.md`) |
| Page, facet and tile transitions | `ps-view-transitions` | nothing, use it as is |
| Native one-page checkout | `ps-one-page-checkout` | only the checkout styling and autofill |
| CSS build pipeline | `ps-css-build` | nothing |
| Core PrestaShop questions | `prestashop-kb` | nothing |

Workflow that worked: run the generic skill on the theme with the project
constraints (tokens, branch, no commit), fix what it finds, verify in the
browser, then record anything reusable in the matching reference here.

## Portability

| Part | Applies to any theme | Notes for other themes |
|---|---|---|
| Method (companion skill, fix, verify, record) | yes | |
| Tokens, z-index scale, focus ring, 16px inputs, type scale | yes | Put them where the theme keeps its variables. In a builder theme, use the builder's global colours and typography first and add only what it cannot express |
| Computed-style diff, `elementFromPoint` overlap check, width sweep | yes | |
| Font files, preload (`typography.md`) | yes | The preload file path is Hummingbird's mechanism. Elsewhere add the `<link rel="preload">` in the head template or hook |
| Autocomplete tokens, birthday mask | yes | Match the form markup of the theme (field `name`s are core, wrappers are not) |
| Drawer, mini cart, loader, 404, cookie card | recipes | Needs the theme's own markup and hooks. Offer them, do not impose them |
| Mega menu, nav fit, sticky aside z-index bug | recipe plus idea | The bug (a sticky panel above the header) exists wherever the theme sticks a summary column. Find it with `elementFromPoint` |
| One-page checkout styling | PS 9.2+ only | Use `ps-one-page-checkout` |

On a builder theme (Elementor-style editors such as ElementFlow) header, footer,
home and product page are built in the editor: run the principle passes against
the theme CSS, add classes to blocks in the editor instead of selecting generated
ids, and do not copy drawers or the mini cart unless asked.

When a project teaches something new, add it to `references/<theme>.md`, and move
it into the shared files only when it held on a second theme.

The skill grew from three passes on one shop: color, surfaces, forms, typography, motion, then a review. Expect to add a `references/<theme>.md` on the next project.

## Order of work in a new theme

1. Tokens first (`references/tokens-surfaces.md`). Every later rule reads them.
2. Buttons and forms (`references/forms.md`): one `.btn` system, one focus ring.
3. Header layering (`references/header-layers.md`) before building menus or drawers.
4. Drawer component (`references/drawer.md`), then everything that slides in
   (filters, size guide, category text, mini cart) reuses it.
5. Typography (`references/typography.md`): font files and the role scale.
6. Motion (`references/motion.md`) as a last pass, on frequent actions only.
7. Review (`references/ui-review.md`): audit every page, verify, then fix.
8. Page-level patterns as needed (`references/patterns.md`).
9. App-feel details, once the base is solid: scrollbars and fades
   (`references/scrollbars.md`), toasts and module modals
   (`references/toasts-modals.md`), the Lottie wishlist like
   (`references/lottie-like.md`, recolour with `scripts/lottie_recolor.py`).
   Hover that glides between menu items or cards: `references/fluid-hover.md`.
   When the user points at another site ("como en X"), study it first with rea
   (`references/replicate-from-web.md`).

## Rules that paid off

- Never hard-code a radius, a control height, a grey or a z-index in a component
  rule. Add the token once, use it everywhere. Two radii only: controls and surfaces.
- Colour, surface and form principles: follow `emil-color`, `emil-surfaces` and
  `emil-forms-and-inputs`. Project outcome: oklch tokens, shadow stacks instead of
  card borders, 16px inputs, one focus ring token.
- Motion: only frequent-safe feedback (button press) and occasional surfaces
  (drawers, cart badge, row removal), always with a reduced-motion branch.
- Every overlay uses a z-index token from one scale. The header must sit above
  Hummingbird's sticky cart and checkout panels (z-index 1020).
- Verify in the browser, not from the CSS: computed styles after a hard reload,
  `document.elementFromPoint` for overlap, a real viewport emulation for mobile.
  Check at 390, 768, 992, 1024, 1200, 1280, 1440 and 1920 px.

## Gotchas (each cost time)

- **No `;` inside a CSS comment** in the theme stylesheet: the theme minifier
  swallows the next rule. Use commas or a full stop.
- **`!important` inside a cascade layer beats `!important` outside it** (order reverses). Change the utility's variable, or write from the same layer (`ui-review.md`).
- **Overflow is invisible on a mobile viewport**: Chrome widens the layout viewport. Compare against the emulated width, not `innerWidth` (`ui-review.md`).
- **Browser keeps the old compiled CSS** under the same hashed name: force it with
  `fetch(url,{cache:'reload'})` and check a new token before measuring (`typography.md`).
- **Compiled asset cache**: after editing `custom.css` or `custom.js`, delete
  `assets/cache/theme-*.css`, `assets/cache/bottom-*.js`, `assets/cache/head-*.js`,
  then `php bin/console cache:clear`. Removing only `theme-*.js` does nothing: the JS bundle is `bottom-<hash>.js`.
  The browser needs a hard reload too (DevTools: reload with cache ignored).
- **Utilities layer**: Bootstrap utilities sit in `@layer utilities`, and module
  rules in `@layer ps-modules` are `!important`. To beat them from the child
  stylesheet use `!important` or a more specific selector from the same layer.
- **A module's heart/modals do nothing**: the custom `footer.tpl` must call
  `{hook h='displayFooter' mod='blockwishlist'}` or the module's Vue apps never mount.
- **Vue edits text in place**: a MutationObserver that only watches `childList`
  misses it. Add `characterData:true` (`lottie-like.md`).
- **Chrome ignores `::-webkit-scrollbar` when `scrollbar-width` or `scrollbar-color`
  is set on the same element** (`scrollbars.md`).
- **Third-party animations**: use LottieFiles "Free" (Lottie Simple License), never
  an app's own asset (`lottie-like.md`).
- **A named `view-transition-name` creates a stacking context**: never name
  `#header` itself (see `ps-view-transitions`).
- **Re-rendered DOM**: facets, cart and checkout replace nodes by AJAX. Delegate
  events on `document` or use a `MutationObserver` on a stable parent, and
  re-apply attributes (autocomplete tokens) after each render.
- **`:focus-visible` is not reliable in headless checks**: a field can be
  `disabled` (checkout address before the email step) and then never takes focus.
  Test on an enabled field and wait for the 150ms transition before reading styles.
- Bootstrap's default focus ring is blue (`rgba(11,105,246,.25)`), off brand on
  most projects: override on `:focus`, not only `:focus-visible`.

## Verifying a pass (checklist)

- [ ] Tokens only: grep the new rules for `px` radius, hex greys, raw z-index.
- [ ] Mobile 390 and 320, tablet 768, laptop 1024 and 1280 look right.
- [ ] Overlays: mega menu over cart aside, drawers over header, mobile menu over
      the WhatsApp button, sticky add-to-cart bar not covered.
- [ ] Reduced-motion branch for every new animation.
- [ ] Focus ring visible on search, newsletter, login, checkout, select.
- [ ] No horizontal scroll, no input smaller than 16px on touch.
