package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ivfisunov/banner-rotator/internal/app"
	"github.com/ivfisunov/banner-rotator/internal/config"
	"github.com/ivfisunov/banner-rotator/internal/logger"
	"github.com/ivfisunov/banner-rotator/internal/server"
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	srv := server.NewServer(app, config.HTTP.Host, config.HTTP.Port)

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals)

		<-signals
		logger.Info("Server is stopping...")
		signal.Stop(signals)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := srv.Stop(ctx); err != nil {
			logger.Error("failed to stop http server: " + err.Error())
		}
	}()

	logger.Info("Banner App started")
	if err := srv.Start(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error("failed to start http server: " + err.Error())
	}
}
