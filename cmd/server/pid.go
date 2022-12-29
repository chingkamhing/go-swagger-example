package server

import (
	"context"
	"fmt"
	"os"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"go-swagger-example/config"
	"go-swagger-example/logger"
)

// note: this server start and shutdown mostly copy from https://aaf.engineering/go-web-application-structure-part-2/
// setup http gateway server and grpc gateway dialer
func InvokePid(lifecycle fx.Lifecycle, cfg *config.Configuration, zapLog *zap.SugaredLogger, log logger.Logger) {
	pidFilename := os.Args[0] + ".pid"
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				// write the pid to default pid file
				pid := os.Getpid()
				data := []byte(fmt.Sprintf("%d", pid))
				err := os.WriteFile(pidFilename, data, 0644)
				if err != nil {
					return fmt.Errorf("fail to save pid file %s: %w", pidFilename, err)
				}
				log.Infof("Save pid %v to file %q", pid, pidFilename)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				err := os.Remove(pidFilename)
				if err == nil {
					log.Infof("Removed pid file %s", pidFilename)
				} else {
					log.Infof("Fail to delete pid file %s.", pidFilename)
				}
				return nil
			},
		},
	)
}
