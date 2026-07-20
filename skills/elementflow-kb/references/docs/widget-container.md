---
source_url: https://elementflow.io/third/posts/flex-container-7066.html
ps_version: [8, 9]
ingested: 2026-07-20
tags: [elementflow, docs, widget]
section: Widgets
---
# Container widget

The flexbox container is the most used widget. The flexbox is the fundamental base of modern web layout — this page shows practical examples of how to use the flexbox container rather than flexbox theory.

### Don't worry about nesting too many containers

Put one container in another container to nest them for the layout, for whatever purpose. You may start questioning "Am I doing right?" after nesting several containers. The answer is YES — nest as many containers as needed.

For an online shop owner, don't think like a developer, forget all the technical rules you know, and forget all the test reports you get; the only thing that matters is whether your customers would like it.

### Flexbox layout VS Sections & Columns layout

The Sections & Columns layout is **deprecated**; what it can do is limited, don't use it anymore.

One major difference is that you have to set a width for every column in the Sections & Columns layout. In the flexbox layout, you barely need to set a width for any container.

### How to put one item on one side, and the other on the other side

In the Sections & Columns layout, you would need to add two columns, put one item in each column, and set the second column's alignment to the right.

It's way easier to achieve the same result: put two items in the same container, and set its **Justify content** to **Space between**.

### How to put some items on one side, the rest on the other side

Two ways to achieve that:

1. Add one container, set its **Direction** to **Row** and **Justify content** to **Space between**, then add two child containers, set their **Widths** to **auto**, and add widgets to the child containers. — **Recommended**
2. Put all widgets in the same container, add a Spacer in the middle to separate them, and set the Spacer's **Size** to **Grow**. Set the divider's **Space** to 0 if needed, and **hide** the spacer on mobile when it's not needed.

### Sort elements

Elements can be sorted by using the **Order** setting — they don't have to stay in the order they were added. Useful for reordering elements differently across devices (e.g. mobile vs desktop).

### Create a full-height panel with the middle part scrollable

Key points for the structure:

1. Set the parent container's height to **100vh** to make the container full-height.
2. Set the scrollable container's **Overflow** to **Auto** to give it a scrollbar, and set its **Size** to **1 1 0** so it takes all empty space and pushes the footer part down to the bottom.
