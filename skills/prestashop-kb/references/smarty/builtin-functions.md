---
source_url: https://smarty-php.github.io/smarty/stable/designers/language-builtin-functions/
topic: smarty
subtopic: control-flow
ingested: 2026-04-18
tags: [smarty, prestashop, control-flow]
ps_version: [8, 9]
---

[Skip to content](#built-in-functions)

## Built-in Functions

Smarty comes with several built-in functions. These built-in functions are the integral part of the smarty template engine. They are compiled into corresponding inline PHP code for maximum performance.

You cannot create your own [custom tags](../language-custom-functions/) with the same name; and you should not need to modify the built-in functions.

A few of these functions have an `assign` attribute which collects the result the function to a named template variable instead of being output; much like the [`{assign}`](language-function-assign/) function.

- [{append}](language-function-append/)
- [{assign} or {$var=...}](language-function-assign/)
- [{block}](language-function-block/)
- [{call}](language-function-call/)
- [{capture}](language-function-capture/)
- [{config\_load}](language-function-config-load/)
- [{debug}](language-function-debug/)
- [{extends}](language-function-extends/)
- [{for}](language-function-for/)
- [{foreach}, {foreachelse}](language-function-foreach/)
- [{function}](language-function-function/)
- [{if}, {elseif}, {else}](language-function-if/)
- [{include}](language-function-include/)
- [{insert}](language-function-insert/)
- [{ldelim}, {rdelim}](language-function-ldelim/)
- [{literal}](language-function-literal/)
- [{nocache}](language-function-nocache/)
- [{section}, {sectionelse}](language-function-section/)
- [{setfilter}](language-function-setfilter/)
- [{strip}](language-function-strip/)
- [{while}](language-function-while/)
