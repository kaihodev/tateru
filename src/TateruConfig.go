package tateru

import (
	"github.com/chunni/fiptoml"
	"github.com/evanw/esbuild/pkg/api"
	tateru "github.com/kaihodev/tateru/src/reflect"
	"log"
	"reflect"
	"unsafe"
)

type RunConfig struct {
	extends *string
	outDir *string
	outFile *string
	inputs []string
	ejs bool
	cjs bool
	bundle bool
	target *string
	platform *string
	minify bool
	tsconfig *string
	outExtension map[string]string
}

type Config struct {
	globalPreset string
	builds       CfgMapT
	RunConfig
}
const (
	DefaultConfigName = "default"
	EmptyConfigName = "empty"
)

type CfgMapT = map[string]*RunConfig
var ConfigPresets = CfgMapT{
	EmptyConfigName: &RunConfig{},
	DefaultConfigName: &RunConfig{
		extends:  tateru.String(EmptyConfigName),
		bundle:   true,
		target:   tateru.String("esnext"),
		minify:   true,
		tsconfig: tateru.String("tsconfig.json"),
	},
}

var AllowedOptions = tateru.StructKeys(reflect.ValueOf(*ConfigPresets[EmptyConfigName]), "global_preset")

type OutType = bool
const (
	File OutType = false
	Dir          = !File
)
func (c *RunConfig) OutType() OutType { return c.outDir == nil }
func (c *RunConfig) OutPath() *string { if c.OutType() == Dir { return c.outDir } else { return c.outFile } }

func (c *RunConfig) OutFormat() api.Format {
	if c.cjs { return api.FormatCommonJS }
	if c.ejs { return api.FormatESModule }
	return api.FormatDefault
}

func FromToml(t *fiptoml.Toml, modules []string) *Config {
	cfg := &Config{}
	var m string

	globalPreset, missing := t.GetStringEx("global_preset")
	if missing != nil { globalPreset = DefaultConfigName
	}
	cfg.globalPreset = globalPreset
	preset := ConfigPresets[globalPreset]
	if preset != nil { MergeConfig(preset, (*RunConfig) (unsafe.Pointer(cfg))) }

	builds := make(CfgMapT)
	if modules == nil || len(modules) == 0 {
		d := tateru.ExposeTomlDict(t)
		modules = make([]string, len(d))
		i := 0
		for k := range d {
			if AllowedOptions[k] { continue }
			modules[i] = k
			i++
		}
		modules = modules[:i]
	}
	for i, L := 0, len(modules); i != L; i++ {
		m = modules[i]
		module, err := t.GetTableToml(m)
		if err != nil { log.Fatalf("Unable to find module to build: %s", m) }
		builds[m] = MakeRunConfigFromToml(module)
	}
	cfg.builds = builds

	return cfg
}

func FromTomlFile(loc string, modules []string) *Config {
	data := ReadConfig(loc)
	if data == nil {
		log.Panicf("could not resolve config path %s", loc)
	}
	log.Printf("[tateru] Loaded config, %v bytes", len(*data))
	toml, err := fiptoml.Parse(*data)
	if err != nil { log.Panic(err) }
	return FromToml(toml, modules)
}

func MakeRunConfigFromToml(t *fiptoml.Toml) *RunConfig {
	cfg := &RunConfig{}
	SetRunConfigFromToml(cfg, t)
	return cfg
}
