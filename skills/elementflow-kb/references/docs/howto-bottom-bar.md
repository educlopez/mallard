---
source_url: https://elementflow.io/third/posts/how-to-create-a-bottom-bar-6877.html
ps_version: [8, 9]
ingested: 2026-07-20
tags: [elementflow, docs, howto, bottom-bar]
section: Features
---
# How to create a bottom bar

[Check this demo](https://elementflow.io/en/?bridge_device=mp&link=https://elementflow.io/en/pages/womenswear-38.html), it has a bottom bar, how to create one?

1. Go to BO > ELement Flow > Content Builder > Hooks.
2. Add a new record, and select the **displayBeforeBodyClosingTag** hook., this is important because whatever is added to the displayBeforeBodyClosingTag hook is a direct child of the body tag, which makes sure that the added content can be placed at the very bottom of the page.
3. Add a container to the canvas, and set its **Position** to Fixed, set **bottom, left, and right** to 0, the very last step is to set a high z-index to make sure the bottom is on top of other elements.
