---
topic: smarty
subtopic: readme
ingested: 2026-04-18
tags: [smarty, prestashop, readme]
ps_version: [8, 9]
---

# Smarty — Template authoring KB

Template-authoring docs for PrestaShop `.tpl` files (Panda theme + modules).

## Purpose

PrestaShop's front templates are written in Smarty. The PS devdocs cover PS-specific helpers only — Smarty core syntax lives on `smarty.net` / `smarty-php.github.io`. This folder fills that gap, scoped to what actually appears in real `.tpl` files.

## Scope

**Included (what you see in PS templates):**

- Variable output + inline modifiers: `{$var|escape}`, `{$var|date_format:'...'}`, etc.
- Control flow: `{if}`, `{foreach}`, `{for}`, `{while}`, `{section}` (legacy)
- Assignments: `{assign}`, `{capture}`
- Includes and inheritance: `{include}`, `{extends}`, `{block}`, `{$smarty.block.parent}`
- Special vars: `$smarty.foreach.*`, `$smarty.section.*`, `$smarty.const.*`, `$smarty.now`
- PrestaShop extensions: `{l}`, `{hook}`, `{widget}`, `{form_field}`, `{url}`, `{render}`, `classname` modifier, and PS-provided context vars (`$urls`, `$link`, `$customer`, `$cart`, `$page`, etc.)

**Excluded:**

- Smarty PHP-side API (assign, caching, security config) — not template authoring
- Writing custom Smarty plugins in PHP — belongs in module development
- Smarty 2 legacy syntax (PS uses Smarty 3/4)
- Performance tuning / internals

## Sources

- **Smarty core** (variables, modifiers, control flow, includes, special vars): `https://smarty-php.github.io/smarty/stable/designers/` — official Smarty docs, designer section. Current stable renders cleanly with `defuddle` and covers Smarty 3+/4/5 syntax. PS uses Smarty 3/4; the designer-level syntax is stable across those versions.
- **PrestaShop extensions** (`{l}`, `{hook}`, `{widget}`, `{form_field}`, `{render}`, `{url}`, PS modifiers): synthesized in [[prestashop-extensions]] from `https://devdocs.prestashop-project.org/9/development/components/smarty-extensions/`. Not scraped — written to stay concise and cross-link to the already-ingested `themes/` and `hooks/` folders.

## Start here

- [[_index|Index of all files]] — grouped by subtopic
- [[prestashop-extensions|PrestaShop-specific tags and context vars]] — the ones Smarty docs don't cover
- [[foreach]], [[if]], [[extends]], [[block]], [[modifier-escape]], [[special-vars]] — the most-used pages in practice

## Related folders

- [[../themes/_index|themes/]] — where `.tpl` files live and how PS discovers them (parent/child theme cascade, layouts)
- [[../themes/template-inheritance|themes/template-inheritance]] — the PS-side explanation of how `{extends}`/`{block}` maps to the override cascade
- [[../hooks/_index|hooks/]] — the names you pass to `{hook h='...'}`
- [[../modules/_index|modules/]] — module templates also use all the syntax documented here

## Re-ingest

When PrestaShop bumps Smarty to a newer major version or PS ships new helpers. Current snapshot: **2026-04-18**.
