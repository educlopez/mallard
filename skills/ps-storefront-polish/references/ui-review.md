# UI review pass (find what the eye misses)

Generic method and rule catalog: see `emil-ui-review`. This file holds the audits
that worked on a PrestaShop storefront, the cascade and layout traps they found,
and the fixes. Run it after the building passes, before a merge.

Pages to cover, mobile 390 and desktop 1280: home, header and mega menu, footer,
category (with the filters drawer), product, cart and mini cart, checkout, login
and register, 404. Verify every finding in the browser before reporting it: the
first draft of a report on Pies Sanos had 2 false positives (an image link that
did have an `alt`, and a menu "cut off" that was a stale compiled CSS).

## Static sweep (seconds, no browser)

```bash
grep -n "transition:\s*all" custom.css                 # unbounded property animation
grep -no "[a-z-]*:[^;}]*ease-in[^-;}]*" custom.css     # ease-in on UI (loops are fine)
grep -no "transition:[^;}]*\(width\|height\|margin\|top\|left\|padding\)" custom.css
grep -no "z-index:[0-9]\{3,\}" custom.css              # raw z-index, should be tokens
grep -c "prefers-reduced-motion" custom.css            # and one per animation name
grep -n "matchMedia\|setInterval" custom.js            # autoplay respects reduced motion? pauses?
```

## Browser audits (run per page)

1. **Contrast, computed not guessed.** For each text node: take the computed colour,
   blend it with its alpha, find the effective background by walking ancestors (skip
   elements with a background image), compute the WCAG ratio, flag under 4.5 (3 for
   text of 24px, or 18.66px bold). Convert any colour syntax (oklch included) by
   painting it on a 1px canvas and reading the pixel. This caught an amber warning at
   1.63:1 and a semi-transparent footnote at 3.93:1 that no one had measured.
2. **Targets under 44px.** List visible `a, button, input, select, summary` with a
   box under 44 in either direction, skip links inside running text.
3. **Hit areas, not boxes.** A box audit cannot see a `::before` hit area. Instead
   sample `document.elementFromPoint` at the top and bottom edge of a 44px area
   centred on the element (set `scroll-behavior:auto` first, or `scrollIntoView` is
   asynchronous and every sample lands off screen).
4. **Names.** Interactive elements whose accessible name is empty (aria-label, labelledby,
   text, image alt, `label[for]`). Icon-only widgets injected by modules (a WhatsApp
   chat button) fail here.
5. **Images:** missing `alt`, missing `width`/`height`. **Headings:** one `h1`, no skipped levels.
6. **Horizontal overflow, measured correctly.** On a mobile viewport Chrome widens the
   layout viewport to fit overflowing content, so `scrollWidth > innerWidth` is always
   false. Compare against the emulated width (390), or list non-fixed elements whose
   `getBoundingClientRect().right` exceeds it. This found a 12px overflow that had shipped.
7. **Layout shift:** read `layout-shift` entries (observer with `buffered:true`).

## Findings and their fixes

### Layout and grid
- A Bootstrap `.row` with a big custom gutter (`--bs-gutter-x:48px`) inside a container
  with 12px padding overflows by `(gutter / 2 - 12)px` on phones. Use the standard
  24px gutter below lg.

### Cascade layers
- Bootstrap utilities live in `@layer utilities` and some modules ship `@layer ps-modules`.
  For `!important` declarations the order is REVERSED: a layered `!important` beats an
  unlayered `!important`. Two consequences:
  - A utility such as `.text-warning` cannot be beaten by a more specific selector.
    Override the variable it reads (`--bs-warning-rgb`) instead.
  - A module that sizes a control with `!important` inside a layer (the wishlist heart)
    can only be changed from the same layer: `@layer ps-modules{ ... !important }`, later in the file.

### Touch targets (`@media (pointer:coarse)`, desktop unchanged)
- Text link or summary: an invisible, larger pseudo-element, no layout shift:
  ```css
  a.my-link{position:relative}
  a.my-link::before{content:"";position:absolute;inset:-.625rem -.25rem}
  ```
  Do not use it on elements that already use `::before`/`::after` for something else, and
  do not set `position:relative` on an element that is already absolute (a link pinned
  with `top`).
- Fixed-size controls (stepper buttons, close buttons, trash, share icons, pager): grow
  to `--pies-btn-h`. A stepper whose border eats 2px: draw the border as
  `outline:1px solid var(--border);outline-offset:-1px` (outlines paint above children) so
  the buttons keep the full 44px.
- Carousel dots: 24px hit area with an 8px dot drawn by `radial-gradient`. Nine dots at
  44px do not fit in 390px; 24px is the WCAG 2.5.8 AA minimum.
- Autoplaying content (rotating info bar, hero): stop under reduced motion, pause on
  hover and on focus (WCAG 2.2.2), and give the fade its own reduced-motion rule.

### Small things worth a line each
- Gallery arrows as 44px white round controls with the floating shadow, visible on any photo.
- Password toggle in an `input-group`: restore the square inner corners lost to a global
  `border-radius`, and style it as an outline button, not a dark block.
- A disabled primary button: pale background and muted text (readable), not white on grey.

## Checkout: a slim footer

The full footer (link columns, newsletter, socials) pulls customers away from paying.
Render a short one on the checkout only: help (phone, email, hours), three legal links
(returns, payment methods, privacy), payment badges, copyright, plus any module hook the
shop needs (Google badge). Pies Sanos: 1000px to 327px on mobile, 179px on desktop.

```smarty
{if $page.page_name == 'checkout'}
  <div class="footer pies-footer pies-footer--checkout"> ... {url entity='cms' id=10} ... </div>
{else}
  ...full footer...
{/if}
```

Use `{url entity='cms' id=N}` for CMS pages and check the ids exist on the other
environments. Layout: `display:flex;flex-wrap:wrap;justify-content:space-between` with the
legal links stacked in a column so the three groups fit one row at 1280.

## Data problems are not theme bugs

The review will surface things the theme should not "fix": junk imported data (a product
`unity` field holding a price, printed as "(0,87 € 59.95)"), English copy in a third-party
widget (a reviews widget configured in its own dashboard). Report them with the cause and
who owns them, and leave them alone unless the owner asks. Never run mass `UPDATE`s on a
shop database as part of a polish pass.

## Reporting format that worked

Verdict (Blocked, Approve with changes, Ship it), then Blocking, Should fix and Nits, each
with the measured value (`1.63:1`, `24x24`, `402px`), the selector or file and line, and the
fix with its measured result. Add a "Discarded after checking" list: it shows the review was
verified and saves the next person re-investigating.
