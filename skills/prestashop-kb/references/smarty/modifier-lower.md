---
source_url: https://smarty-php.github.io/smarty/stable/designers/language-modifiers/language-modifier-lower/
topic: smarty
subtopic: modifiers
ingested: 2026-04-18
tags: [smarty, prestashop, modifiers]
ps_version: [8, 9]
---

## lower

This is used to lowercase a variable. This is equivalent to the PHP [`strtolower()`](https://www.php.net/strtolower) function.

## Basic usage

```
{$myVar|lower}
```

## Examples

```
<?php

$smarty->assign('articleTitle', 'Two Convicts Evade Noose, Jury Hung.');
```

Where template is:

This will output:

```
Two Convicts Evade Noose, Jury Hung.
two convicts evade noose, jury hung.
```

See also [`upper`](../language-modifier-upper/) and [`capitalize`](../language-modifier-capitalize/).
