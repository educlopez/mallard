---
source_url: https://devdocs.prestashop-project.org/9/modules/concepts/hooks/
ps_version: [8, 9]
ingested: 2026-04-18
section: Hooks > Overview (concept, naming, registration, execution)
topic: hooks
tags: [prestashop, hooks, concept, registration, display, action]
---

[Improve this page](https://github.com/PrestaShop/docs/edit/9.x/modules/concepts/hooks/_index.md "Improve this page") [How](/9/contribute/documentation/how)

[Hooks](https://devdocs.prestashop-project.org/9/development/components/hook/) are a way to associate your code to some specific PrestaShop events.

Most of the time, they are used to insert content in a page. The place it will be added (header, footer, left or right column …) will depend on the hook you choose.

Hooks can also be used to perform specific actions under certain circumstances (i.e. sending an e-mail to the client on an order creation).

## Naming scheme

Hook names are prefixed with “action” or “display”. This prefix indicates if a hook is triggered by an event or if it’s used to display content:

**action<Something>**

Triggered by specific events that take place in PrestaShop.

**display<Something>**

Result in something being displayed, either in the front-end or the back-end.

## Using hooks

### Registration

Every hook you want to use must be registered first. This is usually done during the installation of your module, by calling the method `Module::registerHook($hookName)`.

If you do not know where you can register, [a list of available hooks](https://devdocs.prestashop-project.org/9/modules/concepts/hooks/list-of-hooks/) is available.

### Execution

For each registered hook, you must create a non-static public method, starting with the “hook” keyword followed by the name of the hook you want to use (starting with either “display” or “action”).

This method receives one (and only one) argument: an array of the contextual information sent to the hook.

Remember, in order for a module to respond to a hook call, it must be registered within PrestaShop.

## Triggering a hook

### In a controller (legacy)

It is easy to call a hook from within a controller: you simply have to use its name with the `Hook::exec($hook_name, $hook_args = array())` method. Some parameters can be sent as well.

For instance:

```php
$this->context->smarty->assign(

    'HOOK_LEFT_COLUMN',

    Hook::exec('displayLeftColumn')

);
```

### In a controller (Symfony)

In a Symfony controller, please use the dispatchHook method of the inherited `FrameworkBundleAdminController` class:

```php
protected function dispatchHook($hookName, array $parameters);
```

### In a theme, with Smarty

It is easy to call a hook from within a template file (`.tpl`): you simply have to use its name with the hook function. You can add the name of a module that you want the hook execute.

Basic call of a hook:

```fallback
{hook h='hookName'}
```

Call of a hook for a specific module:

```fallback
{hook h='hookName' mod='modulename'}
```

### In a theme, with Twig

It is easy to call a hook from within a twig template file (`.html.twig`): you simply have to use its name with the renderHook twig function. You can add params as a second argument.

Basic call of a hook:

```fallback
{{ renderHook('hookName', { params }) }}
```

## Going further: Creating your own hook

You can create new PrestaShop hooks by adding a new record in the Hook table. This can be done with the Hook class, which inherit ObjectModel features:

```php
$hook = new Hook();

$hook->name = 'displayAtSpecificPlace';

$hook->title = 'The name of your hook';

$hook->description = 'This is a custom hook!';

$hook->position = 1;

$hook->add(); // return true on success

You can check if hook exists before this with Hook::getIdByName('hook_name')
```

… but PrestaShop enables you to do it the easy way:

```php
$this->registerHook('displayAtSpecificPlace');
```

If the hook “displayAtSpecificPlace” doesn’t exist, PrestaShop will create it for you but be careful: this will also plug the current module to the hook.

[Next: Use hooks on modern pages ›](https://devdocs.prestashop-project.org/9/modules/concepts/hooks/use-hooks-on-modern-pages/ "Use hooks on modern pages")

## Related in KB

- [[list-of-hooks|List of all hooks]]
- [[../themes/theme-hooks|Theme hooks reference]]
- [[../modules/widgets|Widgets]]
