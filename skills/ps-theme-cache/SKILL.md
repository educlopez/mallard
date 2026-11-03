---
name: ps-theme-cache
description: >
  Makes a PrestaShop CSS/JS change actually reach the browser when the theme
  combines assets into themes/<theme>/assets/cache/theme-*.css and bottom-*.js
  (CCC, Panda child themes). Use when a custom.css or theme JS edit is saved
  but the page still looks unstyled, when the stylesheet loaded is
  assets/cache/theme-<hash><n>.css, or right after deploying a theme change to
  PRE/PROD ("el CSS no se actualiza", "no se ve el cambio", "borra la caché del
  tema", "PS_CCCCSS_VERSION", "limpiar caché CSS tras el deploy").
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# PS theme cache — que el CSS/JS nuevo se sirva

## Cuándo ocurre

El tema no sirve `assets/css/custom.css` directamente: lo agrupa en
`themes/<theme>/assets/cache/theme-<hash><version>.css` (y `bottom-*.js`). Si cambias
`custom.css` o un `.tpl` con clases nuevas, el HTML ya sale con el markup nuevo pero el CSS agrupado es el
viejo. Síntoma típico: el elemento aparece **sin estilo** (por ejemplo, un badge como texto plano).

Confirmar en el navegador:
```js
[...document.styleSheets].map(s => s.href).filter(h => h && /assets\/cache/.test(h))
```
El nombre acaba en la versión (`theme-ff491a1214.css` → versión `1214`). Si el estilo nuevo no está en
`getComputedStyle`, es esto.

## Local (Lando)

La versión vive en `{prefix}configuration`, `name = 'PS_CCCCSS_VERSION'`, **una fila por tienda**
(`id_shop` NULL + cada tienda). Subirla genera un fichero nuevo:

```bash
cat > q.sql <<'EOF'
UPDATE {prefix}configuration SET value = value + 1 WHERE name = 'PS_CCCCSS_VERSION';
EOF
lando ssh -s database -c "mysql -uroot <recipe> -e 'source /app/q.sql'"; rm q.sql
lando php bin/console cache:clear
```
Recarga la página (con una pausa: si el JS corre antes de que cargue, `querySelector` devuelve null).

## PRE / PROD

Después de un pull que toque `assets/css/` o `templates/` con clases nuevas, hay que regenerar la caché:

1. **Preferido:** Backoffice → Parámetros avanzados → Rendimiento → **Borrar caché**. En multitienda,
   hazlo en *todas las tiendas* (cada una tiene su versión).
2. Alternativa por servidor: subir `PS_CCCCSS_VERSION` de esa tienda en +1. Es una escritura en la DB del
   servidor: pedir confirmación antes, sobre todo en PROD.

No borres a mano los ficheros de `assets/cache/`: es salida generada y se regenera sola.

## Verificar

Abrir la ficha en el navegador, esperar ~3 s y comprobar que el `href` del CSS cambió de versión y que el
estilo nuevo se aplica (`getComputedStyle(el).backgroundColor`, etc.).

## Gotchas

- `themes/<theme>/assets/cache/` suele aparecer como no trackeado en `git status`: **no lo commitees**.
- La herramienta JS del navegador puede bloquear páginas cuyo URL lleve query string
  (`?controller=product&id_product=…`): usa la URL amigable del producto.
- Si el cambio es solo `.tpl` sin CSS nuevo, basta `bin/console cache:clear` (Smarty); no hace falta tocar CCC.
- Esto no sustituye al pipeline de `ps-css-build` (ese es para proyectos con `_dev/css`); aquí el CSS se edita
  a mano en `custom.css` y solo falta invalidar la caché.
