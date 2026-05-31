---
source_url: https://smarty-php.github.io/smarty/stable/designers/language-builtin-functions/language-function-while/
topic: smarty
subtopic: control-flow
ingested: 2026-04-18
tags: [smarty, prestashop, control-flow]
ps_version: [8, 9]
---

## {while}

`{while}` loops in Smarty have much the same flexibility as PHP [while](https://www.php.net/while) statements, with a few added features for the template engine. Every `{while}` must be paired with a matching `{/while}`. All [operators](../../language-basic-syntax/language-syntax-operators/) are recognized, such as *\==*,*||*, *or*, *&&*, *and*, etc and you can use modifiers as functions, such as *is\_array()*.

## Examples

```
{while $foo > 0}
  {$foo--}
{/while}
```

The above example will count down the value of $foo until 1 is reached.

See also [`{foreach}`](../language-function-foreach/),[`{for}`](../language-function-for/) and [`{section}`](../language-function-section/).
