---
source_url: https://devdocs.prestashop-project.org/9/themes/reference/template-inheritance/
ps_version: [8, 9]
ingested: 2026-04-18
section: Themes > Template Inheritance
topic: themes
tags: [prestashop, themes, smarty, templates, inheritance]
---

[Improve this page](https://github.com/PrestaShop/docs/edit/9.x/themes/reference/template-inheritance/_index.md "Improve this page") [How](/9/contribute/documentation/how)

PrestaShop relies a lot on template inheritance in order to create the most consistent theme possible, while heavily reducing the amount of duplicated code.

## The principle

Template inheritance allow you to extend a parent template and only redefine the block you need.

The picture below illustrates the example of a specific product page extending a generic one. Say you have three files: `product-pack.tpl` extending `product.tpl`, itself extending `page.tpl`.

![](img/template-inheritance.png)

Schema for template inheritance

The `product-pack.tpl` file will **only** contain the product description part. Everything else will be exactly the same as product page. Even better, the product page will only define the main content of the page, everything else will be taken from its own parent template (ie page.tpl).

The wrong way of doing this would be to extract header, footer and reusable part of the template and *include* them in each file.

The [official Smarty documentation](https://www.smarty.net/inheritance) has a nice and simple example. In their own words:

> Template inheritance is an approach to managing templates that resembles object-oriented programming techniques. Instead of the traditional use of {include …} tags to manage parts of templates, you can inherit the contents of one template to another (like extending a class) and change blocks of content therein (like overriding methods of a class.) This keeps template management minimal and efficient, since each template only contains the differences from the template it extends.

## PrestaShop real life example

In a PrestaShop theme, many pages are very similar, for example template listing products: categories, new products, search results, and so on. All of them display a list of products, so they all extend `catalog/listing/product-list.tpl` (which extends the main layout).

The template will show a title and a list of products underneath. For category page, we want a nice description with an cover image. So we can simply override the *product\_list\_header*

This reduce code to the minimum, without any repetition.

## Related in KB

- [[parent-child-feature|Parent/child themes]]
- [[templates|Templates reference]]
- [[smarty-helpers|Smarty helper functions]]
