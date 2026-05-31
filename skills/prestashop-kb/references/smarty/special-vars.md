---
source_url: https://smarty-php.github.io/smarty/stable/designers/language-variables/language-variables-smarty/
topic: smarty
subtopic: special-vars
ingested: 2026-04-18
tags: [smarty, prestashop, special-vars]
ps_version: [8, 9]
---

## {$smarty} reserved variable

The PHP reserved `{$smarty}` variable can be used to access several environment and request variables. The full list of them follows.

## Request variables

The [request variables](https://www.php.net/reserved.variables) such as `$_GET`, `$_POST`, `$_COOKIE`, `$_SERVER`, `$_ENV` and `$_SESSION` can be accessed as demonstrated in the examples below:

> **Note**
> 
> For historical reasons `{$SCRIPT_NAME}` is shorthand for `{$smarty.server.SCRIPT_NAME}`.
> 
> ```
> <a href="{$SCRIPT_NAME}?page=smarty">click me</a>
> <a href="{$smarty.server.SCRIPT_NAME}?page=smarty">click me</a>
> ```
> 
> **Note**
> 
> Although Smarty provides direct access to PHP super globals for convenience, it should be used with caution. Directly accessing super globals mixes underlying application code structure with templates. A good practice is to assign specific needed values to template vars.

## {$smarty.now}

The current [timestamp](https://www.php.net/function.time) can be accessed with `{$smarty.now}`. The value reflects the number of seconds passed since the so-called Epoch on January 1, 1970, and can be passed directly to the [`date_format`](../../language-modifiers/language-modifier-date-format/) modifier for display. Note that [`time()`](https://www.php.net/function.time) is called on each invocation; eg a script that takes three seconds to execute with a call to `$smarty.now` at start and end will show the three-second difference.

```
{* use the date_format modifier to show current date and time *}
{$smarty.now|date_format:'%Y-%m-%d %H:%M:%S'}
```

## {$smarty.const}

You can access PHP constant values directly.

```
<?php
// the constant defined in php
define('MY_CONST_VAL','CHERRIES');
```

Output the constant in a template with

```
{$smarty.const.MY_CONST_VAL}
```

> **Note**
> 
> Although Smarty provides direct access to PHP constants for convenience, it is typically avoided as this is mixing underlying application code structure into the templates. A good practice is to assign specific needed values to template vars.

## {$smarty.capture}

Template output captured via the built-in [`{capture}..{/capture}`](../../language-builtin-functions/language-function-capture/) function can be accessed using the `{$smarty.capture}` variable. See the [`{capture}`](../../language-builtin-functions/language-function-capture/) page for more information.

## {$smarty.config}

`{$smarty.config}` variable can be used to refer to loaded [config variables](../language-config-variables/). `{$smarty.config.foo}` is a synonym for `{#foo#}`. See the [{config\_load}](../../language-builtin-functions/language-function-config-load/) page for more info.

## {$smarty.section}

The `{$smarty.section}` variables can be used to refer to [`{section}`](../../language-builtin-functions/language-function-section/) loop properties. These have some very useful values such as `.first`, `.index`, etc.

> **Note**
> 
> The `{$smarty.foreach}` variable is no longer used with the new [`{foreach}`](../../language-builtin-functions/language-function-foreach/) syntax, but is still supported with Smarty 2.x style foreach syntax.

## {$smarty.template}

Returns the name of the current template being processed (without the directory).

## {$smarty.template\_object}

Returns the template object of the current template being processed.

## {$smarty.current\_dir}

Returns the name of the directory for the current template being processed if it is loaded from the filesystem (the default).

## {$smarty.version}

Returns the version of Smarty the template was compiled with.

## {$smarty.block.child}

Returns block text from child template. See [Template inheritance](../../../api/inheritance/).

## {$smarty.block.parent}

Returns block text from parent template. See [Template inheritance](../../../api/inheritance/)

## {$smarty.ldelim}, {$smarty.rdelim}

These variables are used for printing the left-delimiter and right-delimiter value literally, the same as [`{ldelim},{rdelim}`](../../language-builtin-functions/language-function-ldelim/).

See also [assigned variables](../language-assigned-variables/) and [config variables](../language-config-variables/)
