#!/usr/bin/env bash
# Usage: ./encode.sh <frames-dir> <output-without-extension> [poster-frame]
# Outputs <output>.mp4 (og:video), <output>.gif (og:image) and <output>-poster.png (twitter:image), at 1200x630.
# Also use the poster as the static og:image on stacks without a dynamic OG route.
set -euo pipefail
frames="$1"; out="$2"; poster="${3:-0075}"
ffmpeg -v error -y -framerate 30 -i "$frames/%04d.png" -vf "scale=1200:630:flags=lanczos" \
  -an -c:v libx264 -crf 18 -preset slow -pix_fmt yuv420p -movflags +faststart "$out.mp4"
ffmpeg -v error -y -i "$out.mp4" -vf "fps=15,palettegen=max_colors=96:stats_mode=diff" "$out-pal.png"
ffmpeg -v error -y -i "$out.mp4" -i "$out-pal.png" \
  -lavfi "fps=15[x];[x][1:v]paletteuse=dither=sierra2_4a:diff_mode=rectangle" -loop 0 "$out.gif"
rm "$out-pal.png"
ffmpeg -v error -y -i "$frames/$poster.png" -vf "scale=1200:630:flags=lanczos" "$out-poster.png"
wc -c "$out.mp4" "$out.gif" "$out-poster.png"
