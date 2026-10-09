---
name: ps-view-transitions
description: >
  Adds View Transitions to a PrestaShop child theme: soft fade between pages,
  animated product-list swap when filtering/sorting/paginating (facets AJAX),
  and the shared-element morph from a product tile image to the product page
  photo and back. Pure CSS plus small vanilla JS, no libraries. Use when the
  user asks for view transitions, page transitions, "transiciones entre
  paginas", animated filters, "que los productos no salten al filtrar", a
  miniatura-to-ficha image animation, or a more polished/app-like storefront
  UX on a PS 8/9 theme (Hummingbird, Classic, Panda, ElementFlow children).
  Also use to debug a transition that silently does nothing or aborts.
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# PS View Transitions

Montado y probado en `ps9-espaciopiessanos` (PS 9.2, child de Hummingbird, Chrome 154).
Tres capas independientes: puedes aplicar solo las que quieras.

## Antes de empezar

1. Confirma que el tema usa `core.js` de PS para facetas (emite `updateProductList`). Hummingbird y Classic sí; temas con listado propio (Panda/ElementFlow con filtros de módulo) **no**: capa 2 no aplica tal cual.
2. Localiza y apunta estos selectores del proyecto (los de abajo son los de Hummingbird):
   - lista: `#js-product-list` (en Hummingbird `Theme.selectors.listing.list`)
   - tarjeta: `.js-product-miniature` con `data-id-product`
   - imagen de tarjeta: `.product-miniature__image`
   - foto principal de ficha: `.page-product .product__carousel .carousel-item.active img`
   - header sticky: `#header`
3. Respeta `prefers-reduced-motion` en todo.
4. Dónde va el código: CSS en el `custom.css` del child (o su partial si hay build, ver `ps-css-build`), JS en `assets/js/custom.js`. **No toques el core ni el padre.**

## Capa 1: entre páginas (CSS, sin JS)

```css
@view-transition{navigation:auto}
#header{view-transition-name:site-header}
::view-transition-group(site-header){z-index:1}
::view-transition-old(root),::view-transition-new(root){animation-duration:.18s}
@media (prefers-reduced-motion:reduce){
  @view-transition{navigation:none}
  ::view-transition-group(*),::view-transition-old(*),::view-transition-new(*){animation:none!important}
}
```

El header con nombre propio evita que se desvanezca con la página. Sin soporte del navegador no pasa nada: navegación normal.

## Capa 2: filtrar/ordenar/paginar sin saltos (JS + CSS)

El core hace el swap del listado de forma **síncrona** dentro del evento `updateProductList`. Se envuelve `prestashop.emit`:

```js
(function () {
  if (!document.startViewTransition || matchMedia('(prefers-reduced-motion: reduce)').matches) return;
  function init() {
    var p = window.prestashop;
    if (!p || typeof p.emit !== 'function' || p.__vt) return;
    p.__vt = true;
    var listSel = (window.Theme && Theme.selectors && Theme.selectors.listing && Theme.selectors.listing.list) || '#js-product-list';
    var emit = p.emit;
    var tag = function () {
      document.querySelectorAll(listSel + ' .js-product-miniature').forEach(function (el) {
        el.style.viewTransitionName = 'p-' + el.getAttribute('data-id-product');
      });
    };
    p.emit = function (ev) {
      var self = this, args = arguments;
      if (ev !== 'updateProductList' || !document.querySelector(listSel)) return emit.apply(self, args);
      var root = document.documentElement;
      root.classList.add('vt-list');
      tag();
      var t = document.startViewTransition(function () { emit.apply(self, args); tag(); });
      var done = function () { root.classList.remove('vt-list'); };
      t.finished.then(done, done);
      return self;
    };
  }
  if (document.readyState === 'loading') document.addEventListener('DOMContentLoaded', init);
  else init();
})();
```

```css
/* solo se animan las tarjetas, el resto de la pagina queda quieto */
.vt-list::view-transition-old(root),.vt-list::view-transition-new(root){animation:none}
.vt-list::view-transition-group(*){animation-duration:.3s;animation-timing-function:cubic-bezier(.22,1,.36,1)}
```

- Los nombres de las tarjetas se ponen **por JS y solo dentro del listado**, no en la plantilla: si el mismo producto sale dos veces en una página (carrusel en home), el nombre duplicado aborta la transición.
- Si hay **drawer + velo** de filtros, ver Gotchas 2 y 3: hay que nombrarlos y fijar su z-index.

## Capa 3: miniatura ↔ ficha (CSS + JS + script en head)

**Ir a la ficha.** Al salir, se nombra la imagen de la tarjeta pulsada (`pageswap`); la ficha nombra su foto **por CSS**:

```js
(function () {
  if (!('PageSwapEvent' in window) || matchMedia('(prefers-reduced-motion: reduce)').matches) return;
  var clicked = null;
  document.addEventListener('click', function (e) {
    var link = e.target.closest('.js-product-miniature a[href]');
    clicked = link && !(e.metaKey || e.ctrlKey || e.shiftKey || e.button) ? link.closest('.js-product-miniature') : null;
  }, true);
  addEventListener('pageswap', function (e) {
    if (!e.viewTransition || !clicked) return;
    var img = clicked.querySelector('.product-miniature__image');
    clicked = null;
    if (!img) return;
    img.style.viewTransitionName = 'product-image';
    e.viewTransition.finished.finally(function () { img.style.viewTransitionName = ''; });
  });
})();
```

```css
.page-product .product__carousel .carousel-item.active img{view-transition-name:product-image}
::view-transition-group(product-image){animation-duration:.4s;animation-timing-function:cubic-bezier(.22,1,.36,1)}
::view-transition-old(product-image),::view-transition-new(product-image){height:100%;object-fit:cover}
```

`object-fit:cover` evita deformar la foto cuando la tarjeta (cuadrada) y la ficha (vertical) tienen proporciones distintas.

**Volver al listado.** Script **inline en `templates/_partials/head.tpl`** del child (dentro de `{literal}`), porque `pagereveal` se dispara en el primer render, antes de los scripts de abajo:

```js
(function () {
  if (!('PageRevealEvent' in window) || matchMedia('(prefers-reduced-motion: reduce)').matches) return;
  var KEY = 'vtTile';
  document.addEventListener('click', function (e) {
    var tile = e.target.closest('.js-product-miniature');
    if (tile && e.target.closest('a[href]')) sessionStorage.setItem(KEY, tile.getAttribute('data-id-product'));
  }, true);
  addEventListener('pagereveal', function (e) {
    var id = sessionStorage.getItem(KEY);
    var nav = performance.getEntriesByType('navigation')[0];
    if (!e.viewTransition || !id || !nav || nav.type !== 'back_forward') return;
    sessionStorage.removeItem(KEY);
    var img = document.querySelector('#js-product-list .js-product-miniature[data-id-product="' + id + '"] .product-miniature__image');
    if (!img) return;
    img.style.viewTransitionName = 'product-image';
    e.viewTransition.finished.finally(function () { img.style.viewTransitionName = ''; });
  });
})();
```

La clave solo se consume en la rama `back_forward`; si se borrara antes, la propia llegada a la ficha la gastaría.

## Gotchas (cada uno costó tiempo)

1. **Dos elementos con el mismo `view-transition-name` abortan TODA la transición**, sin error visible (`InvalidStateError: Snapshot capture failed` en `transition.ready`). Ojo con elementos ocultos: en Hummingbird conviven el drawer de escritorio y el offcanvas móvil.
2. **Orden de capas:** los grupos que solo existen en el estado nuevo (tarjetas que entran) se pintan **encima** de los antiguos. Si hay velo/drawer, dales nombre (solo al que está abierto: `.drawer.is-open`, `.offcanvas.show`, y el velo) y `z-index` explícito en `::view-transition-group(...)`: header 1, velo 2, drawer 3.
3. **`pagereveal` va antes que los scripts del final del body.** Todo lo que necesite la página nueva en el primer render: CSS o script inline en `<head>`.
4. **No metas assets nuevos por `theme.yml`** si el tema ya está instalado: PS cachea su config en `config/themes/<tema>/shop1.json` (gitignored) y no surte efecto hasta resetear el tema. Por eso el script del head va inline en `head.tpl`.
5. **El fichero compilado CCC conserva el nombre** (`theme-xxxx.css`) aunque cambie su contenido: recarga dura al probar. Borra `themes/<child>/assets/cache/*` y carga una página en el navegador para regenerar (con `curl` no se regenera).
6. **La transición entre documentos tiene tope de ~4 s:** si la página nueva tarda más en renderizar, se salta. Local lento (Lando) puede enmascarar el efecto.
7. **Las navegaciones por CDP/automatización no disparan la transición entre documentos.** Pruébalo con clics reales (o `a.click()` desde la página).
8. `lando php bin/console cache:clear` se queda sin memoria en el warmup: `lando php -d memory_limit=2G bin/console cache:clear --no-warmup` (ambos entornos).

## Cómo probar

- **Listado:** hacer clic en un filtro; `document.startViewTransition` envuelto debe registrar 1 llamada, `ready` y `finished` sin error, y `document.documentElement.className` debe volver a estar sin `vt-list`.
- **Capas:** para ver un fotograma intermedio, alargar la duración (`::view-transition-group(*),::view-transition-old(*),::view-transition-new(*){animation-duration:8s!important}`) y capturar a mitad. Con el drawer abierto, ninguna tarjeta debe verse por encima del velo.
- **Nombres duplicados:** con la clase `vt-list` puesta, listar los elementos con `getComputedStyle(el).viewTransitionName !== 'none'` y comprobar que no hay repetidos.
- **Miniatura → ficha:** clic real en una tarjeta, comprobar que existe exactamente un elemento con `product-image` en la ficha. Volver con `history.back()` y comprobar que la clave `vtTile` desaparece.

## Costes a tener en cuenta (no es gratis del todo)

- Navegadores sin soporte: se ve la navegación normal, sin efecto.
- Cada elemento con nombre es una capa: con listados largos (>40 tarjetas con nombre) vigila memoria; limita a las visibles si hace falta.
- Mantenimiento por proyecto: los selectores y los drawers cambian. Revisar la sección "Antes de empezar" en cada tema.
