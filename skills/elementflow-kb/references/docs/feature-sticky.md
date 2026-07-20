---
source_url: https://elementflow.io/third/posts/sticky-6789.html
ps_version: [8, 9]
ingested: 2026-07-20
tags: [elementflow, docs, feature]
section: Features
---
# Sticky

Here are some usages of the sticky feature.

### #1 - JS sticky

[Click here to open the sticky header in the editor](https://elementflow.io/bo/index.php?controller=AdminStSiteBuilderEditor&id_lang=1&id=310&ffo=1&id_shop=3)

Go to the **Advanced** tab of every widget, and scroll down to see the **Sticky** section.

Pages

For example, if you use the header for several pages, but you just like to make it sticky on the homepage, then select the homepage here.

Sticky to

This setting is what makes the JS sticky feature powerful. It has 3 options:

1. Body
2. Parent
3. Custom

**Body:** Select this to make an element sticky to the top of the browser window when it's scrolled to the top of the browser window no matter where the element is.

**Parent:** This is pretty much the same as Position: Sticky, but here you've more options. Check the **Position: Sticky** section below for more info.

**Custom:** For example, on the category page, you have a banner in the left column, you like to set the banner only (not the entire left column) to be sticky relative to the right column, then use this option.

Show

The **Hide when the page is scrolling down** is perfect for a sticky header. Hide the sticky header when the page is scrolling down to save spacing because when people keep scrolling down, they tend to keep scrolling to see more info, when people scroll the page up, they probably want to scroll up to the menu, it's the perfectly to show them the sticky header. This behavior is more and more popular.

[![](https://elementflow.io/second/img/cms/document/1724.jpg)](https://elementflow.io/second/img/cms/document/1724.jpg)

### #2 - Position: Sticky

[Click here to open the sticky production info section in the editor](https://elementflow.io/bo/index.php?controller=AdminStSiteBuilderEditor&id_lang=1&id=605&ffo=1&id_shop=3)

Go to the **Advanced** tab of every widget, and find the position setting in the **Layout** section.

Position

Set the position to Sticky, you can set different values for different pages.

Setting the Position to sticky won't make it sticky. **A placement position** has to be set to it, this can be any of the left, right, bottom, or top.

If the sticky still doesn't work, check containers wrapping the current element, if any of the parent containers using **overflow** with a value of **Hidden**.

NOTE a sticky element's movement is limited to the boundaries of its parent container.

Try to scroll inside this box to understand how sticky positioning works.

[![](https://elementflow.io/second/img/cms/document/1726.jpg)](https://elementflow.io/second/img/cms/document/1726.jpg)

### #3 - Position: Fixed

This is pretty similar to **Position:Sticky**. Technically, it's not sticky, it's fixed, but the result is kind of similar. The difference is that it doesn't require the scrolling of a page.
