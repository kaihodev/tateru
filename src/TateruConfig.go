package tateru

import (
	"github.com/chunni/fiptoml"
	tateru "github.com/kaihodev/tateru/src/reflect"
	"log"
	"os"
	"reflect"
	"unsafe"
)

type RunConfig struct {
	extends      *string
	outDir       *string
	outFile      *string
	inputs       []string
	ejs          bool
	cjs          bool
	write        bool
	bundle       bool
	target       *string
	platform     *string
	minify       bool
	tsconfig     *string
	outExtension OutExtT

	name string
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
type OutExtT = map[string]string

var ConfigPresets = CfgMapT{
	EmptyConfigName: &RunConfig{name: EmptyConfigName},
	DefaultConfigName: &RunConfig{
		extends:  tateru.String(EmptyConfigName),
		write:    true,
		bundle:   true,
		platform: tateru.String("node"),
		target:   tateru.String("esnext"),
		minify:   true,
		tsconfig: tateru.String("tsconfig.json"),
		name:     DefaultConfigName,
	},
}

var AllowedOptions = tateru.StructKeys(reflect.ValueOf(*ConfigPresets[EmptyConfigName]), "global_preset")

type OutType = bool
const (
	File OutType = false
	Dir          = !File
)

func FromToml(t *fiptoml.Toml, modules []string) *Config {
	cfg := &Config{}
	var m string
	d := tateru.ExposeTomlDict(t)
	globalPreset, ok := d["global_preset"]
	if !ok { globalPreset = DefaultConfigName }
	cfg.globalPreset = globalPreset.(string)
	preset := ConfigPresets[cfg.globalPreset]

	p := (*RunConfig) (unsafe.Pointer(cfg))
	SetRunConfigFromToml(p, t)
	if preset != nil { MergeConfig(preset, p) }

	builds := make(CfgMapT)
	if modules == nil || len(modules) == 0 {
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
		if err != nil || module == nil { log.Fatalf("Unable to find module to build: %s", m) }
		c := MakeRunConfigFromToml(module)
		MergeConfig(p, c)
		c.name = m
		builds[m] = c
	}
	cfg.builds = builds

	return cfg
}

func FromTomlFile(loc string, modules []string) *Config {
	data := ReadConfig(loc)
	if data == nil {
		log.Printf("could not locate a .taterurc config... %s\n", loc)
		os.Exit(1)
	}
	toml, err := fiptoml.Parse(*data)
	if err != nil { log.Panic(err) }
	return FromToml(toml, modules)
}

func MakeRunConfigFromToml(t *fiptoml.Toml) *RunConfig {
	cfg := &RunConfig{}
	SetRunConfigFromToml(cfg, t)
	return cfg
}
