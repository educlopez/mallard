---
tags: [kb, elementflow, prestashop, page-builder]
created: 2026-07-20
---

# ElementFlow Knowledge Base

KB local para el skill `elementflow-kb`. Markdown plano portable, zero vendor lock-in — mismo patrón que `panda-kb`.

## Estructura

| Carpeta | Contenido | Fuente |
| --- | --- | --- |
| `docs/` | 50 docs oficiales scrapeados desde el sitio demo/docs de ElementFlow | https://elementflow.io/third/posts/ |

ElementFlow no tiene un dominio de documentación separado — el propio sitio demo (`elementflow.io`, una tienda PrestaShop corriendo el módulo) sirve de docs vía su blog module (`st_blog`/posts), cada guía es un post bajo `/third/posts/`.

## Convenciones

- **Frontmatter obligatorio** en todos los ficheros:
  ```yaml
  ---
  source_url: https://elementflow.io/third/posts/...
  ps_version: [8, 9]
  ingested: 2026-07-20
  tags: [elementflow, docs, <getting-started|page-builders|widget|feature>]
  section: <Getting started|Page builders|Widgets|Features>
  ---
  ```
- **No editar a mano** ficheros de `docs/` — son ingesta automática scrapeada. Anotaciones propias van en el `SKILL.md` (sección "Key facts") o en notas aparte.
- Cada post-fuente incluye un widget de navegación lateral con ~18 enlaces a otros widgets — eso se descartó durante la ingesta, no forma parte del contenido real de cada página.

## Honestidad sobre la fuente

La mayoría de estos posts son **tips FAQ cortos**, no referencias exhaustivas de settings — el sitio de ElementFlow documenta "cómo resolver X problema concreto" más que "esta es la lista completa de opciones del widget Y". Para el listado real de controles/opciones de un widget, la fuente de verdad es el código del módulo `stsitebuilder` en el proyecto actual (`modules/stsitebuilder/libs/elementor/includes/widgets/*.php`), no esta KB.

## Fuentes canónicas

- Producto/demo/docs: https://elementflow.io/
- Docs índice completo (vía nav renderizado, no crawleable estáticamente): cualquier post bajo `/third/posts/` expone el menú lateral completo al renderizar en browser.

## Report inicial

Scrapeado 2026-07-20 vía `defuddle` (CLI) + navegación en browser real para descubrir el menú completo (el HTML servidor-side de cada post individual no incluye el nav completo — solo aparece renderizado por JS).
