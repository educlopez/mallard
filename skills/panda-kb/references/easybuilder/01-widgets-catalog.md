---
parent: "[[00-overview]]"
widget_count: 31
source_dir: <source-archive>/
source_elements_dir: <source-archive>/
tags: [easybuilder, widgets, elementor]
---

# Easy Builder — Widgets Catalog

> [!info] Scope
> This catalog inventories the widgets shipped natively with `steasybuilder/elementor/includes/widgets/`. These are the Elementor-stock widgets ported to PrestaShop. **PrestaShop-specific widgets** (product card, category listing, price, add-to-cart, etc.) are NOT in this folder — they are contributed by the bridge module [[../easybuilder-bridge/|steasy_trans_panda]] via the `actionEasyBuilderInit` / `actionOtherAddEasyBuilderData` hooks. Task #8 catalogs those.

## Structural elements (`elementor/includes/elements/`)

Not widgets per se — these are the containers Elementor uses to lay out widgets. Included for completeness:

| Folder | Element | Role |
| --- | --- | --- |
| `elements/section.php` | Section | Full-width row, holds one or more columns. |
| `elements/column.php` | Column | Vertical slot inside a section, holds widgets. |
| `elements/repeater.php` | Repeater | Internal repeatable-field helper (used by widgets that accept lists, e.g. Tabs, Accordion, Icon List). Not placeable directly. |

## Widgets (`elementor/includes/widgets/`) — 31 total

> [!note] Reading the table
> - **Name** = Elementor internal ID (`get_name()`), used in `widgetType` of `_elementor_data` JSON.
> - **Title** = user-facing label (`get_title()`).
> - **Category** = Elementor panel category (`get_categories()`). When empty the widget falls back to Elementor's default category (`general`).
> - **Icon** = `eicon-*` font icon shown in the editor panel.
> - **Key controls** = first controls declared via `add_control()` / `add_responsive_control()` — the knobs users see first in the inspector.
> - **Controls count** = total settings exposed (rough complexity indicator).

| File | Name | Title | Category | Icon | Controls (ctrl+resp+group) | Key controls (first 5-8) |
| --- | --- | --- | --- | --- | --- | --- |
| `accordion.php` | `accordion` | Accordion | general | `eicon-accordion` | 15+3+2 | tabs, view, selected_icon, selected_active_icon, title_html_tag, border_width, border_color, title_background |
| `alert.php` | `alert` | Alert | general | `eicon-alert` | 8+0+2 | alert_type, show_dismiss, view, background, border_color, border_left-width, title_color, description_color |
| `audio.php` | `audio` | SoundCloud | general | `eicon-headphones` | 14+0+0 | link, visual, sc_options, sc_auto_play, sc_buying, sc_liking, sc_download, sc_show_artwork |
| `button.php` | `button` | Button | **basic** | `eicon-button` | 14+2+4 | button_type, align, size, selected_icon, icon_align, icon_indent, view, button_css_id |
| `common.php` | `common` | *(base)* | general | *(none)* | 15+13+6 | _title, _margin, _padding, _z_index, _element_id, _css_classes, _animation |
| `counter.php` | `counter` | Counter | general | `eicon-counter` | 10+0+2 | starting_number, ending_number, prefix, suffix, duration, thousand_separator, view |
| `divider.php` | `divider` | Divider | **basic** | `eicon-divider` | 20+7+1 | style, separator_type, pattern_spacing_flag, pattern_round_flag, width, align, look, icon |
| `google-maps.php` | `google_maps` | Google Maps | **basic** | `eicon-google-maps` | 4+1+2 | address, zoom, height, view, hover_transition |
| `heading.php` | `heading` | Heading | **basic** | `eicon-t-letter` | 16+1+4 | size, header_size, align, view, color_gradient, title_color, first_color, second_color |
| `html.php` | `html` | HTML | general | `eicon-code` | 0+0+0 | *(raw HTML textarea, no Elementor controls)* |
| `icon-box.php` | `icon-box` | Icon Box | general | `eicon-icon-box` | 26+5+5 | selected_icon, view, shape, position, content_vertical_alignment, title_size, primary_color, secondary_color |
| `icon-list.php` | `icon-list` | Icon List | general | `eicon-bullet-list` | 15+4+1 | view, selected_icon, icon_list, link_click, space_between, icon_align, divider, divider_style |
| `icon.php` | `icon` | Icon | **basic** | `eicon-favorite` | 12+3+0 | selected_icon, view, shape, link, align, primary_color, secondary_color, hover_primary_color |
| `image-box.php` | `image-box` | Content Box | general | `eicon-image-box` | 37+8+12 | image, position, title_size, content_max_width, horizontal_position, view, text_align |
| `image-carousel.php` | `image-carousel` | Image Carousel | general | `eicon-slider-push` | 29+3+3 | carousel, slides_to_show, slides_to_scroll, image_stretch, navigation, open_lightbox, view, autoplay |
| `image-gallery.php` | `image-gallery` | Basic Gallery | general | `eicon-gallery-grid` | 12+0+3 | wp_gallery, gallery_columns, gallery_link, open_lightbox, gallery_rand, view, image_spacing |
| `image.php` | `image` | Image | **basic** | `eicon-image` | 17+7+7 | image, align, caption_source, link_to, open_lightbox, lightbox_group_open, lightbox_group, lazyloading |
| `menu-anchor.php` | `menu-anchor` | Menu Anchor | general | `eicon-anchor` | 2+0+0 | anchor, anchor_note |
| `progress.php` | `progress` | Progress Bar | general | `eicon-skill-bar` | 11+0+2 | progress_type, percent, display_percentage, view, bar_color, bar_bg_color, bar_height |
| `read-more.php` | `read-more` | Read More | general | `eicon-post-excerpt` | 2+0+0 | theme_support, link_text |
| `shortcode.php` | `shortcode` | Shortcode | general | `eicon-shortcode` | 1+0+0 | shortcode |
| `sidebar.php` | `sidebar` | Sidebar | general | `eicon-sidebar` | 1+0+0 | sidebar |
| `social-icons.php` | `social-icons` | Social Icons | general | `eicon-social-icons` | 16+6+1 | social_icon, link, item_icon_color, social_icon_list, shape, columns, align, view |
| `spacer.php` | `spacer` | Spacer | **basic** | `eicon-spacer` | 1+1+0 | space, view |
| `star-rating.php` | `star-rating` | Star Rating | general | `eicon-rating` | 8+4+1 | rating_scale, rating, star_style, unmarked_star_style, align, title_color, icon_size |
| `tabs.php` | `tabs` | Tabs | general | `eicon-tabs` | 12+0+2 | tabs, view, type, navigation_width, border_width, border_color, background_color |
| `testimonial.php` | `testimonial` | Testimonial | general | `eicon-testimonial` | 9+0+5 | testimonial_image, testimonial_image_position, testimonial_alignment, view, content_content_color |
| `text-editor.php` | `text-editor` | Text Editor | **basic** | `eicon-text` | 11+3+2 | drop_cap, text_columns, column_gap, align, text_color, text_a_color, text_a_hover_color |
| `toggle.php` | `toggle` | Toggle | general | `eicon-toggle` | 15+4+3 | tabs, view, selected_icon, selected_active_icon, title_html_tag, border_width, space_between |
| `video.php` | `video` | Video | **basic** | `eicon-youtube` | 40+2+3 | video_type, youtube_url, vimeo_url, dailymotion_url, insert_url, hosted_url, external_url, start |
| `wordpress.php` | `wp-widget-` | *(base)* | `pojo` | `eicon-pojome` | 1+0+0 | wp — WordPress-widget host (stub, see note below) |

## Observed patterns

- **Total widgets**: 31 PHP files (one file per widget, Elementor convention).
- **Category breakdown**:
  - **`basic`** (9): Button, Divider, Google Maps, Heading, Icon, Image, Spacer, Text Editor, Video. These are the first-tier widgets shown at the top of the editor panel.
  - **`general`** (default, 20): everything else (Accordion, Alert, Audio, Counter, HTML, Icon Box, Icon List, Image Box, Image Carousel, Image Gallery, Menu Anchor, Progress, Read More, Shortcode, Sidebar, Social Icons, Star Rating, Tabs, Testimonial, Toggle).
  - **`pojo`** (1): `wordpress.php` — legacy WordPress-widget bridge category (Pojo = Elementor's original company name).
- **Complexity leaders** (by control count): `video` (40+2+3), `image-box` (37+8+12), `image-carousel` (29+3+3), `icon-box` (26+5+5), `divider` (20+7+1).
- **Stubs / placeholders** (controls ≤ 2): `html` (raw HTML passthrough, no controls), `menu-anchor`, `read-more`, `shortcode`, `sidebar`, `spacer`, `wordpress`.

### PrestaShop-aware vs generic

> [!warning] All 31 widgets are generic — none are PrestaShop-aware
> After static inspection, **no widget in this folder references PrestaShop objects** (`Product`, `Category`, `Manufacturer`, `Cart`, `Customer`, `Context::getContext()->shop`, etc.). They are the stock Elementor widgets carrying WordPress concepts: `wp_gallery` (image-gallery), `shortcode` (WP shortcodes), `sidebar` (WP sidebar widget areas), `wordpress.php` (WP-widget host). Even `image-gallery` uses `wp_gallery` as its data source field name.
>
> PrestaShop catalog integration (product carousels, category grids, add-to-cart buttons, etc.) happens through one of:
> 1. The bridge module [[../easybuilder-bridge/|steasy_trans_panda]] which registers extra widgets via `actionEasyBuilderInit`.
> 2. Dynamic tags in `elementor/core/dynamic-tags/` (ported) and `elementor/modules/dynamic-tags/`.
> 3. The `shortcode` widget can invoke PS modules via custom shortcodes wired by the bridge.
>
> Confirm in task #8.

### WP-only code likely to break on PS

- `wordpress.php` — registers itself per WordPress widget class (scans `$wp_widget_factory->widgets`). On PrestaShop there is no `$wp_widget_factory`, so this widget is effectively inert unless the ported `wp-includes/` shim provides a stub.
- `sidebar.php` — depends on WP `dynamic_sidebar()` function; again stubbed in `wp-includes/`.
- `image-gallery.php` — `wp_gallery` field expects WordPress attachment IDs. In PS context this probably maps to PS image IDs via a shim; worth verifying if gallery editing feels broken.
- `shortcode.php` — WP `do_shortcode()` call; the `actionEasyBuilderInit` hook is likely the injection point used by the bridge to register PS-flavored shortcodes.

## Cross-links to Panda KB

Recommended wikilinks to add from Panda module fiches:

- [[../modules/stmegamenu]] — mega-menu module; its demos often embed Easy Builder templates inside mega-menu panels. Link back to `tabs`, `image-box`, `icon-list`.
- [[../modules/stthemeeditor]] — theme editor; bundles Easy Builder license (skips purchase-code check). Mention this coupling.
- [[../modules/stblog]] — optional blog; enables blog-specific builder types (`StBlogArticleTop`, `StBlogContent`, etc.). Link from stblog fiche.
- [[../modules/steasycontent]] — free sibling content manager; contrast feature matrix (TinyMCE vs Elementor).
- Any Panda hook fiches listing `FullWidthTop2`, `NavFullWidth`, `FooterBefore`, `ProductRightColumn` should cross-link here as "Easy Builder emits into this hook".
