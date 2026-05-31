---
source_url: https://smarty-php.github.io/smarty/stable/designers/language-modifiers/language-modifier-default/
topic: smarty
subtopic: modifiers
ingested: 2026-04-18
tags: [smarty, prestashop, modifiers]
ps_version: [8, 9]
---

## default

This is used to set a default value for a variable. If the variable is unset or an empty string, the given default value is printed instead. Default takes the one argument.

## Basic usage

```
{$myVar|default:"(none)"}
```

## Parameters

| Parameter | Type | Required | Default | Description |
| --- | --- | --- | --- | --- |
| 1 | string | No | *empty* | This is the default value to output if the variable is empty. |

## Examples

```
<?php

    $smarty->assign('articleTitle', 'Dealers Will Hear Car Talk at Noon.');
    $smarty->assign('email', '');
```

Where template is:

```
{$articleTitle|default:'no title'}
{$myTitle|default:'no title'}
{$email|default:'No email address available'}
```

Will output:

```
Dealers Will Hear Car Talk at Noon.
no title
No email address available
```

See also the [default variable handling](../../../appendixes/tips/#default-variable-handling) and the [blank variable handling](../../../appendixes/tips/#blank-variable-handling) pages.
