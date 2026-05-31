---
source_url: https://smarty-php.github.io/smarty/stable/designers/language-builtin-functions/language-function-for/
topic: smarty
subtopic: control-flow
ingested: 2026-04-18
tags: [smarty, prestashop, control-flow]
ps_version: [8, 9]
---

## {for}

The `{for}{forelse}` tag is used to create simple loops. The following different formats are supported:

- `{for $var=$start to $end}` simple loop with step size of 1.
- `{for $var=$start to $end step $step}` loop with individual step size.

`{forelse}` is executed when the loop is not iterated.

## Attributes

| Attribute | Required | Description |
| --- | --- | --- |
| max | No | Limit the number of iterations |

## Option Flags

| Name | Description |
| --- | --- |
| nocache | Disables caching of the `{for}` loop |

## Examples

The above example will output:

```
<ul>
    {for $foo=3 to $to max=3}
        <li>{$foo}</li>
    {/for}
</ul>
```

The above example will output:

The above example will output:

```
no iteration
```

See also [`{foreach}`](../language-function-foreach/),[`{section}`](../language-function-section/) and [`{while}`](../language-function-while/)
