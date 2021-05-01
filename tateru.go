package main

import (
	tateru "github.com/kaihodev/tateru/src"
	"log"
)

func main() {
	println("[tateru] starting...")
	cfg := tateru.FromTomlFile("", nil)
	log.Printf("[tateru] resolved config: %v", cfg)
}