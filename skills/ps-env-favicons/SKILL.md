---
name: ps-env-favicons
description: >
  Favicons distintos por entorno (local, PRE, producción) en un tema hijo de
  PrestaShop, sin romper la subida nativa de favicon del BO. Genera insignias de
  color con la marca del cliente (Python + Pillow), sobrescribe el bloque
  head_icons de _partials/head.tpl del tema hijo y decide por HTTP_HOST: local y
  PRE pintan su insignia; producción y cualquier host desconocido usan
  {$shop.favicon} (el que se sube en Diseño › Tema y logotipo). Incluye cómo
  cargar el favicon de producción "a la nativa" por SSH en un servidor que luego
  pasará a producción. Usar cuando el usuario pida favicon por entorno, "distinguir
  local/pre/prod en la pestaña", "favicon de PrestaShop", "subir el favicon en
  PRE", o al montar un proyecto PS 8/9 con tema hijo Hummingbird/Classic. Es la
  versión PrestaShop del patrón env-favicons (que es build-time para Vite/Next).
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# PS Env Favicons — favicon por entorno en PrestaShop

Montado y probado en `ps8-animalmax` (PS 8.2, Hummingbird 1.0.1 + `animalmax-child`).
La skill `env-favicons` (Vite/Next) decide el entorno en **build**; PrestaShop no
tiene build de plantillas, así que aquí se decide en **render** (Smarty) por host.

## Principio

| Entorno | Host (ejemplo) | Favicon |
|---|---|---|
| Local | `*.lndo.site`, `localhost` | insignia verde (`favicon-local.png`) |
| PRE | `*.pre-testing.com` | insignia dorada (`favicon-pre.png`) |
| Producción / host desconocido | cualquier otro | **`{$shop.favicon}`**, el nativo del BO |

- El color va en el **fondo** de la insignia; la marca siempre blanca (se lee a 16 px).
  Tonos bien separados en matiz y distintos del color de marca de producción.
- Producción **no se toca**: sigue usando el favicon que se sube en el BO. Por eso
  "subir el favicon de PrestaShop" siempre funciona. Un host nuevo cae a producción
  (fallo seguro: nunca enseña una insignia de dev en prod).
- En el servidor que *será* producción, el favicon de prod se carga ya por la vía
  nativa (ver paso 4) y se verá solo cuando cambie el dominio.

## 1 · Generar los assets

Script reutilizable: `scripts/gen-favicons.py` (dentro de esta skill). Copiarlo a
`themes/<child>/scripts/`, ajustar `COLORS` a la paleta del cliente y ejecutar:

```bash
uv run --with pillow python themes/<child>/scripts/gen-favicons.py <logo-completo.png>
```

- Con un logo completo (marca + texto) **recorta la marca** (primer bloque ancho de
  columnas con tinta), elimina el ® y motas por componentes conexos y guarda
  `_dev/brand/<mark>.png`. Sin argumento reutiliza esa marca.
- Salida en `themes/<child>/assets/img/favicons/`:
  `favicon-local.png`, `favicon-pre.png` (64 px) y `favicon-production.ico`
  (multi-tamaño 16–64, para el BO).
- La marca puede ser apaisada (p. ej. "AM"): el script ajusta por **ancho** (~74 %).
  Revisar siempre el PNG recortado antes de generar; si el logo no tiene huecos entre
  marca y texto, recortar a mano.
- Pillow no suele estar en el sistema: `uv run --with pillow` evita instalarlo.

## 2 · Sobrescribir `head_icons` en el tema hijo

Hummingbird define el favicon en el bloque `head_icons` de `_partials/head.tpl`, pero
ese partial se **incluye**, no se extiende: no se puede sobrescribir solo el bloque.
Copiar `themes/hummingbird/templates/_partials/head.tpl` a
`themes/<child>/templates/_partials/head.tpl` y sustituir el bloque:

```smarty
{block name='head_icons'}
  {assign var='amHost' value=$smarty.server.HTTP_HOST|default:''|lower}
  {assign var='amEnvFav' value=''}
  {if strpos($amHost, '.lndo.site') !== false || strpos($amHost, 'localhost') !== false}
    {assign var='amEnvFav' value='local'}
  {elseif strpos($amHost, 'pre-testing.com') !== false}
    {assign var='amEnvFav' value='pre'}
  {/if}
  {if $amEnvFav}
    <link rel="icon" type="image/png" sizes="64x64" href="{$urls.base_url}themes/<child>/assets/img/favicons/favicon-{$amEnvFav}.png">
    <link rel="shortcut icon" type="image/png" href="{$urls.base_url}themes/<child>/assets/img/favicons/favicon-{$amEnvFav}.png">
  {else}
    <link rel="icon" type="image/vnd.microsoft.icon" href="{$shop.favicon}?{$shop.favicon_update_time}">
    <link rel="shortcut icon" type="image/x-icon" href="{$shop.favicon}?{$shop.favicon_update_time}">
  {/if}
{/block}
```

Ajustar los dominios de local/PRE de cada proyecto.

## 3 · Verificar (sin editar la plantilla para "simular prod")

- `lando php bin/console cache:clear` (o el equivalente) y mirar los `<link rel="icon">`
  de la página.
- **No** probar la rama de producción editando temporalmente el `.tpl`. Renderizar el
  mismo bloque con hosts distintos desde un script PHP:

```php
require __DIR__.'/config/config.inc.php';
preg_match("/\{block name='head_icons'\}(.*?)\{\/block\}/s", file_get_contents('<head.tpl>'), $m);
foreach (['x.lndo.site','localhost','x.pre-testing.com','tienda.cliente.es',''] as $h) {
    $_SERVER['HTTP_HOST'] = $h; $s = Context::getContext()->smarty;
    $s->assign(['urls'=>['base_url'=>'https://X/'],'shop'=>['favicon'=>'/img/favicon.ico','favicon_update_time'=>'1']]);
    echo str_pad($h ?: '(vacío)', 28), preg_match('/<link rel="icon"[^>]*href="([^"]+)"/', $s->fetch('string:'.$m[1]), $o) ? $o[1] : 'SIN ICON', "\n";
}
```

- Comprobar que el PNG responde 200 en cada entorno.

## 4 · Cargar el favicon de producción "a la nativa" en un servidor

Equivale a subirlo en el BO (Diseño › Tema y logotipo › Favicon). Replica
`PrestaShop\PrestaShop\Core\Shop\LogoUploader::updateFavicon()`:

```bash
ssh <alias> 'cd <webroot> && src=themes/<child>/assets/img/favicons/favicon-production.ico \
  && cp "$src" img/favicon.ico && cp "$src" img/favicon-<id_shop>.ico'
```

y por PHP (config.inc.php cargado):

```php
$shopId = (int) Context::getContext()->shop->id;
Configuration::updateValue('PS_FAVICON', 'favicon-'.$shopId.'.ico');
Configuration::updateGlobalValue('PS_FAVICON', 'favicon.ico');   // el flujo nativo termina así
Configuration::updateValue('PS_IMG_UPDATE_TIME', time());        // cache-bust de {$shop.favicon_update_time}
```

Después limpiar caché (prod y dev). Si se prefiere, subir el `.ico` a mano en el BO.
Ojo: mirar antes `ls -la img/favicon*` (el de serie pesa ~34 KB) y que el servidor sea
el correcto (en servidores compartidos comprobar `git remote get-url origin`).

## Gotchas

- **`$urls.theme_assets` apunta al tema padre** (Hummingbird) aunque el activo sea el
  hijo: el favicon da 404. Usar `{$urls.base_url}themes/<child>/assets/...`.
- **PRE enseña siempre su insignia**, aunque se suba otro favicon en el BO: el favicon
  del BO solo se ve en hosts "de producción". Es lo esperado.
- **Un servidor PRE que pasará a producción**: cargar el favicon de prod nativo ya (paso
  4); al cambiar el dominio, `head.tpl` deja de pintar la insignia y usa `$shop.favicon`.
- **Los navegadores cachean el favicon muy fuerte**: ver el cambio con recarga forzada
  o ventana privada.
- **`strpos` en `{if}` funciona** en el Smarty de PrestaShop 8 (no hay security policy);
  si un proyecto la activa, mover la detección a un hook/módulo que asigne la variable.
- **`head.tpl` copiado diverge del padre**: con padre fijado (Hummingbird 1.0.1) es poco
  riesgo; si se actualiza el padre, revisar el diff del `head.tpl`.
- **`apply-*` (Laravel) u otras webs del cliente** tienen su propio favicon; este patrón
  es solo para la tienda PrestaShop.
- `HTTP_HOST` solo se compara con `strpos`, nunca se imprime: no hay XSS por cabecera Host.
