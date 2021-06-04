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
	smap := flag.String("smap", "_", "mode to enable source maps on all builds")
	ll := flag.Int("loglevel", 0, "log level: 0 = off, 1 = verbose, 2 = debug, 3 = info, 4 = warning, 5 = error")
	flag.Parse()
	modules := flag.Args()
	if len(modules) == 0 {
		log.Print("[tateru] building all modules in toml...\n\n")
	} else {
		log.Printf("[tateru] building modules: %v...\n", modules)
	}

	tateru.WatchOpts = &api.WatchMode{OnRebuild: Handle}

	cfg := tateru.FromTomlFile(*loc, modules)
	builds, L, names := cfg.GetBuilds(*watchAll, *smap)
	hasWatch := false
	if L > 5 {
		wg := sync.WaitGroup{}
		wg.Add(L)
		for i := 0; i != L; i++ {
			opts := *builds[i]
			opts.LogLevel = api.LogLevel(*ll)
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
			opts.LogLevel = api.LogLevel(*ll)
			if opts.Watch == nil {
				result := api.Build(opts)
				if result.Errors != nil {
					log.Fatalf("[tateru] unable to compile: %v", result.Errors)
				}
				log.Printf("[tateru (%d)] success %s\n", i, names[i])
			} else {
				hasWatch = true
				go api.Build(opts)
			}
		}
	}
	if hasWatch || *watchAll {
		log.Println("[tateru][watch] watch-all has been enabled.")
		<-make(chan struct{})
	}
}
