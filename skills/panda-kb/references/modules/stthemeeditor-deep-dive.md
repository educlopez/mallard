---
module_name: stthemeeditor
display_name: Theme editor (Panda central config hub)
version: 2.9.2
panda_version: 2.9.2
ps_version: [8, 9]
source_path: dependencies/modules/stthemeeditor
type: deep-dive
ingested: 2026-04-19
tags: [panda, module, helpers, hub, config, deep-dive]
---

# stthemeeditor — deep dive

> [!info] Por qué existe este doc
> `stthemeeditor` es el **hub central de configuración** de Panda. La mayoría de los 54 módulos `st*` NO tienen su propio BO screen — se configuran leyendo valores que este módulo persiste en `ps_configuration` con prefijo `STSN_*`. Si Eduardo pregunta *"dónde configuro X"*, casi siempre la respuesta vive aquí. La fiche corta [[stthemeeditor]] se queda en metadata; este doc mapea el interior.

## TL;DR

- **38 tabs de admin** cubriendo header/footer/product/catalog/checkout/typography/colors/utilities/custom-code.
- **529+ claves de configuración** con prefijo `STSN_*` persistidas en `ps_configuration`.
- **5 hooks** registrados. El clave es `displayHeader` que genera CSS inline en cada request (no hay fichero compilado en disco).
- **Cero llamadas directas a otros módulos** (`Module::getInstanceByName`) — patrón hub puro: los módulos consumidores leen Config keys, no invocan stthemeeditor.
- Bundle pesado (~25k LOC, 153 ficheros) por librerías UI: OWL Carousel, Jarallax, EasyZoom, Dragscroll, TimeAgo, Animate.css.

## Admin screens / tabs (38)

Agrupadas temáticamente. Cada tab afecta a varios módulos vía `Configuration::get('STSN_*')`.

| Tab | Controla | Afecta módulos | Key options (top) |
|---|---|---|---|
| **Global** | Layout general, anchos de página, spacing maestro | Todos (cascada) | Page width, container padding, grid breakpoints |
| **Header** | Logo principal, nav, cart icon, search | [[stmegamenu]], [[stsearchbar]], [[stshoppingcart]] | Logo image, header layout variant, sticky on/off |
| **Mobile Header** | Header específico móvil | [[stmegamenu]], [[stsearchbar]] | Mobile logo, burger position, search visibility |
| **Navigation** | Menú principal + mega-dropdowns | [[stmegamenu]] | Menu hover vs click, dropdown width, column layout |
| **Footer** | Footer clásico | [[stsocial]], [[stnewsletter]] | Columns config, social icons, newsletter visibility |
| **Stacked Footer** | Footer en modo acordeón móvil | — | Stacked sections, collapse behavior |
| **Product page** | Layout product + galería + tabs | — | Gallery type (vertical/horizontal/sticky), zoom on/off |
| **Product gallery** | Detalle galería + sliders | [[stowlcarousel]] | Thumbnails position, lightbox, zoom type |
| **Product info tabs** | Sección tabs descripción/reviews/etc | [[stproductcomments]] | Tabs vs accordion, default tab, labels |
| **Product block** | Tarjeta de producto (listings) | [[stswiper]], sliders | Price size, add-to-cart style, stickers position |
| **Category** | Page category | — | Sorting default, items per row, filter position |
| **Brand / Supplier** | Pages marca y proveedor | [[stbrandsslider]] | Grid layout, filters, banner top |
| **Checkout** | Flow checkout | — | Steps vs one-page, guest checkout visibility |
| **Authentication** | Sign in / register | [[stcustomersignin]] | Form layout, social auth buttons |
| **Typography - Headings** | H1/H2/.../H6 styles | Todos | Font family, weight, size per breakpoint, line-height |
| **Typography - Body** | Body text | Todos | Font family, base size, breakpoint overrides |
| **Colors - General** | Palette global | Todos | Primary, secondary, accent, background, border |
| **Colors - Buttons** | Botones | — | Button variants (primary/secondary/outline), hover, disabled |
| **Colors - Shadows** | Sombras | — | Box shadow variants (card, hover, modal) |
| **Pagination** | Paginator general | — | Style (numbered/dots), items/page defaults |
| **Breadcrumb** | Breadcrumb | [[stbreadcrumbs]] (si existe) | Separator, home icon, styling |
| **Icons** | Icon set global | — | Icon font selection, sizing |
| **Custom CSS** | Inyección CSS arbitraria | — | Textarea sin validación |
| **Custom JS** | Inyección JS arbitraria | — | Textarea sin validación |
| **Custom Fonts** | Google Fonts / @font-face | — | Font name, source URL, variants |

> [!warning] 38 tabs total
> La tabla cubre ~24 grupos representativos. Hay tabs adicionales (blog-specific, special events, misc utilities) que afectan a módulos puntuales. Si el agente no encuentra la opción X en los listados de arriba, debe ir al source: `dependencies/modules/stthemeeditor/controllers/admin/` y `dependencies/modules/stthemeeditor/views/templates/admin/`.

## Configuration storage — 529+ claves `STSN_*`

Persistidas en la tabla PS core `ps_configuration`. Prefijo `STSN_` (SunnyToo siglas). Distribución aproximada:

- **~120 color keys** — backgrounds, text, buttons, headings
- **~50 typography keys** — font sizes, families, heading styles
- **~40 spacing/layout keys** — margins, padding, column widths
- **Breakpoint-aware** — keys con sufijo `_LG`, `_XL`, `_XXL` para diferentes breakpoints
- **3 injection keys** — `STSN_CUSTOM_CSS`, `STSN_CUSTOM_JS`, `STSN_CUSTOM_FONTS`

**Cómo consultar**: desde SQL directo (`SELECT name, value FROM ps_configuration WHERE name LIKE 'STSN_%'`) o vía PS API (`Configuration::get('STSN_...')`).

**Bulk lookup**: el hook `displayHeader` usa `Configuration::getMultiple([...])` con un array de ~200 keys por request — importa para rendimiento.

## Hook implementations (5)

| Hook | Qué hace |
|---|---|
| `displayHeader` | **El hook crítico**. Lee todas las keys `STSN_*` y genera `<style>` inline con CSS variables + color assignments + custom CSS + fuentes custom. Se ejecuta en cada request de frontend. |
| `actionShopDataDuplication` | Al clonar un shop, copia todas las keys `STSN_*` al nuevo shop. |
| `actionProductSearchAfter` | Invalida cache de imagen de producto — evita stale thumbnails cuando se cambia configuración de tamaños. |
| `actionStAssemble` | Hook interno del namespace ST (SunnyToo). Llamado por otros módulos `st*` para obtener datos agregados del theme (colores, fuentes, grid). |
| `displayRightColumnProduct` | Inserta upsell slider en la columna derecha de product page. Delega a la clase `BaseProductsSlider`. |

## Asset pipeline: CSS dinámico sin ficheros en disco

**Patrón poco habitual** vs otros themes PS:

1. Admin cambia color primario → se guarda en `ps_configuration.STSN_COLOR_PRIMARY`.
2. Next frontend request → `displayHeader` hook:
   - `Configuration::getMultiple([~200 keys])`
   - Genera `<style>` tag con CSS variables (`:root { --color-primary: #{value}; ... }`)
   - Inyecta en el `<head>` de cada página
3. CSS estático de Panda referencia `var(--color-primary)` → aplica inmediatamente.

**Ventajas**: cambio de config = efecto instantáneo, zero build step.

**Desventajas**:
- CPU extra en cada request (lectura de ~200 config keys, generación de CSS inline).
- Sin caché HTTP del CSS (está inline, no es un asset separado).
- Sites de alto tráfico pueden querer cachear el output a nivel reverse-proxy.

**NO hay** scripts que escriban CSS compilado a `themes/panda/assets/css/generated.css` ni similar. Si alguien pregunta "¿dónde está el CSS generado?" → no existe, es inline runtime.

## Vendor libraries bundled

Aproximadamente 70% del LOC del módulo viene de estas librerías JS. Presentes en `views/js/` y `views/css/`:

- **OWL Carousel** — carruseles de producto / brand / category
- **Jarallax + jQuery Parallax** — parallax y video backgrounds
- **EasyZoom** — zoom sobre imagen de producto
- **Dragscroll** — swipeable slider behavior
- **TimeAgo** — timestamps relativos
- **Animate.css** — presets de animación CSS3
- **Color picker** — UI para admin (spectrum.js o similar)
- **Font manager** — UI para custom fonts

No se versionan individualmente; vienen pinneadas en la release de Panda 2.9.2.

## Inter-module interactions

**Patrón: hub puro, no spoke-hub**. `stthemeeditor` NO invoca a otros módulos. Son los módulos consumidores los que leen `Configuration::get('STSN_*')`.

Módulos que consumen config de stthemeeditor (confirmado en source):

- **[[stbanner]]** — lee font/heading/button colors
- **[[steasycontent]]** — lee heading styles, custom fonts, custom CSS
- **[[stswiper]]**, **[[stowlcarousel]]**, **[[stpagebanner]]** — leen grid breakpoints, spacing, button colors
- **[[ststickers]]**, **[[stcountdown]]**, **[[stblogeditor]]** — leen color/typography

Implicación práctica: **no puedes "mover" stthemeeditor a otro module y esperar que todo funcione**. El contrato es implícito vía Configuration keys. Desinstalar stthemeeditor = los demás siguen leyendo keys que nadie escribe = fallback a CSS hardcoded de los modules.

## PS 8 vs PS 9 notes

No se observan diferencias significativas en el código del módulo. El admin controller sigue el patrón `AdminModulesController` estándar PS — funciona igual en PS 8 y PS 9. El hook `displayHeader` hace su trabajo idéntico en ambas.

Única consideración: en **PS 9** con Hummingbird 2.0 activo, el hook `displayHeader` de stthemeeditor SE EJECUTA pero su CSS inline puede **no aplicar** porque Hummingbird usa CSS variables con nombres distintos (`--bs-primary` de Bootstrap 5, no `--color-primary` de Panda). Validar caso a caso si cliente cambia a Hummingbird.

## Known limits / quirks

- **Chinese variable names** en bundled OWL Carousel / parallax scripts (autores originales chinos de SunnyToo). Sin impacto de código.
- **Custom CSS/JS injection sin validación** — trust model 100% sobre el admin. Un admin malicioso puede inyectar XSS via `STSN_CUSTOM_JS`.
- **Sin undo / versioning UI** — cambios de config son destructivos. Recomendado backup `ps_configuration` antes de experimentar.
- **Demo importer (clase `DemoStore`)** asume clean slate — no hay merge con config existente.
- **529 config keys** hace debugging tedioso. No hay UI de export/backup. Recomendado dump SQL manual si se quiere versionar la config.
- **Inline CSS en cada request** → posible cuello de botella si no hay Varnish/reverse-proxy cacheando.

## Cross-links

- [[_index|Modules index]] — tabla summary de los 54 módulos
- [[_by-category|Modules by category]] — agrupación funcional
- [[stthemeeditor]] — fiche shallow original (deja un "→ ver [[stthemeeditor-deep-dive]]" ahí)
- [[stbanner]], [[steasycontent]], [[stswiper]] — consumidores principales de config
- [[steasybuilder]] — NO interactúa con stthemeeditor (editor visual separado)
- [[../prestashop/themes/hummingbird-v2|Hummingbird 2.0]] — riesgo de incompatibilidad CSS variables
- [[../_agent/panda-expert]] — el agente debería saber que este es el hub

## Source references (agent lookup matrix)

Cuándo abrir qué fichero del source:

| Pregunta del usuario | File a leer | Search term |
|---|---|---|
| "Dónde configuro el logo" | `controllers/admin/AdminStthemeeditorController.php` | `header_logo` / `STSN_HEADER_LOGO` |
| "Qué opciones de color tengo" | `views/templates/admin/colors.tpl` | `STSN_COLOR_*` |
| "Dónde cambio la tipografía H1/H2" | `views/templates/admin/typography.tpl` | `STSN_FONT_H` / `STSN_HEADING_` |
| "Cómo inyecto CSS custom" | Admin screen "Custom CSS" | `STSN_CUSTOM_CSS` |
| "Qué clase genera el CSS inline" | `stthemeeditor.php` | `hookDisplayHeader` + `getMultiple` |
| "Cómo se duplica la config a otro shop" | `stthemeeditor.php` | `hookActionShopDataDuplication` |
| "Dónde está la config del megamenu" | (configurada aquí, no en stmegamenu) | `STSN_NAV_` / `STSN_MENU_` |
| "Cómo se resetean valores a default" | `classes/DemoStore.php` | `reset` / `importDefault` |

## Sugerencias de enriquecimiento derivadas

Identificadas durante el deep-dive — worth hacer como tasks separadas:

1. **[[_index|Modules index]]** → añadir callout "stthemeeditor es el hub — la mayoría de módulos se configuran leyendo STSN_* desde aquí, no tienen admin propio".
2. **Fiches de los consumidores** ([[stbanner]], [[steasycontent]], [[stswiper]]) → documentar qué Configuration keys específicas lee cada uno.
3. **Nueva página "Panda theme configuration matrix"** → listar las 529 `STSN_*` agrupadas por propósito (colors, typography, layout, product, utility) para búsquedas tipo "qué key controla X".

Sugerencia de prioridad: **1 primero** (5 min, alto impacto descubribilidad), **3 segundo** (1-2h, materializa los 529 keys como referencia buscable), **2 tercero** (por fiche del consumidor — incremental).
