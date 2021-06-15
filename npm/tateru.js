#!/usr/bin/env node

const { ensure_cache, install, install_if_missing } = require('./src/install');
const { spawnSync } = require('child_process');
const kUnknownErr = 'Dame dayo onii-chan!~ if you see this, an unknown error occured ｡ﾟヽ(ﾟ´Д｀)ﾉﾟ｡: ｡･ﾟ';
const kInstall = '--install', kInstallIfMissing = '--install-if-missing',
  kInstallOnly = '--install-only', kInstallIfMissingOnly = '--install-if-missing-only';

const opts = { [kInstall]: false, [kInstallIfMissing]: true, [kInstallOnly]: false, [kInstallIfMissingOnly]: false },
  rest = [], full = ensure_cache();
(async () => {
  run: {
    parse: {
      const A = process.argv, L = A.length;
      if (L === 2) break parse;
      if (L === 0 || L === 1) break run;
      for (let e, i = 2; i !== L; ++i) if (opts[e = A[i]] !== undefined) opts[e] = true; else rest[rest.length] = e;
    }

    if (opts[kInstallOnly]) {
      await install(full); process.exit(0);
    }
    if (opts[kInstall]) await install(full);

    if (opts[kInstallIfMissingOnly]) {
      install_if_missing(full); process.exit(0);
    }
    if (opts[kInstallIfMissing]) install_if_missing(full);

    spawnSync(full, rest, { stdio: 'inherit' });
    process.exit(0);
  }
  throw Error(kUnknownErr);
})();
