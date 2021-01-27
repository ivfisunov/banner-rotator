package main

import (
	"flag"
	"log"

	"github.com/ivfisunov/banner-rotator/internal/app"
	"github.com/ivfisunov/banner-rotator/internal/config"
	"github.com/ivfisunov/banner-rotator/internal/logger"
	"github.com/ivfisunov/banner-rotator/internal/storage"
	_ "github.com/lib/pq"
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

	stor, err := storage.New(config.Storage.Dsn)
	if err != nil {
		log.Fatal(err)
	}

	app := app.New(logger, stor)

	app.Logger.Info("App started")
	err = app.Storage.Connect()
	if err != nil {
		log.Fatal(err)
	}
	app.Logger.Info("Connected to Postgres")
	err = app.Storage.Close()
	if err != nil {
		log.Fatal(err)
	}
}
