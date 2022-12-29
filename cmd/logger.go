package cmd

import (
	"go.uber.org/zap"

	"go-swagger-example/config"
	"go-swagger-example/logger"
)

// setup zap log output
func ProvideSugaredLogger(cfg *config.Configuration) (*zap.SugaredLogger, error) {
	log, err := logger.New(
		logger.WithFormat(cfg.Log.Format),
		logger.WithLevel(cfg.Log.Level),
		logger.WithIsConsole(cfg.Log.Console),
		logger.WithFilename(cfg.Log.Filename),
	)
	if err != nil {
		return nil, err
	}
	return log, nil
}

// setup log output
func ProvideLogger(log *zap.SugaredLogger) logger.Logger {
	return log
}
