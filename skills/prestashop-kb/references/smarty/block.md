---
source_url: https://smarty-php.github.io/smarty/stable/designers/language-builtin-functions/language-function-block/
topic: smarty
subtopic: includes
ingested: 2026-04-18
tags: [smarty, prestashop, includes]
ps_version: [8, 9]
---

[Skip to content](#block)

## {block}

`{block}` is used to define a named area of template source for template inheritance. For details see section of [Template Inheritance](../../../api/inheritance/).

The `{block}` template source area of a child template will replace the corresponding areas in the parent template(s).

Optionally `{block}` areas of child and parent templates can be merged into each other. You can append or prepend the parent `{block}` content by using the `append` or `prepend` option flag with the child's `{block}` definition. With `{$smarty.block.parent}` the `{block}` content of the parent template can be inserted at any location of the child `{block}` content. `{$smarty.block.child}` inserts the `{block}` content of the child template at any location of the parent `{block}`.

`{blocks}'s` can be nested.

## Attributes

| Attribute Name | Required | Description |
| --- | --- | --- |
| name | yes | The name of the template source block |
| assign | no | The name of variable to assign the output of the block to. |

> **Note**
> 
> The assign attribute only works on the block that actually gets executed, so you may need to add it to each child block as well.

## Option Flags (in child templates only):

| Name | Description |
| --- | --- |
| append | The `{block}` content will be appended to the content of the parent template `{block}` |
| prepend | The `{block}` content will be prepended to the content of the parent template `{block}` |
| hide | Ignore the block content if no child block of same name is existing. |
| nocache | Disables caching of the `{block}` content |

## Examples

parent.tpl

```
<html>
      <head>
        <title>{block name="title"}Default Title{/block}</title>
        <title>{block "title"}Default Title{/block}</title>  {* short-hand  *}
      </head>
    </html>
```

child.tpl

```
{extends file="parent.tpl"} 
    {block name="title"}
    Page Title
    {/block}
```

The result would look like

parent.tpl

```
<html>
      <head>
        <title>{block name="title"}Title - {/block}</title>
      </head>
    </html>
```

child.tpl

```
{extends file="parent.tpl"} 
    {block name="title" append}
        Page Title
    {/block}
```

The result would look like

```
<html>
      <head>
        <title>Title - Page Title</title>
      </head>
    </html>
```

parent.tpl

```
<html>
      <head>
        <title>{block name="title"} is my title{/block}</title>
      </head>
    </html>
```

child.tpl

```
{extends file="parent.tpl"} 
    {block name="title" prepend}
    Page Title
    {/block}
```

The result would look like

```
<html>
      <head>
        <title>Page title is my titel</title>
      </head>
    </html>
```

parent.tpl

```
<html>
      <head>
        <title>{block name="title"}The {$smarty.block.child} was inserted here{/block}</title>
      </head>
    </html>
```

child.tpl

```
{extends file="parent.tpl"} 
    {block name="title"}
        Child Title
    {/block}
```

The result would look like

```
<html>
      <head>
        <title>The Child Title was inserted here</title>
      </head>
    </html>
```

parent.tpl

child.tpl

```
{extends file="parent.tpl"} 
    {block name="title"}
        You will see now - {$smarty.block.parent} - here
    {/block}
```

The result would look like

```
<html>
      <head>
        <title>You will see now - Parent Title - here</title>
      </head>
    </html>
```

See also [Template Inheritance](../../../api/inheritance/),[`$smarty.block.parent`](../../language-variables/language-variables-smarty/#smartyblockparent-languagevariablessmartyblockparent),[`$smarty.block.child`](../../language-variables/language-variables-smarty/#smartyblockchild-languagevariablessmartyblockchild), and [`{extends}`](../language-function-extends/)
