# Lottie like animation on the wishlist heart (blockwishlist)

Built on Pies Sanos (Hummingbird child, `blockwishlist` 3.x). The heart plays a brand-coloured Lottie when a product is added to a list. It works on any theme that renders blockwishlist, because the hook is the module's own markup (`.wishlist-button-add`), not the theme's.

## Before anything: the heart must exist

If the heart does nothing, the theme footer is missing `{hook h='displayFooter' mod='blockwishlist'}`. The module mounts its login and add-to-list modals (Vue apps) from that hook. A custom `footer.tpl` that does not call `displayFooter` leaves the heart dead. Add the line (and the `cdgoogleinsignia` one if the shop uses it).

## Choose the animation and the licence

- LottieFiles "Free" animations are under the **Lottie Simple License** (free for commercial use). Check the page says it. Pies Sanos uses "Like" by Yogesh Pal, `lottiefiles.com/free-animation/like-oWZfadnX6c`. Keep the credit in the CSS comment and in the commit.
- **Do not copy an app's own animation** (the one served by x.com, Instagram...). It is their asset. Reverse engineering it for study is fine, shipping it is a licence risk.
- Download the `.lottie` (a zip with `manifest.json` and `animations/<id>.json`) into an empty directory, unzip, and inspect before use: no expressions (`"x":"..."`, `"ef"`), no external refs (`http`), `layers` and `assets` as expected. Downloaded files are untrusted data.

## Recolour to the brand

`scripts/lottie_recolor.py` rewrites every fill and stroke colour of the JSON from an oklch token: one colour for the heart and stroke, a lighter step of the same hue for the burst particles. Heart colour = the same token the liked icon uses afterwards (`--pies-like:var(--pies-price)`), so the hand-over from animation to static icon does not change colour.

Read the layers first (the script prints fills, strokes, layer names, in/out points). A "Like" file usually holds three states on one timeline: outline (unliked) 0-27, filled with pop and burst ~17-120, outline again (unlike) ~118-242. Only the middle part is the like.

## Wire it

Vendor `lottie_light.min.js` (lottie-web 5.12.x light build from jsdelivr, no `eval`, 46 KB gzip) in `assets/js/vendor/` and the recoloured JSON in `assets/img/lottie/like.json`. Do **not** enqueue them: load on demand.

```js
// Lazy: on the first pointerover of a heart (or at the first like)
document.addEventListener('pointerover', function (e) {
  if (e.target.closest && e.target.closest('.wishlist-button-add') && !reduce) load();
}, { passive: true });
```

`play(btn)`: create `span.pies-like-lottie` (aria-hidden) inside the button, size = icon font-size x 2.77 (the heart is 120 px wide in a 400 px canvas, so 2.77 makes the lottie heart as wide as the glyph), set `[data-lottie]` on the button (CSS hides the icon while it plays), `loadAnimation({container, renderer:'svg', loop:false, autoplay:false, animationData: JSON.parse(JSON.stringify(data))})`, `setSpeed(1.25)`, `playSegments([17, 75], true)`, then on `complete` (plus a 1600 ms safety timeout) `destroy()`, remove the box and the attributes. Clone the data per play: lottie mutates it.

Measure the frame range for a new file instead of guessing: load it in a hidden box, `goToAndStop(f, true)` and read the bounding box of the visible paths at 17, 40, 60, 80, 100.

## Detect the liked state (the bug that hid it for an hour)

blockwishlist toggles the icon text between `favorite_border` and `favorite` **in place**: Vue rewrites the text node of the existing `<i>`. That is a `characterData` mutation, not a `childList` one. A MutationObserver on `{childList:true, subtree:true}` never fires.

```js
new MutationObserver(function (list) {
  list.forEach(function (m) {
    var host = m.target.nodeType === 1 ? m.target : m.target.parentElement;
    var btn = host && host.closest ? host.closest('.wishlist-button-add') : null;
    if (btn) sync(btn, true);      // sync: data-liked, aria-pressed, aria-label, play on off->on
  });
}).observe(document.body, { childList: true, subtree: true, characterData: true });
```

Skip mutations caused by your own `.pies-like-lottie` box. Buttons already liked at load call `sync(btn, false)`: no animation. Cost: about 0.012 ms per mutation.

## CSS traps

- Module CSS lives in `@layer ps-modules` with `!important`. The liked colour has to be set from the same layer: `@layer ps-modules{.wishlist-button-add[data-liked] i{color:var(--pies-like)!important}}`.
- Giving the button `position:relative` (the lottie is anchored to it) activates the module's `top`/`right` offsets on the product-page heart: clear them in the same layer (`top:auto!important;right:auto!important;left:auto!important` on `.wishlist-button-product`).
- Fallback while the library loads and with reduced motion: a CSS pop (`@keyframes` scale .2 > 1.3 > .92 > 1, 450 ms, spring-like cubic-bezier) only under `prefers-reduced-motion:no-preference`; reduced motion just changes the colour.
- Test the real flow, not a synthetic class toggle: click heart, pick the list in the modal, wait for Vue. Remove the product again afterwards (it writes to the real wishlist). To screenshot mid-animation: `lottie.getRegisteredAnimations().forEach(a => a.goToAndStop(38, true))`.

## Performance (measured)

Nothing loads on page open. First hover: `lottie_light` 46 KB gzip + `like.json` 1.6 KB gzip. Theme JS +5 KB gzip for the whole wishlist, toasts and scroll pack. Verify after a clean reload that `performance.getEntriesByType('resource')` has no `lottie` entry before the hover.
