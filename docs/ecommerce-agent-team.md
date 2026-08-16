# Equipo de agentes ecommerce — investigación y plan

Propuesta para extender Mallard más allá del crew de **desarrollo** (theme, módulos, layout) hacia un crew de **operación comercial**: orquestación, ventas y optimización.

Mallard hoy es un toolkit de desarrollo PrestaShop. Este documento no implementa agentes nuevos: define qué existe en el mercado, qué no hay que reinventar, y qué equipo tiene sentido construir encima de PrestaShop MCP + el patrón actual de skills / commands / subagents.

## 1. Qué hay ya montado (investigación)

El patrón “equipo de agentes + orquestador” **ya existe** en 2025–2026. Casi nadie lo tiene bien resuelto en PrestaShop de forma abierta y reutilizable. Lo que hay se agrupa en cuatro capas.

### 1.1 Plataformas que venden el crew cerrado

| Producto | Qué montaron | Encaje con nosotros |
|----------|--------------|---------------------|
| [Shopify Agents + UCP](https://shopify.dev/docs/agents) | MCP de catálogo/cart/checkout, Agentic Storefronts, Universal Commerce Protocol (con Google), pagos AP2 | Referencia de protocolos. No portable a PS. |
| [VTEX AI Workspace](https://www.vtex.com/en-us/vtex-vision) | 4 agentes nativos: catálogo, promociones, search, BI. Orquestación insight → decisión → ejecución | El analog más cercano a “OS de ecommerce”. Cerrado, enterprise. |
| [Salesforce Agentforce](https://www.salesforce.com/agentforce/) | Merchandiser, Personal Shopper, Buyer Agent; Flows + MuleSoft | Útil si el cliente vive en Salesforce. No es PrestaShop. |
| Adobe Commerce / Brand Concierge | Agentes de discovery y catálogo | Mismo problema: stack distinto. |

Lección: el mercado enterprise ya convergió en **agentes especializados + workspace de aprobación**, no en un chatbot único.

### 1.2 PrestaShop: hay puente, no hay equipo

| Pieza | Estado | Qué cubre |
|-------|--------|-----------|
| [PrestaShop MCP Server](https://docs.mcp.prestashop.com/en/0-getting-started/introduction/) (oficial, gratis, PS 8.2/9) | Producción | Tools de productos, stock, combinaciones, pedidos, clientes. Módulos terceros pueden declarar `#[McpTool]`, prompts y resources. |
| [MCP Tools Plus](https://faq.businesstech.fr/en/faq/616-what-is-mcp-tools-plus) (BusinessTech) | Comercial | Reportes, roturas de stock, auditoría SEO, promociones, hilos de CS, approve/reject de acciones IA, anonimización GDPR. |
| [Busony / Claude-Powered PS](https://busony.com/en/sites-ia/prestashop) | Servicio | Un Claude orquesta BO + agentes FO (shopping + support) vía MCP/WebMCP. |
| [Artículo de Nicolas Dabène](https://nicolas-dabene.fr/en/blog/the-end-of-the-lone-coder-why-future-developers-will-be-ai-orchestrators-and-how-to-get-started-with-prestashop) | Ensayo | Constelación: merchandising (embeddings), fraude, soporte. MCP como piedra angular. |
| [MCPBundles — workflows PS](https://www.mcpbundles.com/blog/prestashop-ai-store-operations-workflows) | Docs | Insiste en **workflows** (pre-launch audit, Friday order triage, localization), no lookups sueltos. ~76 tools en su skill. |

Lección: PrestaShop ya expone **acciones**. Falta el **equipo con roles, guardrails y orquestación**. Eso es el hueco de Mallard.

### 1.3 Open source y frameworks

| Proyecto | Stack | Roles |
|----------|-------|-------|
| [Atlas Mercator](https://github.com/hhdhh/atlas-mercator) | LangGraph supervisor + RAG | `ListingOptimizer`, `SupportAgent`, `MarketingCopilot`, `IntelScout` + orquestador THOUGHT→PLAN→EXECUTE→SYNTHESIZE. Demo/ERP mock; Shopify real en roadmap. |
| [IBM + CrewAI retail](https://github.com/IBM/ibmdotcom-tutorials/blob/main/crew-ai-projects/crewAI-multiagent-retail-example.md) | CrewAI + watsonx | Crew de “retail advisors” para shelf optimization. Tutorial, no tienda. |
| CrewAI Shopify integration | CrewAI Enterprise | Un agente “E-commerce Manager” con tools de productos/pedidos/analytics. Un rol, no un equipo. |
| [Nagent](https://nagent.ai/solutions/ecommerce), [CREAO](https://creao.ai/for-marketers/ecommerce-manager) | SaaS | Catálogo, pricing, markdown, SEO, inventory alerts, weekly briefing. Plantillas, no orquestación propia. |

Patrón recurrente de consultoría ([Decision Crew](https://alexgenovese.com/multi-agent-ai-for-ecommerce-how-agentic-systems-are-reshaping-conversion-in-2026/)):

```
Orchestrator
  ├── Profile / CRM
  ├── Inventory
  ├── Pricing
  └── Logistics
```

Los lifts de conversión que se publican (25–40%, AOV 2.3x) son **anecdóticos de vendors**. Servir de hipótesis, no de KPI prometido.

### 1.4 Protocolos que hay que conocer (no implementar todos)

| Protocolo | Para qué | ¿Lo necesitamos ya? |
|-----------|----------|---------------------|
| **MCP** | Tools/resources hacia la tienda | Sí. Es el bus. |
| **A2A** | Agente ↔ agente sin pasar siempre por el supervisor | Más adelante. |
| **UCP** | Discovery + cart + checkout para agentes de compra | Solo si queremos storefront agentic (ChatGPT/Gemini). Shopify-first hoy. |
| **AP2** | Mandatos de pago de un agente | No, hasta UCP. |

### 1.5 Conclusión de mercado

1. **No reinventar el conector.** PrestaShop MCP Server + (opcional) MCP Tools Plus ya son el I/O.
2. **No copiar VTEX/Shopify.** Su ventaja es el dato nativo. La nuestra es PrestaShop + Panda/ElementFlow + el crew de desarrollo que ya existe.
3. **El diferencial de Mallard** es unir **dev crew + commerce crew** en el mismo CLI: el mismo equipo que arregla el child theme puede pedir un briefing de ventas o una auditoría de catálogo.
4. **Nadie publicó un crew PrestaShop open-source** con roles, guardrails y comandos listos. Hay puente (MCP) y hay servicios (Busony). No hay “Mallard Commerce”.

## 2. Qué tiene Mallard hoy (y qué no)

Crew actual — solo desarrollo:

| Agente | Dominio |
|--------|---------|
| `prestashop-expert` | Core PS 8/9 |
| `panda-expert` | Panda + `st*` + Easy Builder |
| `layout-builder` | Layout/CSS child theme |
| `task-classifier` | Enruta tareas Intervals → skill/agente |

No hay agente que lea pedidos, márgenes, stock, cart rules o GA4. `task-classifier` ya es un **orquestador mínimo** (clasifica y enruta, no ejecuta). Ese patrón se reutiliza.

## 3. Principios de diseño

1. **Un objetivo por agente.** El de pricing no escribe fichas. El de pedidos no toca SEO. Los fallos quedan aislados.
2. **Read-first.** Fase 1 solo lee y recomienda. Escribir precio, stock o cart rule exige aprobación humana (el approve/reject de MCP Tools Plus es el modelo).
3. **Orquestador reconcilia, no adivina.** Si inventario dice “poco stock” y pricing dice “margen alto”, el orquestador elige escasez, no descuento. Desacuerdos por encima de umbral → humano.
4. **MCP es el bus; Mallard es el cerebro.** Skills y agentes viven aquí. Las mutaciones van a la tienda por MCP, no por SQL ni back-office scraping.
5. **Misma forma que el resto de Mallard.** `claude/agents/*.md` + skill + comando `/…`. No meter CrewAI/LangGraph en el binario Go salvo que un flujo lo pida de verdad (estado largo, HITL persistente).
6. **GDPR y AI Act.** Anonimizar PII antes de mandar a un LLM público. Decisiones que afectan al consumidor (precio dinámico, oferta personalizada) deben ser auditables y, en UE, transparentes.
7. **No prometer lifts.** Medir contra baseline de la tienda.

## 4. El equipo propuesto

Dos crews. El de desarrollo no se toca. El comercial es nuevo.

```
                    ┌─────────────────────────┐
                    │  commerce-orchestrator  │
                    │  (enruta + reconcilia)  │
                    └───────────┬─────────────┘
          ┌──────────┬──────────┼──────────┬──────────┐
          ▼          ▼          ▼          ▼          ▼
     sales-      merch-     inventory-  pricing-   order-
     analyst     andiser    ops         advisor    triage
          │          │          │          │          │
          └──────────┴────┬─────┴──────────┴──────────┘
                          ▼
                   PrestaShop MCP
              (tools / prompts / resources)
```

### 4.1 Núcleo (construir primero)

| Agente | Misión | Lee | Escribe (solo con OK) | Señal de éxito |
|--------|--------|-----|------------------------|----------------|
| `commerce-orchestrator` | Clasifica el pedido (“briefing semanal”, “auditoría pre-lanzamiento”, “viernes de pedidos”) y reparte. Reconcilia conflictos. | Salidas de los especialistas | Nunca muta la tienda | Ruta correcta + decisión explicable |
| `sales-analyst` | KPIs: ventas, conversión, AOV, top/bottom SKU, vs periodo anterior | Pedidos, productos, (luego) analytics | Nada | Briefing accionable en <5 min de lectura |
| `merchandiser` | Salud de catálogo: sin foto, categoría vacía, ficha pobre, combinación no comprable, related/cross-sell | Productos, combinaciones, categorías, imágenes | Textos/asignación de categoría | Checklist priorizado pre-lanzamiento |
| `inventory-ops` | Roturas, sell-through, variantes ocultas con stock 0, proveedores inactivos en SKU vivos | Stock, combinaciones, proveedores | Ajustes de stock / alertas de compra | Cero “activo pero no comprable” sorpresa |
| `order-triage` | Cola operativa: impagos viejos, pedidos atascados, reembolsos, hilos CS abiertos | Pedidos, historial, carriers, customer threads | Cambio de estado / nota | Cola del viernes ordenada por urgencia |

### 4.2 Dinero (segunda oleada, siempre con techo)

| Agente | Misión | Guardrail |
|--------|--------|-----------|
| `pricing-advisor` | Specific prices, suelo de margen, watch de competencia | **Recomienda**. Aplicar solo si margen ≥ suelo y descuento ≤ techo (p.ej. 15%). |
| `promo-optimizer` | Cart rules vivas/caducadas, abandono, escasez vs dto | No apilar descuentos que rompan margen. Preferir escasez si stock bajo + margen alto. |

### 4.3 Crecimiento (tercera oleada)

| Agente | Misión | Nota |
|--------|--------|------|
| `seo-listing` | Titles, metas, copy de categoría, interlinking | Encaja con `ps-translate` y KBs de theme. |
| `cx-support` | RAG de políticas + lookup de pedido | Customer-facing; otro canal y otro riesgo. |
| `intel-scout` | Precios/catálogo competencia | Empieza read-only (URLs), sin scraping agresivo. |

### 4.4 Lo que no es un agente

- Un “Ecommerce God” que hace de todo.
- Un chatbot FO sin MCP ni políticas.
- Repricing autónomo 24/7.
- Sustituto de GA4/Looker: el analyst **narra y prioriza**, no reemplaza BI.

## 5. Flujos de orquestación (los que importan)

Copiar la idea de MCPBundles: workflows de ops, no “list products”.

| Flujo | Trigger | Agentes | Output |
|-------|---------|---------|--------|
| **Weekly briefing** | Lunes / `/briefing` | analyst → inventory → merchandiser → orchestrator | Un markdown: qué vender, qué se acaba, qué ficha está rota |
| **Pre-launch audit** | Antes de campaña | merchandiser + inventory + promo | Checklist: fotos, stock comprable, cart rules vivas |
| **Friday order triage** | Viernes / `/order-triage` | order-triage | Cola rankeada (impago, CS abierto, reembolso) |
| **Cart-rule sanity** | Al crear promo | promo + pricing + inventory | “¿Esta regla pierde dinero o pisa otra?” |
| **Listing refresh** | SKU nuevo o ficha pobre | seo-listing + merchandiser | Copy + metas; humano publica |
| **Abandon / recover** (luego) | Evento carrito | orchestrator + inventory + pricing | Oferta o espera; no dto ciego |

El orquestador es el primo de `task-classifier`:

```
TYPE: briefing | catalog-audit | order-triage | pricing-review | promo-review | listing | unknown
ROUTE: <agent/skill>
WHY: …
CONFIDENCE: high|medium|low
AUTONOMY: read | propose | apply-with-cap | escalate
```

## 6. Cómo encaja en Mallard (sin CrewAI el día 1)

Reutilizar el mecanismo que ya documenta `docs/adding-skills.md`.

```
skills/commerce-kb/          # playbooks: KPIs, guardrails, workflows
skills/ps-mcp/               # cómo conectar MCP Server, scopes, dry-run
claude/agents/
  commerce-orchestrator.md
  sales-analyst.md
  merchandiser.md
  inventory-ops.md
  order-triage.md
claude/commands/
  commerce.md                # /commerce  → orquestador
  briefing.md                # /briefing
  catalog-audit.md
  order-triage.md
```

`task-classifier` gana tipos `commerce-briefing`, `catalog-audit`, `order-triage` cuando el módulo Intervals sea tienda/ops y no Disseny.

**¿CrewAI o LangGraph?** No en v1. Un subagent Claude + MCP cubre briefing y auditorías. Subir a LangGraph solo si hace falta checkpoint / HITL de horas (campaña de repricing por oleadas). CrewAI encaja si un cliente quiere un servicio Python aparte; no pertenece al binario Go.

## 7. Fases

### Fase 0 — Bus (sin agentes de negocio)

- Documentar setup: `ps_mcp_server` + webservice + conector Claude/Cursor.
- Skill `ps-mcp`: scopes, dry-run, “nunca desactivar auth en prod”.
- Decidir si MCP Tools Plus entra en el stack de agencia (reportes + approve/reject + GDPR) o se reimplementa lo mínimo.

### Fase 1 — Read-only crew

- Agentes: orchestrator, sales-analyst, merchandiser, inventory-ops, order-triage.
- Comandos `/briefing`, `/catalog-audit`, `/order-triage`.
- KB `commerce-kb` con playbooks y umbrales por defecto (stock bajo, pedido impago >14 días, etc.).
- Piloto en **una** tienda Lando/staging. Cero writes.

### Fase 2 — Writes con techo

- `pricing-advisor` + `promo-optimizer`.
- Toda mutación: diff → humano → apply.
- Techos por tienda en un `commerce.toml` gitignored del proyecto cliente (suelo de margen, dto máx, quién aprueba).

### Fase 3 — Crecimiento

- `seo-listing` (reusa `ps-translate` / image regen).
- Extender classifier Intervals.
- Métricas: tiempo de briefing, % de recomendaciones aplicadas, incidentes de write.

### Fase 4 — Solo si hay demanda de cliente

- `cx-support` / shopping agent (FO, WebMCP).
- UCP si un marketplace agentic lo pide.
- A2A entre crews (dev ↔ commerce): p.ej. merchandiser detecta “sin foto” → abre tarea → `ps-image-regen`.

## 8. Riesgos

| Riesgo | Mitigación |
|--------|------------|
| LLM interpreta mal un tool y cambia un precio | Writes off by default; approve/reject; log de cada tool call |
| PII a Claude/ChatGPT | Anonimizar (Tools Plus o capa propia); minimizar campos |
| Multi-shop: mismo SKU, otro stock/precio | El prompt obliga `id_shop`; auditorías por tienda |
| Coste 2–4× vs un solo modelo | Empezar con 2–3 agentes; modelos baratos en analyst, mejores en orquestador |
| “IA God” que pisa al experto PS | Commerce crew **no** toca theme/módulos; eso sigue en prestashop-expert / panda-expert |
| Promesas de conversión | Baseline + A/B; no vender el plan con % de vendors |

## 9. Decisión que hay que tomar antes de codear

1. **¿Mallard Commerce es producto de agencia (operar tiendas de clientes) o feature del CLI de desarrollo?** Recomendación: las dos, pero el CLI solo trae agentes + playbooks; la tienda trae MCP.
2. **¿MCP Tools Plus o solo el server oficial?** Plus acelera reportes y HITL. El oficial basta para Fase 1 si aceptamos armar los playbooks.
3. **¿Primera tienda piloto?** Sin tienda real (o staging) el crew es teatro.
4. **¿Primer flujo?** Recomendación: `/briefing` + `/catalog-audit`. Valor alto, riesgo bajo, encaja con el día a día de ops.

## 10. Fuentes

- [PrestaShop MCP Server docs](https://docs.mcp.prestashop.com/en/0-getting-started/introduction/)
- [PrestaShop — hablar con la tienda](https://prestashop.com/blog/tech-en/taking-your-first-steps-with-ai-how-to-simply-talk-to-your-prestashop-store/)
- [MCP Tools Plus FAQ](https://faq.businesstech.fr/en/faq/616-what-is-mcp-tools-plus)
- [MCPBundles — store operations are workflows](https://www.mcpbundles.com/blog/prestashop-ai-store-operations-workflows)
- [Nicolas Dabène — AI orchestrators + PrestaShop](https://nicolas-dabene.fr/en/blog/the-end-of-the-lone-coder-why-future-developers-will-be-ai-orchestrators-and-how-to-get-started-with-prestashop)
- [Busony Claude-Powered PrestaShop](https://busony.com/en/sites-ia/prestashop)
- [Shopify Agents / UCP](https://shopify.dev/docs/agents)
- [VTEX Vision / AI Workspace](https://www.vtex.com/en-us/vtex-vision)
- [Salesforce Agentforce](https://www.salesforce.com/agentforce/)
- [Atlas Mercator](https://github.com/hhdhh/atlas-mercator)
- [Alex Genovese — Decision Crew](https://alexgenovese.com/multi-agent-ai-for-ecommerce-how-agentic-systems-are-reshaping-conversion-in-2026/)
- [Vortex IQ — A2A orchestration](https://www.vortexiq.ai/resources/blog/agent-to-agent-orchestration-next-frontier)
- [CrewAI 1.10 MCP + A2A](https://www.datapath.ai/blog/crewai-1-10-sistemas-multi-agente-mcp-a2a-2026)
