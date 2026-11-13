#!/usr/bin/env python3
"""
Genera los favicons de Animal Max por entorno (patrón /env-favicons adaptado a PrestaShop).

  uv run --with pillow python themes/animalmax-child/scripts/gen-favicons.py [logo-completo.png]

- Con un logo completo (AM + texto) recorta la marca «AM» y la guarda en _dev/brand/am-mark.png.
- Sin argumentos usa el _dev/brand/am-mark.png ya guardado.
- El color del entorno va en el FONDO de la insignia; la marca siempre es blanca (se lee mejor a 16 px).

Salida (assets/img/favicons/):
  favicon-production.ico  → para SUBIR en el BO (Diseño › Tema y logotipo › Favicon) en PRE y PROD
  favicon-pre.png         → lo pinta head.tpl cuando el host es *.pre-testing.com
  favicon-local.png       → lo pinta head.tpl cuando el host es *.lndo.site / localhost
"""
import sys
from pathlib import Path

from PIL import Image, ImageDraw

ROOT = Path(__file__).resolve().parent.parent
MARK = ROOT / '_dev' / 'brand' / 'am-mark.png'
OUT = ROOT / 'assets' / 'img' / 'favicons'

# Tonos bien separados: producción (marca), PRE (dorado oscuro), local (verde).
COLORS = {
    'production': (14, 25, 38),    # --ink  #0e1926
    'pre': (153, 107, 45),         # --gold-dark #996b2d
    'local': (30, 142, 78),        # #1e8e4e
}


def extract_mark(src):
    """Recorta la marca AM (a la izquierda del texto) de un logo completo, en RGBA negro/alpha."""
    im = Image.open(src).convert('RGBA')
    px = im.load()
    w, h = im.size
    has_alpha = any(px[x, y][3] < 250 for x in range(0, w, 7) for y in range(0, h, 7))

    def ink(x, y):
        r, g, b, a = px[x, y]
        return a if has_alpha else int((255 - (r + g + b) / 3) * (a / 255))

    cols = [any(ink(x, y) > 128 for y in range(h)) for x in range(w)]
    # primer hueco de columnas vacías ancho tras la marca = separación con el texto
    runs, start = [], None
    for x, c in enumerate(cols + [False]):
        if c and start is None:
            start = x
        if not c and start is not None:
            runs.append((start, x))
            start = None
    # A y M van pegadas (hueco pequeño); el texto «ANIMAL MAX» queda tras un hueco mayor.
    # La marca es el primer grupo ancho de columnas con tinta (el ® flota sobre la A, sin columnas propias).
    groups = []
    for r in runs:
        if groups and r[0] - groups[-1][1] < h * 0.2:
            groups[-1] = (groups[-1][0], r[1])
        else:
            groups.append(r)
    x0, x1 = next(g for g in groups if g[1] - g[0] > h * 0.5)
    mask = Image.new('L', (x1 - x0, h))
    mp = mask.load()
    for x in range(x0, x1):
        for y in range(h):
            mp[x - x0, y] = min(255, max(0, ink(x, y)))
    # quitar el ® y cualquier mota: se conservan solo los componentes conexos grandes (la A y la M)
    mp = mask.load()
    mw, mh = mask.size
    seen = [[False] * mh for _ in range(mw)]
    comps = []
    for sx in range(mw):
        for sy in range(mh):
            if mp[sx, sy] > 128 and not seen[sx][sy]:
                stack, pts = [(sx, sy)], []
                seen[sx][sy] = True
                while stack:
                    cx, cy = stack.pop()
                    pts.append((cx, cy))
                    for nx, ny in ((cx + 1, cy), (cx - 1, cy), (cx, cy + 1), (cx, cy - 1)):
                        if 0 <= nx < mw and 0 <= ny < mh and not seen[nx][ny] and mp[nx, ny] > 128:
                            seen[nx][ny] = True
                            stack.append((nx, ny))
                comps.append(pts)
    biggest = max(len(c) for c in comps)
    near = set()  # píxeles a ≤2 px de un componente grande (conserva el antialiasing del borde)
    for c in comps:
        if len(c) >= biggest * 0.2:
            for (cx, cy) in c:
                for dx in range(-2, 3):
                    for dy in range(-2, 3):
                        near.add((cx + dx, cy + dy))
    for x in range(mw):
        for y in range(mh):
            if (x, y) not in near:
                mp[x, y] = 0
    # recorte ajustado
    bbox = mask.point(lambda v: 255 if v > 128 else 0).getbbox()
    mask = mask.crop(bbox)
    mark = Image.new('RGBA', mask.size, (0, 0, 0, 255))
    mark.putalpha(mask)
    MARK.parent.mkdir(parents=True, exist_ok=True)
    mark.save(MARK)
    print('marca extraída:', MARK.relative_to(ROOT), mark.size)


def badge(bg, size, glyph_w=0.74):
    """Cuadrado redondeado de color + marca blanca centrada (la marca es apaisada: se ajusta por ancho)."""
    mark = Image.open(MARK).convert('RGBA')
    mw, mh = mark.size
    side = 1024
    img = Image.new('RGBA', (side, side), (0, 0, 0, 0))
    ImageDraw.Draw(img).rounded_rectangle((0, 0, side, side), radius=int(side * 0.22), fill=bg + (255,))
    gw = int(side * glyph_w)
    gh = int(mh * gw / mw)
    white = Image.new('RGBA', (gw, gh), (255, 255, 255, 255))
    white.putalpha(mark.split()[3].resize((gw, gh), Image.LANCZOS))
    img.paste(white, ((side - gw) // 2, (side - gh) // 2), white)
    return img.resize((size, size), Image.LANCZOS)


def main():
    if len(sys.argv) > 1:
        extract_mark(sys.argv[1])
    if not MARK.exists():
        sys.exit('Falta _dev/brand/am-mark.png: pasa el logo completo como argumento.')
    OUT.mkdir(parents=True, exist_ok=True)
    badge(COLORS['pre'], 64).save(OUT / 'favicon-pre.png')
    badge(COLORS['local'], 64).save(OUT / 'favicon-local.png')
    # .ico multi-tamaño para el favicon de PrestaShop (el que se sube en el BO)
    ico = badge(COLORS['production'], 256)
    ico.save(OUT / 'favicon-production.ico', sizes=[(16, 16), (24, 24), (32, 32), (48, 48), (64, 64)])
    for f in sorted(OUT.iterdir()):
        print('generado:', f.relative_to(ROOT), f.stat().st_size, 'B')


if __name__ == '__main__':
    main()
