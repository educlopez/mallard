---
source_url: https://smarty-php.github.io/smarty/stable/designers/language-variables/
topic: smarty
subtopic: variables
ingested: 2026-04-18
tags: [smarty, prestashop, variables]
ps_version: [8, 9]
---

## Variables

Smarty has several types of variables. The type of the variable depends on what symbol it is prefixed or enclosed within.

- [Variables assigned from PHP](language-assigned-variables/)
- [Variables loaded from config files](language-config-variables/)
- [{$smarty} reserved variable](language-variables-smarty/)

Variables in Smarty can be either displayed directly or used as arguments for [tags](../language-basic-syntax/language-syntax-tags/),[attributes](../language-basic-syntax/language-syntax-attributes/) and [modifiers](../language-modifiers/), inside conditional expressions, etc. To print a variable, simply enclose it in the [delimiters](../language-basic-syntax/language-escaping/) so that it is the only thing contained between them.

## Scopes

You can assign variables to specific [variable scopes](language-variable-scopes/).

> **Note**
> 
> An easy way to examine assigned Smarty variables is with the [debugging console](../chapter-debugging-console/).
