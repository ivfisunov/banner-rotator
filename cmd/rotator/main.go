package main

import (
	"flag"
	"log"

	"github.com/ivfisunov/banner-rotator/internal/app"
	"github.com/ivfisunov/banner-rotator/internal/config"
	"github.com/ivfisunov/banner-rotator/internal/logger"
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

	logger, err := logger.New(config.Env, config.Logger.Path, config.Logger.Level)
	if err != nil {
		log.Fatal(err)
	}

	app := app.New(logger)

	app.Logger.Info("App started")
}
