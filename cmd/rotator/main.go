package main

import (
	"banner-rotator/internal/config"
	"flag"
	"fmt"
	"log"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./configs/config.toml", "Path to configuration file")
}

func main() {
	flag.Parse()

	config, err := config.NewConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", config)
}
