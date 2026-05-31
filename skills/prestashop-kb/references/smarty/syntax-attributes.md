---
source_url: https://smarty-php.github.io/smarty/stable/designers/language-basic-syntax/language-syntax-attributes/
topic: smarty
subtopic: variables
ingested: 2026-04-18
tags: [smarty, prestashop, variables]
ps_version: [8, 9]
---

## Attributes

Most of the [tags](../language-syntax-tags/) take attributes that specify or modify their behavior. Attributes to Smarty functions are much like HTML attributes. Static values don't have to be enclosed in quotes, but it is required for literal strings. Variables with or without modifiers may also be used, and should not be in quotes. You can even use PHP function results, plugin results and complex expressions.

Some attributes require boolean values (TRUE or FALSE). These can be specified as `true` and `false`. If an attribute has no value assigned it gets the default boolean value of true.

## Examples

> **Note**
> 
> Although Smarty can handle some very complex expressions and syntax, it is a good rule of thumb to keep the template syntax minimal and focused on presentation. If you find your template syntax getting too complex, it may be a good idea to move the bits that do not deal explicitly with presentation to PHP by way of plugins or modifiers.
