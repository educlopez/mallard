---
source_url: https://elementflow.io/third/posts/js-events-8246.html
ps_version: [8, 9]
ingested: 2026-07-20
tags: [elementflow, docs, feature, js-events]
section: Features
---
# JS events

### 1. How to make the scroll to top button appear only after scrolling 50% of the page.

The scroll to top button is generally added by using the displayBeforeBodyClosingTag hook.

1) Adjust the button for a bit, add a **stsb_display_none** classname to the button. Set it to be the listener in the JS event section, see the first picture below.

2) In the same record, add a container, set it to be at 50% of the page by using the following settings:

- Content Width: Full width
- Width: 1px
- Height: 1px
- Position: Absolute
- Left: 0
- Top: 50%

Set it to be the trigger in the JS event section, see the second picture below.

3) Add this css code to the custom css code field.

```
body{height:auto;position:relative;}
```
