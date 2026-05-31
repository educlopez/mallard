---
source_url: https://devdocs.prestashop-project.org/9/development/components/smarty-extensions/
topic: smarty
subtopic: prestashop-extensions
ingested: 2026-04-18
tags: [smarty, prestashop, prestashop-extensions]
ps_version: [8, 9]
---

# PrestaShop-specific Smarty extensions

These tags and modifiers are **not** part of Smarty core — they are plugins registered by PrestaShop. You will see them constantly in `.tpl` files in the Panda theme and in module templates, but you will not find them in [[modifiers|Smarty core modifier docs]] or [[builtin-functions|Smarty core function docs]].

Source: PrestaShop devdocs — Smarty Extensions reference.

## Functions

### `{l}` — translations

The most common PS-specific tag. Inserts a translated string.

```smarty
{l s="Add to cart" d="Shop.Theme.Actions"}
{l s="Hello %name%" d="Shop.Theme.Customer" sprintf=['%name%' => $customer.firstname]}
{l s="Legacy module text" mod="mymodule"}
```

| Arg | Required | Description |
|-----|----------|-------------|
| `s` | yes | Source message to translate |
| `d` | no  | Translation domain (new translation system, e.g. `Shop.Theme.Actions`, `Modules.Mymodule.Shop`) |
| `mod` | no | Module name (legacy translation system) |
| `sprintf` | no | Placeholder replacement array |

Escape when embedding in JS:

```smarty
var msg = '{l|escape:"javascript" s="Don't do this" d="Shop.Theme.Global"}';
```

### `{hook}` — invoke a hook

Executes a display hook and injects its output.

```smarty
{hook h="displayHome"}
{hook h="displayFooter"}
{hook h="displayProductAdditionalInfo" mod="ps_sharebuttons"}   {* specific module only *}
{hook h="displayHome" excl="ps_banner,ps_imageslider"}          {* exclude modules *}
```

Hook names are documented separately — see [[../hooks/_index|hooks index]]. This page only documents the Smarty **syntax**.

### `{widget}` / `{widget_block}` — render module widgets

Renders a module's widget output without needing a hook.

```smarty
{widget name="ps_contactinfo"}
{widget name="ps_contactinfo" hook="displayFooter"}
```

`{widget_block}` lets you override the module's template inline:

```smarty
{widget_block name="ps_linklist"}
  {foreach $linkBlocks as $linkBlock}
    <ul>
      {foreach $linkBlock.links as $link}
        <li><a href="{$link.url}">{$link.title}</a></li>
      {/foreach}
    </ul>
  {/foreach}
{/widget_block}
```

> PS devdocs explicitly warn: avoid `{widget}` where possible — it couples the template to a module.

### `{form_field}` — render a form field

```smarty
{form_field field=$field}
```

`$field` is an associative array with `name`, `type`, `required`, `label`, `value`, `availableValues`, `errors`. Markup lives in `_partials/form-fields.tpl` — override it to customize.

### `{render}` — render a sub-template with UI data

Used mostly for customer/checkout forms.

```smarty
{render file="customer/_partials/login-form.tpl" ui=$login_form}
```

### `{url}` — generate URLs

Handles scheme, domain, base URI, rewriting.

```smarty
<a href="{url entity='product' id=$product.id_product}">{$product.name}</a>
```

Prefer `{url}` over `$link->getXxx(...)`. The `$link` object still exists for backward compatibility but PS devdocs recommend against it. Also see the pre-built `$urls` dataset (below) for common page URLs.

## Modifiers

Provided by PS on top of Smarty core:

### `classname`
Slugifies a string into a valid CSS class name (lowercases, transliterates accents, replaces non-alphanumeric chars with `-`, collapses dashes).

```smarty
{assign var=attr value='Hérè-Is_a-Clàssnåme--@#$$ˆ*(&-----'}
{$attr|classname}   {* → here-is_a-classname *}
```

For the complete PS modifier list, consult the [source reference](https://devdocs.prestashop-project.org/9/development/components/smarty-extensions/). Most modifiers used in `.tpl` files are Smarty core — see [[modifiers]].

## Common PrestaShop context variables

Available in front templates by default (injected by the front controller and `FrontController::assignGeneralPurposeVariables`):

| Var | Contents | Example |
|-----|----------|---------|
| `$urls` | Pre-built URLs dataset | `{$urls.base_url}`, `{$urls.pages.cart}`, `{$urls.pages.my_account}`, `{$urls.img_url}`, `{$urls.theme_assets}` |
| `$link` | `Link` object (legacy — prefer `{url}`) | `{$link->getPageLink('cms', true, null, 'id_cms=1')}` |
| `$configuration` | Shop configuration values | `{$configuration.display_taxes_label}` |
| `$customer` | Current customer data | `{$customer.firstname}`, `{$customer.is_logged}` |
| `$cart` | Current cart | `{$cart.products}`, `{$cart.totals.total.value}` |
| `$language` | Active language | `{$language.iso_code}`, `{$language.is_rtl}` |
| `$shop` | Current shop | `{$shop.name}`, `{$shop.logo}` |
| `$page` | Current page metadata | `{$page.page_name}`, `{$page.body_classes}`, `{$page.meta.title}` |
| `$currency` | Active currency | `{$currency.iso_code}`, `{$currency.sign}` |

On product pages: `$product`. On category pages: `$category`, `$listing`, `$products`. On CMS pages: `$cms`.

## Cross-links

- Override cascade (how PS decides which `.tpl` wins): [[../themes/parent-child-feature|Parent/child themes]]
- Template inheritance with `{extends}` + `{block}`: [[../themes/template-inheritance|PS template inheritance]] and core [[extends]], [[block]]
- Hook names and concept: [[../hooks/_index|hooks index]] and [[../hooks/hooks-overview|hooks overview]]
- Template organization and layout system: [[../themes/templates|templates reference]]
- PS devdocs helper summary (intentionally thin): [[../themes/smarty-helpers|smarty-helpers]]

## What is NOT here

- Registering custom Smarty plugins from PHP (module development) — belongs under modules.
- Smarty caching / security config — PHP-side, out of scope.
- Full modifier catalog — see source URL above if you hit an exotic one.
