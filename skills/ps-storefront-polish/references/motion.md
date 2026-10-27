# Motion rules and the snippets that shipped

Reference: `assets/css/custom.css` ("Buttons", "Motion polish", mini cart,
checkout, mobile menu) and `custom.js` mini cart. The method was: list every
candidate, reject by frequency and purpose, ship the few that pass. A shop is a
task UI: people want to buy, not to be entertained.

## Gate

Decide what to animate with `find-animation-opportunities` (frequency, purpose,
speed, function) and polish existing motion with `improve-animations`. Project
curves: `--pies-ease-out` `cubic-bezier(.22,1,.36,1)` for presses and entrances,
`cubic-bezier(.32,.72,0,1)` for drawers. Exit shorter than enter (200 vs 280ms).

## What shipped

1. **Button press** (150ms, `scale(.97)` on `:active`)
```css
:where(.btn,.pies-mini-qty__btn,.pies-mini-line__remove){transition-property:transform,background-color,border-color,color;transition-duration:150ms;transition-timing-function:var(--pies-ease-out)}
.btn:active:not(:disabled):not(.disabled){transform:scale(.97)}
@media (prefers-reduced-motion:reduce){.btn:active{transform:none}}
```
Module buttons that cannot take `.btn` (cookie accept/reject) are listed in the same rule.

2. **Cart badge bump** when the count changes (220ms). The widget node is
replaced on every refresh, so compare the count across renders in JS:
```css
.header-block__badge.is-bumped{animation:pies-badge-bump 220ms var(--pies-ease-out)}
@keyframes pies-badge-bump{0%{transform:scale(1)}45%{transform:scale(1.18)}100%{transform:scale(1)}}
@media (prefers-reduced-motion:reduce){.header-block__badge.is-bumped{animation:none}}
```
```js
function bump() {
  var b = document.querySelector('.blockcart .header-block__badge'); if (!b) return;
  var n = b.textContent.trim();
  if (lastCount !== null && n !== lastCount) {
    b.classList.add('is-bumped');
    b.addEventListener('animationend', function () { b.classList.remove('is-bumped'); }, { once: true });
  }
  lastCount = n;
}
```

3. **Mini-cart row removal** (180ms): the row fades and slides 8px, then the
cart refresh runs, so the list does not jump.
```css
.pies-mini-line.is-removing{opacity:0;transform:translateX(8px);pointer-events:none;transition:opacity 180ms var(--pies-ease-out),transform 180ms var(--pies-ease-out)}
```
JS: add the class, wait 180ms (`new Promise`), then emit the refresh. Quantity
changes only dim the row (`is-loading`).

4. **Checkout sections fade-in** (200ms) when delivery and payment blocks re-render after a choice:
```css
body#checkout .opc-delivery-methods,body#checkout .opc-payment-methods{animation:pies-opc-in 200ms var(--pies-ease-out) both}
@keyframes pies-opc-in{from{opacity:0;transform:translateY(6px)}to{opacity:1;transform:none}}
@keyframes pies-opc-fade{from{opacity:0}to{opacity:1}}
@media (prefers-reduced-motion:reduce){body#checkout .opc-delivery-methods,body#checkout .opc-payment-methods{animation-name:pies-opc-fade}}
```
Not verified in the reference: whether it replays when the carrier changes (the
module may patch the node instead of replacing it).

5. **Mobile menu sheet** on the drawer curve, 280ms in and 200ms out
(`transform:translateX(-100%)` closed). Was plain `ease`.

6. **Category tiles** image zoom: 350ms and 3% instead of 700ms, and only
under `@media (hover:hover) and (pointer:fine)`.

7. **Cookie card** entrance 320ms, `translateY(16px)` to 0 plus opacity (`pies-cookie-in`).

## What was rejected (so nobody adds it back)

- Mega menu item stagger or slide: opened dozens of times a day.
- Product tile hover scale: 50 uses a day, and it fights the tile-to-product morph.
- Search overlay animation: user wants to type now.
- Anything on keyboard shortcuts.

## Reduced motion

Every animation needs a `@media (prefers-reduced-motion:reduce)` branch. For
transforms, set `none`. For the drawers, set the duration tokens to `1ms`
(keeps the open/closed state logic intact). For the page loader and 404 trail,
show the static final frame.
