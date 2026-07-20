---
source_url: https://elementflow.io/third/posts/svg-6980.html
ps_version: [8, 9]
ingested: 2026-07-20
tags: [elementflow, docs, feature]
section: Features
---
# Svg

To add an SVG icon to a page is as easy as adding an image, but **not all SVG files are ready for PrestaShop**, some are not even ready for the web.

If you have an SVG that causes errors on your site, try exporting it again in a software. Google for a proper tutorial for your software on how to export SVG for the web.

If it still doesn't work, open it up in a text editor to simplify it. Here is an example.

The { and } in the **style** tag is what causes errors, the **style** tag has to be removed. Don't worry, colors can't be set later in the editor.

[![](https://elementflow.io/img/cms/document/1793.jpg)](https://elementflow.io/img/cms/document/1793.jpg)

```
<svg xmlns viewBox>
  <path/>
</svg>
```

Some files have **multiple paths**, and some may need the **style** attribute. Here is the final result in code and file.

[![](https://elementflow.io/img/cms/document/1795.jpg)](https://elementflow.io/img/cms/document/1795.jpg)
