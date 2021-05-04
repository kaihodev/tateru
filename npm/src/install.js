const { join } = require('path');
const { chmodSync, mkdirSync, statSync } = require('fs');
const { writeFile } = require('fs/promises');
const { name, cache, fname } = require('./util');
const fetch = require('./http');

fetch(name, handle);
if (!statSync(cache, { throwIfNoEntry: false })) mkdirSync(cache);

const full = join(cache, fname);

async function handle(raw) {
  await Promise.all([
    writeFile(full, raw),
    writeFile('./tateru.js',
      `#!/usr/bin/env node\nrequire('child_process').spawnSync('${full}', process.argv.slice(2), { stdio: 'inherit' });`),
  ]);
  chmodSync(full, 0o777);
}
