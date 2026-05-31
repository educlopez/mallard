---
ingested: 2026-04-18
source: https://devdocs.prestashop-project.org/
scope: platform-level PrestaShop docs (NOT Panda-specific)
tags: [prestashop, kb, devdocs, platform]
---

# PrestaShop Platform KB

> [!info] Purpose
> This subfolder ingests the **official PrestaShop developer documentation** (`devdocs.prestashop-project.org`) so the Panda Expert agent has **platform-level** context — theme cascade, hooks mechanics, module lifecycle, PS 8 → 9 differences.
>
> It complements the Panda-specific KB in `../modules/`, `../_agent/`, and `../demos/`.

## Directory map

```
prestashop/
├── README.md                 # this file
├── _errors.md                # pages that failed to fetch/parse
├── themes/                   # theme dev: organization, theme.yml, child themes, assets
│   ├── _index.md
│   └── hummingbird-v2.md     # PS 9 default theme context (not scraped, written from findings)
├── modules/                  # module dev: structure, controllers, templating, overrides
│   └── _index.md
├── hooks/                    # hook concept + full list of available hooks
│   └── _index.md
├── smarty/                   # Smarty template-authoring docs (for .tpl files)
│   └── _index.md
├── install/                  # install workflows (not scraped, written from real findings)
│   └── source-vs-distribution.md
└── version-migration/        # PS 8 ↔ PS 9 differences, legacy vs modern architecture
    └── _index.md
```

## Conventions

- **PS 9 is primary.** Canonical URLs fetched from `/9/...`. PS 9 docs cover back-compat with PS 8 in-line.
- **When PS 8 differs materially**, the file frontmatter shows `ps_version: [8]` or `ps_version: [9]`. If `ps_version: [8, 9]`, content applies to both.
- **Not ingested (out of scope):** full API class reference, payment modules, i18n internals, back-office admin tutorials (those live on the other docs site, `docs.prestashop-project.org`), full DB schema, CLI tooling unrelated to theme/module dev.
- **Smarty docs** — Smarty core syntax (variables, modifiers, control flow, template inheritance) lives in [[smarty/_index|smarty/]], scraped from `smarty-php.github.io/smarty/stable/designers/`. PS-specific extensions (`{l}`, `{hook}`, `{widget}`, `{form_field}`, `{url}`) are in [[smarty/prestashop-extensions]].
- **Re-ingest** when PrestaShop releases a new major version. Current snapshot: **2026-04-18**.

## Topic indexes

- [[themes/_index|Themes index]] — 11 scraped pages + `hummingbird-v2.md` comparative
- [[modules/_index|Modules index]] — 10 pages
- [[hooks/_index|Hooks index]] — 2 pages
- [[smarty/_index|Smarty index]] — 36 pages (Smarty core + PS extensions)
- [[install/source-vs-distribution|Install — source vs distribution]] — PS 9 git clone vs zip differences
- [[version-migration/_index|Version migration index]] — 7 pages

## Most important pages for Eduardo

1. [[themes/parent-child-feature|Parent/child themes]] — the convention Eduardo flagged as critical.
2. [[themes/theme-yml|theme.yml reference]] — full spec including `parent:` and `use_parent_assets`.
3. [[themes/hummingbird-v2|Hummingbird 2.0 context]] — PS 9 default theme + Panda-vs-Hummingbird decision matrix.
4. [[install/source-vs-distribution|PS 9 source vs distribution]] — why `install-dev/` and `admin-dev/` exist in clones.
5. [[hooks/hooks-overview|Hooks overview]] — naming scheme (display/action), registration, Smarty `{hook h='...'}` usage.
6. [[themes/theme-organization|Theme organization]] — directory structure of a PrestaShop theme.
7. [[version-migration/architecture-modern|Modern architecture (Symfony + Twig in BO)]] — what changed in PS 9.

## Related

- [[../_agent/panda-expert|Panda Expert agent]]
- [[../README|Panda KB root]]
- [[../SETUP-NEW-MACHINE|Panda setup on a new machine]]
- [[../modules/README|Panda module fiches README]]
