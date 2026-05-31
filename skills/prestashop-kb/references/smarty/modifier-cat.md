---
source_url: https://smarty-php.github.io/smarty/stable/designers/language-modifiers/language-modifier-cat/
topic: smarty
subtopic: modifiers
ingested: 2026-04-18
tags: [smarty, prestashop, modifiers]
ps_version: [8, 9]
---

## cat

This value is concatenated to the given variable.

## Basic usage

```
{$myVar|cat:' units'}
```

## Parameters

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| 1 | string | No | This value to concatenate to the given variable. |

## Examples

```
<?php

    $smarty->assign('articleTitle', "Psychics predict world didn't end");
```

Where template is:

```
{$articleTitle|cat:' yesterday.'}
```

Will output:

```
Psychics predict world didn't end yesterday.
```
