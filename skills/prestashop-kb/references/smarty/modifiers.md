---
source_url: https://smarty-php.github.io/smarty/stable/designers/language-modifiers/
topic: smarty
subtopic: modifiers
ingested: 2026-04-18
tags: [smarty, prestashop, modifiers]
ps_version: [8, 9]
---

## Variable Modifiers

Variable modifiers can be applied to [variables](../language-variables/), [custom tags](../language-custom-functions/) or strings. To apply a modifier, specify the value followed by a `|` (pipe) and the modifier name. A modifier may accept additional parameters that affect its behavior. These parameters follow the modifier name and are separated by a `:`(colon).

Modifiers can be applied to any type of variables, including arrays and objects.

## Examples

```
{* apply modifier to a variable *}
{$title|upper}

{* modifier with parameters *}
{$title|truncate:40:"..."}

{* apply modifier to a function parameter *}
{html_table loop=$myvar|upper}

{* with parameters *}
{html_table loop=$myvar|truncate:40:"..."}

{* apply modifier to literal string *}
{"foobar"|upper}

{* using date_format to format the current date *}
{$smarty.now|date_format:"%Y/%m/%d"}

{* apply modifier to a custom function *}
{mailto|upper address="smarty@example.com"}

{* using  php's str_repeat *}
{"="|str_repeat:80}

{* php's count *}
{$myArray|@count}

{* this will uppercase the whole array *}
<select name="name_id">
{html_options output=$my_array|upper}
</select>
```

## Combining Modifiers

You can apply any number of modifiers to a variable. They will be applied in the order they are combined, from left to right. They must be separated with a `|` (pipe) character.

```
<?php

$smarty->assign('articleTitle', 'Smokers are Productive, but Death Cuts Efficiency.');
```

where template is:

The above example will output:

```
Smokers are Productive, but Death Cuts Efficiency.
S M O K E R S   A R ....snip....  H   C U T S   E F F I C I E N C Y .
s m o k e r s   a r ....snip....  b u t   d e a t h   c u t s...
s m o k e r s   a r e   p r o d u c t i v e ,   b u t . . .
s m o k e r s   a r e   p. . .
```

## Using modifiers in expressions

Modifiers can also be used in expressions. For example, you can use the [isset modifier](language-modifier-isset/) to test if a variable holds a value different from null.

```
{if $varA|isset}
    <b>variable A is set</b>
{/if}
```

You can also use modifiers in expressions in a PHP-style syntax:

```
{if isset($varA)}
    <b>variable A is set</b>
{/if}
```

See also [`registerPlugin()`](../../programmers/api-functions/api-register-plugin/), [combining modifiers](../language-combining-modifiers/). and [extending smarty with plugins](../../api/extending/introduction/)
