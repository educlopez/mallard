---
source_url: https://smarty-php.github.io/smarty/stable/designers/language-modifiers/language-modifier-upper/
topic: smarty
subtopic: modifiers
ingested: 2026-04-18
tags: [smarty, prestashop, modifiers]
ps_version: [8, 9]
---

## upper

This is used to uppercase a variable. This is equivalent to the PHP [`strtoupper()`](https://www.php.net/strtoupper) function.

## Basic usage

```
{$myVar|upper}
```

## Examples

```
<?php
$smarty->assign('articleTitle', "If Strike isn't Settled Quickly it may Last a While.");
```

Where template is:

Will output:

```
If Strike isn't Settled Quickly it may Last a While.
IF STRIKE ISN'T SETTLED QUICKLY IT MAY LAST A WHILE.
```

See also [`lower`](../language-modifier-lower/) and [`capitalize`](../language-modifier-capitalize/).
