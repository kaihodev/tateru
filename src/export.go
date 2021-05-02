package tateru

import (
	"github.com/evanw/esbuild/pkg/api"
	"strings"
	"unsafe"
)

func (c *RunConfig) OutType() OutType { return c.outDir == nil }
func (c *RunConfig) OutPath() *string { if c.OutType() == Dir { return c.outDir } else { return c.outFile } }

func (c *RunConfig) Format() api.Format {
	if c.cjs { return api.FormatCommonJS }
	if c.ejs { return api.FormatESModule }
	return api.FormatDefault
}

func (c *RunConfig) Write() bool { return c.write }

func (c *RunConfig) Platform() api.Platform {
	switch *c.platform {
	case "node":
		return api.PlatformNode
	case "browser":
		return api.PlatformBrowser
	default:
		return api.PlatformNeutral
	}
}

func (c *RunConfig) Inputs() []string { return c.inputs }

func (c *RunConfig) Target() api.Target {
	switch strings.ToLower(*c.target) {
	case "node":
	case "esnext":
		break

	case "es5":
		return api.ES5
	case "es6":
	case "es2015":
		return api.ES2015

	case "es11":
	case "es2020":
		return api.ES2020

	case "es10":
	case "es2019":
		return api.ES2019
	case "es9":
	case "es2018":
		return api.ES2018
	case "es8":
	case "es2017":
		return api.ES2017
	case "es7":
	case "es2016":
		return api.ES2016
	}
	return api.ESNext
}

func (c *RunConfig) Minify() bool { return c.minify }

func (c *RunConfig) MakeESBOptions() *api.BuildOptions {
	opts := &api.BuildOptions{}
	if c.OutType() == File {
		opts.Outfile = *c.OutPath()
	} else {
		opts.Outdir = *c.OutPath()
	}
	opts.Write = c.Write()
	opts.Platform = c.Platform()
	opts.Format = c.Format()
	opts.EntryPoints = c.Inputs()
	opts.Target = c.Target()
	if c.Minify() {
		opts.MinifyWhitespace = true
		opts.MinifySyntax = true
		opts.MinifyIdentifiers = true
	}
	return opts
}

func (c *Config) GetBuilds() ([]*api.BuildOptions, int) {
	b := c.builds
	L := len(b)
	if L == 0 {
		return []*api.BuildOptions{c.MakeESBOptions()}, 1
	}
	p := (*RunConfig) (unsafe.Pointer(c))
	i, result := 0, make([]*api.BuildOptions, L)
	for k := range b {
		cfg := b[k]
		MergeConfig(p, cfg)
		result[i] = cfg.MakeESBOptions()
		i++
	}
	return result, L
}
