---
source_url: https://smarty-php.github.io/smarty/stable/designers/language-modifiers/language-modifier-replace/
topic: smarty
subtopic: modifiers
ingested: 2026-04-18
tags: [smarty, prestashop, modifiers]
ps_version: [8, 9]
---

## replace

A simple search and replace on a variable. This is equivalent to the PHP's [`str_replace()`](https://www.php.net/str_replace) function.

## Basic usage

```
{$myVar|replace:"foo":"bar"}
```

## Parameters

| Parameter Position | Type | Required | Description |
| --- | --- | --- | --- |
| 1 | string | Yes | This is the string of text to be replaced. |
| 2 | string | Yes | This is the string of text to replace with. |

## Examples

```
<?php

$smarty->assign('articleTitle', "Child's Stool Great for Use in Garden.");
```

Where template is:

```
{$articleTitle}
{$articleTitle|replace:'Garden':'Vineyard'}
{$articleTitle|replace:' ':'   '}
```

Will output:

```
Child's Stool Great for Use in Garden.
Child's Stool Great for Use in Vineyard.
Child's   Stool   Great   for   Use   in   Garden.
```

See also [`regex_replace`](../language-modifier-regex-replace/) and [`escape`](../language-modifier-escape/).
