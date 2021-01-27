package app

import (
	"github.com/ivfisunov/banner-rotator/internal/storage"
	"github.com/sirupsen/logrus"
)

type App struct {
	Storage storage.Storage
	Logger  *logrus.Logger
}

func New(logger *logrus.Logger, storage storage.Storage) *App {
	return &App{Logger: logger, Storage: storage}
}
