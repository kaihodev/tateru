package tateru

import (
	"github.com/chunni/fiptoml"
	tateru "github.com/kaihodev/tateru/src/reflect"
	"path/filepath"
)

func SetRunConfigFromToml(c *RunConfig, t *fiptoml.Toml) {
	var modules, target string
	var paths []string
	var ejs, cjs bool

	modules, _ = t.GetStringEx("modules")
	target, _ = t.GetStringEx("target")
	paths, _ = filepath.Glob(modules)
	ejs, _ = t.GetBoolEx("ejs")
	cjs, _ = t.GetBoolEx("cjs")

	c.inputs = paths
	c.ejs = ejs
	c.cjs = cjs

	c.target = tateru.String(target)
}

func MergeConfig(def *RunConfig, o *RunConfig) {
	if def.target != nil { o.target = def.target }
	if def.platform != nil { o.platform = def.platform }
	if def.bundle { o.bundle = true }
	if def.minify { o.minify = true }
	if def.tsconfig != nil { o.tsconfig = def.tsconfig }
}
