---
topic: migration
tags: [prestashop, migration, ps8-ps9, index]
ingested: 2026-04-18
---

# PrestaShop Version Migration (PS 8 ↔ PS 9) — Index

Architecture shifts, breaking changes, and upgrade paths. PS 9 is the canonical reference; PS 8 material is noted when it differs.

| File | Section | ps_version | Source URL |
|------|---------|------------|------------|
| [[keeping-up-to-date]] | Update vs Migration — two distinct processes | [9] | `/9/basics/keeping-up-to-date/` |
| [[migration-process]] | Migration process (PS 8 → PS 9) | [9] | `/9/basics/keeping-up-to-date/migration/` |
| [[upgrade-faq]] | Upgrade FAQ | [9] | `/9/faq/upgrade/` |
| [[architecture-introduction]] | Architecture introduction (PS as a whole) | [8, 9] | `/9/development/architecture/introduction/` |
| [[architecture-modern]] | Modern architecture (Symfony + Twig in BO) | [9] | `/9/development/architecture/modern/` |
| [[architecture-legacy]] | Legacy architecture (pre-Symfony core, Smarty BO) | [8, 9] | `/9/development/architecture/legacy/` |
| [[migration-guide-bo]] | Back Office migration guide (legacy → Symfony) | [8, 9] | `/9/development/architecture/migration-guide/` |

## Key takeaways (from ingested content)

- PS 9 is a **migration**, not an update — data/modules/themes are not guaranteed to carry over. Custom modules must be explicitly verified.
- **Modern architecture** = Symfony controllers + Twig in back office. **Legacy** = old ObjectModel + Smarty BO (still supported in modules).
- Theme inheritance (`parent:` + `use_parent_assets`) is **fully working starting from 8.0** — earlier versions had a broken implementation.

## Related

- [[../themes/_index|Themes index]]
- [[../modules/_index|Modules index]]
- [[../README|PrestaShop KB]]
