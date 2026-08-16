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
skills/ecommerce-ux-kb/          # SKILL.md = router de superficies, no dump
  references/
    homepage-and-navigation.md   # home, nav, PLP, product card
    product-page.md
    cart-and-checkout.md         # en PS el checkout SÍ es nuestro (Hydrogen lo delega)
    search-and-filters.md
    mobile.md                    # tap, sticky CTA, drawers — sin mezclar LCP
    _index.md

skills/ecommerce-perf-kb/         # SKILL.md = presupuesto + cómo priorizar
  references/
    cwv.md                       # LCP / INP / CLS (umbrales Google)
    images.md
    js-and-third-parties.md      # sliders st*, reviews, pixels
    how-to-measure.md            # Lighthouse local; JSON fuera del contexto
    _index.md
```

Reglas de la KB:

- SKILL.md enruta (“esto es PDP → lee product-page.md”). No vuelques toda la KB.
- Cada guideline: título imperativo + **Source** + **Why** + **En este stack** (Panda / ElementFlow / child checkout). Microformato de Hydrogen, sin su código Shopify.
- Principio primero, plataforma después.
- Nada de lifts inventados. Hydrogen lo dice y luego cita “+35% conversión”: no copiamos números. Si no hay fuente pública o caso vuestro, se marca opinión.
- Checkout/pago/a11y = “high stakes”: se dice claro, sin % inventado.
- Un hallazgo repetido en dos clientes → se documenta.

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
- KBs mínimas: las 5 superficies UX de arriba + `cwv.md` + `how-to-measure.md`. Oficio vuestro en **En este stack**, no un dump de blogs ni de Hydrogen.
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

## Cotilleo GitHub (agosto 2026)

El molde “skill/agente experto” ya existe. El combo **UX de tienda + perf de tienda + CWD PrestaShop/Panda + worker que maqueta** no.

### Lo más parecido a `ecommerce-ux-expert`

| Repo | Stars | Qué es | Por qué no lo copiamos |
|------|------:|--------|------------------------|
| [dnh33/webshop-ux-expert](https://github.com/dnh33/webshop-ux-expert) | 0 | Plugin Claude: consultor de webshop, KB Baymard/Amazon, audita URL con Playwright | 0 stars, mezcla agentic-commerce y un scout de joyería DK. Idea sí: consultor ≠ implementador |
| [hmtkyn/hydrogen-ecommerce-ux](https://github.com/hmtkyn/hydrogen-ecommerce-ux) | 2 | Skill Claude: PDP/cart/checkout/search + notas Hydrogen. Estructura de KB casi idéntica a la nuestra | Atado a Shopify Hydrogen. Cotillear carpetas, no el runtime |
| [nexscope-ai/eCommerce-Skills](https://github.com/nexscope-ai/eCommerce-Skills) | ~700 | 157 skills de *seller* (Amazon, Shopify, PPC, checkout genérico) | Marketing/ops de marketplace, no criterio sobre un child theme |
| [Seance1723/UXCraft](https://github.com/Seance1723/UXCraft) | 0 | UX general con capítulo ecommerce | Demasiado amplio; no es oficio de tienda |
| [PrestaShop/skills](https://github.com/PrestaShop/skills) | 5 | Skills oficiales: update, check, rollback | Ops de tienda, cero UX/perf |

### Lo más parecido a `ecommerce-perf-expert`

| Repo | Stars | Qué es | Encaje |
|------|------:|--------|--------|
| [addyosmani/web-quality-skills](https://github.com/addyosmani/web-quality-skills) | ~2.6k | Skills Lighthouse / CWV / a11y / SEO (web genérico) | **El de cotillear para perf.** No sabe de sliders `st*` ni peso de cart PS |
| [mohitkale/web-vitals-auditor](https://github.com/mohitkale/web-vitals-auditor) | 0 | Plugin: Lighthouse + bundle/images/fonts, subagentes | Buen patrón “medir en local”. No es ecommerce |
| [ualiyou/web-performance-audit](https://github.com/ualiyou/web-performance-audit), [EVEDensity/web-perf-audit](https://github.com/EVEDensity/web-perf-audit) | bajos | Auditores CWV genéricos | Misma familia; no storefront |

### Qué no hay

Nadie junta: experto UX que **no** maqueta + experto perf que **exige** Lighthouse + classifier que ya habla con `layout-builder` / `panda-expert` / `prestashop-expert` + KB que crece con clientes PS.

Cotillear en serio esas dos: abajo, [qué extraemos](#qué-extraemos-de-las-dos).

## Qué extraemos de las dos

Leídas enteras (SKILL + references). No se vendoriza el repo; se copia el *mecanismo*.

### Hydrogen UX — qué sirve

1. **KB por superficie de compra**, no por tecnología. Home+nav+PLP juntos; PDP; cart+checkout; search+filters; mobile. Search no es un extra: en Baymard público es first-class. Lo habíamos dejado flojo.
2. **SKILL.md es un router.** Identifica la superficie por el path (`templates/checkout`, `product.tpl`, builder) y lee **un** reference. “Do not dump every guideline at once.”
3. **Microformato de guideline** (título que es una regla, no un tema):

   ```
   ### Offer guest checkout; never require account creation
   **Source:** Baymard (top-3 abandonment)
   **Why it matters:** una frase
   **En este stack:** Panda / ElementFlow / `templates/checkout` …
   ```

   Su campo `Hydrogen implementation` lo sustituimos por notas al `panda-kb` / `elementflow-kb`. El experto UX no reescribe cómo funciona Easy Builder.
4. **Citar el porqué.** “Baymard: guest checkout” > “sé más usable”.
5. **No inventar porcentajes.** Lo dicen ellos; varios números del propio repo huelen a blog. Nosotros: umbrales Google (CWV) sí; lifts Baymard premium no.
6. **High stakes.** Checkout, pago, a11y: el experto lo marca; no lo diluye en un checklist de home.
7. **PS invierte su nota de checkout.** Hydrogen *no* posee el pago (`cart.checkoutUrl`). PrestaShop/Panda/ElementFlow **sí**. Por eso el experto UX aquí vale más que en Hydrogen: el child theme *es* el checkout.

### Hydrogen UX — qué no

- Código `CartForm`, `useOptimisticCart`, Predictive Search API.
- Mezclar LCP/a11y dentro del archivo de mobile (rompe el split de expertos).
- Copiar guidelines verbatim (Baymard premium + su LICENSE). Principios públicos + oficio vuestro.
- “Show code, don't just describe” para el experto UX. Aquí quien escribe CSS es `layout-builder`. El experto enseña *qué* y *dónde*; el worker el *cómo*.

### Addy web-quality-skills — qué sirve

1. **Un experto, varias lentes.** Ellos instalan 6 skills (audit / performance / cwv / a11y / seo / best-practices). Nosotros no. `ecommerce-perf-expert` carga `cwv.md` o `images.md` según el trigger, igual que el UX carga una superficie. `web-quality-audit` ≈ nuestro classifier cuando el ticket es “la web va mal” sin más.
2. **CWV no es otra persona.** Es un reference (`LCP.md` / nuestro `cwv.md`) con tabla Good / Needs work / Poor (umbrales Google, 75º percentil).
3. **Presupuesto en tabla**, luego retocado a storefront (pixels, sliders `st*`, reviews). Forma de Addy; números nuestros cuando midamos.
4. **Severidad: Critical / High / Medium / Low.** Prioridad: fallo CWV y barrera a11y antes que estilo. El output del experto perf es una lista rankeada, no un ensayo.
5. **SKILL.md < 500 líneas; references cargables solos.** Progressive disclosure. Hydrogen peca de files de 200+ líneas; Addy pide ~200. Nosotros: un tema por file, TOC arriba.
6. **Medir fuera del contexto.** `scripts/analyze.sh`: logs en stderr, JSON en stdout. El experto perf no “cree” el LCP: corre Lighthouse en la URL Lando (o pide el JSON). Fase 0 puede ser “cómo medir”; el script puede esperar.
7. **Bad / Good en perf.** Un `<img>` sin width/height vs con `width` `height` `fetchpriority`. Más útil que prosa. En UX el par es hallazgo en *este* tpl vs prescripción, no un snippet React.
8. **Agnóstico primero.** Vanilla/HTML, luego nota de framework. Igual que nosotros: principio → Panda/EF.

### Addy — qué no

- Early Hints, Speculation Rules, `getServerSideProps` como consejo por defecto en un child PS.
- Skill `best-practices` (CSP, HSTS) como experto de ecommerce. Eso no es conversión.
- Grep de HTML estático como auditoría de una tienda Smarty/builder (su `analyze.sh` no entiende `.tpl`).
- Instalar las 6 skills en paralelo a Mallard: duplicarían al classifier.

### Encaje en Fase 0

| De ellos | En Mallard |
|----------|------------|
| Router SKILL.md | `ecommerce-ux-kb/SKILL.md` y el agent apuntan a una superficie |
| 5 files de superficie | lista de KBs de arriba; search entra ya |
| Microformato Source/Why/Stack | plantilla de cada guideline |
| CWV table + budgets | `cwv.md` + presupuesto en el SKILL de perf |
| Critical/High/Medium/Low | formato de salida de ambos expertos |
| Script Lighthouse | `how-to-measure.md` ahora; script luego |
| UX prescribe, otro implementa | ya era la regla; Hydrogen tentaba a romperla con “show code” |

## Fuentes de oficio (para las KBs, no como runtime)

- Oficio vuestro de clientes (primera fuente).
- [Baymard](https://baymard.com/) — checkout, mobile, search (criterio, no copiar papers enteros).
- [web.dev / CWV](https://web.dev/explore/learn-core-web-vitals) — medición.
- KBs ya en Mallard: `prestashop-kb`, `panda-kb`, `elementflow-kb` para el *cómo* en PS.
- Estructura a mirar (no a vendorizar): [hydrogen-ecommerce-ux](https://github.com/hmtkyn/hydrogen-ecommerce-ux), [web-quality-skills](https://github.com/addyosmani/web-quality-skills).
