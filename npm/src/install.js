const { join } = require('path');
const { writeFileSync, chmodSync, mkdirSync, existsSync } = require('fs');
const { name, cache, fname } = require('./util');
const fetch = require('./http');

const ensure_cache = exports.ensure_cache = function ensure_cache() {
  mkdirSync(cache, { recursive: true });
  const full = join(cache, fname).replace(/\\/g, '\\\\');
  console.log('[tateru] using cache', full);
  return full;
};

const install = exports.install = function install(f) {
  return new Promise(res => {
    fetch(name, handle);

    const full = f || ensure_cache();

    function handle(raw) {
      writeFileSync(full, raw);
      chmodSync(full, 0o777);
      console.log('[tateru] downloaded new binary to', full);
      res(full);
    }
  });
};

exports.install_if_missing = function install_if_missing(f) {
  const full = f || ensure_cache();
  if (existsSync(full)) return console.log('[tateru] binary found at', full);
  return install(full);
};
