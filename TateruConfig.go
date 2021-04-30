package tateru

import (
	"github.com/chunni/fiptoml"
	"github.com/evanw/esbuild/pkg/api"
	tateru "github.com/kaihodev/tateru/reflect"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

type CfgMapT = map[string]*RunConfig
var ConfigPresets = make(CfgMapT)

type RunConfig struct {
	extends *string
	outDir *string
	outFile *string
	inputs []string
	ejs bool
	cjs bool
}

type Config struct {
	globalPreset string
	builds CfgMapT
	RunConfig
}
var AllowedOptions = map[string]bool{
	"outDir": true,
	"outFile": true, "ejs": true, "cjs": true}

type OutType = bool
const (
	File OutType = false
	Dir = !File
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
	if missing != nil { globalPreset = "default" }
	cfg.globalPreset = globalPreset
	preset := ConfigPresets[globalPreset]
	if preset {
		SetRunConfigFromToml()
	}

	builds := make(CfgMapT)
	if len(modules) == 0 {
		d := *tateru.ExposeTomlDict(&t)
		modules := make([]string, len(d))
		i := 0
		for k := range d {
			modules[i] = k
			i++
		}
	}
	for i, L := 0, len(modules); i != L; i++ {
		m = modules[i]
		if AllowedOptions[m] { continue }
		module, err := t.GetTableToml(m)
		if err != nil { log.Fatalf("Unable to find module to build: %s", m) }
		builds[m] = MakeRunConfigFromToml(module)
	}

	return cfg
}

func FromTomlFile(loc string, modules []string) *Config {
	data := ReadConfig(loc)
	if data == nil {
		log.Panicf("could not resolve config path %s", loc)
	}
	log.Printf("[tateru] Loaded %v bytes", len(*data))
	toml, err := fiptoml.Parse(*data)
	log.Printf("contents %v | err %v", toml, err)
	return FromToml(toml, modules)
}

func MakeRunConfigFromToml(t *fiptoml.Toml) *RunConfig {
	cfg := &RunConfig{}
	SetRunConfigFromToml(cfg, t)
	return cfg
}
func SetRunConfigFromToml(c *RunConfig, t *fiptoml.Toml) {
	var modules string
	var paths []string
	var ejs, cjs bool

	modules, _ = t.GetStringEx("modules")
	paths, _ = filepath.Glob(modules)
	ejs, _ = t.GetBoolEx("ejs")
	cjs, _ = t.GetBoolEx("cjs")

	c.inputs = paths
	c.ejs = ejs
	c.cjs = cjs
}

func ReadConfig(loc string) *[]byte {
	targetPath := filepath.Dir(loc)
	fileName := filepath.Base(loc)
	if targetPath == "" { targetPath, _ = os.Getwd() }
	if fileName == "." { fileName = ".taterurc" }
	basePath := "/"
	log.Printf("Target and file: %s || %s | %s", loc, targetPath, fileName)
	for {
		rel, _ := filepath.Rel(basePath, targetPath)
		if rel == "." { break }
		p := path.Join(targetPath, fileName)
		log.Printf("Looking at %s", p)
		if !exists(p) {
			targetPath = path.Join(targetPath, "..")
			continue
		}
		bytes, err := ioutil.ReadFile(p)
		if err != nil { log.Fatalln(err) }
		return &bytes
	}
	return nil
}

func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}