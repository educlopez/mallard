---
name: og-preview
description: >
  Builds a site's link preview: the static og:image and, optionally, an animated
  one (GIF + MP4 + poster) from the site's real visuals. Next, Laravel or
  PrestaShop. Use when the user says "haz la OG", "OG animada" or "og:image",
  or when a new web project is created.
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# og-preview

Makes the image a link shows when it is shared (X, LinkedIn, Slack, WhatsApp, iMessage). It is not a social post video and not a favicon.

## 0. New project: ask first

When a new web project is being created, ask once: "¿Le hacemos OG estática, y también animada?" Static is the default for every site. Build the animated one only when the user says yes.

## 1. Read the site's real visuals

1. Open the live site (or the dev server) and take a full-page screenshot. Headless Chrome without a GPU can crash pages that use WebGL: launch with `--enable-unsafe-swiftshader --use-angle=swiftshader`.
2. Collect, from the code, not from the screenshot:
   - colors and fonts from the CSS tokens and computed styles, never the average color of the logo;
   - the backgrounds that sit under the hero and the illustrations (gradients, dot grids, blurred scenes): copy their CSS as-is;
   - logos and icons from `public/` or the components. Never draw or invent a logo; a missing one is asked for;
   - real product content: screenshots, cards, commands with their real output, names from the README.
3. Check every claim on the card against the README or the page copy (`og:title`, `description`).

Done when you can list, with its source file, the background, the colors, the font, the logo and the content the card uses.

Taste rules that came from rejected drafts: the site's own background beats a saturated full-bleed brand color; the brand color goes on one accent (a pill, a word, a line). Flat, no glow, no black filler.

## 2. Static og:image (always)

| Stack | How |
|---|---|
| Next (App Router) | `ImageResponse` route with `dynamic = 'force-static'`. Satori reads TTF/OTF only: commit the font files. Satori has no CSS blur: pre-render blurred backgrounds to a JPG and embed it as a data URL. Shared renderer for home and per-page cards. |
| Laravel, PrestaShop, any other | Build the card as HTML (step 3), render one still with `scripts/stills.mjs`, save it as `og-image.png` (1200x630, under 600 KB for WhatsApp). |

Done when the image returns 200 locally and you have looked at it at full size.

## 3. Animated preview (only when asked)

1. Write `index.html` at 1200x630: the CSS is the rest state. `anim.js` builds every animation with the Web Animations API, pauses them and exposes `window.DURATION` and `window.seek(ms)`, so each frame is deterministic. 5 to 6 s; frame 0 must equal the last frame.
2. Render and encode:
   ```bash
   cd <skill>/scripts && npm i
   node stills.mjs <page>/index.html /tmp/og 0 2000 4000   # check stills first
   node frames.mjs <page>/index.html <out>/frames
   ./encode.sh <out>/frames <out>/og-animated 0120          # poster = frame 120
   ```
3. Check the loop: compare the first and last frames (a few hundred antialiasing pixels on rotated layers is acceptable). Check stills at full size, by zone.
4. Weight: photos moving all the time make huge GIFs (12.5 MB measured). Move in steps with rests; aim for a GIF under 3 MB.

Done when the MP4, GIF and poster exist and the user has seen them (`open <file>`).

## 4. Serve and wire the metadata

1. Put the files in `public/social/` (Next, Laravel) or the theme's `assets/img/social/` (PrestaShop).
2. Read `robots.txt` (and the code that generates it). Any `Disallow` covering the image paths blocks X, LinkedIn and Slack, which honour robots. Check the home and one inner page of each section, not only the home.
3. Add the tags from [references/metas.md](references/metas.md) for the stack. `og:video` must be an absolute URL (Next does not resolve videos against `metadataBase`).

Done when the dev server serves the tags and every file returns 200 with the right `content-type`.

## 5. After deploy

```bash
curl -sL https://<site> | grep -oE '<meta[^>]+(og:image|og:video|twitter:image)[^>]+>'
for f in og-animated.gif og-animated.mp4 og-animated-poster.png; do curl -sI https://<site>/social/$f | head -1; done
```

Done when production serves the tags and the three files return 200. Platforms cache previews: an old card after a deploy is their cache.

Does not commit, open PRs, merge or deploy without the user asking.
