package main

import (
	"flag"
	"log"
	"sync"

	"github.com/evanw/esbuild/pkg/api"

	tateru "github.com/kaihodev/tateru/src"
)

func Handle(result api.BuildResult) {
	if len(result.Errors) != 0 {
		log.Printf("[tateru][watch] unable to compile: %v", result.Errors)
	}
	oF := result.OutputFiles
	for i, L := 0, len(oF); i != L; i++ {
		file := oF[i]
		log.Printf("[tateru][watch] built successfully: %s\n", file.Path)
	}
}

func main() {
	loc := flag.String("config", "", "path to taterurc config file")
	watchAll := flag.Bool("watch", false, "true/false flag to enable watch mode on all builds")
	flag.Parse()
	modules := flag.Args()
	if len(modules) == 0 {
		log.Print("[tateru] building all modules in toml...\n\n")
	} else {
		log.Printf("[tateru] building modules: %v...\n", modules)
	}

	tateru.WatchOpts = &api.WatchMode{OnRebuild: Handle}

	cfg := tateru.FromTomlFile(*loc, modules)
	builds, L, names := cfg.GetBuilds(*watchAll)
	hasWatch := false
	if L > 5 {
		wg := sync.WaitGroup{}
		wg.Add(L)
		for i := 0; i != L; i++ {
			opts := *builds[i]
			if opts.Watch != nil { hasWatch = true }
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
			if opts.Watch != nil {
				hasWatch = true
				result := api.Build(opts)
				if result.Errors != nil {
					log.Fatalf("[tateru] unable to compile: %v", result.Errors)
				}
				log.Printf("[tateru (%d)] success %s\n", i, names[i])
			} else {
				go api.Build(opts)
			}
		}
	}
	if hasWatch || *watchAll {
		log.Println("[tateru][watch] watch-all has been enabled.")
		<-make(chan struct{})
	}
}
