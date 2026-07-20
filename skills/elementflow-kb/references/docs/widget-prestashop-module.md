---
source_url: https://elementflow.io/third/posts/prestashop-module-widget-6857.html
ps_version: [8, 9]
ingested: 2026-07-20
tags: [elementflow, docs, widget]
section: Widgets
---
# PrestaShop module widget

Element Flow can work with about **99% of PrestaShop modules**. Some of them can work **even better** with Element Flow, because a module can only be displayed where its hooks are located — moving it around the page normally is hard, but with Element Flow it can be placed exactly where you like.

### Module

Select a module from the drop-down list, or fill in its name in the **The name of your module** field. The name shown on the back office (e.g. **Wishlist**) is the **display name — don't use it**.

**Right-click** (not click) on the module's **Configure** button to open its link address and read the real module name from the URL. For example, the Wishlist module's real name/folder is **blockwishlist**.

### Hook

Select a hook from the **Hook** dropdown menu, or fill in one in **The name of your hook** field.

**How to find which hooks to use:**

1. Check the module's documentation or contact its developers.
2. Find them in the back office: go to BO > Design > Positions, filter the list with the **Show** dropdown to show all hooks the module is transplanted on, and look for hooks starting with **display** (e.g. `displayFooter` for the native Link list module, ps_linklist).
3. Find them in the module's main file: search for `function hookDisplay...`, and if nothing, search for `function renderWidget`.

If a site is built with Element Flow **partially**, a module can sometimes appear twice (e.g. once where added via the widget, once via its native hook like `displayFooter`). To fix it, remove the module from the native hook (`displayFooter`) and use `renderWidget` in the PrestaShop module widget instead.

### Parameters

Most hooks don't need parameters. Some hooks on the product page / product miniature require the `product` variable — if you select a hook like **displayProductAdditionalInfo**, the `product` variable is added by default. Add more variables via the **Custom parameters** field, e.g. `category|$category`. The resulting Smarty code looks like:

```
{hook h='displayProductAdditionalInfo' product=$product category=$category}
```

If you type a hook name manually instead of selecting one, you must also add any required variables (e.g. `product|$product`) yourself via **Custom parameters**.

### PrestaShop module integration

A few modules need extra integration steps to work with Element Flow — see the dedicated PrestaShop module integration page for details.
