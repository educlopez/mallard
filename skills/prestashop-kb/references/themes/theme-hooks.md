---
source_url: https://devdocs.prestashop-project.org/9/themes/reference/hooks/
ps_version: [8, 9]
ingested: 2026-04-18
section: Themes > Theme Hooks reference
topic: themes
tags: [prestashop, themes, hooks, display]
---

[Improve this page](https://github.com/PrestaShop/docs/edit/9.x/themes/reference/hooks/_index.md "Improve this page") [How](/9/contribute/documentation/how)

## Create hooks available in Front Office

This section of the documentation is only about front office hooks: [display and action](/9/development/components/hook/).

## Creating a dynamic hook

When your module or theme calls a hook, PrestaShop executes it.

This is how it is called from a PHP file:

```php
<?php

Hook::exec('MyCustomHook');
```

This is how it is called from a Smarty template:

```html
{hook h='MyCustomHook'}
```

## Register the hook to make it visible and reusable

If you add a hook call, it is better to register it.

This will enable Back Office user to:

- see it in the hooks list
- be able to plug some modules on it (in Position page)
- allow other modules to listen to this hook being called and add some extra behavior

You can register your hook from your theme’s [theme.yml file](/9/themes/getting-started/theme-yml/):

You can also register your hook from a module:

## Related in KB

- [[../hooks/hooks-overview|Hooks overview]]
- [[../hooks/list-of-hooks|List of available hooks]]
