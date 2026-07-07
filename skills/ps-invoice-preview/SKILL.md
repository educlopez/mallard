---
name: ps-invoice-preview
description: >
  Generates a local invoice PDF viewer for PrestaShop projects, deployed as
  invoice-preview/index.php (subfolder, ROOT = dirname(__DIR__)). Lists the
  last 50 orders (searchable by reference/customer/id), renders the real PDF
  invoice for any order straight from classes/pdf/HTMLTemplateInvoice.php +
  pdf/*.tpl (respecting any themes/{child}/pdf/*.tpl override PS itself
  resolves) inline in the browser via an iframe, with a "Generar factura"
  button for orders that don't have one yet. Zero placeholders to fill in —
  admin dir, Lando URL, DB credentials and shop name are all auto-detected
  at runtime from config/parameters.php and the PS_SHOP_NAME configuration
  value. Gated behind backoffice employee auth (Cookie class, always denies
  on failure). Geist-font UI matching the ps-email-preview visual style.
  Use this skill whenever the user wants to preview, inspect, or debug
  PrestaShop invoice/receipt PDF templates without going through the
  backoffice order page each time, asks "cómo se ve la factura", "quiero
  ver el PDF de la factura", "invoice preview", "visor de facturas",
  "previsualizar recibo/rebut", or is iterating on pdf/invoice*.tpl,
  pdf/footer.tpl, or HTMLTemplateInvoice overrides. Trigger proactively when
  the user is about to edit invoice/receipt PDF layout.
version: "1.0.0"
metadata:
  author: Eduardo Calvo
---

# PS Invoice Preview

Genera un visor de facturas PDF para PrestaShop en un solo archivo PHP.
Renderiza el PDF real (mismo motor que usa el backoffice: `HTMLTemplateInvoice`
+ `pdf/*.tpl`) para cualquier pedido, directo en el navegador, sin pasar por
el backoffice. Pensado para iterar rápido sobre cambios de maquetación de
factura (`pdf/invoice*.tpl`, `pdf/footer.tpl`, overrides de tema).

A diferencia de `ps-email-preview`, esta skill **no necesita placeholders**:
todo se detecta en runtime dentro del propio PHP (carpeta de admin, URL,
credenciales DB, nombre de tienda). El paso de generación es una copia
directa del asset.

## Características del visor

- **Sidebar con últimos 50 pedidos** — buscador en tiempo real por referencia, nombre/apellido de cliente o ID de pedido
- **Estado del pedido** — badge con el nombre del estado actual (`order_state_lang`)
- **Pedido con factura generada** → botón "Ver factura #N" que carga el PDF real en un iframe
- **Pedido sin factura** → botón "Generar factura" que llama a `$order->setInvoice(true)` y recarga
- **PDF real, no una aproximación HTML** — usa `PDF::TEMPLATE_INVOICE` + `Order::getInvoicesCollection()`, el mismo pipeline que el backoffice (`AdminOrders` → generar factura PDF), así que respeta cualquier override en `themes/{child}/pdf/*.tpl` o clases `override/classes/pdf/HTMLTemplateInvoice.php` tal cual estarían en producción
- **Auto-detección de carpeta admin** — busca la primera carpeta de primer nivel con `init.php` + `bootstrap.php` (cubre admins renombrados por seguridad, ej. `panelAmics`)
- **Nombre de tienda dinámico** — lee `PS_SHOP_NAME` de la config PS para el `<title>` y el encabezado del sidebar
- **Sin caché** — cada carga re-renderiza el PDF desde cero contra la BD actual

---

## Paso 1 — Detectar el contexto del proyecto (opcional, informativo)

```bash
# Child theme (para saber dónde el usuario debe poner sus overrides de pdf/)
grep -rl "^parent:" themes/*/config/theme.yml 2>/dev/null

# Nombre del proyecto Lando → URL de acceso
grep "^name:" .lando.yml | head -1
```

- **`LANDO_URL`** = `https://{nombre-lando}.lndo.site` (mismo cálculo que en `ps-email-preview`: los puntos de versión se eliminan, ej. `v9.1` → `v91`)
- **`CHILD_THEME`** — solo se usa para comunicarle al usuario dónde crear sus overrides (`themes/{child}/pdf/*.tpl`); el visor mismo no lo necesita como variable porque PS resuelve la ruta internamente vía `HTMLTemplate::getTemplate()`.

## Paso 2 — Verificar la carpeta de admin

El visor necesita saber si hay sesión de empleado activa. Detecta automáticamente
la carpeta de admin buscando `init.php` + `bootstrap.php`, pero si el proyecto
tiene una convención rara, verifica manualmente:

```bash
find . -maxdepth 1 -type d -iname "*admin*"
```

No hace falta configurar nada — es solo para poder decirle al usuario a qué URL
ir a loguearse si el visor le muestra "Acceso restringido".

## Paso 3 — Generar invoice-preview/index.php

Copia el asset **tal cual** (no hay placeholders que sustituir):

```bash
mkdir -p invoice-preview/
cp {ruta-de-este-skill}/assets/invoice-preview.php invoice-preview/index.php
```

El archivo usa `ROOT = dirname(__DIR__)` para subir un nivel y acceder a PS.

### Nota PS8 vs PS9

El asset bootea el contexto con el contenedor legacy `'front'`
(`PrestaShop\PrestaShop\Adapter\ContainerBuilder::getContainer('front', _PS_MODE_DEV_)`),
porque es el que expone `prestashop.core.localization.locale.repository`
(necesario para que `Tools`/`Locale` formateen precios dentro de los `.tpl`).
Esto funciona en PS8 y PS9. Si el proyecto es PS9 con `AdminKernel` disponible
y se prefiere ese camino, se puede reemplazar `bootPsContext()` por:

```php
require_once ROOT . '/autoload.php';
$kernel = new AdminKernel(_PS_ENV_, _PS_MODE_DEV_);
$kernel->boot();
Context::getContext()->container = $kernel->getContainer();
```

No es necesario salvo que el proyecto ya use ese patrón en otro lado.

## Paso 4 — Verificar y comunicar al usuario

Confirma al usuario:
- La URL de acceso: `{LANDO_URL}/invoice-preview/`
- Que necesita estar logueado en el backoffice (mismo guard que `ps-email-preview`)
- Dónde van sus overrides de plantilla de factura: `themes/{CHILD_THEME}/pdf/{nombre}.tpl` (ej. `invoice.tpl`, `invoice.product-tab.tpl`, `footer.tpl`, `header.tpl`) — PS resuelve child → core automáticamente (para PDFs **no** hay nivel intermedio de parent theme, a diferencia de emails: ver `classes/pdf/HTMLTemplate.php::getTemplate()`)
- Que "Ver factura" pide un pedido con `invoice_number` ya generado; si no lo tiene, aparece "Generar factura"

## Notas importantes

- El archivo `invoice-preview/index.php` es solo para desarrollo local — no subir a producción.
- Antes de tocar un `.tpl` de factura, recomienda al usuario limpiar el caché
  de clases de PS si acaba de tocar un `override/classes/pdf/*.php`:
  `var/cache/{dev,prod}/class_index.php` — si no, Smarty puede fallar con
  "unknown modifier" al no cargar el override que registra el modifier.
- Los templates de factura son compartidos entre tipos de documento cuando
  aplica: `pdf/footer.tpl`, `pdf/header.tpl` y `pdf/pagination.tpl` los usan
  TODOS los PDFs (factura, albarán, abono, devolución), no solo la factura.
  Si el usuario pide "quitar X del pie de la factura" y X está en `footer.tpl`,
  avisa que el cambio afectará también a los demás documentos PDF, salvo que
  se prefiera duplicar lógica condicional por tipo de documento dentro del tpl.
- Para forzar quitar una columna de la tabla de productos (ej. impuestos) sin
  tocar la config global (`PS_TAX`), reutiliza el layout que el core YA aplica
  cuando `$isTaxEnabled` es false (ver `computeLayout()` en
  `classes/pdf/HTMLTemplateInvoice.php`): en el override del `.tpl`, fuerza
  `{assign var='widthColProduct' value=$layout.product.width+$layout.tax_code.width}`
  y quita los bloques `{if $isTaxEnabled}` correspondientes, sin envolverlos
  en condicional — así el resto de cálculos de impuestos sigue intacto.
