# Toasts (Sonner port) and module modals

Toasts and modals are the "app feel" details. Both were done on Pies Sanos and apply to any PrestaShop 8/9 theme, because they hook the markup the modules already render.

## Toasts: Sonner in vanilla JS

Port of Sonner 1.7 (`sonner.emilkowal.ski`) with no dependency. Styles use Sonner's own attributes (`data-sonner-toaster`, `data-sonner-toast`) so the CSS of the reference can be compared 1:1.

API: `window.piesToast(message, {description, duration, id, action, cancel, closeButton})` plus `.success .error .info .warning .dismiss`. `duration:Infinity` keeps it. Behaviour:

- newest in front, 3 visible, 14px gap, stack expands on hover, touch or focus
- timers pause while expanded and while the tab is hidden
- close button appears on hover, swipe down or right to dismiss (threshold 20px, plus velocity), Alt+T focuses the stack
- bottom-right, 356px, 24px offsets; mobile full width with 16px offsets and lifted above the sticky add-to-cart bar: `body:has(.pies-sticky-atc.is-visible) [data-sonner-toaster]{--mobile-offset-bottom:calc(74px + 16px + env(safe-area-inset-bottom))}`
- `prefers-reduced-motion`: no slide, opacity only
- Time on screen scales with the text (about 4 s base), not a flat 2 s.

Mirror the toasts the platform already raises instead of replacing the module code:

| Source | How |
|---|---|
| `blockwishlist` (Vue) | MutationObserver on `.wishlist-toast` (class `isActive`, types success/error/basic), read its text, call `piesToast`, hide the original with `.wishlist-toast{display:none!important}` |
| Hummingbird `#js-toast-container` | observe additions, mirror, `display:none!important` on the container |

Mount point: one `<section aria-label="Notifications" tabindex="-1">` with an `<ol>` appended to `body` on first use. `role="status"` on each toast and `aria-live="polite"` on the section.

Traps: toasts must sit under modals and drawers in the z-index scale but above the sticky bars; keep the swipe handler passive-safe (`touch-action:none` only on the toast); use CSS transitions, not keyframes, for enter/exit so a toast replaced mid-animation retargets.

## Module modals (blockwishlist and similar)

The module renders them with its own Vue app, so only the **markup classes** can be styled. Do it once for all of them (add to list, create, rename, delete, share, login) so they share one dialog language:

- `.wishlist-modal .modal-content`: radius `--pies-radius-lg` (8px), floating shadow token, `border:0`, `overflow:hidden`; `.modal-dialog{max-width:26rem}`
- 24px gutters everywhere. Bootstrap's `px-*` utilities on the header need `@layer utilities{.wishlist-modal .modal-header{padding-inline:var(--space)!important}}`
- list rows 44px with a hover tint (`--pies-bg-soft`), the "create new list" row in the primary colour (also from `@layer utilities`)
- inputs and buttons at the shared control height (44px), cancel as an outline button, primary as `.btn`
- entry: `@keyframes pies-dialog-in{from{opacity:0;transform:translateY(8px) scale(.97)}to{opacity:1;transform:none}}`, 220ms ease-out, only under `prefers-reduced-motion:no-preference`. **No exit animation**: Vue hides the modal immediately, a CSS exit never runs.

Verify in a real session (logged in): opening the modals writes nothing, but picking a list does. Open, look, cancel.

## Header dropdown under the nav (found while testing)

`.pies-header` and `.pies-nav` shared `z-index:1025`, so the account dropdown opened under the nav. Fix: `.pies-header{z-index:calc(var(--pies-z-nav) + 1)}`. Find this kind of bug with `document.elementFromPoint(x, y)` over the open menu, not by eye.
