# Equipo de expertos ecommerce — plan

Mallard ya tiene expertos de **plataforma y maquetación**. Falta el oficio que mueve ventas sin ser ops de pedidos: **UX/CRO** y **performance** de tienda (y, luego, SEO y accesibilidad).

Mismo molde que `prestashop-expert` / `panda-expert` / `layout-builder`: knowledge base que crece según trabajas + agente CWD-first + el classifier enruta. Sin MCP ajeno, sin `store.yaml`, sin leer pedidos ni aplicar precios.

## Qué estás montando (el resultado)

Un **equipo de expertos** en el mismo Claude/Cursor de siempre. No operan la tienda. **Diagnosticán y prescriben** oficio de storefront: PDP, cart, checkout, búsqueda, móvil, LCP. El CWD puede ser un theme Shopify, Woo, Magento, Hydrogen, un child PS o un front custom. Quien implementa sigue siendo `layout-builder` (o el worker del stack) o tú.

El *cómo* en una plataforma concreta (Panda, Dawn, Flatsome, Hydrogen) no vive en estos expertos: vive en las KBs de plataforma que ya hay, o en una nota **En este stack**.

La unidad portable es la **skill** (7 harnesses). El subagente es candado de herramientas y contexto, y hoy Mallard solo lo entrega a Claude y Cursor. Detalle: [Harnesses y Mallard](#harnesses-y-mallard-agosto-2026).

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
         Mira el CWD: templates/código de checkout (da igual el motor).
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
         Mira el storefront: hero pesado, related above-the-fold, imágenes 2400px, JS de reviews/pixels.
         Prescripción:
         - hero: una imagen, width/height, fetchpriority
         - no cargar related en above-the-fold
         - third-party: diferir el de reviews
         Verificar: Lighthouse en la URL local (o JSON), no “yo creo que va mejor”.
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
| `ecommerce-ux-expert` | Conversión y oficio de tienda: checkout, PDP, PLP, cart, búsqueda, confianza, móvil, merchandising visual | Templates, CSS, builder, flujos reales del CWD | No maqueta (sin Edit/Write en el subagente; en skill, instrucción). No instala apps ni toca catálogo/pedidos |
| `ecommerce-perf-expert` | Velocidad que afecta a venta: LCP/INP/CLS, imágenes, JS/CSS budget, third-parties, peso de cart/checkout | Assets, sliders, fuentes, requests, URL local | No “optimiza” a ciegas. Prescribe + cómo medir. Bash para Lighthouse, no para reescribir el theme |
| `task-classifier` | Tipos nuevos: `ux`, `perf` | Título / notas / módulo | Sigue sin hacer el trabajo |

Cada experto es CWD-first, como los otros: primero el proyecto, luego la KB.

### Segunda oleada (cuando el núcleo se use)

| Agente | Oficio |
|--------|--------|
| `ecommerce-seo-expert` | IA de categorías, fichas, facetas, indexación, canónicos. No copy barato |
| `ecommerce-a11y-expert` | EAA / teclado / contraste / carrito y checkout usables. El expert PS ya menciona EAA; este lo posee |

### Relación con lo que ya hay

- **UX prescribe, `layout-builder` ejecuta.** “El CTA queda bajo el fold” → el builder mueve el bloque. Si el experto UX se pone a editar 15 CSS, está roto (mismo espíritu `delegate_only` de gentle-ai).
- **Perf prescribe, tú o un skill concreto ejecutan.** Regen de imágenes, build de CSS, CDN: el experto nombra el *qué*; el skill/worker del stack hace el *cómo*. No inventa un pipeline.
- **Expertos de plataforma** (PS, Panda, o el que toque) son la fuente de “cómo se hace aquí”. El UX no reescribe esas KBs; dice *qué* hay que conseguir y pregunta *con qué pieza del theme*.
- Playbooks **agnósticos** (un checkout es un checkout). **En este stack** es una nota al CWD, no el oficio.

## Knowledge bases (lo que de verdad acumulamos)

El valor no es el YAML del agente. Es la KB que vas llenando con oficio de clientes, igual que `panda-kb`.

```
skills/ecommerce-ux-kb/          # SKILL.md = router de superficies, no dump
  references/
    homepage-and-navigation.md   # home, nav, PLP, product card
    product-page.md
    cart-and-checkout.md         # checkout propio vs hosted (Shopify Checkout vs Woo/Magento/PS/custom)
    search-and-filters.md
    mobile.md                    # tap, sticky CTA, drawers — sin mezclar LCP
    _index.md

skills/ecommerce-perf-kb/         # SKILL.md = presupuesto + cómo priorizar
  references/
    cwv.md                       # LCP / INP / CLS (umbrales Google)
    images.md
    js-and-third-parties.md      # sliders, reviews, pixels, tag managers
    how-to-measure.md            # Lighthouse local; JSON fuera del contexto
    _index.md
```

Reglas de la KB:

- SKILL.md enruta (“esto es PDP → lee product-page.md”). No vuelques toda la KB.
- Cada guideline: título imperativo + **Source** + **Why** + **En este stack** (solo si el CWD lo pide: Dawn, Flatsome, Panda, Hydrogen…). Microformato de Hydrogen, sin su código Shopify.
- Principio primero, plataforma después.
- Nada de lifts inventados. Hydrogen lo dice y luego cita “+35% conversión”: no copiamos números. Si no hay fuente pública o caso vuestro, se marca opinión.
- Checkout/pago/a11y = “high stakes”: se dice claro, sin % inventado.
- Un hallazgo repetido en dos clientes → se documenta.

## Cómo se usa (por harness)

Mallard no es solo Claude. Lo que el usuario *teclea* cambia:

| Qué quieres | Claude Code | Cursor | OpenCode | Codex / Gemini / Windsurf |
|-------------|-------------|--------|----------|---------------------------|
| Oficio UX/perf | skill se autoinvoca **y** subagente | skill (`/` + nombre) **y** subagente en `~/.cursor/agents` | skill tool + `/ux` si hay command | **solo la skill** (description = router) |
| Invocar a mano | `/ux` `/perf` (command → subagente o skill) | `/ecommerce-ux` si la skill se llama así; no hay `CommandsDir` | `/ux` (command). Hoy Mallard **no** enlaza agents de OpenCode | `$ecommerce-ux-kb` (Codex) o descripción |
| `/task` → tipo ux/perf | sí (`task-classifier`) | no hay `/task` (command Claude/OpenCode) | sí el command, no el subagente classifier | no |

Consecuencia: **la KB skill es el experto en 7 de 7 harnesses.** El fichero `claude/agents/*.md` es lujo de aislamiento (Claude + Cursor hoy). Si el oficio vive solo en el agente, Codex no lo ve.

Nombres (mismo molde que `panda-kb` / `panda-expert` / `/panda`):

```
skills/ecommerce-ux-kb/      +  claude/agents/ecommerce-ux-expert.md  +  /ux
skills/ecommerce-perf-kb/    +  claude/agents/ecommerce-perf-expert.md +  /perf
```

No hace falta `mallard commerce`. `mallard update` enlaza. Cero Go para Fase 0.

## Principios

1. **Criterio ≠ implementación.** El experto diagnostica; el worker (o el humano) cambia el theme.
2. **CWD-first.** La home de *este* cliente, no un checklist genérico de internet.
3. **Agnóstico de plataforma en el oficio, concreto en el cómo.** El agente UX no se llama `ps-ux-expert` ni `shopify-ux-expert`. La investigación tampoco se filtra por “¿encaja en un child PS?”.
4. **Medir cuando sea perf.** Sin Lighthouse/network, el experto perf solo puede hablar de riesgos, no de victoria.
5. **Sin conector de tienda.** No leemos stock ni pedidos. El código y el front son la fuente.
6. **Sin MCP de terceros.** Si un día el experto perf lanza Lighthouse, es un binario local (`scripts/` de la skill), no un bus ajeno.
7. **La skill es el producto; el agente es el candado.** Oficio y references viven en `skills/`. El agente es fino: CWD-first, carga la KB, formato de salida, sin Edit/Write (UX) o con Bash para medir (perf). Si duplicamos el oficio en el `.md` del agente, Codex/Gemini se quedan fuera.

## Harnesses y Mallard (agosto 2026)

Investigación de cómo cada runtime carga skills / commands / agents, contra lo que Mallard **realmente** enlaza. Esto cambia Fase 0.

### Qué enlaza Mallard hoy

Código: `internal/agents/*.go`. `mallard update` hace symlink desde el repo.

| Adapter | Detect | Skills | Commands (`/foo`) | Agents (subagente) |
|---------|--------|--------|-------------------|-------------------|
| Claude Code | `~/.claude` o `claude` en PATH | `~/.claude/skills` y `<ws>/.claude/skills` | sí | sí |
| Cursor | `~/.cursor` | `~/.cursor/skills` | **vacío** | sí (`~/.cursor/agents`) |
| OpenCode | `~/.config/opencode` o `opencode` | `~/.config/opencode/skills` | sí | **vacío** (el producto ya los tiene) |
| Codex | `~/.codex` o `codex` | `~/.codex/skills` | vacío | vacío |
| Gemini CLI | `~/.gemini` o `gemini` | `~/.gemini/skills` | vacío | vacío |
| Windsurf | `~/.codeium/windsurf` | `~/.codeium/windsurf/skills` | vacío | vacío |
| Generic | `~/.agents` | `~/.agents/skills` | vacío | vacío |

`docs/adding-skills.md` dice “agents = Claude only”. Eso **ya es falso**: Cursor recibe agents. OpenCode documenta `~/.config/opencode/agents/*.md` y Mallard no los toca.

### El estándar (agentskills.io)

Una skill es un directorio con `SKILL.md` (frontmatter `name` + `description`) + `references/` + `scripts/` + `assets/`. Divulgación progresiva:

1. Al arrancar: solo `name` + `description` (~100 tokens) de **todas** las skills.
2. Al activar: el body del `SKILL.md` (recomendado < 500 líneas / < 5k tokens).
3. A demanda: `references/` y `scripts/`.

Eso **es** el diseño de `panda-kb` y el que ya copiamos de Hydrogen/Addy. No es un detalle de Claude.

`description` máx 1024 caracteres y tiene que decir **qué** y **cuándo**. En Codex/Gemini/Windsurf esa frase **es el classifier**. Si no menciona checkout, PDP, LCP, conversión, Shopify, Woo, Magento, Hydrogen, el experto no salta. `panda-expert` dice explícitamente que **no** se dispare en Shopify/Woo: el hueco lo tiene que ocupar esta description.

Campo `compatibility`: para perf, más adelante `Requires Lighthouse or npx lighthouse`. No en Fase 0.

`allowed-tools` es experimental y no lo usan todos. El candado real de “UX no edita CSS” es el subagente sin Write, no la skill.

### Tres primitivas (no son tamaños distintos)

Regla de Anthropic / Cursor, útil aquí:

| Primitiva | Quién la dispara | Dónde corre | Para Mallard UX/perf |
|-----------|------------------|-------------|----------------------|
| **Skill** | El modelo, por `description` (o `/nombre` en Cursor) | Hilo actual | El oficio + KB. Portable. |
| **Command** | Tú, `/ux` | Hilo actual (o lanza subagente) | Atajo. Claude + OpenCode. |
| **Subagente** | El modelo o el command | **Otro** contexto; devuelve el resultado | Aislar el audit (mucho Grep/Lighthouse) y **quitar Edit**. Claude + Cursor hoy. |

Un subagente no es “una skill más gorda”. Es otra habitación. Cursor lo dice igual: changelog → skill; research ruidoso → subagente.

UX/perf **sí** merecen subagente donde exista: el audit ensucia el hilo (templates, network, JSON de Lighthouse) y el UX **no debe** tener Write. En Codex no hay habitación: la skill corre inline y el “no edites CSS” es instrucción.

### Claude Code

Skills en `~/.claude/skills`. Commands en `~/.claude/commands`. Subagentes en `~/.claude/agents` (Task tool).

Deuda actual: `/panda` y `/ps` invocan `prestashop-experts:panda-expert` (namespace de **plugin**). Los agentes del repo usan `${CLAUDE_PLUGIN_ROOT}/skills/...`. Mallard **también** symlinkea esos mismos `.md` a `~/.claude/agents`. Doble distribución (plugin freelance + Mallard equipo).

**Los expertos nuevos no usan namespace de plugin ni `CLAUDE_PLUGIN_ROOT`.** Cargan `ecommerce-ux-kb` como skill Mallard (`~/.claude/skills/ecommerce-ux-kb` tras `mallard update`). Si un día salen en el plugin, se adapta; no al revés.

`/ux` no debe ser un one-liner “Invoke the plugin subagent”. Cuerpo:

```
Carga la skill ecommerce-ux-kb y síguela.
Si existe el subagente ecommerce-ux-expert, delega (contexto aislado, sin Edit).
No edites CSS; prescribe y pasa a layout-builder.
```

Así el command no se rompe el día que no hay plugin.

### Cursor (2026)

Skills: `.cursor/skills`, `~/.cursor/skills`, y por compat `.claude/skills`, `.codex/skills`, `.agents/skills`. Slash `/skill-name` **es** la skill. Commands clásicos (`.cursor/commands`) se migran a skills con `disable-model-invocation: true`. Mallard no enlaza commands a Cursor: **correcto**, si la skill tiene buen nombre.

Subagentes: `~/.cursor/agents` y compat `~/.claude/agents`. Mallard ya enlaza agents a Cursor. El `.md` estilo Claude (frontmatter `name`, `description`, `tools`) le vale.

Gotcha: `mallard install --all` deja la misma skill en `~/.cursor/skills` **y** Cursor también lee `~/.claude/skills`. Puede listar duplicados. No lo resolvemos en este plan; no instalar de más o aceptar el dup.

Para `/ux` en Cursor: o la skill se llama `ux` (corto, chocará) o el usuario escribe `/ecommerce-ux-kb`. Preferible **no** inventar un alias `ux` skill. El command `/ux` es Claude/OpenCode; en Cursor el disparo es descripción + subagente.

### Codex

Oficio oficial: skills. USER path documentado: `$HOME/.agents/skills`. Repo: `.agents/skills` walking-up. Mallard enlaza `~/.codex/skills` (sigue existiendo) **y**, si detecta `~/.agents`, el adapter generic también. Codex invoca con `$nombre` o por description. Prompts/commands van a skills.

No hay subagentes Codex en el adapter. Cursor docs mencionan `.codex/agents` como compat; no es el producto Codex CLI. **No diseñamos agentes Codex.**

`agents/openai.yaml` (UI, `allow_implicit_invocation`) es extensión OpenAI, no el spec. No lo necesitamos.

### OpenCode

Skills: nativo `skill({ name })`. Lee `.opencode/skills`, `~/.config/opencode/skills`, **y** `.claude/skills` / `.agents/skills`. Commands: `~/.config/opencode/commands` — Mallard sí los enlaza.

**Agents (2026):** markdown en `~/.config/opencode/agents/` con `mode: subagent` y `permission: edit: deny`. Es el candado que queremos para UX. Mallard **`AgentsDir` está vacío.** Hueco del CLI, no de los expertos.

Fase 0 de expertos: commands `/ux` `/perf` que cargan la **skill** (OpenCode tiene `skill` tool). No esperar al adapter.

Más adelante (otro PR, Go): `opencodeAdapter.AgentsDir() = ~/.config/opencode/agents` + workspace `.opencode/agents`. El formato no es el de Claude (`tools: Read, Grep` vs `permission: edit: deny`). Traducir 1:1 es mentira; o un template OpenCode al lado, o un generador. No mezclar frontmatters.

Command OpenCode puede poner `agent: ecommerce-ux-expert` + `subtask: true` **cuando** el agent exista. Hasta entonces, el body carga la skill.

### Gemini CLI y Windsurf

Solo skills. Gemini: `~/.gemini/skills`. Windsurf: skills bajo `~/.codeium/windsurf/skills`. Mallard **no** mapea workflows `.windsurf/workflows` (formato distinto). Correcto: no son el experto.

### Generic `~/.agents`

Path “universal” que Codex y Cursor también miran. Mallard lo rellena si existe `~/.agents`. Bien para el experto-skill.

### Implicaciones (lo que cambia el diseño)

1. **Fase 0 skill-first.** Dos KBs con `description` de disparo, no dos agentes gordos.
2. **Agentes finos**, `tools: Read, Grep, Glob, Bash` **sin** Edit/Write en UX. Perf puede Bash (Lighthouse). `layout-builder` sigue con Write.
3. **Commands duales** (skill, y subagente si existe). Cero `prestashop-experts:`.
4. **`/task` + classifier** es extra de Claude/OpenCode, no el router universal.
5. **Script Lighthouse** en `ecommerce-perf-kb/scripts/` (spec), no en el prompt.
6. **No Go** para nacer. Go solo si un día queremos agents OpenCode.
7. **Handshake UX → layout-builder:** el entregable del experto es una lista de cambios (archivo, qué, por qué) para que el worker no re-diagnostique. Eso es routing gentle-ai (`delegate_only`) en la práctica.
8. **No** `/storefront-audit` paralelo UX+perf en Fase 0. Claude/Cursor pueden spawnear los dos subagentes después, si se usa.

## Fases

### Fase 0 — Skills portables + agentes finos

Orden: KB primero (7 harnesses), agente después (Claude/Cursor), command dual, classifier.

- `skills/ecommerce-ux-kb/` y `skills/ecommerce-perf-kb/` con `SKILL.md` router y `description` que dispare en **cualquier** storefront (checkout, PDP, cart, LCP, conversión, Shopify, Woo, Magento, Hydrogen — el hueco que `panda-expert` rechaza).
- References mínimas: 5 superficies UX + `cwv.md` + `how-to-measure.md`. Oficio en las guidelines; **En este stack** solo si el CWD lo pide.
- `references/output-format.md` compartido de hecho: Critical/High/Medium/Low + evidencia vs hipótesis (mardab) + Quick Wins / High-Impact / Test Ideas (Corey). UX añade bloque **Implementación → layout-builder** (archivos, qué, no el CSS).
- Agentes finos: `claude/agents/ecommerce-ux-expert.md` y `ecommerce-perf-expert.md`. CWD-first, “carga la KB”, formato de salida. UX: `tools: Read, Grep, Glob, Bash` — **sin Edit/Write**. Perf: + Bash para medir. No copiar la KB al body del agente.
- `/ux` y `/perf`: “carga la skill; si hay subagente, delega”. Sin namespace de plugin.
- `task-classifier`: tipos `ux` y `perf` + señales. Solo ayuda a `/task`.
- Cero Go. Cero adapter OpenCode. Cero script Lighthouse todavía (el `how-to-measure.md` dice cómo; `scripts/` en Fase 1 si duele).

### Fase 1 — Que se peleen con proyectos reales

- Un ticket de checkout y uno de home lenta en un storefront real (el que esté abierto), en **más de un harness** si podéis (Claude + Cursor o Codex). Si solo funciona en Claude, la skill está mal y el oficio está en el agente.
- El experto UX debe **delegar** el CSS a `layout-builder` (o no tocar archivos).
- Añadir a la KB solo lo que hayáis tenido que decidir de verdad.
- Si medir a mano cansa: `ecommerce-perf-kb/scripts/` que deje JSON en stdout (Addy).

### Fase 2 — Solo si el núcleo se usa

- SEO y a11y, mismo molde skill-first.
- Opcional, **otro** trabajo: `opencodeAdapter.AgentsDir` + template `permission: edit: deny` (no es este crew).
- Opcional: command `/storefront-audit` que lance UX y perf en paralelo (Claude/Cursor).

## Encaje con gentle-ai (sin SDD)

Sigue valiendo el routing, no la ceremonia:

- Inline: “el logo pisa el buscador” → `layout-builder` (skill/agente con Write).
- Delegate: “no convierten en checkout” → experto UX **sin Write**, luego worker. En Claude/Cursor eso es subagente; en Codex es la misma skill con la regla escrita.
- El classifier no ejecuta (como `task-classifier` hoy). En harnesses sin `/task`, la `description` de la skill hace de classifier.

No hay recibo de precios ni `apply`. El “gate” de siempre: el experto **draft**, tú revisas en local, no deploya.

## Línea apartada (ops)

La investigación anterior (briefing semanal, triage de pedidos, adapters, recibo/`apply`, MCP de PS/Shopify) era un **OS de operaciones**. Encaja en “orquestación y ventas” si se lee como back-office. No es lo que querías.

Si un día hace falta, es otro namespace y otro momento. No mezclarlo con estos expertos: un agente que lee pedidos y otro que opina de LCP no son el mismo oficio.

Notas y fuentes de esa línea (mercado, adapters, gentle-ai RDD) se pueden recuperar del historial de git de este doc si hicieran falta.

## Cotilleo GitHub (agosto 2026)

Filtro: **oficio de tienda genérico** (PDP, cart, checkout, search, CWV de storefront). No “¿sirve para un child PS?”.

El molde “skill/agente experto” ya existe. El combo **UX de storefront + perf de storefront + diagnostica y no maqueta** casi no.

### Lo más parecido a `ecommerce-ux-expert` (oficio de tienda)

| Repo | Stars | Qué es | Por qué no lo copiamos entero |
|------|------:|--------|-------------------------------|
| [hmtkyn/hydrogen-ecommerce-ux](https://github.com/hmtkyn/hydrogen-ecommerce-ux) | 2 | KB por superficie de compra (home/nav/PLP, PDP, cart+checkout, search, mobile) + notas Hydrogen | El mapa de superficies **es** el oficio genérico. El runtime Hydrogen no |
| [dnh33/webshop-ux-expert](https://github.com/dnh33/webshop-ux-expert) | 0 | Consultor de webshop, KB Baymard, audita URL con Playwright | 0 stars, mezcla agentic-commerce. Idea sí: consultor ≠ implementador |
| [lgboim/ux-builder](https://github.com/lgboim/ux-builder) | 2 | 500+ reglas Baymard/NN/g; `references/ecommerce.md` = checkout/cart/PDP/search | Checklist genérico de producto, no CWD. Números de lift dudosos |
| [mardab96/ecommerce-claude-skills](https://github.com/mardab96/ecommerce-claude-skills) | 3 | Pack **agnóstico de plataforma**: Checkout Friction Finder + Product Page Conversion Review (screenshots/URLs). No muta la tienda | El resto (ads, margen, LTV, feeds) es ops de merchant. Shape de 1 y 2: sí |
| [finsilabs/awesome-ecommerce-skills](https://github.com/finsilabs/awesome-ecommerce-skills) | 45 | 178 skills Shopify/Woo/BC/Magento/headless. Storefront & UI (PDP, facetas, search, mega-menu) + checkout-flow-optimization | Setup de admin/apps (“instala Rebuy”). No diagnostica el CWD. Mapa de temas útil |
| [nexscope-ai/eCommerce-Skills](https://github.com/nexscope-ai/eCommerce-Skills) | ~700 | 157 skills de *seller* (Amazon, PPC, checkout genérico) | Marketplace/ops, no criterio sobre un storefront |
| [40RTY-ai/shopify-admin-skills](https://github.com/40RTY-ai/shopify-admin-skills) | 177 | `conversion-optimization` = informes de admin (abandonment, discounts) | Ops Shopify, no UX de ficha/checkout |

### Lo más parecido a `ecommerce-perf-expert`

| Repo | Stars | Qué es | Encaje |
|------|------:|--------|--------|
| [addyosmani/web-quality-skills](https://github.com/addyosmani/web-quality-skills) | ~2.6k | Skills Lighthouse / CWV / a11y / SEO (web genérico) | **El de cotillear para perf.** No sabe de peso de cart/checkout |
| [finsilabs](https://github.com/finsilabs/awesome-ecommerce-skills) `infrastructure-performance` | 45 | CDN de imágenes, cache, edge, load test | Infra de tienda, no LCP de una PDP |
| Auditores CWV sueltos (`web-vitals-auditor`, etc.) | bajos | Lighthouse genérico | Misma familia; no storefront |

### Qué no hay

Nadie junta, con estrellas: experto UX de **superficies de compra** que **no** maqueta + experto perf que **exige** medir + classifier que ya habla con un worker de layout.

Cotillear en serio Hydrogen y Addy: [qué extraemos](#qué-extraemos-de-las-dos). High-star y genéricos: [reinvestigación high-star](#reinvestigación-high-star).

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
   **En este stack:** Dawn / Flatsome / Panda / el checkout del CWD …
   ```

   Su campo `Hydrogen implementation` lo sustituimos por **En este stack** (Dawn, Flatsome, Panda, Liquid, Smarty…). El experto UX no reescribe cómo funciona el builder de turno.
4. **Citar el porqué.** “Baymard: guest checkout” > “sé más usable”.
5. **No inventar porcentajes.** Lo dicen ellos; varios números del propio repo huelen a blog. Nosotros: umbrales Google (CWV) sí; lifts Baymard premium no.
6. **High stakes.** Checkout, pago, a11y: el experto lo marca; no lo diluye en un checklist de home.
7. **Checkout propio vs hosted.** Hydrogen *no* posee el pago (`cart.checkoutUrl` → Shopify Checkout). Woo, Magento, PS, BigCommerce theme y muchos custom **sí**. El experto UX tiene que preguntar si el checkout es editable en este CWD (mardab lo hace igual). Donde sí lo es, vale más: el theme *es* el checkout.

### Hydrogen UX — qué no

- Código `CartForm`, `useOptimisticCart`, Predictive Search API.
- Mezclar LCP/a11y dentro del archivo de mobile (rompe el split de expertos).
- Copiar guidelines verbatim (Baymard premium + su LICENSE). Principios públicos + oficio vuestro.
- “Show code, don't just describe” para el experto UX. Aquí quien escribe CSS es `layout-builder`. El experto enseña *qué* y *dónde*; el worker el *cómo*.

### Addy web-quality-skills — qué sirve

1. **Un experto, varias lentes.** Ellos instalan 6 skills (audit / performance / cwv / a11y / seo / best-practices). Nosotros no. `ecommerce-perf-expert` carga `cwv.md` o `images.md` según el trigger, igual que el UX carga una superficie. `web-quality-audit` ≈ nuestro classifier cuando el ticket es “la web va mal” sin más.
2. **CWV no es otra persona.** Es un reference (`LCP.md` / nuestro `cwv.md`) con tabla Good / Needs work / Poor (umbrales Google, 75º percentil).
3. **Presupuesto en tabla**, luego retocado a storefront (pixels, sliders, reviews, tag managers). Forma de Addy; números nuestros cuando midamos.
4. **Severidad: Critical / High / Medium / Low.** Prioridad: fallo CWV y barrera a11y antes que estilo. El output del experto perf es una lista rankeada, no un ensayo.
5. **SKILL.md < 500 líneas; references cargables solos.** Progressive disclosure. Hydrogen peca de files de 200+ líneas; Addy pide ~200. Nosotros: un tema por file, TOC arriba.
6. **Medir fuera del contexto.** `scripts/analyze.sh`: logs en stderr, JSON en stdout. El experto perf no “cree” el LCP: corre Lighthouse en la URL local (o pide el JSON). Fase 0 puede ser “cómo medir”; el script puede esperar.
7. **Bad / Good en perf.** Un `<img>` sin width/height vs con `width` `height` `fetchpriority`. Más útil que prosa. En UX el par es hallazgo en *este* tpl vs prescripción, no un snippet React.
8. **Agnóstico primero.** Vanilla/HTML, luego nota de framework. Igual que nosotros: principio → **En este stack**.

### Addy — qué no

- Early Hints, Speculation Rules, `getServerSideProps` como consejo por defecto en un theme clásico (Liquid, Smarty, PHP). En Hydrogen/Next sí pueden aplicar: **En este stack**.
- Skill `best-practices` (CSP, HSTS) como experto de ecommerce. Eso no es conversión.
- Grep de HTML estático como auditoría de un storefront con templates (`.liquid`, `.tpl`, `.phtml`).
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
| Tipo de página primero (Corey) | ya lo cubre el router de superficies |
| Quick Wins / High-Impact / Test Ideas (Corey) | capa extra sobre Critical/High/Medium/Low: fácil vs caro vs hipótesis |
| Heurísticas + dark patterns (Wondelai) | lentes en checkout/móvil; no un skill Nielsen aparte |
| Métricas antes de grep (Vercel Optimize) | el experto perf no barre el repo hasta tener Lighthouse/JSON |
| Evidence vs hypothesis (mardab) | si no hay screenshot/URL/dato, se marca hipótesis |
| Mapa de temas Finsi (PDP, facetas, checkout) | confirma las 5 superficies; no 178 skills de admin |

## Reinvestigación high-star

Agosto 2026. Filtro corregido: **ecommerce genérico** (PDP, cart, checkout, search, LCP de storefront), no “¿encaja en PrestaShop?”.

No hay un repo de muchas estrellas que sea ese oficio como dos expertos diagnose-and-delegate. Las estrellas grandes son UI genérica, CRO de landing SaaS, índices y herramientas. El oficio de *tienda* vive en repos medianos/chicos.

### Dónde están las estrellas (adyacente, no tienda)

| Stars | Repo | Qué es de verdad | Encaje |
|------:|------|------------------|--------|
| 170k | [anthropics/skills](https://github.com/anthropics/skills) | Formato oficial. `frontend-design` = identidad visual | Molde SKILL.md. **No** es CRO de checkout |
| 117k | [nextlevelbuilder/ui-ux-pro-max-skill](https://github.com/nextlevelbuilder/ui-ux-pro-max-skill) | Design system / “taste” UI. Tiene página `checkout` como layout visual | No es conversión de tienda. Si alguien pide “que se vea menos plantilla”, no es `/ux` |
| 73k / 30k | awesome-\* de skills | Índices | Ruido |
| 44.5k | [coreyhaines31/marketingskills](https://github.com/coreyhaines31/marketingskills) | **El CRO con más estrellas.** Landing/SaaS: headline, CTA, pricing, forms. Purchase aparece; PDP/cart/checkout de tienda no | Forma de salida + tipo de página. No el embudo signup→trial |
| 31k / 95k / 8.6k | Lighthouse, Playwright, web-vitals | Herramientas | El experto las **invoca**. No son el experto |
| 30k | [vercel-labs/agent-skills](https://github.com/vercel-labs/agent-skills) | `vercel-optimize`: métricas antes de grep; `file:line` | Doctrina útil. Recetas Next/Vercel = **En este stack**, no el oficio |
| 2.6k | [addyosmani/web-quality-skills](https://github.com/addyosmani/web-quality-skills) | Pack CWV | Sigue siendo el mejor pack high-star de **perf web**. No de cart/checkout |
| 1.9k | [wondelai/skills](https://github.com/wondelai/skills) | Heurísticas + CRE | Lentes. No 50 libros |

### Oficio de tienda (genérico). Pocas estrellas, este es el cotilleo

| Stars | Repo | Oficio | Qué robamos | Qué no |
|------:|------|--------|-------------|--------|
| 45 | [finsilabs/awesome-ecommerce-skills](https://github.com/finsilabs/awesome-ecommerce-skills) | 178 skills multi-plataforma (Shopify/Woo/BC/Magento/headless). Storefront & UI + checkout-flow-optimization | Mapa de temas: PDP, facetas, search, mega-menu, responsive, checkout. Guest, costes tarde, express pay, trust junto al pago | “Instala Rebuy / Checkout X / CartFlows”. Eso es ops de merchant, no CWD |
| 177 | [40RTY-ai/shopify-admin-skills](https://github.com/40RTY-ai/shopify-admin-skills) | Informes de admin (abandonment, discounts) | — | Ops. No UX de ficha |
| 86 | [Weaverse/shopify-hydrogen-skills](https://github.com/Weaverse/shopify-hydrogen-skills) | Cómo construir Hydrogen | — | Implementación de plataforma |
| 3 | [mardab96/ecommerce-claude-skills](https://github.com/mardab96/ecommerce-claude-skills) | **El shape más cercano.** Agnóstico: screenshots/URLs, no muta la tienda. Checkout Friction Finder + Product Page Conversion Review | Evidence vs hypothesis; preguntar si el checkout es editable; ranked friction; no prometer lift | Skills 3–20 (ads, margen, LTV, feeds) son otra línea |
| 2 | [hmtkyn/hydrogen-ecommerce-ux](https://github.com/hmtkyn/hydrogen-ecommerce-ux) | Superficies de compra + Hydrogen | El mapa. Ya extraído | Runtime Shopify |
| 2 | [lgboim/ux-builder](https://github.com/lgboim/ux-builder) | Baymard/NN/g en `ecommerce.md` | Confirma checkout/cart/PDP/search como dominio | Lifts tipo “+28%”; no CWD |
| 699 | nexscope eCommerce-Skills | Amazon/seller | — | No storefront |
| 0 | [dnh33/webshop-ux-expert](https://github.com/dnh33/webshop-ux-expert) | Consultor webshop + Playwright | Consultor ≠ implementador | 0 stars, UCP mezclado |

`topic:agent-skills ecommerce`: el techo con estrellas es seller/ops (nexscope 699, 40RTY 177). El techo de **criterio de storefront** es Finsi 45, y el shape que queremos es mardab/Hydrogen (2–3★).

### Qué sirve (sin vendorizar)

**Finsi storefront + checkout (45★, genérico de verdad)**

1. Plataformas en tabla: Shopify (checkout a menudo **hosted**), Woo/BC/custom (**editable**). El experto pregunta el control, no asume PS ni Hydrogen.
2. Temas de Storefront & UI que ya teníamos: PDP, facetas, search autocomplete, mega-menu, responsive. Quick-view / 360 / wishlist = extra, no Fase 0.
3. Checklist de checkout público: guest, campos de más, shipping cost tarde, express above the form, trust junto al pago, validate on blur.
4. **No** el modo “recomienda la app del marketplace”.

**mardab (3★, agnóstico)**

1. Input: URL, screenshot, políticas. No hace falta Admin API.
2. Evidence tag + confidence. Sin dato → `hypothesis` / `needs_data`.
3. “¿El checkout se puede editar o lo bloquea la plataforma?”
4. Guardrail: no mutar, no prometer lift. Ya era nuestra regla.

**Corey `cro` (44.5k)**

Tipo de página primero (nosotros: superficie de compra, no pricing SaaS). Salida Quick Wins / High-Impact / Test Ideas. No el embudo landing→signup→trial.

**Wondelai (1.9k)**

Investigar antes de tips; objeciones; dark patterns; 44px; un solo scale de severidad.

**Vercel Optimize (30k)**

Métricas antes de grep; `file:line`. Recetas Fluid/ISR/Observability = solo si el CWD es Vercel.

**Addy (2.6k) + Lighthouse**

CWV como lente. Medir fuera del prompt. Playwright live-audit: opcional.

### Qué no hacer con esto

- Filtrar repos por “¿esto habla de PrestaShop?”. El oficio no es de plataforma.
- Vendorizar las 178 skills de Finsi ni el pack marketing de Corey.
- Mezclar CRO de landing SaaS con superficies de tienda.
- Un segundo experto “heurísticas” aparte del UX.
- Tratar índices awesome-\* o toolkits Magento/Woo/Hydrogen como el experto UX.

High stars ≠ el producto. Oficio de tienda = Hydrogen (superficies) + mardab (diagnóstico sin mutar) + Finsi (mapa multi-plataforma). Medición = Addy. De los grandes: forma (Corey), ética (Wondelai), métricas-primero (Vercel).

## Fuentes de oficio (para las KBs, no como runtime)

- Oficio de clientes / CWD (primera fuente).
- [Baymard](https://baymard.com/) — checkout, mobile, search (criterio, no copiar papers enteros).
- [web.dev / CWV](https://web.dev/explore/learn-core-web-vitals) — medición.
- Superficies: [hydrogen-ecommerce-ux](https://github.com/hmtkyn/hydrogen-ecommerce-ux).
- Diagnóstico sin mutar: [mardab96/ecommerce-claude-skills](https://github.com/mardab96/ecommerce-claude-skills) (friction + PDP review).
- Mapa multi-plataforma (no las apps): [finsilabs/awesome-ecommerce-skills](https://github.com/finsilabs/awesome-ecommerce-skills) storefront-ui + checkout-flow-optimization.
- Medición: [web-quality-skills](https://github.com/addyosmani/web-quality-skills).
- Forma de salida CRO (no el embudo SaaS): [marketingskills/skills/cro](https://github.com/coreyhaines31/marketingskills).
- Heurísticas / objeciones: [wondelai/skills](https://github.com/wondelai/skills).
- Métricas antes de grep: doctrina de [vercel-optimize](https://github.com/vercel-labs/agent-skills/tree/main/skills/vercel-optimize).
- KBs de plataforma en Mallard (PS/Panda/EF, etc.) solo para el *cómo* cuando el CWD las pide.

Harnesses (para el plan de empaquetado, no para las KBs):

- [Agent Skills spec](https://agentskills.io/specification) — `SKILL.md`, progressive disclosure, `description` ≤ 1024.
- [Cursor Skills](https://cursor.com/docs/skills) y [Subagents](https://cursor.com/docs/subagents) — slash = skill; agents en `~/.cursor/agents`.
- [Codex skills](https://developers.openai.com/codex/skills) — `$HOME/.agents/skills`, invocación `$`.
- [OpenCode skills](https://opencode.ai/docs/skills) y [agents](https://opencode.ai/docs/agents) — `permission: edit: deny`; Mallard aún no enlaza agents.
- Mallard adapters: `internal/agents/*.go`. `docs/adding-skills.md` está desfasado (dice agents = Claude only).
