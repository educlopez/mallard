import { chromium } from 'playwright-core';
import { pathToFileURL } from 'node:url';
import { mkdirSync } from 'node:fs';
const [file, out, fps = '30'] = process.argv.slice(2);
mkdirSync(out, { recursive: true });
const b = await chromium.launch({ channel: 'chrome' });
const p = await b.newPage({ viewport: { width: 1200, height: 630 }, deviceScaleFactor: 2 });
await p.goto(pathToFileURL(file).href, { waitUntil: 'networkidle' });
await p.waitForTimeout(500);
const D = await p.evaluate(() => window.DURATION);
const n = Math.round(D / 1000 * +fps);
for (let i = 0; i < n; i++) {
  await p.evaluate(t => window.seek(t), i * 1000 / +fps);
  await p.screenshot({ path: `${out}/${String(i).padStart(4, '0')}.png` });
}
console.log('frames', n);
await b.close();
