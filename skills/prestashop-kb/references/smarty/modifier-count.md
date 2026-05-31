---
source_url: https://smarty-php.github.io/smarty/stable/designers/language-modifiers/language-modifier-count/
topic: smarty
subtopic: modifiers
ingested: 2026-04-18
tags: [smarty, prestashop, modifiers]
ps_version: [8, 9]
---

## count

Returns the number of elements in an array (or Countable object). Will return 0 for null. Returns 1 for any other type (such as a string).

If the optional mode parameter is set to 1, count() will recursively count the array. This is particularly useful for counting all the elements of a multidimensional array.

## Basic usage

```
{if $myVar|count > 3}4 or more{/if}
{if count($myVar) > 3}4 or more{/if}
```

## Parameters

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| 1 | int | No | If set to 1, count() will recursively count the array. |
