# Tateru

### Build your TypeScript projects with speed and ease.

[NPM](https://npm.im/tateru) |
[Docs](https://kaihodev.github.io/tateru) |
[Github](https://github.com/kaihodev/tateru/tree/develop)

[![CD: Canary](https://github.com/kaihodev/tateru/actions/workflows/cd-canary.yml/badge.svg)](https://github.com/kaihodev/tateru/actions/workflows/cd-canary.yml)

Tateru is released under the [MIT license](https://github.com/kaihodev/tateru/blob/develop/LICENSE) & supports modern environments.<br>

## Global install

**Use any of the following.**
```shell
$ npm i -g tateru@latest
$ yarn global add tateru
$ pnpm install --global tateru
```
Note: add `--save` if you are using npm < 5.0.0

**You may also install via --save-dev and equivalent options. In those cases, make sure to invoke the executable via npx.**

Set up your tateru toml:

1. Create .taterurc
2. Add your build config
3. Run `$ tateru`

Running the tateru script standalone will invoke the build process. Optionally, you may pass CLI args for advanced usage. Try it out with `$ tateru --help` to display available options.

You may also add tateru to your scripts object in package.json. From there, you can invoke the build via yarn build. For more details, see how we use it in [hikidashi](https://npm.im/hikidashi).

Example partial toml:

```toml
out_dir = 'dist'
cjs = true

[submodule]
modules = 'src/sub/**/*.ts'
out_dir = 'dist/sub'
ejs = true
mjs = true

[dist_bundle]
modules = 'src/index.ts'

[tests]
modules = 'tests/**/*.ts'
out_dir = 'tests/dist'
watch = true
```

Then, modify your package.json to use yarn build:src and yarn build:test.
```json
{...
  "scripts": {
    "build:src": "tateru submodule dist_bundle",
    "build:test": "tateru tests"
  }
...}
```

<div align="center">
<img src="https://files.yande.re/sample/8eb9828da10f4cee45fa2297dce552a8/yande.re%20415409%20sample%20akashi_%28azur_lane%29%20animal_ears%20azur_lane%20heels%20nishina_hima.jpg" width="720" />
</div>
