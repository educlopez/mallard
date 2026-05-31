---
source_url: https://smarty-php.github.io/smarty/stable/designers/language-builtin-functions/language-function-extends/
topic: smarty
subtopic: includes
ingested: 2026-04-18
tags: [smarty, prestashop, includes]
ps_version: [8, 9]
---

## {extends}

`{extends}` tags are used in child templates in template inheritance for extending parent templates. For details see section of [Template Inheritance](../../../api/inheritance/).

- The `{extends}` tag must be on the first line of the template.
- If a child template extends a parent template with the `{extends}` tag it may contain only `{block}` tags. Any other template content is ignored.
- Use the syntax for [template resources](../../../api/resources/) to extend files outside the [`$template_dir`](../../../programmers/api-variables/variable-template-dir/) directory.

## Attributes

| Attribute | Required | Description |
| --- | --- | --- |
| file | Yes | The name of the template file which is extended |

> **Note**
> 
> When extending a variable parent like `{extends file=$parent_file}`, make sure you include `$parent_file` in the [`$compile_id`](../../../programmers/api-variables/variable-compile-id/). Otherwise, Smarty cannot distinguish between different `$parent_file` s.

## Examples

```
{extends file='parent.tpl'}
{extends 'parent.tpl'}  {* short-hand *}
```

See also [Template Inheritance](../../../api/inheritance/) and [`{block}`](../language-function-block/).
