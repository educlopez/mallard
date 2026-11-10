import { chromium } from 'playwright-core';
import { pathToFileURL } from 'node:url';
const [file, out, ...ts] = process.argv.slice(2);
const b = await chromium.launch({ channel: 'chrome' });
const p = await b.newPage({ viewport: { width: 1200, height: 630 }, deviceScaleFactor: 1 });
await p.goto(pathToFileURL(file).href, { waitUntil: 'networkidle' });
await p.waitForTimeout(500);
for (const t of ts) { await p.evaluate(t => window.seek(t), +t); await p.screenshot({ path: `${out}-${t}.png` }); }
await b.close();
