# Equipo de agentes ecommerce — plan propio

Producto **nuestro**: un kernel de operación comercial (orquestación, ventas, optimización) que habla un modelo canónico. Cada tienda entra por un **adaptador que escribimos nosotros**. Sirve para cualquier ecommerce (PrestaShop, Shopify, Woo, Magento, custom, marketplace).

## Decisión (esta iteración)

- **No** dependemos del MCP Server de PrestaShop, MCP Tools Plus, Shopify MCP/UCP, Busony, VTEX, Agentforce, CrewAI Cloud, n8n ni ningún bus de terceros.
- **No** es un skill más de Mallard pegado a una tienda PS. Mallard sigue siendo el toolkit de *desarrollo* PrestaShop. Esto es un **producto hermano**: ops comercial, platform-agnostic.
- Los agentes **nunca** llaman a la API nativa de una plataforma. Solo ven el contrato canónico + políticas + cola de aprobación.
- Un conector de tienda **sí** hace falta (webservice, Admin API, REST propia). Eso no es “plataforma extra”: es el I/O de la tienda. El conector es código nuestro, fino, sustituible.
- Si algún día exponemos MCP, será **una fachada opcional de *nuestras* tools**, no una dependencia. El runtime habla CLI/JSON (y luego HTTP) que controlamos.

La investigación de mercado (Shopify, VTEX, PS MCP, CrewAI, Atlas Mercator) queda en el [apéndice](#apéndice-qué-hay-ya-montado). Sirve para no copiar su lock-in. El diseño de abajo no los usa.

## Por qué este corte

Cualquier crew “universal” que se ate a un MCP ajeno hereda el modelo, los scopes y el ritmo de esa plataforma. PrestaShop MCP no sirve en Shopify. Shopify UCP no sirve en una tienda custom. El patrón que sí escala —y que ya usan los kernels serios— es **un contrato + N adaptadores**.

Mallard ya piensa así por otro lado: `internal/agents.Adapter` es un conector por herramienta de coding. Aquí el analog es un conector por *motor de tienda*.

```
  /briefing  /catalog-audit  /order-triage
                 │
                 ▼
        ┌─────────────────┐
        │  Orchestrator   │  enruta, reconcilia, aplica políticas
        │  + playbooks    │
        └────────┬────────┘
                 │  solo tipos canónicos
                 ▼
        ┌─────────────────┐
        │  Commerce API   │  catalog / inventory / orders / sales / pricing / promo
        │  (nuestro)      │
        └────────┬────────┘
                 │
     ┌───────────┼───────────┬────────────┐
     ▼           ▼           ▼            ▼
  adapter.ps  adapter.shopify  adapter.woo  adapter.fake
  (WS nuestro) (Admin API)     (REST)       (fixtures)
     │           │           │            │
     ▼           ▼           ▼            ▼
  cualquier tienda que exponga un API que sepamos mapear
```

## Principios

1. **Contrato primero.** Product, Variant, Order, Stock, Price, Promo son *nuestros* tipos. El adaptador traduce. Si un motor no tiene “combinación”, el adaptador inventa una variante 1:1 y lo declara en `capabilities`.
2. **Capabilities, no ifs de plataforma en los agentes.** `inventoryRealtime`, `promos`, `orderNotes`, `writePrice`… El playbook se degrada, no se reescribe.
3. **Un objetivo por agente.** Pricing no escribe fichas. Pedidos no tocan SEO.
4. **Read-first.** Fase 1 solo lee y propone. Mutar precio/stock/promo = propuesta → humano → `apply`.
5. **Playbooks deterministas donde se pueda.** Auditoría de catálogo y cola de pedidos son reglas + datos. El LLM narra y reconcilia; no es el bus.
6. **PII mínima.** El kernel anonimiza antes de mandar nada a un modelo. Coste, email, teléfono no salen del adaptador sin flag.
7. **Cero promesa de lift.** Se mide contra el baseline de esa tienda.
8. **Sin frameworks de orquestación ajenos.** El orquestador es código nuestro (router + políticas). No CrewAI, no LangGraph Cloud.

## Contrato canónico (v0)

Lo que todo adaptador debe implementar o marcar `unsupported`.

| Área | Tipos | Operaciones |
|------|-------|-------------|
| Identidad | `Store` (`id`, `platform`, `locale`, `currency`, `timezone`) | `Ping`, `Capabilities` |
| Catálogo | `Product`, `Variant`, `Category`, `Media`, `SEO` | `ListProducts`, `GetProduct`, `ListCategories` |
| Inventario | `Stock` (por variant + location opcional) | `ListStock`, `StockGaps` (derivable en kernel si el adapter solo da raw) |
| Precios | `Price` (list, sale, cost opcional, margin) | `ListPrices`, `ProposePrice` / `ApplyPrice` (gated) |
| Promos | `Promo` (tipo, ventana, stack, techo) | `ListPromos`, `PromoSanity` |
| Pedidos | `Order`, `Line`, `PaymentState`, `FulfillmentState`, `Refund` | `ListOrders`, `GetOrder`, `ProposeStatus` (gated) |
| Ventas | agregados, no filas PII | `SalesKpis(from, to, groupBy)` — el adapter puede calcular o el kernel agrega pedidos |
| Clientes | `CustomerRef` (id opaco) | Solo ids + segmento; PII off by default |

`Capabilities` ejemplo:

```yaml
reads: [catalog, inventory, orders, prices, promos, sales]
writes: []          # vacío en fase 1
notes: false
multiLocation: false
variants: true
```

Un adaptador incompleto es válido. El orquestador no llama lo que no existe.

## Runtime (nuestro, local-first)

Producto hermano de Mallard: un binario (Go, mismo ADN: un CLI, cero daemon SaaS).

```
commerce                  # CLI
  store add|ls|ping
  pull                    # snapshot normalizado vía adapter
  briefing
  catalog-audit
  order-triage
  propose ls|show
  apply <id>              # solo si policy + humano
  audit tail
```

Estado por tienda (en disco, nuestro):

```
~/.commerce/stores/<id>/
  store.yaml              # adapter, endpoint, secret ref, policies
  snapshot/               # JSON canónico (cache, no source of truth)
  proposals/
  audit.jsonl
```

`store.yaml` (forma, no implementación):

```yaml
id: acme-prod
adapter: prestashop        # o shopify | woo | magento | http | fake
endpoint: https://shop.example.com
auth: env:ACME_API_KEY
locale: es-ES
currency: EUR
policies:
  readOnly: true
  marginFloor: 0.35
  maxDiscount: 0.15
  unpaidDays: 14
```

Los “agentes” son paquetes nuestros: prompt + tools permitidas + playbook. Se invocan desde el CLI o, si queremos DX tipo Mallard, desde un comando que **llama al CLI**, no a un MCP ajeno.

```
claude/commands/briefing.md  →  exec: commerce briefing --store $STORE
```

Eso es opcional y no acopla el kernel a Claude.

## Equipo (igual de especialización, otro bus)

```
                commerce-orchestrator
                 (router + políticas)
     ┌──────────┬──────────┼──────────┬──────────┐
     ▼          ▼          ▼          ▼          ▼
 sales-     merch-    inventory-  pricing-   order-
 analyst    andiser   ops         advisor    triage
     │          │          │          │          │
     └──────────┴────┬─────┴──────────┴──────────┘
                     ▼
              Commerce API (nuestra)
```

| Agente | Misión | Éxito |
|--------|--------|-------|
| `orchestrator` | Clasifica el trabajo, reparte, reconcilia conflictos, decide `read \| propose \| apply-with-cap \| escalate` | Ruta correcta + decisión explicable |
| `sales-analyst` | Ventas, AOV, top/bottom SKU, vs periodo | Briefing <5 min de lectura |
| `merchandiser` | Sin media, categoría vacía, ficha pobre, variant no comprable | Checklist priorizado |
| `inventory-ops` | Roturas, sell-through, activo-pero-no-comprable | Cero sorpresas de stock |
| `order-triage` | Impagos, atascados, reembolsos | Cola rankeada |
| `pricing-advisor` | Recomienda precio dentro de suelo/techo | Nunca aplica solo en v1 |
| `promo-optimizer` | Promos que se pisan o rompen margen | “esta regla pierde dinero” |

Más adelante (no v1): `seo-listing`, `cx-support`, `intel-scout`. Customer-facing es otro producto de riesgo.

Lo que **no** es un agente: un “Ecommerce God”, un chatbot FO, repricing 24/7, un reemplazo de GA4.

## Flujos (ops, no “list products”)

| Flujo | Quién | Output |
|-------|-------|--------|
| Weekly briefing | analyst → inventory → merchandiser → orchestrator | Qué vender, qué se acaba, qué ficha está rota |
| Pre-launch audit | merchandiser + inventory + promo | Fotos, comprable, promos vivas |
| Order triage | order-triage | Cola por urgencia |
| Promo sanity | promo + pricing + inventory | ¿Pisa otra regla? ¿rompe margen? |
| Listing refresh | seo-listing + merchandiser | Copy; humano publica |

Salida del orquestador (estable, testeable):

```
TYPE: briefing | catalog-audit | order-triage | pricing-review | promo-review | unknown
ROUTE: <agent>
WHY: …
CONFIDENCE: high|medium|low
AUTONOMY: read | propose | apply-with-cap | escalate
```

## Adaptadores (nuestro código, una plataforma cada uno)

El kernel **no** conoce SKUs de PrestaShop ni GraphQL de Shopify.

| Adapter | Habla con | Cuándo |
|---------|-----------|--------|
| `fake` | Fixtures JSON canónicos | Día 1. Tests, demos, CI. Obligatorio. |
| `http` | REST/JSON mapeado por `mapping.yaml` | Tienda custom o headless rara |
| `prestashop` | Webservice nativo **nuestro** (no el módulo MCP) | Primera tienda real, donde ya tenemos oficio |
| `shopify` | Admin API | Segunda plataforma, valida que el contrato no es “PS disfrazado” |
| `woocommerce` / `magento` | REST | Cuando haya cliente |

Regla: un adaptador nuevo = implementación del contrato + test de conformidad contra las mismas fixtures. Si el test de `fake` pasa y el de `shopify` no mapea `Variant.buyable`, se arregla el adapter, no el merchandiser.

No scrapeamos back-offices. No SQL directo a la BD de la tienda (rompe upgrades y GDPR).

## Relación con Mallard

| | Mallard | Este producto |
|--|---------|----------------|
| Job | Desarrollar temas/módulos PS | Operar y optimizar **cualquier** tienda |
| Usuario | Dev | Ops / growth / nosotros en clientes |
| I/O | repo git, Lando | APIs de tienda vía adapters |
| Agentes | prestashop-expert, panda-expert… | sales-analyst, merchandiser… |

Puente futuro (opcional): el merchandiser marca “sin foto” → abre tarea Intervals → Mallard `ps-image-regen`. Eso es A2A *nuestro*, no un protocolo externo.

**No** meter el kernel dentro de `skills/` de Mallard. Contamina el CLI de desarrollo y acopla el modelo a PrestaShop. Repo hermano (o monorepo con `cmd/commerce` claramente separado). Recomendación: **repo hermano**, mismo estilo de release (Go, un binario).

## Fases

### Fase 0 — Kernel vacío + fake

- Tipos canónicos + interface `Adapter`.
- CLI: `store add`, `ping`, `pull`.
- Adapter `fake` con un catálogo/pedidos de fixture.
- Policy engine: `readOnly`, techos.
- Audit log.
- Cero LLM. Cero red de tienda real.

### Fase 1 — Playbooks read-only sobre el fake (y un adapter real)

- `catalog-audit`, `order-triage`, `sales.kpis` como código (reglas).
- Un adapter real (recomendado: PrestaShop webservice *nuestro*, porque lo sabemos operar) **o** `http` si el piloto no es PS.
- Briefing: el LLM solo redacta a partir de JSON que ya calculó el kernel.
- Piloto en staging. `writes: []`.

### Fase 2 — Segundo adapter

- Shopify (u otra distinta a la primera). Si el contrato cruje, se corrige ahora, no a la décima tienda.
- Cola de `proposals` + `apply` con diff.

### Fase 3 — Dinero con techo

- `pricing-advisor`, `promo-optimizer`.
- Writes solo con `marginFloor` / `maxDiscount` y aprobación.

### Fase 4 — DX y crecimiento

- Comandos opcionales en el coding agent que llaman al CLI.
- Más adapters.
- Fachada MCP **nuestra** solo si hace falta enchufar Claude Desktop. Sigue siendo I/O nuestro.

## Riesgos

| Riesgo | Mitigación |
|--------|------------|
| El contrato se vuelve “PrestaShop con otros nombres” | Segundo adapter distinto en Fase 2; tests de conformidad |
| LLM muta un precio | Writes off; `apply` explícito; audit |
| PII al modelo | Anonimizar en el kernel; sales solo agregados |
| Multi-store / multi-currency | `Store` es la unidad; un briefing = una store |
| “Plataforma extra” se cuela (SaaS de agents) | Prohibido en runtime. LLM es un vendor de inferencia, como ya usamos para codear; no es el orquestador |
| Scope creep FO / checkout agentic | Fuera de v1. Eso es UCP/AP2 y otro producto |

## Qué ya no hay que decidir

- ¿MCP Tools Plus o MCP oficial? **Ninguno.**
- ¿CrewAI / LangGraph? **No.**
- ¿Agnóstico o solo PS? **Agnóstico, adapters nuestros.**
- ¿Mallard feature o producto? **Producto hermano.**

## Qué sí hay que decidir para codear Fase 0

1. **Nombre y repo** — hermano de Mallard (recomendado) vs `cmd/commerce` en este repo.
2. **Primer adapter real** — PrestaShop webservice nuestro (oficio) vs `http` genérico si el piloto no es PS. `fake` se hace igual.
3. **Tienda piloto / fixtures** — sin un JSON de tienda real (anonimizado) o un staging, Fase 1 es teatro.

Cuando esas tres estén claras, Fase 0 es mecánica: contrato, `fake`, CLI, tests.

## Apéndice — qué hay ya montado

El patrón “equipo + orquestador” existe. Casi todo está atado a una plataforma o a un bus ajeno. Por eso no lo usamos como runtime.

| Cosa | Qué es | Por qué no es nuestro bus |
|------|--------|---------------------------|
| Shopify Agents + UCP | MCP catálogo/cart/checkout | Lock-in Shopify |
| VTEX AI Workspace | Agentes catálogo/promo/search/BI | Cerrado, enterprise |
| Salesforce Agentforce | Merchandiser, Personal Shopper | Vive en Salesforce |
| PrestaShop MCP Server + Tools Plus | Tools oficiales / BusinessTech | Solo PS; approve/reject y GDPR son ideas a **reimplementar** |
| Busony | Claude + MCP/WebMCP sobre PS | Servicio de terceros |
| Atlas Mercator | LangGraph listing/support/intel | Demo; ERP mock |
| CrewAI Shopify | Un “Ecommerce Manager” | Un rol + SaaS |
| Nagent / CREAO | Plantillas SaaS | No es orquestación nuestra |
| Omnix / thorprovider adapters | Contrato + N adapters | Confirma el patrón; no lo adoptamos como dep |

Protocolos (conocer, no depender): MCP (fachada opcional *nuestra*), A2A (más tarde, entre *nuestros* crews), UCP/AP2 (solo si un día hacemos storefront agentic).

Fuentes: [PS MCP](https://docs.mcp.prestashop.com/en/0-getting-started/introduction/), [Tools Plus](https://faq.businesstech.fr/en/faq/616-what-is-mcp-tools-plus), [Shopify Agents](https://shopify.dev/docs/agents), [VTEX Vision](https://www.vtex.com/en-us/vtex-vision), [Atlas Mercator](https://github.com/hhdhh/atlas-mercator), [Decision Crew](https://alexgenovese.com/multi-agent-ai-for-ecommerce-how-agentic-systems-are-reshaping-conversion-in-2026/), [Omnix gateway](https://github.com/OmnixHQ/omnix-gateway).
