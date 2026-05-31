---
parent: "[[00-overview]]"
preset_count: 3
source_dir: <source-archive>/
tags: [easybuilder, presets, seed-data, bootstrap]
---

# Easy Builder — Presets / Import Data Catalog

> [!info] What lives in `import_data/`
> Unlike typical preset galleries (one folder per page-template ready to import with one click), `steasybuilder/import_data/` contains just **3 bootstrap files** loaded at install time by `installDefaultData()` / `import_base_data()` in `steasybuilder.php`. There is **no per-client preset gallery shipped in the open source code** — clients get a demo seed + base SQL only. A richer preset gallery is likely fetched remotely from sunnytoo.com (see `AdminStEasyBuilderEditorController::processElementorLibraryDirectActions()` which routes to Elementor's template library) or bundled inside `stthemeeditor`.

## Files

| File | Size | Type | Purpose |
| --- | --- | --- | --- |
| `base_easy_builder.sql` | 953 B | SQL seed | Schema/seed run during `installDefaultData()` when `stthemeeditor` is not present. Loaded via `import_base_data()`. |
| `easy_builder_postmeta.json` | 9.5 KB | JSON (array of records) | Initial rows for the `easy_builder_postmeta` table. Contains demo Elementor JSON for a handful of posts (IDs 2-5) plus `_elementor_easymenu_page` metadata for menu bootstrap. |
| `settings.json` | 77 KB | JSON (Kit styleGlobals) | Default Elementor **Kit / Style Globals**: system colors (Primary `#6EC1E4`, Secondary `#444444`, Text `#7A7A7A`, Accent `#61CE70`), system typography (Vollkorn as Primary font), and kit schemes. Imported into the global styles on install. |

## Component breakdown per file

### `base_easy_builder.sql`

Schema/seed file (empty/near-empty in the extracted copy — no `INSERT INTO` statements detected). Actual schema lives in PHP inside `installDB()` of `steasybuilder.php`. This file is a hook for custom seed data per release — may be populated at install time by the license server or left empty when `stthemeeditor` is the bundle source.

### `easy_builder_postmeta.json`

Demo Elementor page-meta records. Decoded structure:

| `post_id` | Keys present | Notes |
| --- | --- | --- |
| 1 | `_elementor_easymenu_page` | Main menu seed (position 0, id_parent 0). Serialized PHP array. |
| 2 | `_elementor_template_type` (`wp-page`), `_elementor_version` (`3.0.11`), `_elementor_edit_mode` (`builder`), `_elementor_page_settings`, `_elementor_data` | Demo page — contains a single-section / single-column layout with one `image` widget pointing to `/img/cms/demo12parallaxbg.jpg`. |
| 3 | `_elementor_template_type`, `_elementor_version` | Empty page skeleton. |
| 4 | `_elementor_template_type`, `_elementor_version` | Empty page skeleton. |
| 5 | `_elementor_easymenu_page`, `_elementor_template_type`, `_elementor_version` | "Categories" menu item seed (multilingual: `menu_title_zh: Categories`, `menu_title_en: 0000`). |

> [!note] Elementor version baseline
> All `_elementor_version` values are `3.0.11` — this is the Elementor version the port tracks. Useful data point for upstream-compatibility questions (Elementor is now at v3.2x in WordPress).

**Widgets referenced in demo JSON**: only `image` is confirmed in the visible sample. Rest of the 77 KB JSON likely exercises more widgets — kept unparsed here (catalog scope).

### `settings.json` — Elementor Kit style globals

Keys at top level (observed):

- `styleGlobals.system_colors[]` — 4 tokens:
  - `primary` → `#6EC1E4` (cyan)
  - `secondary` → `#444444` (dark gray)
  - `text` → `#7A7A7A` (medium gray)
  - `accent` → `#61CE70` (green)
- `styleGlobals.custom_colors[]` — empty in stock.
- `styleGlobals.system_typography[]` — 1+ entries:
  - `primary` → `custom`, `Vollkorn`, weight `400`, per-breakpoint sizes empty (inherits).
- `styleGlobals.custom_typography[]` — empty in stock.
- `styleGlobals.default_generic_fonts` → `Sans-serif`.
- `styleGlobals.activeItemIndex` — editor UI state.

This is a direct port of Elementor's **Site Settings → Global Fonts / Global Colors** defaults. The values (`#6EC1E4` primary, Vollkorn font, accent `#61CE70`) are **Elementor's factory defaults**, not Panda-themed — Panda's orange/brand palette does NOT appear here. Clients must configure their brand colors per install.

> [!warning] Brand mismatch out of the box
> Fresh installs of Easy Builder carry Elementor's default color palette (cyan / gray / green), which clashes with Panda's usual warm palette. Expect every client deployment to need an initial pass on Site Settings → Global Colors before building pages. This is a likely onboarding checklist item.

## Summary

- **3 bootstrap files total**, no per-template gallery.
- Presets are delivered differently from "typical Elementor asset packs": only a minimal demo (a few empty pages + one image widget) + global style kit.
- The full preset / template gallery is probably fetched at runtime from the SunnyToo / Elementor template library, not shipped in the source tarball. Verify by hitting the Editor UI and watching network traffic.
- Seed SQL hook (`base_easy_builder.sql`) exists but is empty in the extracted copy.

## Cross-links

- [[00-overview#Licensing / upgrade path]] — install flow invoking these seeds.
- [[03-admin-controllers#AdminStEasyBuilderEditorController]] — owns the template-library routing (`processElementorLibraryDirectActions`).
- [[01-widgets-catalog#image-gallery.php | image-gallery widget]] — uses `wp_gallery` source, related to demo seed.
