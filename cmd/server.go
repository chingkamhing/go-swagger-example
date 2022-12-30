package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"go.uber.org/fx"

	"go-swagger-example/cmd/server"
	"go-swagger-example/config"
	"go-swagger-example/logger"
)

// Web cli command settings
var cmdServer = &cobra.Command{
	Use:   "server",
	Short: "Start Telesales Admin Web Server",
	Run:   runServer,
}

func init() {
	basename := filepath.Base(os.Args[0])
	// add local parameter
	cmdServer.Flags().StringSlice("server.schemes", []string{"http", "https"}, "comma-seperated schemes to enable: http, https, unix")
	cmdServer.Flags().Duration("server.cleanup-timeout", time.Duration(10*time.Second), "grace period for which to wait before killing idle connections")
	cmdServer.Flags().Duration("server.graceful-timeout", time.Duration(15*time.Second), "grace period for which to wait before shutting down the server")
	cmdServer.Flags().Int("server.max-header-size", 1048576, "controls the maximum number of bytes the server will read parsing the request header's keys and values, including the request line. It does not limit the size of the request body.")
	cmdServer.Flags().String("server.socket-path", fmt.Sprintf("/var/run/%s.sock", basename), "the unix socket to listen on")
	cmdServer.Flags().String("server.host", "localhost", "the IP to listen on")
	cmdServer.Flags().Int("server.port", 80, "the port to listen on for insecure connections, defaults to a random value")
	cmdServer.Flags().Int("server.listen-limit", 0, "limit the number of outstanding requests")
	cmdServer.Flags().Duration("server.keep-alive", time.Duration(3*time.Minute), "sets the TCP keep-alive timeouts on accepted connections. It prunes dead TCP connections ( e.g. closing laptop mid-download)")
	cmdServer.Flags().Duration("server.read-timeout", time.Duration(30*time.Second), "maximum duration before timing out read of the request")
	cmdServer.Flags().Duration("server.write-timeout", time.Duration(60*time.Second), "maximum duration before timing out write of the response")
	cmdServer.Flags().String("server.tls-host", "", "the IP to listen on for tls, when not specified it's the same as --host")
	cmdServer.Flags().Int("server.tls-port", 443, "the port to listen on for secure connections, defaults to a random value")
	cmdServer.Flags().String("server.tls-certificate", "", "the certificate to use for secure connections")
	cmdServer.Flags().String("server.tls-key", "", "the private key to use for secure connections")
	cmdServer.Flags().String("server.tls-ca", "", "the certificate authority file to be used with mutual tls auth")
	cmdServer.Flags().Int("server.tls-listen-limit", 0, "limit the number of outstanding requests")
	cmdServer.Flags().Duration("server.tls-keep-alive", time.Duration(3*time.Minute), "sets the TCP keep-alive timeouts on accepted connections. It prunes dead TCP connections ( e.g. closing laptop mid-download)")
	cmdServer.Flags().Duration("server.tls-read-timeout", time.Duration(30*time.Second), "maximum duration before timing out read of the request")
	cmdServer.Flags().Duration("server.tls-write-timeout", time.Duration(60*time.Second), "maximum duration before timing out write of the response")
	cmdServer.Flags().String("server.static-path", "dist", "static file server root path")
	cmdServer.Flags().Int("server.compress-level", 5, "maximum duration before timing out write of the response")

	cmdServer.Flags().Int("top.number", 100, "Number of top telesales to track")
	cmdServer.Flags().Duration("top.interval", time.Duration(10*time.Minute), "Cron job interval in second to poll the top telesales promotion status")

	cmdServer.Flags().Duration("session.life-time", time.Duration(180*time.Minute), "session life time")
	cmdServer.Flags().Duration("session.idle-timeout", time.Duration(60*time.Minute), "session idle timeout")
	cmdServer.Flags().String("session.name", basename, "session cookie name")
	cmdServer.Flags().Bool("session.persist", true, "cookie persist across browser close")
	cmdServer.Flags().String("session.store", "mem", "cookie storage of: redis, file or mem")
	cmdServer.Flags().String("session.network", "tcp", "(if store is redis) redis network of: tcp")
	cmdServer.Flags().String("session.address", "localhost:6379", "(if store is redis) redis ip address")
	cmdServer.Flags().String("session.password", "", "(if store is redis) redis password")
	cmdServer.Flags().Int("session.retry", 10, "(if store is redis) Session connection retry count")
	cmdServer.Flags().Duration("session.interval", time.Duration(3*time.Second), "(if store is redis) Session connection retry interval")
	cmdServer.Flags().String("session.file-path", ".session", "(if store is file) cookie storage directory path")
	cmdServer.Flags().Int("session.file-size", 0, "(if store is file) specify value log file size in MB (0 for default 2GB)")

	cmdServer.Flags().String("log.format", "text", "log output format (text, json)")
	cmdServer.Flags().String("log.level", "error", "log output level (debug, info, error)")
	cmdServer.Flags().Bool("log.console", false, "log output to console")
	cmdServer.Flags().String("log.filename", "", "log output file")

	// add command
	rootCmd.AddCommand(cmdServer)
}

// note: this server start and shutdown mostly copy from https://aaf.engineering/go-web-application-structure-part-2/
func runServer(cmd *cobra.Command, args []string) {
	fx.New(
		fx.Supply(cmd),
		fx.Provide(
			ProvideSugaredLogger,
			fx.Annotate(ProvideSugaredLogger, fx.As(new(logger.Logger))),
		),
		fx.Invoke(logVersion),
		config.Module,
		server.Module,
	).Run()
}
