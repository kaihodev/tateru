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
	_ = api.BuildOptions{Write: true}
}