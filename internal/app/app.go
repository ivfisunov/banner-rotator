package app

import "github.com/sirupsen/logrus"

type App struct {
	Logger  *logrus.Logger
	Storage Storage
}

type Storage struct{}

func New(logger *logrus.Logger) *App {
	return &App{Logger: logger, Storage: Storage{}}
}
