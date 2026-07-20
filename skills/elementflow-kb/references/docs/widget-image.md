---
source_url: https://elementflow.io/third/posts/image-widget-6870.html
ps_version: [8, 9]
ingested: 2026-07-20
tags: [elementflow, docs, widget]
section: Widgets
---
# Image widget

The image widget is straightforward, but has a few less obvious features:

### #1 — How to add an external image

The widget supports pointing to an externally-hosted image, not just images uploaded to PrestaShop.

### #2 — How to nicely show images that are in different sizes

When mixing a landscape and a portrait image side by side, use the **Object Fit** setting to make them look consistent. The **Object Fit** setting only shows once a height is set on the image — in most cases select the **Cover** option. Note that a large width value is sometimes required to make sure the images show at the same width.

### #3 — Different image sources for various screen sizes

You can upload **multiple** sources for a single image so users always receive the most appropriate image for their device's screen size. Use the **Mobile button** on the top bar to set a different image (e.g. a different crop, or a portrait version) for mobile.

- The desktop image can remain the same, just resized down, or
- A different image (e.g. portrait) can be loaded specifically for mobile, which often looks better there.

This improves page load times and UX — only the most appropriate image gets loaded by the browser; the others are not downloaded.
