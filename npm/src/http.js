const tls = require('tls');

const PROTO = 'HTTP/1.1';
const HOST = 'raw.githubusercontent.com';

const offset = `${PROTO} 200 OK\r\nConnection: keep-alive\r\nContent-Length: `.length;
const offend = offset + 9;

module.exports = function run(name, cb) {
  const PATH = `/kaihodev/tateru/canary-builds/${name}`;

  const client = tls.connect(
    443, HOST, { servername: HOST },
    () => client.write(`GET ${PATH} ${PROTO}\r\nHost: ${HOST}\r\n\r\n`),
  );
  const m = {};
  client.once('data', respHeaders(m));
  client.on('close', finish(cb, m));
};

function respHeaders(m) {
  return function handle(data) {
    const region = data.slice(offset, offend);
    const str = String.fromCharCode.apply(null, region.toJSON().data);
    const len = +str.substr(0, str.indexOf('\r'));
    console.log('[tateru] sugoi... now downloading %d bytes~', len);
    this.on('data', aggregate(len, m)); // eslint-disable-line no-invalid-this
  };
}

function aggregate(L, m) {
  const buf = m.buf = Buffer.allocUnsafe(L);
  let ptr = 0;
  return function handle(data) {
    data.copy(buf, ptr);
    ptr += data.byteLength;
    if (ptr < L) return true;
    if (ptr === L) return this.end(); // eslint-disable-line no-invalid-this
    console.log('[tateru] yabai!!!  ｡ﾟヽ(ﾟ´Д｀)ﾉﾟ｡: ｡･ﾟ  invalid chunks received...');
    return process.reallyExit(1);
  };
}

function finish(cb, m) {
  return function complete() {
    console.log('[tateru] download completed!~');
    cb(m.buf);
  };
}
