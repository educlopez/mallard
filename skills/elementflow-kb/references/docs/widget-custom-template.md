---
source_url: https://elementflow.io/third/posts/custom-template-widget-7227.html
ps_version: [8, 9]
ingested: 2026-07-20
tags: [elementflow, docs, widget]
section: Widgets
---
# Custom template widget

The custom template widget is powerful for customization — it lets you drop arbitrary Smarty template code into the page. Example scenario: a module's documentation tells you to use a specific snippet of code to display its content at a custom location. To do this in Element Flow:

1. Go to the `/modules/stsitebuilder/views/templates/front/custom_templates/` folder via FTP.
2. Find the `empty_template.tpl` file and make a copy of it to create a new file; rename it and put the code in it. **NOTE**: `empty_template.tpl` isn't actually empty — it contains a license notice which can be removed or kept. If that file doesn't exist, use any `.tpl` file already in the folder as a base.
3. Go back to the editor, **reload it**, and select the new file in the custom template widget.
