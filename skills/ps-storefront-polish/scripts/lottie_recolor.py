#!/usr/bin/env python3
"""Inspect and recolour a Lottie JSON to brand oklch colours.

Usage:
  lottie_recolor.py INPUT.json OUTPUT.json "L,C,H" "L,C,H"
    first oklch = heart / strokes (the dominant colour), second = secondary (burst particles)
  lottie_recolor.py INPUT.json            (inspect only: size, frames, layers, fill and stroke colours)

Rules: a solid colour whose green channel is above 0.40 (a light tint) becomes the secondary colour,
everything else the main one. Adjust the SPLIT constant when a file does not fit.
Reads untrusted files: it only parses JSON, never executes anything. It refuses files with expressions
or external references.
"""
import json
import math
import sys

SPLIT = 0.40


def oklch_to_srgb(l, c, h):
    a = c * math.cos(math.radians(h))
    b = c * math.sin(math.radians(h))
    l_ = (l + 0.3963377774 * a + 0.2158037573 * b) ** 3
    m_ = (l - 0.1055613458 * a - 0.0638541728 * b) ** 3
    s_ = (l - 0.0894841775 * a - 1.2914855480 * b) ** 3
    rgb = (
        4.0767416621 * l_ - 3.3077115913 * m_ + 0.2309699292 * s_,
        -1.2684380046 * l_ + 2.6097574011 * m_ - 0.3413193965 * s_,
        -0.0041960863 * l_ - 0.7034186147 * m_ + 1.7076147010 * s_,
    )
    f = lambda x: 12.92 * x if x <= 0.0031308 else 1.055 * x ** (1 / 2.4) - 0.055
    return [round(max(0.0, min(1.0, f(v))), 4) for v in rgb] + [1]


def walk(node, fn, path=""):
    if isinstance(node, dict):
        fn(node, path)
        for k, v in node.items():
            walk(v, fn, f"{path}/{k}")
    elif isinstance(node, list):
        for i, v in enumerate(node):
            walk(v, fn, f"{path}/{i}")


def main():
    if len(sys.argv) < 2:
        sys.exit(__doc__)
    data = json.load(open(sys.argv[1]))
    raw = json.dumps(data)
    if '"ef"' in raw or "http" in raw or '"x":"' in raw:
        sys.exit("refused: effects or external references found, inspect by hand")
    print("size", data["w"], "x", data["h"], "fps", data["fr"], "frames", data["ip"], "-", data["op"])
    for layer in data["layers"]:
        print("layer", layer.get("nm"), "ip", layer.get("ip"), "op", layer.get("op"))
    found = []

    def collect(node, path):
        if node.get("ty") in ("fl", "st") and "c" in node and not node["c"].get("a"):
            found.append((node["ty"], path, node["c"]["k"]))

    walk(data, collect)
    for ty, path, k in found:
        print(ty, path, k)
    if len(sys.argv) < 5:
        return
    main_c = oklch_to_srgb(*map(float, sys.argv[3].split(",")))
    second_c = oklch_to_srgb(*map(float, sys.argv[4].split(",")))

    def paint(node, path):
        if node.get("ty") in ("fl", "st") and "c" in node and not node["c"].get("a"):
            node["c"]["k"] = second_c if node["c"]["k"][1] > SPLIT else main_c

    walk(data, paint)
    json.dump(data, open(sys.argv[2], "w"), separators=(",", ":"))
    print("written", sys.argv[2], "main", [round(v * 255) for v in main_c[:3]], "second", [round(v * 255) for v in second_c[:3]])


main()
