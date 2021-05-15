package tateru

import (
	"strings"

	"github.com/chunni/fiptoml"
	"github.com/mattn/go-zglob"

	tateru "github.com/kaihodev/tateru/src/reflect"
)

func SetRunConfigFromToml(c *RunConfig, t *fiptoml.Toml) {
	var modules, target string
	var paths []string
	var ejs, cjs, mjs, watch bool
	d := tateru.ExposeTomlDict(t)
	modules, _ = d["modules"].(string)
	target, _ = d["target"].(string)
	paths, _ = zglob.GlobFollowSymlinks(modules)
	ejs, _ = d["ejs"].(bool)
	cjs, _ = d["cjs"].(bool)
	mjs, _ = d["mjs"].(bool)
	watch, _ = d["watch"].(bool)

	if minify, ok := d["minify"]; ok {
		c.minify = minify.(bool)
	}
	if outDir, ok := d["out_dir"]; ok {
		c.outDir = tateru.String(outDir.(string))
	} else {
		outFile, ok := d["out_file"]
		if ok { c.outFile = tateru.String(outFile.(string)) }
	}
	if tsconfig, ok := d["tsconfig"]; ok {
		c.tsconfig = tateru.String(tsconfig.(string))
	}

	c.inputs = paths
	c.ejs = ejs
	c.cjs = cjs
	c.watch = watch

	if outExt, ok := d["out_extensions"]; ok {
		ext := make(OutExtT)
		var s string
		A := outExt.([]string)
		for i, L := 0, len(A); i != L; i++ {
			s = A[i]
			kv := strings.Split(s, "=")
			ext[kv[0]] = ext[kv[1]]
		}
		c.outExtension = ext
	}
	if mjs {
		if c.outExtension == nil { c.outExtension = make(OutExtT) }
		c.outExtension[".js"] = ".mjs"
	}

	c.target = tateru.String(target)
}

func MergeConfig(def *RunConfig, o *RunConfig) {
	if def.write { o.write = def.write }
	if def.bundle { o.bundle = def.bundle }
	if o.target == nil { o.target = def.target }
	if o.platform == nil { o.platform = def.platform }
	if def.minify { o.minify = true }
	if o.tsconfig == nil { o.tsconfig = def.tsconfig }
	if o.outDir == nil { o.outDir = def.outDir }
	if o.outFile == nil { o.outFile = def.outFile }
	if def.cjs { o.cjs = true }
	if def.ejs { o.ejs = true }
}
