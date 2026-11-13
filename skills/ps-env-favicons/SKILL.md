---
name: ps-env-favicons
description: >
  Favicons distintos por entorno (local, PRE, producción) en PrestaShop 1.7.7–9
  sin romper la subida nativa de favicon del BO. Incluye un módulo ligero
  (module/envfavicon): un hook de salida sustituye los <link rel="icon"> en local y
  PRE por una insignia de color con la marca del cliente; producción y cualquier
  host desconocido conservan el favicon nativo. Dominios por entorno, colores y logo
  se configuran en el BO (se genera todo desde UN logo con GD y puede cargar el
  favicon de producción como nativo). Alternativa sin módulo: override de head.tpl
  en el tema hijo con script Python. Usar cuando el usuario pida favicon por
  entorno, "distinguir local/pre/prod en la pestaña", "favicon de PrestaShop",
  "subir el favicon en PRE", dominios de PRE que no son pre-testing (ecommercloud,
  etc.), o al montar un proyecto PS con tema hijo. Es la versión PrestaShop del
  patrón env-favicons (que es build-time para Vite/Next).
version: "0.2.0"
metadata:
  author: Eduardo Calvo
---

# PS Env Favicons — favicon por entorno en PrestaShop

Montado y probado en `ps8-animalmax` (PS 8.2, Hummingbird 1.0.1 + `animalmax-child`).
La skill `env-favicons` (Vite/Next) decide el entorno en **build**; PrestaShop no
tiene build de plantillas, así que aquí se decide en **render**, por host.

## Principio

| Entorno | Host | Favicon |
|---|---|---|
| Local | lista del BO (`lndo.site, localhost, .test…`) | insignia verde |
| PRE | lista del BO (`pre-testing.com, ecommercloud…`) | insignia dorada |
| Producción / host desconocido | cualquier otro | **favicon nativo de PrestaShop** |

- El color va en el **fondo**; la marca siempre blanca (se lee a 16 px). Tonos bien
  separados en matiz y distintos del color de marca de producción.
- Producción **no se toca**: sigue el favicon que se sube en el BO (Diseño › Tema y
  logotipo). Un host nuevo cae a producción: nunca se enseña una insignia de dev en prod.
- Los PRE no siempre son `pre-testing.com` (hay `ecommercloud`, subdominios del
  cliente…): por eso los patrones son **configurables**, no van en código.

## Opción A (recomendada) · módulo `envfavicon`

Carpeta `module/envfavicon/` de esta skill. Ligero: un solo hook
(`actionOutputHTMLBefore`) que solo actúa en local/PRE y solo toca los
`<link rel="icon">` de la `<head>` (conserva `apple-touch-icon`). Sin tablas, sin JS,
sin CSS, sin plantillas; no copia `head.tpl` del padre.

```bash
cp -R <skill>/module/envfavicon <proyecto>/modules/envfavicon
lando php bin/console prestashop:module install envfavicon   # o instalar desde el BO
```

En el BO (Módulos › Env Favicon · favicon por entorno):

1. **Dominios LOCAL y PRE**: fragmentos de host separados por coma (se comparan con
   `strpos`; el panel muestra cómo se detecta el host actual).
2. **Colores** de local, PRE y producción.
3. **Logo o marca** (PNG/JPG/WEBP, ≤3 MB): *un solo archivo*. Con «Recortar solo la
   marca» usa el primer bloque a la izquierda y quita el ® y las motas (logos
   "marca + texto"). Guardar regenera `favicon-local.png`, `favicon-pre.png` y
   `favicon-production.ico` (16–64 px, PNG embebido) con GD en ~0,2 s.
4. **Cargar el de producción como favicon nativo**: copia el `.ico` a `img/favicon.ico`
   y `img/favicon-<id_shop>.ico` y fija `PS_FAVICON`/`PS_IMG_UPDATE_TIME` igual que el BO
   (`LogoUploader::updateFavicon`). Útil en un servidor PRE que luego pasará a producción.

Seguridad/info: los dominios viven en `ps_configuration` (BD), **no en ficheros
descargables** del tema; el host solo se compara (nunca se imprime: no hay XSS por
cabecera Host); la subida valida con `ImageManager::validateUpload` y se guarda con nombre
fijo (`source.png|jpg|webp`).

Requisitos: PS ≥ 1.7.7, PHP ≥ 7.2, extensión GD (WEBP opcional).

## Opción B · sin módulo (override de `head.tpl` + script)

Si no se quiere un módulo. Script `scripts/gen-favicons.py` (Pillow vía
`uv run --with pillow`) y override del tema hijo.

1. Generar los assets (ajustar `COLORS`):
   ```bash
   uv run --with pillow python themes/<child>/scripts/gen-favicons.py <logo-completo.png>
   ```
   Salida en `themes/<child>/assets/img/favicons/`: `favicon-local.png`,
   `favicon-pre.png` y `favicon-production.ico` (para subir en el BO).
2. Hummingbird define el favicon en el bloque `head_icons` de `_partials/head.tpl`, pero
   ese partial se **incluye** (no se extiende): copiar `themes/hummingbird/templates/_partials/head.tpl`
   al tema hijo y sustituir el bloque:

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

Inconvenientes frente al módulo: los dominios quedan en un `.tpl` **descargable**
(`/themes/<child>/templates/_partials/head.tpl` responde 200 en PrestaShop), hay que
editar código para cada dominio nuevo y el `head.tpl` copiado diverge del padre.

## Verificar

- Limpiar caché y mirar los `<link rel="icon">` de la página en cada entorno; el PNG
  debe responder 200.
- Probar la rama de producción **sin editar el código**: con el módulo, llamar al hook
  desde un script PHP con `$_SERVER['HTTP_HOST']` distintos y el HTML **por referencia**
  (`$p = ['html' => &$html]`, como hace el core); con la opción B, renderizar el bloque
  `head_icons` con `$smarty->fetch('string:'.$bloque)` y hosts distintos.

## Cargar el favicon nativo por SSH (sin BO)

Equivale a subirlo en el BO; replica `LogoUploader::updateFavicon()`:

```bash
ssh <alias> 'cd <webroot> && cp <favicon-production.ico> img/favicon.ico && cp <favicon-production.ico> img/favicon-<id_shop>.ico'
```
```php
$shopId = (int) Context::getContext()->shop->id;
Configuration::updateValue('PS_FAVICON', 'favicon-'.$shopId.'.ico');
Configuration::updateGlobalValue('PS_FAVICON', 'favicon.ico');   // el flujo nativo termina así
Configuration::updateValue('PS_IMG_UPDATE_TIME', time());        // cache-bust de {$shop.favicon_update_time}
```

Luego limpiar caché. Mirar antes `ls -la img/favicon*` (el de serie pesa ~34 KB) y que el
servidor sea el correcto (en servidores compartidos, `git remote get-url origin`).

## Gotchas

- **El módulo solo sirve el favicon de local/PRE**: en PRE se ve siempre la insignia aunque
  se suba otro favicon en el BO; ese se verá en hosts de producción.
- **Servidor PRE que pasará a producción**: cargar ya el favicon nativo; al cambiar el
  dominio (y no casar con las listas) se usa solo.
- **`$urls.theme_assets` apunta al tema padre** aunque el activo sea el hijo (opción B):
  usar `{$urls.base_url}themes/<child>/assets/...`.
- **Un `html` del hook por valor no cambia nada**: el core lo pasa por referencia; en
  pruebas hay que hacer lo mismo.
- **Los navegadores cachean el favicon muy fuerte**: recarga forzada o ventana privada. El
  módulo añade `?v=<n>` y lo incrementa al regenerar.
- **Si había override de `head.tpl` (opción B) y se instala el módulo**, gana el módulo (el
  hook sustituye los iconos ya emitidos); conviene quitar el override para no duplicar lógica.
- **Las webs `apply-*` (Laravel) u otras** tienen su propio favicon; esto es solo para la
  tienda PrestaShop.
- Los archivos generados (`views/img/source.*`, `favicon-*`) no se versionan: se regeneran
  desde el BO en cada entorno.
