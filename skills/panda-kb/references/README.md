---
tags: [kb, panda, prestashop, expert-agent]
created: 2026-04-18
---

# Panda Knowledge Base

KB local del agente [[Panda Expert Agent]]. Markdown plano portable, zero vendor lock-in.

## Estructura

| Carpeta | Contenido | Fuente |
| --- | --- | --- |
| `docs/` | Panda docs oficiales scrapeados (shallow, sin imágenes) | https://panda2.sunnytoo.com/doc/ |
| `modules/` | 54 fiches generadas del source Panda | Distribución descomprimida |
| `demos/` | Cross-ref demos ↔ módulos (pending Task 3) | https://panda2.sunnytoo.com/select-demo.html |
| `easybuilder/` | Track paralelo Easy Builder + bridge (paid opcional) | steasybuilder + trans_panda source |
| `prestashop/` | **Docs oficiales PS platform** (themes, hooks, modules, migración 8↔9) | https://devdocs.prestashop-project.org |
| `_agent/` | Canonical del system prompt (fuente de verdad) | Mantenido a mano |
| `forum/` | Hilos foro / zona cliente (deferred Fase 2) | sunnytoo.com/my-account |

## Convenciones

- **Frontmatter obligatorio** en todos los ficheros:
  ```yaml
  ---
  source_url: https://panda2.sunnytoo.com/...
  ps_version: [8, 9]
  ingested: 2026-04-18
  tags: [panda, docs]  # o [panda, demo], [panda, module], etc.
  ---
  ```
- **Tag dual** `ps_version: [8, 9]`: Panda apenas cambia entre PS 8 y 9. Sólo anotar inline (callout `[!warning]`) donde exista una diferencia real.
- **Callouts + wikilinks** estilo Obsidian nativo.
- **No editar a mano** ficheros de `docs/` — son ingesta automática. Anotaciones propias van en `modules/` o notas aparte.

## Fuentes canónicas

- Producto: https://www.sunnytoo.com/product/panda-creative-responsive-prestashop-theme
- Docs: https://panda2.sunnytoo.com/doc/
- Demos: https://panda2.sunnytoo.com/select-demo.html
- PrestaShop DevDocs: https://devdocs.prestashop-project.org

## Proyecto

Ver [[Panda Expert Agent]] para el plan completo, fases, y decisiones de arquitectura.

## Report inicial

[[00-Research-Report]] — reconocimiento hecho 2026-04-18 antes de ingerir.

## Otros docs operativos

- [[SETUP-NEW-MACHINE]] — cómo replicar el agente en otro Mac.
- [[_agent/README]] — fuente de verdad del system prompt.
