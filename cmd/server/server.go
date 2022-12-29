package server

import (
	"context"

	"github.com/go-chi/chi/middleware"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/flagext"
	"github.com/justinas/alice"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"go-swagger-example/config"
	"go-swagger-example/gen/restapi"
	"go-swagger-example/gen/restapi/operations"
	"go-swagger-example/logger"
)

// Module provided to fx
var Module = fx.Options(
	fx.Invoke(InvokeHttpServer),
	fx.Invoke(InvokePid),
)

// note: this server start and shutdown mostly copy from https://aaf.engineering/go-web-application-structure-part-2/
// setup http gateway server and grpc gateway dialer
func InvokeHttpServer(lifecycle fx.Lifecycle, cfg *config.Configuration, zapLog *zap.SugaredLogger, log logger.Logger) {
	// init swagger server
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatal(err)
	}
	api := operations.NewExampleAPI(swaggerSpec)
	api.Logger = log.Infof
	server := restapi.NewServer(api)
	server.EnabledListeners = cfg.Server.Schemes
	server.CleanupTimeout = cfg.Server.CleanupTimeout
	server.GracefulTimeout = cfg.Server.GracefulTimeout
	server.MaxHeaderSize = flagext.ByteSize(cfg.Server.MaxHeaderSize)
	server.SocketPath = cfg.Server.SocketPath
	server.Host = cfg.Server.Host
	server.Port = cfg.Server.Port
	server.ListenLimit = cfg.Server.ListenLimit
	server.KeepAlive = cfg.Server.KeepAlive
	server.ReadTimeout = cfg.Server.ReadTimeout
	server.WriteTimeout = cfg.Server.WriteTimeout
	server.TLSHost = cfg.Server.TLSHost
	server.TLSPort = cfg.Server.TLSPort
	server.TLSCertificate = cfg.Server.TLSCertificate
	server.TLSCertificateKey = cfg.Server.TLSCertificateKey
	server.TLSCACertificate = cfg.Server.TLSCACertificate
	server.TLSListenLimit = cfg.Server.TLSListenLimit
	server.TLSKeepAlive = cfg.Server.TLSKeepAlive
	server.TLSReadTimeout = cfg.Server.TLSReadTimeout
	server.TLSWriteTimeout = cfg.Server.TLSWriteTimeout
	server.SetAPI(api)

	// setup middleware
	handler := alice.New(
		middleware.RealIP,
		middleware.Recoverer,
		middleware.Compress(cfg.Server.CompressLevel),
	).Then(api.Serve(nil))
	server.SetHandler(handler)

	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go server.Serve()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				err := server.Shutdown()
				log.Infof("Server http shutdown: %v", err)
				zapLog.Sync()
				return nil
			},
		},
	)
}
