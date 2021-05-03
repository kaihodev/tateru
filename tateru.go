package main

import (
	"flag"
	"github.com/evanw/esbuild/pkg/api"
	tateru "github.com/kaihodev/tateru/src"
	"log"
	"sync"
)


func main() {
	loc := flag.String("config", "", "path to taterurc config file")
	flag.Parse()
	modules := flag.Args()
	if len(modules) == 0 {
		log.Print("[tateru] building all modules in toml...\n\n")
	} else {
		log.Printf("[tateru] building modules: %v...\n", modules)
	}
	cfg := tateru.FromTomlFile(*loc, modules)
	builds, L, names := cfg.GetBuilds()

	if L > 5 {
		wg := sync.WaitGroup{}
		wg.Add(L)
		for i := 0; i != L; i++ {
			opts := *builds[i]
			go func(i int) {
				result := api.Build(opts)
				if result.Errors != nil {
					log.Fatalf("[tateru] unable to compile: %v", result.Errors)
				}
				log.Printf("[tateru (%d)] success %s\n", i, names[i])
				wg.Done()
			}(i)
		}
		wg.Wait()
	} else {
		for i := 0; i != L; i++ {
			opts := *builds[i]
			result := api.Build(opts)
			if result.Errors != nil {
				log.Fatalf("[tateru] unable to compile: %v", result.Errors)
			}
			log.Printf("[tateru (%d)] success %s\n", i, names[i])
		}
	}
}
