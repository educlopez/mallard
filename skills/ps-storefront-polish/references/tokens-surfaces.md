# Tokens, colour, surfaces

Source of truth in the reference theme: `assets/css/custom.css` `:root`
(first 70 lines) and the README section "Design tokens".

## Token block (copy and rename the prefix)

```css
:root{
  /* brand and text. oklch: lightness is perceptual, so contrast repairs only touch L */
  --pies-primary:oklch(0.357 0.03 223.632);
  --pies-primary-hover:oklch(0.297 0.026 225.531);
  --pies-primary-tint:oklch(0.963 0.006 197.05);
  --pies-text:oklch(0.262 0.009 248.19);
  --pies-text-body:oklch(0.44 0.03 223.6);
  --pies-text-muted:oklch(0.52 0.025 223.6);   /* 4.9:1 on white */
  --pies-price:oklch(0.54 0.06 194.3);         /* was a light teal at 1.86:1, darkened for AA */
  --pies-danger:oklch(0.5 0.182 29.513);
  --pies-danger-bg:oklch(0.972 0.012 23.703);
  --pies-danger-border:oklch(0.801 0.109 24.573);
  /* neutrals */
  --pies-border:oklch(0.843 0.01 222.075);      /* control borders */
  --pies-border-soft:oklch(0.924 0.005 106.504);/* dividers */
  --pies-bg-soft:oklch(0.976 0 89.876);         /* panels, hover */
  /* ONE radius scale */
  --pies-radius:4px;        /* controls: buttons, inputs, steppers, chips, thumbs */
  --pies-radius-lg:8px;     /* surfaces: cards, panels, drawers, modals */
  --pies-radius-pill:999px;
  /* control heights shared by buttons, inputs and steppers */
  --pies-btn-h-sm:36px; --pies-btn-h:44px; --pies-btn-h-lg:50px;
  /* surfaces */
  --pies-ring:0 0 0 1px oklch(0 0 0 / .06);
  --pies-elevation-raised:0 0 0 1px oklch(0 0 0 / .06),0 1px 3px oklch(0.357 0.03 223.632 / .06),0 4px 10px -4px oklch(0.357 0.03 223.632 / .06);
  --pies-elevation-floating:0 0 0 1px oklch(0 0 0 / .06),0 2px 6px oklch(0.357 0.03 223.632 / .06),0 12px 28px -6px oklch(0.357 0.03 223.632 / .16);
  --pies-hairline:1px;
  --pies-focus-ring:0 0 0 3px oklch(0.357 0.03 223.632 / .25);
  /* motion */
  --pies-ease-out:cubic-bezier(.22,1,.36,1);
  /* stacking scale, see header-layers.md */
  --pies-z-nav:1025; --pies-z-sticky-bar:1030; --pies-z-mobile-menu:1080; --pies-z-cookie:8005;

  /* feed Bootstrap so every stock component follows the tokens */
  --bs-primary:var(--pies-primary);
  --bs-border-radius:var(--pies-radius);
  --bs-border-radius-sm:var(--pies-radius);
  --bs-border-radius-lg:var(--pies-radius);
}
@media (min-resolution:192dpi){:root{--pies-hairline:.5px}}
```

Leave `--bs-primary-rgb` as comma-separated RGB (Bootstrap needs it for
`rgba(var(--bs-primary-rgb), .5)`), pure `#fff`/`#000`, and colours inside
SVG `url()` data URIs: they cannot read variables.

## Buttons: one system

Use Bootstrap `.btn`, `.btn-primary`, `.btn-outline-primary`, `.btn-sm`, `.btn-lg`
and set the sizes once:

```css
.btn{display:inline-flex;align-items:center;justify-content:center;gap:.5rem;min-height:var(--pies-btn-h);padding:0 1.25rem;border-radius:var(--pies-radius);font-size:.9375rem;font-weight:600;line-height:1.25;text-transform:none}
.btn-lg{min-height:var(--pies-btn-h-lg);padding:0 1.5rem;font-size:1rem}
.btn-sm{min-height:var(--pies-btn-h-sm);padding:0 .875rem;font-size:.8125rem}
.btn-primary{--bs-btn-bg:var(--pies-primary);--bs-btn-border-color:var(--pies-primary);--bs-btn-color:#fff;
  --bs-btn-hover-bg:var(--pies-sage);--bs-btn-hover-border-color:var(--pies-sage);--bs-btn-hover-color:var(--pies-primary)}
```

Module markup that cannot take classes (e.g. `idxcookies` `#idxrcookiesOK`) reads
the same variables. If buttons look different between sections, the cause is a
local rule with its own radius or height: delete it.

## Surfaces

Principles (shadows over borders, nested radii, hairlines, eased gradients,
masks, image outlines, dark mode): see `emil-surfaces`. Project choices:

- Raised card: `box-shadow:var(--pies-elevation-raised)`. Floating (menus, drawers,
  cookie card): `--pies-elevation-floating`. Grey panel (cart and checkout summary):
  `--pies-bg-soft` plus `--pies-ring`. Divider: `var(--pies-hairline) solid var(--pies-border-soft)`.
- Photos on a grey product tile get no outline (the tile frames them). Photos on
  white (category tiles, menu tiles) get `outline:1px solid oklch(0 0 0 / .1);outline-offset:-1px`.
- Category tile scrim: 8-stop eased gradient, alpha .5 to 0 over 45% of the height.
  Clamped text fade: `mask-image` eased gradient, never over scrollable content.
- No dark mode in the reference site.

## Colour

Method, contrast and gamut rules: see `emil-color`. Project numbers: text on
white at lightness 0.54 or less (4.9:1 at 0.52 for muted text). The old light
teal price was 1.86:1, darkened to `--pies-price`. Sage is fill only, never text.
