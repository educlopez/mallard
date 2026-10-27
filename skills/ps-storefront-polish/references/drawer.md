# Shared side drawer (one component, many uses)

Reference: `assets/css/custom.css` "Pies drawer" block and `assets/js/custom.js`
"Pies drawer". Used for the size guide, the filters, the category description
"Read more" and the mini cart. One look, one motion, one JS API.

## Markup contract

```html
<aside class="pies-drawer pies-drawer--right" id="x" data-pies-drawer aria-hidden="true"
       role="dialog" aria-modal="true" aria-labelledby="x-title" tabindex="-1">
  <header class="pies-drawer__header">
    <h2 class="pies-drawer__title" id="x-title">Title</h2>
    <button type="button" class="btn-close" data-pies-drawer-close aria-label="Cerrar"></button>
  </header>
  <div class="pies-drawer__body">...</div>
  <footer class="pies-drawer__footer"><a class="btn btn-primary" href="#">Action</a></footer>
</aside>

<button type="button" data-pies-drawer-open="#x" aria-expanded="false" aria-controls="x">Open</button>
```

- Opener may carry `data-pies-drawer-mq="(min-width: 768px)"`: ignored when it does not match
  (the desktop filter drawer is a drawer, the mobile one is a Bootstrap offcanvas).
- Esc, outside click and the veil close it. Focus goes into the drawer and returns to the opener.
- Fill content lazily with `data-pies-fill-from="#selector"` on the body: the full text
  lives once in the page and is copied on first open (see `patterns.md`, category description).
- Events `pies:drawer:open` and `pies:drawer:close` bubble from the drawer; listen on `document`.
- A drawer whose parent re-renders by AJAX (mini cart): keep the persistent shell in the footer
  partial and copy the fresh content into it with a MutationObserver.

## Decisions baked in

- Slide-in 280ms, slide-out 200ms (exit faster than enter), curve `cubic-bezier(.32,.72,0,1)`.
- `visibility:hidden` when closed (so hidden drawers are not focusable) with a delayed visibility transition.
- Veil is one shared element, same colour and opacity as Bootstrap's `.offcanvas-backdrop`.
- `html.pies-drawer-lock{overflow:hidden}` while any drawer is open, and the WhatsApp button hides.
- Reduced motion sets the durations to 1ms.
- Capture-phase handler on `window`: Bootstrap's data-api listens on `document` in capture, so this is the only way to pre-empt an offcanvas toggle on the same button.

## CSS (verbatim from the reference theme)

```css
/* ===== Pies drawer: ONE common side-drawer component (JS: custom.js "Pies drawer") =====
   Single source of truth for look + motion. Used by: size guide (.pies-drawer), desktop filters (#search_filters_wrapper,
   .pies-drawer) and the mobile filter offcanvas (Bootstrap .offcanvas, same variables, no Pies JS). */
:root{
--pies-drawer-width:400px;--pies-drawer-max-width:92vw;
--pies-drawer-duration:280ms;--pies-drawer-duration-out:200ms;--pies-drawer-ease:cubic-bezier(.32,.72,0,1);
--pies-drawer-z:1045;--pies-drawer-veil-z:1040;
--pies-drawer-bg:#fff;--pies-drawer-border:var(--pies-border-soft);--pies-drawer-pad:1.25rem;--pies-drawer-header-h:57px;
--pies-drawer-veil-color:var(--pies-primary);--pies-drawer-veil-opacity:.3;
--pies-drawer-shadow:-8px 0 24px oklch(0 0 0 / .08);
}
.pies-drawer,.pies-filter-drawer{width:min(var(--pies-drawer-width),var(--pies-drawer-max-width));background:var(--pies-drawer-bg);border-radius:0;box-shadow:var(--pies-drawer-shadow);outline:0;
  --bs-offcanvas-width:min(var(--pies-drawer-width),var(--pies-drawer-max-width));--bs-offcanvas-padding-x:var(--pies-drawer-pad);--bs-offcanvas-padding-y:1rem;
  --bs-offcanvas-transition:transform var(--pies-drawer-duration) var(--pies-drawer-ease);--bs-offcanvas-zindex:var(--pies-drawer-z)}
.pies-drawer{position:fixed;top:0;bottom:0;z-index:var(--pies-drawer-z);display:flex;flex-direction:column;visibility:hidden;border:0 solid var(--pies-drawer-border);
  transition:transform var(--pies-drawer-duration-out) var(--pies-drawer-ease),visibility 0s linear var(--pies-drawer-duration-out)}
.pies-drawer--right{right:0;left:auto;border-left-width:1px;transform:translateX(100%)}
.pies-drawer--left{left:0;right:auto;border-right-width:1px;box-shadow:8px 0 24px oklch(0 0 0 / .08);transform:translateX(-100%)}
.pies-drawer.is-open{visibility:visible;transform:none;transition:transform var(--pies-drawer-duration) var(--pies-drawer-ease),visibility 0s}
.pies-drawer__header,.pies-filter-drawer .offcanvas-header{display:flex;flex:none;align-items:center;justify-content:space-between;box-sizing:border-box;height:var(--pies-drawer-header-h);padding:0 var(--pies-drawer-pad);border-bottom:var(--pies-hairline) solid var(--pies-drawer-border)}
.pies-drawer__title,.pies-filter-drawer .offcanvas-title{margin:0;font-size:1rem;font-weight:700;text-transform:uppercase;letter-spacing:.04em;color:var(--pies-primary)}
.pies-drawer__body,.pies-filter-drawer .offcanvas-body{flex:1 1 auto;min-height:0;overflow-y:auto;padding:var(--pies-drawer-pad)}
.pies-drawer__footer{flex:none;display:flex;gap:.5rem;padding:.875rem var(--pies-drawer-pad);background:var(--pies-drawer-bg);border-top:var(--pies-hairline) solid var(--pies-drawer-border)}
.pies-drawer__footer .btn{flex:1 1 0}
.pies-drawer__footer .btn-primary{background:var(--pies-primary);border-color:var(--pies-primary);color:#fff}
.pies-drawer__footer .btn-primary:hover{background:var(--pies-sage);border-color:var(--pies-sage);color:var(--pies-primary)}
.pies-drawer__footer .btn-outline-primary{background:#fff;border-color:var(--pies-primary);color:var(--pies-primary)}
.pies-drawer__footer .btn-outline-primary:hover{background:var(--pies-primary);color:#fff}
/* one global veil (created by JS) + Bootstrap's backdrop for the mobile offcanvas: same colour/opacity */
.pies-drawer-veil{position:fixed;inset:0;z-index:var(--pies-drawer-veil-z);background:var(--pies-drawer-veil-color);opacity:0;visibility:hidden;transition:opacity var(--pies-drawer-duration-out) var(--pies-drawer-ease),visibility 0s linear var(--pies-drawer-duration-out)}
.pies-drawer-veil.is-open{opacity:var(--pies-drawer-veil-opacity);visibility:visible;transition:opacity var(--pies-drawer-duration) var(--pies-drawer-ease)}
.offcanvas-backdrop{background-color:var(--pies-drawer-veil-color);--bs-backdrop-zindex:var(--pies-drawer-veil-z);--bs-backdrop-opacity:var(--pies-drawer-veil-opacity)}
.offcanvas-backdrop.show{opacity:var(--pies-drawer-veil-opacity)}
html.pies-drawer-lock{overflow:hidden}
html.pies-drawer-lock .whatsapp,body.pies-mm-open .whatsapp,body:has(.offcanvas.show) .whatsapp{display:none}
@media (prefers-reduced-motion:reduce){:root{--pies-drawer-duration:1ms;--pies-drawer-duration-out:1ms}}
```

## JS (verbatim from the reference theme)

```js
/* ===== Pies drawer: ONE common side-drawer component (design/motion in custom.css "Pies drawer") =====
   Markup contract:  <aside class="pies-drawer pies-drawer--right" id="x" data-pies-drawer aria-hidden="true"
                       role="dialog" aria-modal="true" aria-labelledby="x-title" tabindex="-1"> header/body/footer </aside>
   Openers: [data-pies-drawer-open="#x"] (optional data-pies-drawer-mq="(min-width: 768px)": ignored when it does not match).
   Closers: [data-pies-drawer-close] (inside the drawer, or with a "#id" value). Esc, outside click and the veil close it.
   API: window.PiesDrawer.open(id) / .close(id) / .isOpen(id); events pies:drawer:open|close (bubble from the drawer).
   Openers are handled on `window` in the capture phase: Bootstrap's delegated data-api handlers are capture
   listeners on `document`, so this is the only way to pre-empt e.g. an offcanvas toggle on the same button. */
(function () {
  var FOCUSABLE = 'a[href],button:not([disabled]),input:not([disabled]):not([type="hidden"]),select:not([disabled]),textarea:not([disabled]),[tabindex]:not([tabindex="-1"])';
  var stack = [];
  var veil = null;
  function el(ref) {
    if (!ref) return null;
    if (typeof ref !== 'string') return ref;
    return document.getElementById(ref.replace(/^#/, ''));
  }
  function getVeil() {
    if (!veil) {
      veil = document.createElement('div');
      veil.className = 'pies-drawer-veil';
      veil.setAttribute('data-pies-drawer-veil', '');
      veil.setAttribute('aria-hidden', 'true');
      document.body.appendChild(veil);
    }
    return veil;
  }
  function isOpen(ref) { var d = el(ref); return !!d && d.classList.contains('is-open'); }
  function syncOpeners(d, open) {
    document.querySelectorAll('[data-pies-drawer-open]').forEach(function (o) {
      var t = el(o.getAttribute('data-pies-drawer-open'));
      if (t === d) o.setAttribute('aria-expanded', open ? 'true' : 'false');
    });
  }
  function sync() {
    getVeil().classList.toggle('is-open', stack.length > 0);
    document.documentElement.classList.toggle('pies-drawer-lock', stack.length > 0);
  }
  function emit(d, name) { d.dispatchEvent(new CustomEvent('pies:drawer:' + name, { bubbles: true, detail: { id: d.id, drawer: d } })); }
  function open(ref, opener) {
    var d = el(ref);
    if (!d || isOpen(d)) return;
    d._piesOpener = opener || document.activeElement;
    d.classList.add('is-open');
    d.setAttribute('aria-hidden', 'false');
    stack.push(d);
    sync();
    syncOpeners(d, true);
    /* focus() is ignored while the drawer is still visibility:hidden (it turns visible a frame after is-open is
       added; the very first open can need longer): try after two frames and again when the slide ends */
    function focusInto() {
      if (!isOpen(d) || d.contains(document.activeElement)) return;
      var f = d.querySelector('[data-pies-drawer-autofocus]') || d.querySelector('.pies-drawer__body ' + FOCUSABLE.split(',').join(',.pies-drawer__body ')) || d.querySelector(FOCUSABLE);
      (f || d).focus({ preventScroll: true });
    }
    requestAnimationFrame(function () { requestAnimationFrame(focusInto); });
    d.addEventListener('transitionend', function once(ev) {
      if (ev.target !== d || ev.propertyName !== 'transform') return;
      d.removeEventListener('transitionend', once);
      focusInto();
    });
    emit(d, 'open');
  }
  function close(ref) {
    var d = el(ref);
    if (!d || !isOpen(d)) return;
    var inside = d.contains(document.activeElement);
    d.classList.remove('is-open');
    d.setAttribute('aria-hidden', 'true');
    stack = stack.filter(function (x) { return x !== d; });
    sync();
    syncOpeners(d, false);
    var o = d._piesOpener;
    if (inside && o && document.contains(o) && o.focus) o.focus({ preventScroll: true });
    emit(d, 'close');
  }
  window.addEventListener('click', function (e) {
    var t = e.target;
    if (!t || !t.closest) return;
    var opener = t.closest('[data-pies-drawer-open]');
    if (opener) {
      var mq = opener.getAttribute('data-pies-drawer-mq');
      var d = el(opener.getAttribute('data-pies-drawer-open'));
      if (d && (!mq || window.matchMedia(mq).matches)) {
        e.preventDefault();
        e.stopImmediatePropagation();
        if (isOpen(d)) close(d); else open(d, opener);
        return;
      }
    }
    var closer = t.closest('[data-pies-drawer-close]');
    if (closer) {
      var v = closer.getAttribute('data-pies-drawer-close');
      var target = v ? el(v) : closer.closest('[data-pies-drawer]');
      if (target) close(target);
    }
  }, true);
  document.addEventListener('mousedown', function (e) {
    var top = stack[stack.length - 1];
    if (!top) return;
    if (top.contains(e.target) || (e.target.closest && e.target.closest('[data-pies-drawer-open]'))) return;
    close(top);
  });
  document.addEventListener('keydown', function (e) {
    var top = stack[stack.length - 1];
    if (!top) return;
    if (e.key === 'Escape') { close(top); return; }
    if (e.key !== 'Tab') return;
    var items = [].filter.call(top.querySelectorAll(FOCUSABLE), function (n) { return n.offsetParent !== null; });
    if (!items.length) { e.preventDefault(); top.focus(); return; }
    var first = items[0], last = items[items.length - 1];
    if (e.shiftKey && (document.activeElement === first || document.activeElement === top)) { e.preventDefault(); last.focus(); }
    else if (!e.shiftKey && document.activeElement === last) { e.preventDefault(); first.focus(); }
  });
  /* a drawer bound to a media query (e.g. the desktop filter drawer) must not stay open when it stops applying */
  window.addEventListener('resize', function () {
    stack.slice().forEach(function (d) { if (getComputedStyle(d).display === 'none') close(d); });
  });
  window.PiesDrawer = { open: open, close: close, isOpen: isOpen };
})();
```

## Adapting

- Rename `pies-` and `data-pies-` consistently. The z tokens come from the scale in `header-layers.md`.
- Drawers must sit above the header (1045 vs 1025) and below the cookie card.
- Stack: opening a drawer from inside another is supported (`stack`).
