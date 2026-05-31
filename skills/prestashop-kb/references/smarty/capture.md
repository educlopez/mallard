---
source_url: https://smarty-php.github.io/smarty/stable/designers/language-builtin-functions/language-function-capture/
topic: smarty
subtopic: variables
ingested: 2026-04-18
tags: [smarty, prestashop, variables]
ps_version: [8, 9]
---

## {capture}

`{capture}` is used to collect the output of the template between the tags into a variable instead of displaying it. Any content between `{capture name='foo'}` and `{/capture}` is collected into the variable specified in the `name` attribute.

The captured content can be used in the template from the variable [`$smarty.capture.foo`](../../language-variables/language-variables-smarty/#smartycapture-languagevariablessmartycapture) where "foo" is the value passed in the `name` attribute. If you do not supply the `name` attribute, then "default" will be used as the name ie `$smarty.capture.default`.

`{capture}'s` can be nested.

## Attributes

| Attribute Name | Required | Description |
| --- | --- | --- |
| name | Yes | The name of the captured block |
| assign | No | The variable name where to assign the captured output to |
| append | No | The name of an array variable where to append the captured output to |

## Option Flags

| Name | Description |
| --- | --- |
| nocache | Disables caching of this captured block |

## Examples

This example demonstrates the capture function.

```
{capture name=some_content assign=popText}
{capture some_content assign=popText} {* short-hand *}
The server is {$my_server_name|upper} at {$my_server_addr}<br>
Your ip is {$my_ip}.
{/capture}
<a href="#">{$popText}</a>
```

This example also demonstrates how multiple calls of capture can be used to create an array with captured content.

```
{capture append="foo"}hello{/capture}I say just {capture append="foo"}world{/capture}
{foreach $foo as $text}{$text} {/foreach}
```

The above example will output:

```
I say just hello world
```

See also [`$smarty.capture`](../../language-variables/language-variables-smarty/#smartycapture-languagevariablessmartycapture),[`{eval}`](../../language-custom-functions/language-function-eval/),[`{fetch}`](../../language-custom-functions/language-function-fetch/), [`fetch()`](../../../programmers/api-functions/api-fetch/) and [`{assign}`](../language-function-assign/).
