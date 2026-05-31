---
source_url: https://smarty-php.github.io/smarty/stable/designers/language-modifiers/language-modifier-truncate/
topic: smarty
subtopic: modifiers
ingested: 2026-04-18
tags: [smarty, prestashop, modifiers]
ps_version: [8, 9]
---

## truncate

This truncates a variable to a character length, the default is 80. As an optional second parameter, you can specify a string of text to display at the end if the variable was truncated. The characters in the string are included with the original truncation length. By default,`truncate` will attempt to cut off at a word boundary. If you want to cut off at the exact character length, pass the optional third parameter of TRUE.

## Basic usage

```
{$myVar|truncate:40:"..."}
```

## Parameters

| Parameter Position | Type | Required | Default | Description |
| --- | --- | --- | --- | --- |
| 1 | integer | No | 80 | This determines how many characters to truncate to. |
| 2 | string | No | ... | This is a text string that replaces the truncated text. Its length is included in the truncation length setting. |
| 3 | boolean | No | FALSE | This determines whether or not to truncate at a word boundary with FALSE, or at the exact character with TRUE. |
| 4 | boolean | No | FALSE | This determines whether the truncation happens at the end of the string with FALSE, or in the middle of the string with TRUE. Note that if this setting is TRUE, then word boundaries are ignored. |

## Examples

```
<?php
$smarty->assign('articleTitle', 'Two Sisters Reunite after Eighteen Years at Checkout Counter.');
```

where template is:

This will output:

```
Two Sisters Reunite after Eighteen Years at Checkout Counter.
Two Sisters Reunite after Eighteen Years at Checkout Counter.
Two Sisters Reunite after...
Two Sisters Reunite after
Two Sisters Reunite after---
Two Sisters Reunite after Eigh
Two Sisters Reunite after E...
Two Sisters Re..ckout Counter.
```
