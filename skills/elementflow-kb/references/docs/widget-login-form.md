---
source_url: https://elementflow.io/third/posts/login-form-6961.html
ps_version: [8, 9]
ingested: 2026-07-20
tags: [elementflow, docs, widget]
section: Widgets
---
# Login form

It's out of the question that these 3 fields are required:

1. Email
2. Password
3. Submit button

### Where will users be redirected to after logging in?

The key point is the **back** parameter in the URL, e.g.:

`https://example.com/en/login?back=index`
`https://example.com/en/login?back=https%3A%2F%2Fexample.com%2Fen%2F14-coats`

The value of the `back` parameter can be a page name or a page URL; users are redirected according to it.

How the `back` parameter gets set:

1. If the **Sign in** button is added via the **Sign in** widget, by default the `back` parameter is set to the URL of the current page. To avoid this, don't leave the URL field empty — set it to the Sign-in page itself; in that case the `back` parameter won't be added automatically. See the **Sign in** widget page for more info.
2. In the Login form widget, there's a **Fallback redirection** field: if the `back` parameter doesn't exist in the URL, the Fallback redirection takes effect. If neither the `back` parameter nor a fallback redirection is set, the homepage is used.
