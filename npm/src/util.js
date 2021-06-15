/* eslint quote-props: ['error', 'consistent'] */
const { join } = require('path');
const { homedir } = require('os');
const { version } = require('../package.json');

const home = homedir();

const arch = {
  'arm': 'arm',
  'arm64': 'arm64',
  'x32': '386',
  'x64': 'amd64',
  'ppc': 'ppc',
  'ppc64': 'ppc64',
  'mips': 'mips',
  'mipsel': 'mipsle',
}[process.arch];

const os = {
  'win32': 'windows',
  'sunos': 'solaris',
}[process.platform] || process.platform;

const cache = {
  'linux': process.env.XDG_CACHE_HOME ? join(process.env.XDG_CACHE_HOME, 'tateru-bin') : join(home, '.cache', 'tateru-bin'),
  'darwin': join(home, 'Library', 'Caches', 'tateru-bin'),
  'windows': join(home, 'AppData', 'Local', 'Cache', 'tateru-bin'),
}[os] || join(home, '.cache', 'tateru-bin');

const name = `${os}-${arch}-tateru`;

module.exports = {
  arch,
  os,
  cache,
  name,
  version,
  fname: `${version}-${name}${os === 'windows' ? '.exe' : ''}`,
};
