---
source_url: https://smarty-php.github.io/smarty/stable/designers/language-basic-syntax/language-syntax-tags/
topic: smarty
subtopic: variables
ingested: 2026-04-18
tags: [smarty, prestashop, variables]
ps_version: [8, 9]
---

Every Smarty tag either prints a [variable](../language-syntax-variables/) or invokes some sort of function. These are processed and displayed by enclosing the function and its [attributes](../language-syntax-attributes/) within delimiters like so: `{funcname attr1="val1" attr2="val2"}`.

## Examples

- Both [built-in functions](../../language-builtin-functions/) and [custom functions](../../language-custom-functions/) have the same syntax within templates.
- Built-in functions are the **inner** workings of Smarty, such as [`{if}`](../../language-builtin-functions/language-function-if/), [`{section}`](../../language-builtin-functions/language-function-section/) and [`{strip}`](../../language-builtin-functions/language-function-strip/). There should be no need to change or modify them.
- Custom tags are **additional** tags implemented via [plugins](../../../api/extending/introduction/). They can be modified to your liking, or you can create new ones. [`{html_options}`](../../language-custom-functions/language-function-html-options/) is an example of a custom function.

See also [`registerPlugin()`](../../../programmers/api-functions/api-register-plugin/)
