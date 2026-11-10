# Link preview tags per stack

Static only: `og:image` points at the static card and there is no `og:video`. Animated: GIF for `og:image`, MP4 for `og:video`, poster for `twitter:image` (X does not animate GIFs or play `og:video`).

## HTML (any stack)

```html
<meta property="og:image" content="https://<site>/social/og-animated.gif">
<meta property="og:image:type" content="image/gif">
<meta property="og:image:width" content="1200">
<meta property="og:image:height" content="630">
<meta property="og:video" content="https://<site>/social/og-animated.mp4">
<meta property="og:video:type" content="video/mp4">
<meta property="og:video:width" content="1200">
<meta property="og:video:height" content="630">
<meta name="twitter:card" content="summary_large_image">
<meta name="twitter:image" content="https://<site>/social/og-animated-poster.png">
```

## Next (App Router metadata)

```ts
openGraph: {
  images: [{ url: '/social/og-animated.gif', width: 1200, height: 630, type: 'image/gif', alt }],
  // Not resolved against metadataBase: must be absolute.
  videos: [{ url: `${siteUrl}/social/og-animated.mp4`, width: 1200, height: 630, type: 'video/mp4' }],
},
twitter: { card: 'summary_large_image', images: ['/social/og-animated-poster.png'] },
```

A page that sets its own `openGraph` replaces the whole object, so it keeps its own image.

## Laravel (Blade)

Put the tags in the main layout's `<head>`, with `asset('social/og-animated.gif')` for the URLs (absolute when `APP_URL` is right). Allow a `@yield('og_image')` override for pages with their own card.

## PrestaShop

Add the tags to the child theme's `templates/_partials/head.tpl` (or a module on `displayHeader`), with `{$urls.theme_assets}img/social/...` for the URLs. In classic-based themes `product.tpl` already sets `og:image` from the product cover: keep it there and use the card for the home and CMS pages.
