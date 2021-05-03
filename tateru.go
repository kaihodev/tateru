package main

import (
	"flag"
	"github.com/evanw/esbuild/pkg/api"
	tateru "github.com/kaihodev/tateru/src"
	"log"
)


func main() {
	log.Println("[tateru] starting...")
	flag.Parse()
	modules := flag.Args()
	if len(modules) == 0 {
		log.Println("[tateru] building all modules in toml...\n")
	} else {
		log.Printf("[tateru] building modules: %v...\n", modules)
	}
	cfg := tateru.FromTomlFile("", modules)

	builds, L, names := cfg.GetBuilds()
	for i := 0; i != L; i++ {
		opts := *builds[i]
		result := api.Build(opts)
		if result.Errors != nil {
			log.Fatalf("[tateru] unable to compile: %v", result.Errors)
		}
		log.Printf("[tateru] success %s\n", names[i])
	}
}
