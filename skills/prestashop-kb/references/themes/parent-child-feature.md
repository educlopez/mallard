---
source_url: https://devdocs.prestashop-project.org/9/themes/reference/template-inheritance/parent-child-feature/
ps_version: [8, 9]
ingested: 2026-04-18
section: Themes > Parent/Child themes (child theme convention)
topic: themes
tags: [prestashop, themes, child-theme, parent-theme, convention]
---

[Improve this page](https://github.com/PrestaShop/docs/edit/9.x/themes/reference/template-inheritance/parent-child-feature.md "Improve this page") [How](/9/contribute/documentation/how)

PrestaShop relies a lot on template inheritance in order to create the most consistent theme possible, while heavily reducing the amount of duplicated code.

We took it even further introducing the Parent Child theme feature. The point to is to avoid modifying the main theme, so you can update it!

This feature is only useful if you want to slightly modify a theme (to add a block for example). If you need to modify the markup more deeply, modify the theme, don’t create a child theme.

## The principle

So far we talked about extending template within the same theme. But you can also extend templates from another theme.

As a theme developer, you want to create as many blocks as possible so your user can override the minimum amount of code.

## How to create a child theme

First you need to have the theme you want to use as parent in your store `/themes` folder.

Then you can create a very minimal theme, containing only the following files:

```bash
.

  ├── config

  │   └── theme.yml

  └── preview.png
```

Once you have this, you will specify in your child theme `theme.yml` which theme should be used as a parent. The value must be the theme technical name (ie: the theme folder name).

```yaml
parent: classic

name: childtheme

display_name: My first child Theme

version: 1.0.0

assets:

  use_parent_assets: true
```

Find more information about [`use_parent_assets` behavior here](/9/themes/getting-started/theme-yml/#parent--child-settings)

Go ahead, select this theme in your backoffice and you’re all set.

### Overriding templates

Without any further modification the child theme will use every template from the parent theme.

In the following example we’ll modify the category template.

### Re-defining the whole template

Create the category template in your child theme `templates/catalog/listing/category.tpl`. At this point you can do anything you want in this template but most likely you still want to extend product-list template. If so, you don’t have to copy product-list template to your child theme, PrestaShop will use the parent file.

Extend product-list the normal way and override the block you need.

```smarty
{extends file='catalog/listing/product-list.tpl'}
```

### Extending the same template

Another way to overriding the category template would be to extend the parent template and define just the block you need. There is a little difference if you want to override the same template, you need to add the parent resource (note the parent keyword in the example below).

```smarty
{extends file='parent:catalog/listing/category.tpl'}
```

## Related in KB

- [[theme-yml|theme.yml — parent + use_parent_assets]]
- [[template-inheritance|Template inheritance]]
- [[../../_agent/panda-expert|Panda Expert agent]]
- [[../../SETUP-NEW-MACHINE|Panda setup on a new machine]]
