---
source_url: https://elementflow.io/third/posts/text-editor-widget-8169.html
ps_version: [8, 9]
ingested: 2026-07-20
tags: [elementflow, docs, widget]
section: Widgets
---
# Text editor widget

### Why is the gap below the text widget larger.

The text editor wraps the content in a pair of **p** tags; p tags have a 1rem (16px) bottom margin by default. Two solutions to remove the gap:

1. Go to the **Style** tab of the text editor, find the **Paragraph** section, and set the bottom value to 0.
2. Another solution is to change the **p** tags to **div** tags.

Screenshots: [1](https://elementflow.io/img/cms/document/2220.jpg), [2](https://elementflow.io/img/cms/document/2221.jpg)
