---
source_url: https://smarty-php.github.io/smarty/stable/designers/language-modifiers/language-modifier-escape/
topic: smarty
subtopic: modifiers
ingested: 2026-04-18
tags: [smarty, prestashop, modifiers]
ps_version: [8, 9]
---

## escape

`escape` is used to encode or escape a variable to `html`, `url`,`single quotes`, `hex`, `hexentity`, `javascript` and `mail`. By default its `html`.

## Basic usage

```
{$myVar|escape}
```

## Parameters

| Parameter Position | Type | Required | Possible Values | Default | Description |
| --- | --- | --- | --- | --- | --- |
| 1 | string | No | `html`, `htmlall`, `url`, `urlpathinfo`, `quotes`, `hex`, `hexentity`, `javascript`, `mail` | `html` | This is the escape format to use. |
| 2 | string | No | `ISO-8859-1`, `UTF-8`, and any character set supported by [`htmlentities()`](https://www.php.net/htmlentities) | `UTF-8` | The character set encoding passed to htmlentities() et. al. |
| 3 | boolean | No | FALSE | TRUE | Double encode entities from & to &amp; (applies to `html` and `htmlall` only) |

## Examples

```
<?php

$smarty->assign('articleTitle',
                "'Stiff Opposition Expected to Casketless Funeral Plan'"
                );
$smarty->assign('EmailAddress','smarty@example.com');
```

These are example `escape` template lines followed by the output

This snippet is useful for emails, but see also [`{mailto}`](../../language-custom-functions/language-function-mailto/)

```
{* email address mangled *}
<a href="mailto:{$EmailAddress|escape:'hex'}">{$EmailAddress|escape:'mail'}</a>
```

See also [auto-escaping](../../../api/configuring/#enabling-auto-escaping), [escaping smarty parsing](../../language-basic-syntax/language-escaping/),[`{mailto}`](../../language-custom-functions/language-function-mailto/) and the [obfuscating email addresses](../../../appendixes/tips/#obfuscating-e-mail-addresses) pages.
