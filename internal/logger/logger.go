package logger

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

func New(env string, path string, level string) (*log.Logger, error) {
	logger := log.New()
	lvl, err := log.ParseLevel(level)
	if err != nil {
		return nil, err
	}
	logger.Level = lvl
	logger.Formatter = &log.JSONFormatter{}

	logFile, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	if env == "production" {
		logger.Out = logFile
	} else {
		logger.Out = io.MultiWriter(logFile, os.Stdout)
	}

	return logger, nil
}
