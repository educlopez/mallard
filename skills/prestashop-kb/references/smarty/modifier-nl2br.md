---
source_url: https://smarty-php.github.io/smarty/stable/designers/language-modifiers/language-modifier-nl2br/
topic: smarty
subtopic: modifiers
ingested: 2026-04-18
tags: [smarty, prestashop, modifiers]
ps_version: [8, 9]
---

## nl2br

All `"\n"` line breaks will be converted to html `<br />` tags in the given variable. This is equivalent to the PHP\\'s [`nl2br()`](https://www.php.net/nl2br) function.

## Basic usage

```
{$myVar|nl2br}
```

## Examples

```
<?php

$smarty->assign('articleTitle',
                "Sun or rain expected\ntoday, dark tonight"
                );
```

Where the template is:

```
{$articleTitle|nl2br}
```

Will output:

```
Sun or rain expected<br />today, dark tonight
```

See also [`word_wrap`](../language-modifier-wordwrap/),[`count_paragraphs`](../language-modifier-count-paragraphs/) and [`count_sentences`](../language-modifier-count-sentences/).
