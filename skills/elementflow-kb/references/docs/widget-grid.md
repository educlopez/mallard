---
source_url: https://elementflow.io/third/posts/grid-widget-6816.html
ps_version: [8, 9]
ingested: 2026-07-20
tags: [elementflow, docs, widget]
section: Widgets
---
# Grid widget

The grid widget makes layout easy. Practical ways to use it:

### #1 — To have a fixed-width column

Set the **Columns** setting to a custom value like `200px 1fr` to get a 200px-wide left column. The same result is achievable in the flexbox layout, but it requires several settings to avoid the left column getting squeezed down.

### #2 — Put the logo dead center

Set the **Columns** setting to a custom value like `1fr 200px 1fr` to get a 200px-wide center column.

To center a row of submenus that are each 200px wide:

1. Set **Columns** to `repeat(auto-fit, 200px)`
2. Set **Justify Content** to **Middle**

The advantage of this approach is that the 200px width is set in one go; in the flexbox layout, the width has to be set for every submenu individually.

### For your information

Keep **Rows** set to **1** — the default value of 1 works for 99% of cases. Don't change it unless you know what you are doing.
