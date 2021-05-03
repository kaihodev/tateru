package tateru

import (
	"github.com/chunni/fiptoml"
	tateru "github.com/kaihodev/tateru/src/reflect"
	"path/filepath"
	"strings"
)

func SetRunConfigFromToml(c *RunConfig, t *fiptoml.Toml) {
	var modules, target string
	var paths []string
	var ejs, cjs, mjs bool

	modules, _ = t.GetStringEx("modules")
	target, _ = t.GetStringEx("target")
	paths, _ = filepath.Glob(modules)
	ejs, _ = t.GetBoolEx("ejs")
	cjs, _ = t.GetBoolEx("cjs")
	mjs, _ = t.GetBoolEx("mjs")

	if outDir, e := t.GetStringEx("out_dir"); e == nil {
		c.outDir = tateru.String(outDir)
	} else {
		outFile, _ := t.GetStringEx("out_file")
		c.outFile = tateru.String(outFile)
	}

	c.inputs = paths
	c.ejs = ejs
	c.cjs = cjs

	if outExt := t.GetStringArray("out_extensions"); outExt == nil {
		ext := make(map[string]string)
		var s string
		for i, L := 0, len(outExt); i != L; i++ {
			s = outExt[i]
			kv := strings.Split(s, "=")
			ext[kv[0]] = ext[kv[1]]
		}
		c.outExtension = ext
	}
	if mjs {
		if c.outExtension == nil { c.outExtension = make(map[string]string) }
		c.outExtension[".js"] = ".mjs"
	}

	c.target = tateru.String(target)
}

func MergeConfig(def *RunConfig, o *RunConfig) {
	if def.write { o.write = def.write }
	if def.bundle { o.bundle = def.bundle }
	if def.target != nil { o.target = def.target }
	if def.platform != nil { o.platform = def.platform }
	if def.minify { o.minify = true }
	if def.tsconfig != nil { o.tsconfig = def.tsconfig }
}
