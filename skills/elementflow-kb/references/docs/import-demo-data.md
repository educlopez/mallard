---
source_url: https://elementflow.io/third/posts/import-demo-data-6756.html
ps_version: [8, 9]
ingested: 2026-07-20
tags: [elementflow, docs, getting-started]
section: Getting started
---
# Import demo data

There are **4 ways** to import data from the demo sites to your site.

### #1 - Use the Import feature

1. Go to BO > Improve > Element Flow > Settings > Import & export.
2. Hit the **Import** button.
3. Select the **Element Flow demo site** from the **Select a store** dropdown menu.
4. Click the **Fetch data** button.
5. You will get a list of demo sites and pages — select one option.
6. You will get checkboxes of what is included in the selected option.
7. Check any, some, or all to import.
8. When you create a header, it will not show on the front office right away — you need to assign the header to the pages where you'd like it shown (one site can have several different headers). The same applies when you import a header, footer, or page template: if you don't enable the **Enable the imported records** setting, the records will be imported but not assigned for use on the front end. Enabling that setting makes the import feature try to assign the imported records for use.

One nice thing about this way is that you can keep track of what you've imported. When you see a record, you know if it's imported and when, and you can easily delete any unused imported record.

### #2 - Library

Click on the "Folder" icon to open the library in the editor. From there you can import blocks and pages.

### #3 - Copy from the demo site front office

1. Click on the black **Editor** button on the right side of every page.
2. **Copy** and **Editor** buttons will show up; click the copy button of the section you like.
3. Go to the editor on your site and right-click where you'd like to insert the copied data.
4. Select **Paste from another site** and press ctrl/cmd + v. The copied data will soon be added to your site.

The **Paste from another site** feature may not work from some nested widgets, like the tab widget and accordion widget.

### #4 - Copy from the editor

Very similar to the above, but lets you select the exact things to copy (e.g. just a button).

1. Click on the black **Editor** button on the right side of every page.
2. **Copy** and **Editor** buttons will show up; click the editor button of the section you like to copy from.
3. The section opens in the editor — right-click to copy anything you like.
4. Go to the editor of your site, right-click where you'd like to insert the copied data, select **Paste from another site**, and press ctrl/cmd + v. The copied data will soon be added to your site.

The **Paste from another site** feature may not work from some nested widgets, like the tab widget and accordion widget.

### FAQs

**Why does the imported content not look like the demo?**

1. Clear the caches completely — see [FAQ #10](https://elementflow.io/third/posts/faqs-2805.html).
2. The imported image type may not exist on your site. Correct the image types in all **miniature**, **product template**, and **category template** records.
3. The selected category in **Products** widgets doesn't exist on your site — set an existing category for all Products widgets, and the same for **Categories** widgets and **Blog posts** widgets.
4. The **text content** is in English only on the demos. If your site isn't in English, all text fields will be empty after import; you need to fill them with content, otherwise some widgets may not show because they are empty.

**Is my existing data safe?**

Yes — the import feature adds new data to your site; it doesn't delete or modify existing data.

**Quite several global colors got added**

That's expected — if a global color in the selected data doesn't exist on your site, a new global color will be added. The same color is only added once.

**The import feature doesn't work**

Two possible reasons:

1. Temporary network interruption — try again later.
2. Some widgets in the selected records got disabled on your site. Go to BO > Improve > Element Flow > Settings > Widgets manager to enable all widgets, then try importing again.
