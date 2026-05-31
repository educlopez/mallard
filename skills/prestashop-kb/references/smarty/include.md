---
source_url: https://smarty-php.github.io/smarty/stable/designers/language-builtin-functions/language-function-include/
topic: smarty
subtopic: includes
ingested: 2026-04-18
tags: [smarty, prestashop, includes]
ps_version: [8, 9]
---

## {include}

`{include}` tags are used for including other templates in the current template. Any variables available in the current template are also available within the included template.

## Attributes

| Attribute Name | Required | Description |
| --- | --- | --- |
| file | Yes | The name of the template file to include |
| assign | No | The name of the variable that the output of include will be assigned to |
| cache\_lifetime | No | Enable caching of this subtemplate with an individual cache lifetime |
| compile\_id | No | Compile this subtemplate with an individual compile\_id |
| cache\_id | No | Enable caching of this subtemplate with an individual cache\_id |
| scope | No | Define the scope of all in the subtemplate assigned variables: 'parent','root' or 'global' |
| \[var...\] | No | variable to pass local to template |

- The `{include}` tag must have the `file` attribute which contains the template resource path.
- Setting the optional `assign` attribute specifies the template variable that the output of `{include}` is assigned to, instead of being displayed. Similar to [`{assign}`](../language-function-assign/).
- Variables can be passed to included templates as [attributes](../../language-basic-syntax/language-syntax-attributes/). Any variables explicitly passed to an included template are only available within the scope of the included file. Attribute variables override current template variables, in the case when they are named the same.
- You can use all variables from the including template inside the included template. But changes to variables or new created variables inside the included template have local scope and are not visible inside the including template after the `{include}` statement. This default behaviour can be changed for all variables assigned in the included template by using the scope attribute at the `{include}` statement or for individual variables by using the scope attribute at the [`{assign}`](../language-function-assign/) statement. The later is useful to return values from the included template to the including template.
- Use the syntax for [template resources](../../../api/resources/) to `{include}` files outside of the [`$template_dir`](../../../programmers/api-variables/variable-template-dir/) directory.

## Option Flags

| Name | Description |
| --- | --- |
| nocache | Disables caching of this subtemplate |
| caching | Enable caching of this subtemplate |
| inline \* | If set, merge the compile-code of the subtemplate into the compiled calling template |

\* The `inline` option flag is currently not implemented in Smarty v5. Using it will not trigger an error, however.

## Examples

The template above includes the example `links.tpl` below

Variables assigned in the included template will be seen in the including template.
```
{include 'sub_template.tpl' scope=parent}
...
{* display variables assigned in sub_template *}
{$foo}<br>
{$bar}<br>
...
```

The template above includes the example `sub_template.tpl` below

```
...
{assign var=foo value='something'}
{assign var=bar value='value'}
...
```

The included template will not be cached.

```
{include 'sub_template.tpl' nocache}
...
```

In this example included template will be cached with an individual cache lifetime of 500 seconds.

```
{include 'sub_template.tpl' cache_lifetime=500}
...
```

In this example included template will be cached independent of the global caching setting.

```
{include 'sub_template.tpl' caching}
...
```

This example assigns the contents of `nav.tpl` to the `$navbar` variable, which is then output at both the top and bottom of the page.

This example includes another template relative to the directory of the current template.

```
{include 'template-in-a-template_dir-directory.tpl'}
{include './template-in-same-directory.tpl'}
{include '../template-in-parent-directory.tpl'}
```

See also [template resources](../../../api/resources/) and [componentized templates](../../../appendixes/tips/#componentized-templates).
