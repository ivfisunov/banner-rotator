package app

import "github.com/sirupsen/logrus"

type App struct {
	Storage Storage
	Logger  *logrus.Logger
}

type Storage struct{}

func New(logger *logrus.Logger) *App {
	return &App{Logger: logger, Storage: Storage{}}
}
