---
source_url: https://build.prestashop-project.org/news/2026/prestashop-9-1-0-available/
related_source: https://github.com/PrestaShop/hummingbird
ps_version: [9]
ingested: 2026-04-19
section: Hummingbird 2.0 context
topic: themes
tags: [prestashop, themes, hummingbird, ps9, comparative]
---

# Hummingbird 2.0 — context para el Panda Expert

> [!info] Why this exists
> El agente `panda-expert` solo cubre Panda (SunnyToo). Al probar contra un PS 9.1.0 real, apareció Hummingbird 2.0 como theme default. Este doc le da contexto mínimo para responder comparativas sin inventar.

## ¿Qué es Hummingbird 2.0?

Theme front-office **oficial y por defecto desde PS 9.1** (reemplaza a Classic). Mantenido por el equipo de PrestaShop en [github.com/PrestaShop/hummingbird](https://github.com/PrestaShop/hummingbird). Open source, MIT.

No es un producto para vender sitios "visualmente ricos" — está pensado como **base técnica moderna y accesible** que la comunidad y los devs customizan.

## Arquitectura técnica

- **Bootstrap 5** con LTS — componentes consistentes, design system conocido por muchos devs.
- **BEM-based CSS naming** para clases — mantenibilidad, menos colisiones de specificity.
- **SCSS con `@layer`** — cascade ordenada, overrides predecibles.
- **Reorganización del `_dev/`** respecto a Classic — directorio más plano, menos legacy.
- **Reducción de deps third-party** — prioriza los componentes built-in de Bootstrap.
- **Separación clara** Bootstrap vs PrestaShop-specific styles — custom overrides más fáciles.
- **Native dark mode** via SCSS variable (toggle sencillo).

## Accessibility

- **>95% compliance con European Accessibility Act (EAA)** — requerimiento legal desde junio 2025 para tiendas UE.
- Semantic HTML + ARIA attributes + keyboard navigation + color contrast adecuado built-in.

## Build requirements

- Node 20+ y NPM 10+
- `npm ci && npm run build` tras instalación desde source
- El zip distributivo (prestashop_9.1.0.zip de prestashop.com/versions/) ya trae assets compilados
- El clone de GitHub **no** los trae — hay que correr el build

## Panda vs Hummingbird — decision matrix

| Criterio | Panda (SunnyToo) | Hummingbird 2.0 |
|---|---|---|
| Licencia | Pago, por sitio | Open source MIT |
| Target | Tiendas visualmente ricas, clientes que quieren sensación "premium" | Base técnica moderna, devs que customizan desde cero |
| Layouts predefinidos | 4 (full-width, three-columns, left-column, right-column) | Bootstrap 5 grid — layouts libres |
| Módulos del vendor | 54 `st*` con funcionalidades específicas (megamenu, sliders, blog suite, stickers, etc.) | Ninguno — usa módulos PS stock |
| Demos pre-built | 19 demos temáticas (fashion, furniture, sport, creative…) | Ninguna |
| Visual editor compatible | Easy Builder (addon SunnyToo paid) | Ninguno nativo — hay que usar modules PS genéricos |
| Accessibility EAA | No medido oficialmente | >95% out-of-the-box |
| Update cadence | Ligada a releases de SunnyToo (menos frecuente) | Ligada a releases PS (cada menor) |
| Code modernity | PHP/Smarty clásico, algunos hooks legacy | Bootstrap 5, BEM, SCSS @layer |
| Dark mode | No built-in | Built-in via SCSS variable |

## Cuándo elegir cuál

**Panda** cuando:
- El cliente paga por design — espera resultado visual sin desarrollarlo desde cero.
- Usará Easy Builder o quiere sus layouts predefinidos.
- Proyecto en PS 8.x (Hummingbird 2.0 es PS 9 only).
- Necesita demos listas para importar (fashion/furniture/etc.).
- Quiere el ecosistema cerrado SunnyToo (updates, support dedicado).

**Hummingbird** cuando:
- Cliente/tienda con exigencia legal de accessibility (UE, EAA).
- El dev está cómodo con Bootstrap 5 + BEM + SCSS.
- Proyecto en PS 9.x sin interés en theme comercial.
- Customización desde base limpia, sin dependencia de módulos paid.
- Stack moderno prioritario sobre features out-of-the-box.

**Coexistencia**: son temas **excluyentes a nivel activo** (sólo uno en `ps_shop.theme_name`). Pueden coexistir en `themes/` folder sin conflicto pero un shop usa uno a la vez.

## Implicaciones para Eduardo (freelance PS + Panda)

- Si un cliente llega con **PS 9.1 fresh install + sin decisión de theme** → Hummingbird es la opción por defecto, pero **Panda sigue siendo la elección si el cliente pidió Panda explícitamente** (suele ser así en el flujo freelance con SunnyToo).
- Si un cliente está en PS 8.x → Hummingbird no aplica (v2.0 es PS 9 only).
- Para **migración futura PS 9 → PS 10**: Hummingbird tiene más probabilidad de quedar mantenido (es core de PrestaShop). Panda depende del ciclo de SunnyToo.
- Si el cliente menciona **accessibility como requirement** (p.ej. sector público UE) → advertir que Panda **no está auditado para EAA** — puede haber gap legal. Hummingbird sí.

## Qué NO hace este doc

- No es KB operativa de Hummingbird — si un cliente pide Hummingbird en serio, hay que ingestar su documentación específica.
- No cubre cómo customizar Hummingbird (theme.yml, overrides, custom SCSS) — solo contexto comparativo.
- No compara Easy Builder vs alternativas Hummingbird — eso requiere research adicional.

## Related in KB

- [[../../easybuilder/00-overview]] — Easy Builder (addon para customización visual en Panda, sin equivalente nativo Hummingbird)
- [[./parent-child-feature]] — child themes (aplican igual para Hummingbird y Panda)
- [[./theme-yml]] — la spec es la misma (`parent:`, `use_parent_assets:` etc.)
- [[../version-migration/architecture-modern]] — contexto arquitectónico PS 9 donde Hummingbird encaja
