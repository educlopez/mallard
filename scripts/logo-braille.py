#!/usr/bin/env python3
"""Print an SVG as braille art, for the TUI logo in internal/tui/logo.go.

Usage: python3 scripts/logo-braille.py <file.svg> [columns=24] [threshold=140]
Needs rsvg-convert (brew install librsvg) and Pillow.
"""
import io
import subprocess
import sys

from PIL import Image

svg = sys.argv[1]
cols = int(sys.argv[2]) if len(sys.argv) > 2 else 24
thr = int(sys.argv[3]) if len(sys.argv) > 3 else 140
W = cols*2
png = subprocess.run(['rsvg-convert','-w',str(W*8),svg],capture_output=True,check=True,timeout=30).stdout
im = Image.open(io.BytesIO(png)).convert('RGBA')
h = round(im.height*W/im.width)
H = (h+3)//4*4
a = im.getchannel('A').resize((W,h), Image.LANCZOS)
canvas = Image.new('L',(W,H),0); canvas.paste(a,(0,(H-h)//2))
px = canvas.load()
bits=[(0,0,0x01),(0,1,0x02),(0,2,0x04),(1,0,0x08),(1,1,0x10),(1,2,0x20),(0,3,0x40),(1,3,0x80)]
for r in range(H//4):
    line=''
    for c in range(cols):
        v=0
        for dx,dy,b in bits:
            if px[c*2+dx, r*4+dy] >= thr: v|=b
        line+=chr(0x2800+v)
    print(line)
