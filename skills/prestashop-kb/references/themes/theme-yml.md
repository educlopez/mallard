---
source_url: https://devdocs.prestashop-project.org/9/themes/getting-started/theme-yml/
ps_version: [8, 9]
ingested: 2026-04-18
section: Themes > theme.yml reference
topic: themes
tags: [prestashop, themes, theme-yml, config, layouts]
---

[Improve this page](https://github.com/PrestaShop/docs/edit/9.x/themes/getting-started/theme-yml.md "Improve this page") [How](/9/contribute/documentation/how)

The theme’s theme.yml file defines all of the theme’s configuration and meta information, such as its version number, layouts, compatibility range, hook configuration, etc.

## Theme description

The theme’s name MUST match its directory name. For instance, if the theme is named “My Awesome Theme” and its ’name’ value is set to “my-awesome-theme”, then the folder MUST be /my-awesome-theme.

Users will be able to choose the layout for each page from the theme’s settings page. Layouts are automatically parsed from the theme’s /templates/layouts folder, so this configuration key is optional, but it allows designers to provide some more user-friendly info than just a filename.

## Parent / child settings

Find more information about [Theming inheritance - Parent/child themes](/9/themes/reference/template-inheritance/parent-child-feature/)

To set the parent theme, set the parent key to the theme’s name:

```yaml
parent: classic
```

### use\_parent\_assets

Starting from 8.0, the setting `use_parent_assets` is now fully working and correctly loads assets from the parent theme.

```yaml
assets:

  use_parent_assets: true
```

When set to `true`, some additional variables are changed for `Smarty` templates.

- `theme_assets` is replaced with the parent theme assets URI,
- `img_url`, `css_url`, `js_url` are replaced with the parent theme corresponding URIs,

some other variables are available:

- `child_theme_assets` which is the child theme assets URI,
- `child_img_url`, `child_css_url`, `child_js_url` which are the child theme corresponding URIs.

When set to `false`, only `theme_assets`, `img_url`, `css_url` and `js_url` are available, and are the child theme corresponding URIs.

To maintain the behavior of 1.7

and load the child theme’s assets, `use_parent_assets` must be set to `false`.

## Global settings

### Configuration

You can have the theme change the configuration of PrestaShop when the theme is enabled.

```yaml
global_settings:

    configuration:

      PS_QUICK_VIEW: false

      NEW_PRODUCTS_NBR: 4

      PS_PNG_QUALITY: 8
```

You can also add custom configuration variables here (to use on modules/themes).

### Modules

You can have the theme enable, disable or reset modules when the theme is enabled.

```yaml
global_settings:

    modules:

      to_enable:

        # All modules below are enabled when

        # the theme is enabled (and installed if needed).

        # They are disabled when the theme is disabled.

        - my-custom-module

        - yippeeslider

      to_disable:

        # All modules below are disabled when the theme is enabled.

        # They are re-enabled when the theme is disabled.

        - homeslider

        - blockwishlist

      to_reset:

        # All modules below are reset when the theme is enabled.

        - blockreassurance

        - blockwishlist
```

You can also have the theme create hooks and attach modules to custom and existing hooks when the theme is enabled.

```yaml
global_settings:

    hooks:

      custom_hooks:

        - name: displayFooterBefore

          title: displayFooterBefore

          description: Add a widget area above the footer

      modules_to_hook:

        displayHeaderTop:

          # displayHeaderTop will have exactly the following

          # modules hooked to it, in the specified order.

          # Each module in this list will be unhooked

          # from all other display hooks it is hooked to.

          - blocklanguages

          - blockcurrencies

          - blockuserinfo

        displayHeaderMiddle:

          # displayHeaderMiddle will have whatever is currently hooked to it

          # kept hooked to it, and blocksearch will be appended

          # to the list (or moved to the end if already hooked there).

          - ~

          - blocksearch

        displayHeaderBottom:

          # displayHeaderBottom will have blocktopmenu and blockcart

          # prepended to it.

          - blocktopmenu

          - blockcart

          - ~

        displayFooter:

          - blocknewsletter

        displayLeftColumn:

          # blockcategories is hooked on all pages on displayLeftColumn

          - blockcategories

          # blocktags is hooked on displayLeftColumn on all pages

          # except "category" and "index"

          - blocktags:

              except_pages:

                - category

                - index
```

### Image settings

Enabling the theme will remove all the existing image types.

Therefore, themes MUST declare their image types, and what they apply to.

### Theme settings

All the settings below can be changed through an interface in the theme’s back office interface, and only depend on the theme/shop combination.

When the theme.yml file is parsed by PrestaShop, the ’theme\_settings’ configuration key is copied to a file named settings\_n.yml, where ’n’ is the id of the shop where the theme is installed (settings\_123456.yml, for instance).

When the configuration is changed through the back office interface, only the settings\_n.yml file is updated - the theme.yml file remains unchanged.

```yaml
global_settings:

    theme_settings:

      default_layout: layout-full-width

      layouts:

        # Specific layout for some pages

        identity: layout-left-side-column

        order-confirmation: layout-left-side-column
```

You can disable `core.js` loading and provide a custom implementation of the whole logic initially contained in `core.js`. To do this, use the `core_scripts` option, set to false, in `theme_settings` like in this example:

```yaml
theme_settings:

    core_scripts: false
```

## Dependencies

When making a theme you may want to add features with custom modules. It’s important that these modules are installed with your theme. These modules should be declared as dependencies so you’re sure prestashop will export them when creating your theme zipball.

So far themes only have modules dependencies.

```yaml
dependencies:

    modules:

      - xx_customslider

      - xx_customproductpage
```

## Related in KB

- [[theme-organization|Theme organization]]
- [[parent-child-feature|Parent/child themes feature]]
- [[asset-management|Asset management]]
- [[../../_agent/panda-expert|Panda Expert agent]]
