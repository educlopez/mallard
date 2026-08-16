# Equipo de expertos ecommerce — plan

Mallard ya tiene expertos de **plataforma y maquetación**. Falta el oficio que mueve ventas sin ser ops de pedidos: **UX/CRO** y **performance** de tienda (y, luego, SEO y accesibilidad).

Mismo molde que `prestashop-expert` / `panda-expert` / `layout-builder`: knowledge base que crece según trabajas + agente CWD-first + el classifier enruta. Sin MCP ajeno, sin `store.yaml`, sin leer pedidos ni aplicar precios.

## Qué estás montando (el resultado)

Un **equipo de expertos** en el mismo Claude/Cursor de siempre. No operan la tienda. **Diagnosticán y prescriben** sobre el proyecto abierto (el child theme, el checkout, el peso de la home). Quien implementa sigue siendo `layout-builder` o tú.

Hoy el hueco se ve así:

| Pregunta | Quién la coge hoy | Qué falta |
|----------|-------------------|-----------|
| “¿Cómo registro un hook en PS 9?” | `prestashop-expert` | — |
| “¿Qué módulo `st*` pongo en el header?” | `panda-expert` | — |
| “El logo pisa el buscador en móvil” | `layout-builder` | — |
| “El checkout pierde gente en móvil” | Nadie con criterio | Experto UX |
| “La ficha tarda 4 s en LCP” | Nadie con criterio | Experto perf |
| “¿Esta PLP se puede indexar así?” | Nadie | Experto SEO (luego) |
| “¿Esto pasa EAA?” | Apuntes sueltos | Experto a11y (luego) |

`layout-builder` **implementa** CSS. No decide si el checkout debe ser un paso o tres. `prestashop-expert` sabe el hook. No sabe si el campo de teléfono en guest checkout mata la conversión.

### Un día concreto

```
tú:     el checkout en móvil se siente eterno y la gente no paga

Mallard: [classifier → ecommerce-ux-expert]
         Revisa templates/checkout + CSS del child (CWD).
         Hallazgos:
         - 11 campos en guest; Baymard: teléfono y empresa sobran para este vertical
         - CTA bajo el fold en 375px; el resumen tapa el pagar
         - No hay trust (pago / envío) junto al botón
         Implementación: 3 cambios de layout → layout-builder
         No toco precios ni pedidos.
```

```
tú:     la home va a trompicones, LCP > 4s

Mallard: [classifier → ecommerce-perf-expert]
         Mira el child: slider hero, 8 webs de st*, imágenes 2400px, JS del carrito.
         Prescripción:
         - hero: una imagen, width/height, fetchpriority
         - no cargar el slider de related en above-the-fold
         - third-party: diferir el de reviews
         Verificar: Lighthouse local en la URL Lando, no “yo creo que va mejor”.
```

`/task 41953` gana tipos `ux` y `perf`. Si el ticket dice “Disseny, el botón de comprar no se ve”, puede ser layout (CSS) o UX (jerarquía). El classifier elige; si duda, pregunta.

### Qué no es

No es `/briefing` de ventas, ni cola de impagos, ni un conector a la API de la tienda, ni un agente que venda en la web. Eso era la línea de **ops**. Queda [apartada](#línea-apartada-ops): otro producto, otro día.

## El equipo

Orquestación = el `task-classifier` que ya existe, ampliado. No un “Ecommerce God”.

```
                    task-classifier
                 (enruta, no ejecuta)
     ┌──────────┬──────────┼──────────┬──────────┐
     ▼          ▼          ▼          ▼          ▼
 prestashop-  panda-   layout-   ecommerce-  ecommerce-
 expert       expert   builder   ux-expert   perf-expert
 (plataforma) (theme)  (hace CSS) (criterio)  (criterio)
```

### Núcleo (montar primero)

| Agente | Oficio | Mira | No hace |
|--------|--------|------|---------|
| `ecommerce-ux-expert` | Conversión y oficio de tienda: checkout, PDP, PLP, cart, búsqueda, confianza, móvil, merchandising visual | Templates, CSS, builder, flujos reales del CWD | No maqueta él (delega a `layout-builder`). No toca módulos PS de negocio |
| `ecommerce-perf-expert` | Velocidad que afecta a venta: LCP/INP/CLS, imágenes, JS/CSS budget, third-parties, peso de cart/checkout | Assets, sliders, fuentes, requests, Lando | No “optimiza” a ciegas. Prescribe + cómo medir |
| `task-classifier` | Tipos nuevos: `ux`, `perf` | Título / notas / módulo | Sigue sin hacer el trabajo |

Cada experto es CWD-first, como los otros: primero el proyecto, luego la KB.

### Segunda oleada (cuando el núcleo se use)

| Agente | Oficio |
|--------|--------|
| `ecommerce-seo-expert` | IA de categorías, fichas, facetas, indexación, canónicos. No copy barato |
| `ecommerce-a11y-expert` | EAA / teclado / contraste / carrito y checkout usables. El expert PS ya menciona EAA; este lo posee |

### Relación con lo que ya hay

- **UX prescribe, `layout-builder` ejecuta.** “El CTA queda bajo el fold” → el builder mueve el bloque. Si el experto UX se pone a editar 15 CSS, está roto (mismo espíritu `delegate_only` de gentle-ai).
- **Perf prescribe, tú o un skill concreto ejecutan.** Regen de imágenes → `ps-image-regen`. CSS build → `ps-css-build`. El experto no inventa un pipeline.
- **Panda/PS expert** siguen siendo la fuente de “cómo se hace en esta plataforma”. El UX no reescribe `panda-kb`; dice *qué* hay que conseguir y pregunta al experto de theme *con qué módulo*.
- Playbooks **agnósticos** (un checkout es un checkout). Notas “en Panda / ElementFlow / PS 9” viven en la KB y apuntan a las KBs ya existentes.

## Knowledge bases (lo que de verdad acumulamos)

El valor no es el YAML del agente. Es la KB que vas llenando con oficio de clientes, igual que `panda-kb`.

```
skills/ecommerce-ux-kb/references/
  checkout.md          # pasos, guest, campos, trust, errores
  product-page.md      # galería, variant, CTA, envío, reviews
  listing.md           # PLP, facets, empty, sort
  cart.md
  mobile.md
  merchandising.md     # home, above-the-fold, campañas
  _index.md

skills/ecommerce-perf-kb/references/
  cwv.md               # LCP INP CLS en storefront
  images.md
  js-budget.md
  third-parties.md
  checkout-weight.md
  how-to-measure.md    # Lighthouse local, qué no mentir
  _index.md
```

Reglas de la KB:

- Principio primero (“no pidas teléfono si no envías por SMS”), plataforma después (“en PS el guest vive en…, en ElementFlow el checkout es child theme”).
- Nada de lifts inventados. Si no hay fuente (Baymard, web.dev, un caso vuestro), se marca opinión.
- Un hallazgo repetido en dos clientes → se documenta. Así se “amplía a otras plataformas” de verdad: el principio es el mismo; la nota PS/Shopify se añade.

## Cómo se usa (comandos)

```
/ux          → ecommerce-ux-expert
/perf        → ecommerce-perf-expert
/task        → classifier; ahora puede devolver TYPE: ux | perf
```

No hace falta `mallard commerce`. No hay subcomando nuevo. `mallard update` enlaza agentes y skills, como siempre.

## Principios

1. **Criterio ≠ implementación.** El experto diagnostica; el worker (o el humano) cambia el theme.
2. **CWD-first.** La home de *este* cliente, no un checklist genérico de internet.
3. **Agnóstico de plataforma en el oficio, concreto en el cómo.** El agente UX no se llama `ps-ux-expert`.
4. **Medir cuando sea perf.** Sin Lighthouse/network, el experto perf solo puede hablar de riesgos, no de victoria.
5. **Sin conector de tienda.** No leemos stock ni pedidos. El código y el front son la fuente.
6. **Sin MCP de terceros.** Si un día el experto perf lanza Lighthouse, es un binario local, no un bus ajeno.

## Fases

### Fase 0 — Los dos expertos vacíos pero útiles

- `claude/agents/ecommerce-ux-expert.md` y `ecommerce-perf-expert.md` (mismo frontmatter que `layout-builder`).
- KBs mínimas: checkout + PDP + mobile (UX); CWV + images + how-to-measure (perf). Oficio vuestro, no un dump de blogs.
- `/ux`, `/perf`.
- `task-classifier`: tipos `ux` y `perf` + señales (“conversión”, “checkout”, “LCP”, “lento”).
- Cero Go nuevo. Cero adapter.

### Fase 1 — Que se peleen con proyectos reales

- Un ticket de checkout y uno de home lenta en un cliente PS (Panda o ElementFlow).
- El experto UX debe **delegar** el CSS a `layout-builder`.
- Añadir a la KB solo lo que hayáis tenido que decidir de verdad.

### Fase 2 — SEO y a11y, si el núcleo se usa

- Mismo molde. No antes: si UX/perf no se invocan, no hace falta más roster.

## Encaje con gentle-ai (sin SDD)

Sigue valiendo el routing, no la ceremonia:

- Inline: “el logo pisa el buscador” → `layout-builder`.
- Delegate: “no convierten en checkout” → experto UX, luego worker.
- El classifier no ejecuta (como `task-classifier` hoy).

No hay recibo de precios ni `apply`. El “gate” de siempre: el experto **draft**, tú revisas en Lando, no deploya.

## Línea apartada (ops)

La investigación anterior (briefing semanal, triage de pedidos, adapters, recibo/`apply`, MCP de PS/Shopify) era un **OS de operaciones**. Encaja en “orquestación y ventas” si se lee como back-office. No es lo que querías.

Si un día hace falta, es otro namespace y otro momento. No mezclarlo con estos expertos: un agente que lee pedidos y otro que opina de LCP no son el mismo oficio.

Notas y fuentes de esa línea (mercado, adapters, gentle-ai RDD) se pueden recuperar del historial de git de este doc si hicieran falta.

## Fuentes de oficio (para las KBs, no como runtime)

- Oficio vuestro de clientes (primera fuente).
- [Baymard](https://baymard.com/) — checkout, mobile, search (criterio, no copiar papers enteros).
- [web.dev / CWV](https://web.dev/explore/learn-core-web-vitals) — medición.
- KBs ya en Mallard: `prestashop-kb`, `panda-kb`, `elementflow-kb` para el *cómo* en PS.
