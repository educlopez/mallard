---
module_name: steasybuilder
display_name: Easy Builder
version: 1.2.1
panda_version: 2.9.2
ps_version: [8, 9]
ps_compliancy: "1.7.1.0 - ~"
category: page-builder
is_configurable: true
author: SUNNYTOO.COM
source_path: dependencies/modules/steasybuilder
size_approx: 36MB
tags: [panda, module, page-builder, paid, optional, elementor-based]
---

# Easy Builder (`steasybuilder`)

> [!warning] Paid + optional
> **NO está incluido en Panda stock.** Requiere licencia separada SunnyToo y activación. Comprueba con `ls ./modules/steasybuilder` en el CWD antes de proponerlo.

## Purpose

Visual drag-and-drop page builder embebido en PrestaShop. Bajo el capó está construido sobre una versión adaptada de **Elementor**.

## Estructura del módulo (PS 9 + Panda 2.9.2)

```
modules/steasybuilder/
├── steasybuilder.php                # Main module class
├── config.xml + config_es.xml
├── classes/                         # Domain classes (templates, presets, settings)
├── controllers/admin/
│   ├── AdminStEasyBuilderController.php           # Main BO
│   ├── AdminStEasyBuilderEditorController.php     # Visual editor
│   ├── AdminStEasyBuilderGeneralSettingController.php
│   └── AdminStEasyBuilderLinceseController.php    # (sic — typo upstream)
├── elementor/                       # Forked/adapted Elementor core
├── import_data/                     # Sample data and presets
├── other-setting/
├── upgrade/                         # Module version upgrade migrations
├── translations/
└── views/                           # Admin + front templates
```

## Cuándo recomendarlo

✅ Usa Easy Builder **si**:
- El cliente edita layouts **mensualmente o más**.
- Necesitan landings de campaña con composición libre (no solo banner+texto).
- El cliente quiere autonomía total sin pedirte ayuda cada vez.
- Tienes licencia válida o el cliente la pagó.

❌ **NO** uses Easy Builder para:
- Ediciones esporádicas (banners 2 veces al año, página "Sobre nosotros" estática) → usa `stbanner`, `stfeaturedslider`, `steasycontent` (este último viene bundled gratis con Panda).
- Clientes con presupuesto ajustado: la licencia anual no compensa.

## Widgets disponibles

Easy Builder trae **2 catálogos** de widgets:

1. **31 widgets genéricos** del core EB (botón, heading, imagen, columna, etc.) — ver `panda-kb/easybuilder/01-widgets-catalog.md`.
2. **20 widgets PS-aware** vía el bridge module `steasy_trans_panda` (producto, categoría, manufacturer dinámicos) — ver `panda-kb/easybuilder/04-bridge-trans-panda.md`.

**Para widgets PS-aware necesitas ambos** módulos instalados: `steasybuilder` + `steasy_trans_panda`. Solo `steasybuilder` te limita a los genéricos.

## Incompatibilidades conocidas

Easy Builder **no convive** con:
- `jscomposer` (otro page builder de terceros).
- `creativeelements` (Elementor wrapper alternativo).
- `crazyelements`.

Razón: todos forkean Elementor o intentan registrarse como page builder principal. Si detectas alguno de estos en `./modules/`, avisa al cliente antes de instalar EB.

## Recursos relacionados en la KB

- `panda-kb/easybuilder/00-overview.md` — overview del track Easy Builder.
- `panda-kb/easybuilder/01-widgets-catalog.md` — 31 widgets genéricos.
- `panda-kb/easybuilder/02-presets-catalog.md` — presets/templates pre-armados.
- `panda-kb/easybuilder/03-admin-controllers.md` — controllers BO detallados.
- `panda-kb/easybuilder/04-bridge-trans-panda.md` — 20 widgets PS-aware en `steasy_trans_panda`.
- [[modules/steasy_trans_panda]] — bridge module obligatorio para widgets PS-aware.
- [[modules/steasycontent]] — alternativa gratis bundled en Panda para bloques HTML libres.

## Gap conocido

Esta fiche es de alto nivel — el detalle de cada widget, preset, y controller vive en los docs del track `panda-kb/easybuilder/`. Para opciones específicas de cada controller, lee `./modules/steasybuilder/controllers/admin/AdminStEasyBuilder*Controller.php` del CWD.
