# Fluid hover (a highlight that glides between items)

After `fluidfunctionalism.com/docs/fluid-hover`. Plain `:hover` blinks off and on at every gap between items. A fluid highlight is ONE element that travels from item to item, so the eye stays on the task. Built on Pies Sanos as a vanilla component (`custom.js` "Fluid hover", `custom.css` "Fluid hover") that generalises the header nav pill.

## Rules of the reference

- One highlight per container, moved with a **transform**. Width and height only animate between items of different size.
- A fresh entry **fades in where it lands**, it does not travel from the last place it was.
- `nearest`: the item closest to the pointer is lit, so gaps and padding never go dark. Off, only the item under the pointer lights. Axis says how near is measured: `y` lists and menus, `x` strips and tabs, `xy` grids of cards (distance to the rect).
- A click in a gap goes to the lit item (limit it to a few px).
- Keyboard focus moves the highlight too. Reduced motion: it still fades in but stops sliding.

## Recipe

CSS: container `.is-fluid{position:relative;isolation:isolate}`, highlight `.pies-fluid{position:absolute;left:0;top:0;z-index:0;width:var(--fh-w);height:var(--fh-h);border-radius:var(--fh-r);background:var(--fh-bg);opacity:0;pointer-events:none;transform:translate(var(--fh-x),var(--fh-y));transition:transform .26s <ease-out>,width .2s,height .2s,opacity .15s}` and `.is-on{opacity:1}`. Items get `position:relative;z-index:1`. Reduced motion: only `opacity` transitions.

JS: a list of configs `{ root, item, axis, inset, media, gap }`. One delegated `pointermove` (rAF-throttled, mouse and pen only, only when `(hover:hover) and (pointer:fine)`), find the config whose root contains the target, pick the nearest item, write `--fh-x/y/w/h` from `getBoundingClientRect` relative to the root. `inset` is positive when the highlight grows past the item (the card grid uses 8, so a soft halo shows in the gap, the lists use 0). Hide 80 ms after the pointer leaves the root, on resize, on scroll and when the pointer leaves the window. `focusin` with `:focus-visible` lights the focused item.

When the JS is active the per-item `:hover` background is removed (`.is-fluid .item:hover{background:none}`), but keep the colour change. Without JS the plain hover still works.

## Where it fits (and where not)

- Yes: sidebars and menus, tabs, radio and option lists, card groups, table rows, dropdown lists.
- Not on the current page of a nav (keep its own tint) and not on touch (skip below the pointer media query; the mobile account menu is a row of pills).
- **Not inside lists a framework re-renders** (Vue wishlist, React lists): the highlight is a foreign child and the framework may remove it. Wrap such a list from outside or leave it with `:hover`.

## Traps

- Inset sign: positive grows, negative shrinks. A shrunk highlight inside a white card is invisible.
- A highlight tinted with `--bg-soft` on a white page is almost invisible in screenshots. Use the brand tint.
- Cards that also lift on hover double the effect: when fluid, drop the lift and keep a plain shadow.
- Check with real pointer moves (the `hover` action of the browser tool), and with the pointer in a gap, not on an item.
- The header nav pill predates this component and keeps its own code (it positions against the nav for the mega panels). Do not refactor it without a reason.
