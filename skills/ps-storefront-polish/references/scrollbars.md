# Scrollbars, scroll areas and scroll fades

After `fluidfunctionalism.com/docs/scrollbars` (read the whole reference page, the details are in the bar states). Implemented in `custom.css` "Scrollbars and scroll fades" and a 10-line IIFE in `custom.js`. Works on any theme: only the list of scroller selectors changes.

## Three tiers

1. **Baseline, every scroller, fine pointers only.** Thin bar, 8% ink thumb. The page scrollbar (`html`) stays native. Touch keeps the native overlay bar.
2. **Main scrollers: the bar shows only on hover or just after scrolling.** Drawers, mega-menu panels, facet lists, gallery thumbnails, lightbox thumbs, any `overflow:auto` panel. Same behaviour as a ScrollArea component.
3. **Edge fades** where there is more content.

```css
:root{--pies-ink:0 0 0;--scroll-fade-size:48px}
@media (pointer:fine){
  *:not(html){scrollbar-width:thin;scrollbar-color:rgb(var(--pies-ink) / .08) transparent}
  :is(SCROLLERS){scrollbar-width:auto;scrollbar-color:auto}      /* tier 2 resets both, see trap 1 */
  :is(SCROLLERS)::-webkit-scrollbar{width:10px;height:10px}
  :is(SCROLLERS)::-webkit-scrollbar-track,...-corner{background:transparent}
  :is(SCROLLERS)::-webkit-scrollbar-thumb{background-color:transparent;background-clip:content-box;border:3px solid transparent;border-radius:var(--pies-radius-pill);transition:background-color 120ms ease-out 160ms}
  :is(SCROLLERS)::-webkit-scrollbar-thumb:vertical{border-left-width:1px;border-right-width:5px}     /* 10px track, 4px thumb nudged off the edge */
  :is(SCROLLERS):is(:hover,[data-scrolling])::-webkit-scrollbar-thumb{background-color:rgb(var(--pies-ink) / .08);transition-duration:160ms;transition-delay:0s}
  ...thumb:vertical:hover{background-color:rgb(var(--pies-ink) / .12);border-right-width:3px}      /* grows to 6px */
  ...thumb:active{background-color:rgb(var(--pies-ink) / .16)}
  @supports not selector(::-webkit-scrollbar){ /* Firefox: scrollbar-color transparent, 16% ink on hover/scrolling */ }
}
```

Numbers: track 10px, thumb 4px (6px on hover), pill radius, 8 > 12 > 16% ink for rest, hover, press, fade in 160ms, fade out 120ms after a 160ms wait. Make the ink a token (`--pies-ink`) so a dark surface can flip it to white.

`data-scrolling` linger (scroll does not bubble: capture):

```js
document.addEventListener('scroll', function (e) {
  var t = e.target;
  if (!t || t.nodeType !== 1 || !t.matches(SCROLLERS)) return;
  t.setAttribute('data-scrolling', '');
  clearTimeout(t._piesScrollTimer);
  t._piesScrollTimer = setTimeout(function () { t.removeAttribute('data-scrolling'); }, 600);
}, { capture: true, passive: true });
```

## Fades: CSS only, no JS

A `mask-image` gradient whose ends are driven by a scroll timeline. The real start and end stay crisp, the edge that still has content dissolves.

```css
@property --sf-start{syntax:"<number>";inherits:false;initial-value:1}
@property --sf-end{syntax:"<number>";inherits:false;initial-value:1}
@keyframes sf-reveal-start{from{--sf-start:1}to{--sf-start:0}}
@keyframes sf-reveal-end{from{--sf-end:0}to{--sf-end:1}}
@supports (animation-timeline:scroll()){
  .pies-scroll-fade{mask-image:linear-gradient(to bottom,rgb(0 0 0 / var(--sf-start)) 0,#000 var(--scroll-fade-size),#000 calc(100% - var(--scroll-fade-size)),rgb(0 0 0 / var(--sf-end)) 100%);
    animation:sf-reveal-start both linear,sf-reveal-end both linear;animation-timeline:scroll(self),scroll(self);
    animation-range-start:0,calc(100% - var(--scroll-fade-size));animation-range-end:var(--scroll-fade-size),100%}
  .pies-scroll-fade-x{/* same with "to right" and scroll(self inline) */}
  @media (prefers-reduced-motion:reduce){.pies-scroll-fade,.pies-scroll-fade-x{animation:none;mask-image:none}}
}
```

48px for panels, 32px for small ones (thumbnails, product rails). Horizontal rails (`.products` sliders) use the `-x` variant. Opt-in utility classes `.pies-scroll`, `.pies-scroll-fade`, `.pies-scroll-fade-x` for new scrollers.

## Traps

1. **Chrome ignores `::-webkit-scrollbar-*` if `scrollbar-width` or `scrollbar-color` is set on the same element.** The baseline sets them on `*`, so tier 2 must put both back to `auto` or the custom thumb never shows.
2. **Never fade a scroll area whose first or last item must be fully readable** without checking: with the scroll timeline the fade is 0 at the very start and end, so it is safe. A plain gradient overlay is not (it also hides content).
3. **`scrollbar-gutter:stable`** on a scroller whose width is fixed by a grid (gallery thumbs column), plus the track width in the column: otherwise the bar appears on hover and crops the photos or shifts the layout. Pies Sanos: column `calc(5.5rem + 14px)`, list padding 2px.
4. **Shadows and rings cut by `overflow`**: a scroller clips the outer `box-shadow` of its children. Use an inner padding (2px) and draw the active ring as an outer `box-shadow`, or `outline-offset:-2px`.
5. **Headless Chrome hides scrollbars** (overlay), so screenshots from headless lie. Check in a real browser and also with `getComputedStyle(el,'::-webkit-scrollbar-thumb')`.
6. `*:not(html)` selector keeps the page bar native. Do not style `html`/`body` scrollbars (`emil-ui-polish`).
