---
module_name: steasy_trans_panda
display_name: Easy Builder Bridge — Panda PS-aware widgets
version: 1.0.x
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: page-builder-bridge
is_configurable: false
author: SUNNYTOO.COM
source_path: dependencies/modules/steasy_trans_panda
size_approx: 2.7MB
depends_on: [steasybuilder]
tags: [panda, module, page-builder, easy-builder, bridge, ps-aware]
---

# Easy Builder Bridge — Panda PS-aware (`steasy_trans_panda`)

> [!info] Bridge module
> Este módulo **NO se usa solo**. Es el puente que añade widgets **PS-aware** (producto dinámico, categoría, manufacturer, etc.) al editor de Easy Builder. Requiere `steasybuilder` instalado y licenciado.

## Purpose

Extiende Easy Builder con **20 widgets PS-aware** que conectan con datos del catálogo PrestaShop:

- **Productos**: featured products, special offers, related products, viewed products, new arrivals, bestsellers.
- **Categorías**: category slider, category tree, featured categories.
- **Marcas / Manufacturers**: brands slider, manufacturer list.
- **Contenido editorial**: blog posts (si `stblog` está instalado).
- **Componentes UI**: megamenu integration, newsletter signup wrapper.

Ver lista completa en `panda-kb/easybuilder/04-bridge-trans-panda.md`.

## Estructura del módulo

```
modules/steasy_trans_panda/
├── steasy_trans_panda.php
├── config.xml + config_es.xml
├── classes/
├── controllers/                     # Controllers que sirven datos PS al editor EB
├── elementor/                       # Adapters/wrappers para widgets PS-aware
├── custom_templates/                # Template overrides para widgets EB
├── views/
└── translations/
```

## Cuándo lo instalas

- ✅ Si el cliente usa Easy Builder Y quiere mostrar productos, categorías o marcas dinámicas en sus landings.
- ❌ NO lo instales sin `steasybuilder` — fallará al arrancar (depende de la API de widgets de EB).

## Cómo identificar widgets PS-aware en una landing existente

Si abres `./modules/steasybuilder/` y ves widgets cuyo `widget_type` empieza por `panda_` o `st_panda_`, son del bridge. Sin `steasy_trans_panda` instalado/activo, esos widgets aparecen rotos en el editor.

## Recursos relacionados en la KB

- [[modules/steasybuilder]] — módulo padre/dependencia.
- `panda-kb/easybuilder/04-bridge-trans-panda.md` — catálogo completo de los 20 widgets PS-aware.
- `panda-kb/easybuilder/00-overview.md` — overview general del track Easy Builder.

## Gap conocido

Las opciones de cada widget PS-aware (límite de productos, filtros por categoría, sort, etc.) viven en `./modules/steasy_trans_panda/elementor/` del CWD. La KB no las cristaliza por widget — léelas en source cuando necesites detalle.
