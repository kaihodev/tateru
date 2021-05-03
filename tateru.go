package main

import (
	"github.com/evanw/esbuild/pkg/api"
	tateru "github.com/kaihodev/tateru/src"
	"log"
)


func main() {
	println("[tateru] starting...")
	cfg := tateru.FromTomlFile("", nil)

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
