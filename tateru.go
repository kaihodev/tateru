package main

import (
	"github.com/evanw/esbuild/pkg/api"
	tateru "github.com/kaihodev/tateru/src"
	"log"
)

func main() {
	println("[tateru] starting...")
	cfg := tateru.FromTomlFile("", nil)
	log.Printf("[tateru] resolved config: %v", cfg)
	opts := api.BuildOptions{}
	if cfg.OutType() == tateru.File {
		opts.Outfile = *cfg.OutPath()
	} else {
		opts.Outdir = *cfg.OutPath()
	}
	opts.Write = cfg.Write()
	opts.Platform = cfg.Platform()
	opts.Format = cfg.Format()
	opts.EntryPoints = cfg.Inputs()
	log.Println(opts)
}
