package app

import (
	"github.com/ivfisunov/banner-rotator/internal/storage/stortypes"
	"github.com/sirupsen/logrus"
)

type App struct {
	Storage stortypes.Storage
	Logger  *logrus.Logger
}

func New(logger *logrus.Logger, storage stortypes.Storage) *App {
	return &App{Logger: logger, Storage: storage}
}
