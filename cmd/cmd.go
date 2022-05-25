package main

import (
	"log"

	"github.com/chmey/ransomware_canary/canary"
	"github.com/chmey/ransomware_canary/cfg"
)

func main() {
	config, err := cfg.NewConfig("config.toml")
	if err != nil {
		log.Fatal("error:", err)
	}
	c := canary.NewCanary(config)
	c.Start()
}
