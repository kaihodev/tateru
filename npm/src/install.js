const { join } = require('path');
const { chmodSync, mkdirSync } = require('fs');
const { writeFile } = require('fs/promises');
const { name, cache, fname } = require('./util');
const fetch = require('./http');

fetch(name, handle);
mkdirSync(cache, { recursive: true });

const full = join(cache, fname).replace(/\\/g, '\\\\');

async function handle(raw) {
  await Promise.all([
    writeFile(full, raw),
    writeFile('./tateru.js',
      `#!/usr/bin/env node\nrequire('child_process').spawnSync('${full}', process.argv.slice(2), { stdio: 'inherit' });`),
  ]);
  chmodSync(full, 0o777);
}
