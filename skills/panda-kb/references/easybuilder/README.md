---
kind: kb-subfolder-index
product: Easy Builder (steasybuilder)
status: separate-track
tags: [easybuilder, panda, index, paid-addon]
---

# Easy Builder — Knowledge Base (separate track)

> [!info] Why this folder is separate from the main Panda KB
> **Easy Builder is a paid optional SunnyToo addon** (`steasybuilder`) for the Panda theme — technically a fork/port of the WordPress **Elementor** plugin adapted to PrestaShop. Not every Panda client owns it. The main Panda KB ([[../00-overview|Resources/Panda Knowledge Base]]) must stay usable for clients without Easy Builder; references to these notes should only appear as wikilinks when relevant, never as hard dependencies.

## Files in this folder

- [[00-overview]] — product fiche, architecture, integration with Panda, licensing, upgrade path.
- [[01-widgets-catalog]] — 31 Elementor widgets (stock-Elementor ports; no PS-aware widgets here).
- [[02-presets-catalog]] — 3 bootstrap files in `import_data/` (no per-template gallery shipped in source).
- [[03-admin-controllers]] — 4 BO controllers (manager, editor, general settings, license).
- [[04-bridge-trans-panda]] — **bridge module** (`steasy_trans_panda`): 20 PrestaShop-aware widgets (product slider, category slider, brand slider, Instagram, newsletter, Revolution-Slider wrapper, etc.) that plug into Easy Builder via `actionEasyBuilderInit`.

> [!important] The bridge is required for PS catalog editing
> The bridge module (`steasy_trans_panda`, see [[04-bridge-trans-panda]]) is required if the client wants PS-aware editing. **Without it, Easy Builder is just generic Elementor in PrestaShop** — no product widgets, no category sliders, no catalog integration. The 31 widgets in [[01-widgets-catalog]] are all generic (buttons, images, tabs, etc.); every real catalog widget lives in the bridge.

## Related tracks

- [[../00-overview|Panda KB root]] — main knowledge base for the base Panda theme + its 54 bundled modules.

## When to reference this folder from the Panda KB

Reference Easy Builder ONLY if one of the following is true:

1. The user explicitly mentions **Easy Builder**, **steasybuilder**, **Elementor**, **drag and drop editor**, or shows the Elementor UI.
2. The user describes a visual-editing need that Panda's stock content modules (`steasycontent`, `stcustomhtml`, `stthemeeditor`, `stblockreassurance`, etc.) can't cover — typically multi-section landing pages with fine layout control on responsive breakpoints, or custom product-description layouts with carousels / galleries / tabs.
3. The user asks about the **BO menu entries** `AdminStEasyBuilder*` or the PS tables `ps_st_easy_builder*` / `ps_easy_builder_postmeta`.
4. The user asks about a hook in the **Easy Builder-emitted** list (e.g. `FullWidthTop2`, `FullWidthCategoryHeader`, `ProductDescRightColumn`) and the answer benefits from explaining who writes content into that hook.

Otherwise, **keep Easy Builder out of the answer** — assume the client does not have it installed. If unsure, ask the user: *"Do you have the Easy Builder addon installed, or are we working with Panda stock modules only?"*

## One-line recap for system prompts

> `steasybuilder` — paid SunnyToo addon, PrestaShop port of WordPress Elementor. 31 generic widgets (buttons, images, tabs, carousels, etc.), no PS-aware widgets in the core module; PS catalog integration comes from the `steasy_trans_panda` bridge. Config via BO "Improve → Easy Builder". Panda works without it.
