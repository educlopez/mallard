---
source_url: https://smarty-php.github.io/smarty/stable/designers/language-basic-syntax/language-escaping/
topic: smarty
subtopic: variables
ingested: 2026-04-18
tags: [smarty, prestashop, variables]
ps_version: [8, 9]
---

## Escaping Smarty parsing

It is sometimes desirable or even necessary to have Smarty ignore sections it would otherwise parse. A classic example is embedding Javascript or CSS code in a template. The problem arises as those languages use the { and } characters which are also the default [delimiters](../../language-builtin-functions/language-function-ldelim/) for Smarty.

> **Note**
> 
> A good practice for avoiding escapement altogether is by separating your Javascript/CSS into their own files and use standard HTML methods to access them. This will also take advantage of browser script caching. When you need to embed Smarty variables/functions into your Javascript/CSS, then the following applies.

In Smarty templates, the { and } braces will be ignored so long as they are surrounded by white space. This behavior can be disabled by setting the Smarty class variable [`$auto_literal`](../../../programmers/api-variables/variable-auto-literal/) to false.

## Examples

```
<script>
   // the following braces are ignored by Smarty
   // since they are surrounded by whitespace
   function foobar {
    alert('foobar!');
   }
   // this one will need literal escapement
   {literal}
    function bazzy {alert('foobar!');}
   {/literal}
</script>
```

[`{literal}..{/literal}`](../../language-builtin-functions/language-function-literal/) blocks are used for escaping blocks of template logic. You can also escape the braces individually with [`{ldelim}`, `{rdelim}`](../../language-builtin-functions/language-function-ldelim/) tags or [`{$smarty.ldelim}`,`{$smarty.rdelim}`](../../language-variables/language-variables-smarty/#smartyldelim-smartyrdelim-languagevariablessmartyldelim) variables.

Smarty's default delimiters { and } cleanly represent presentational content. However, if another set of delimiters suit your needs better, you can change them with Smarty's `setLeftDelimiter()` and `setRightDelimiter()` methods.

> **Note**
> 
> Changing delimiters affects ALL template syntax and escapement. Be sure to clear out cache and compiled files if you decide to change them.

Where the template is:

```
Welcome <!--{$name}--> to Smarty
    <script>
  var foo = <!--{$foo}-->;
  function dosomething() {
    alert("foo is " + foo);
  }
  dosomething();
</script>
```
