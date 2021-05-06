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
$ npm i -g tateru
$ yarn global add tateru
$ pnpm install --global tateru
```
Note: add `--save` if you are using npm < 5.0.0

**You may also install via --save-dev and equivalent options. In those cases, make sure to invoke the executable via npx.**

Set up your tateru toml:

1. Create .taterurc
2. Add your build config
3. Run `$ tateru`

You may also add tateru to your scripts object in package.json. From there, you can invoke the build via yarn build. For more details, see how we use it in [hikidashi](https://npm.im/hikidashi).

Example partial toml:

```toml
out_dir = 'dist'
cjs = true

[safe]
modules = 'src/safe/**/*.ts'
out_dir = 'dist/safe'
ejs = true
mjs = true

[safe_bundle]
modules = 'src/safe.ts'

[hikidashi_module]
modules = 'src/index.ts'

[tests]
modules = 'tests/**/*.ts'
out_dir = 'tests/dist'
minify = true
```

<div align="center">
<img src="https://files.yande.re/sample/8eb9828da10f4cee45fa2297dce552a8/yande.re%20415409%20sample%20akashi_%28azur_lane%29%20animal_ears%20azur_lane%20heels%20nishina_hima.jpg" width="720" />
</div>
