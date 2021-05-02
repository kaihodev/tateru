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

	builds, L := cfg.GetBuilds()
	for i := 0; i != L; i++ {
		opts := *builds[i]
		log.Println(opts)
		result := api.Build(opts)
		log.Println(result)
	}
}
