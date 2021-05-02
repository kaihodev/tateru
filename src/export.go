package tateru

import "github.com/evanw/esbuild/pkg/api"

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
