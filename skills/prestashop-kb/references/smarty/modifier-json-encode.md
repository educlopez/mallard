---
source_url: https://smarty-php.github.io/smarty/stable/designers/language-modifiers/language-modifier-json-encode/
topic: smarty
subtopic: modifiers
ingested: 2026-04-18
tags: [smarty, prestashop, modifiers]
ps_version: [8, 9]
---

## json\_encode

Transforms a value into a valid JSON string.

## Basic usage

```
{$user|json_encode}
```
Depending on the value of `$user` this would return a string in JSON-format, e.g. `{"username":"my_username","email":"my_username@smarty.net"}`.

## Parameters

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| 1 | int | No | bitmask of flags, directly passed to [PHP's json\_encode](https://www.php.net/json_encode) |

## Examples

By passing `16` as the second parameter, you can force json\_encode to always format the JSON-string as an object. Without it, an array `$myArray = ["a","b"]` would be formatted as a javascript array:

```
{$myArray|json_encode} # renders: ["a","b"]
{$myArray|json_encode:16} # renders: {"0":"a","1":"b"}
```
