package config

import (
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

// default settings
const defaultEnvPrefix string = "EXAMPLE" // default server env prefix

// default config settings in case there is no config file
var defaultConfigs = map[string]interface{}{
	"log.level":   "warn",
	"log.console": true,
	"log.file":    "",
}

var config Configuration

// Configuration is a system wide configuration settings
// note: please refer https://godoc.org/github.com/jessevdk/go-flags#hdr-Available_field_tags for option's tag usage
type Configuration struct {
	Server   ServerConfig
	Database DatabaseConfig
	Session  SessionConfig
	Log      LogConfig
}

// ServerConfig define server config structure
type ServerConfig struct {
	Schemes           []string      `mapstructure:"schemes"`           // the listeners to enable, this can be repeated and defaults to the schemes in the swagger spec
	CleanupTimeout    time.Duration `mapstructure:"cleanup-timeout"`   // grace period for which to wait before killing idle connections
	GracefulTimeout   time.Duration `mapstructure:"graceful-timeout"`  // grace period for which to wait before shutting down the server
	MaxHeaderSize     int           `mapstructure:"max-header-size"`   // controls the maximum number of bytes the server will read parsing the request header's keys and values, including the request line. It does not limit the size of the request body.
	SocketPath        string        `mapstructure:"socket-path"`       // the unix socket to listen on
	Host              string        `mapstructure:"host"`              // the IP to listen on
	Port              int           `mapstructure:"port"`              // the port to listen on for insecure connections, defaults to a random value
	ListenLimit       int           `mapstructure:"listen-limit"`      // limit the number of outstanding requests
	KeepAlive         time.Duration `mapstructure:"keep-alive"`        // sets the TCP keep-alive timeouts on accepted connections. It prunes dead TCP connections ( e.g. closing laptop mid-download)
	ReadTimeout       time.Duration `mapstructure:"read-timeout"`      // maximum duration before timing out read of the request
	WriteTimeout      time.Duration `mapstructure:"write-timeout"`     // maximum duration before timing out write of the response
	TLSHost           string        `mapstructure:"tls-host"`          // the IP to listen on for tls, when not specified it's the same as --host
	TLSPort           int           `mapstructure:"tls-port"`          // the port to listen on for secure connections, defaults to a random value
	TLSCertificate    string        `mapstructure:"tls-certificate"`   // the certificate to use for secure connections
	TLSCertificateKey string        `mapstructure:"tls-key"`           // the private key to use for secure connections
	TLSCACertificate  string        `mapstructure:"tls-ca"`            // the certificate authority file to be used with mutual tls auth
	TLSListenLimit    int           `mapstructure:"tls-listen-limit"`  // limit the number of outstanding requests
	TLSKeepAlive      time.Duration `mapstructure:"tls-keep-alive"`    // sets the TCP keep-alive timeouts on accepted connections. It prunes dead TCP connections ( e.g. closing laptop mid-download)
	TLSReadTimeout    time.Duration `mapstructure:"tls-read-timeout"`  // maximum duration before timing out read of the request
	TLSWriteTimeout   time.Duration `mapstructure:"tls-write-timeout"` // maximum duration before timing out write of the response
	StaticPath        string        `mapstructure:"static-path"`       // Static file server's root path
	CompressLevel     int           `mapstructure:"compress-level"`    // Response compress level of 0 ~ 9, where 0 is no compression
}

// DatabaseConfig define database config structure
type DatabaseConfig struct {
	DBAccounts    string        `mapstructure:"accounts"` // DBAccounts web config string (note: this has higher precedence than other database settings)
	Driver        string        `mapstructure:"driver"`   // Database driver: mssql
	Host          string        `mapstructure:"host"`     // Database host name
	Port          int           `mapstructure:"port"`     // Database port number
	Dbname        string        `mapstructure:"name"`     // Database name
	User          string        `mapstructure:"user"`     // Database username
	Password      string        `mapstructure:"password"` // Database name
	RetryCount    int           `mapstructure:"retry"`    // Database connection retry count
	RetryInterval time.Duration `mapstructure:"interval"` // Database connection retry interval in second
}

// SessionConfig define the http session config structure
type SessionConfig struct {
	Lifetime      time.Duration `mapstructure:"life-time"`    // session life time (in minute)
	IdleTimeout   time.Duration `mapstructure:"idle-timeout"` // session idle timeout (in minute)
	Name          string        `mapstructure:"name"`         // session cookie name
	Persist       bool          `mapstructure:"persist"`      // cookie persist across browser close
	Store         string        `mapstructure:"store"`        // cookie storage of: redis, file or mem
	Network       string        `mapstructure:"network"`      // (if store is redis) redis network of: tcp
	Address       string        `mapstructure:"address"`      // (if store is redis) redis ip address
	Password      string        `mapstructure:"password"`     // (if store is redis) redis password
	RetryCount    int           `mapstructure:"retry"`        // (if store is redis) Session connection retry count
	RetryInterval time.Duration `mapstructure:"interval"`     // (if store is redis) Session connection retry interval in second
	FilePath      string        `mapstructure:"file-path"`    // (if store is file) cookie storage directory path
	FileSize      int           `mapstructure:"file-size"`    // (if store is file) specify value log file size in MB (0 for default 2GB)
}

// LogConfig define the log config structure
type LogConfig struct {
	Format   string // log format: text, json
	Level    string // log levels: debug, info, error
	Console  bool   // log output to console
	Filename string // log output file
}

// Module provided to fx
var Module = fx.Provide(ProvideConfig)

// LoadConfig load config setting from file name stored in flag "config"
func ProvideConfig(cmd *cobra.Command) (*Configuration, error) {
	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		return nil, err
	}

	// set default config settings
	for k, v := range defaultConfigs {
		viper.SetDefault(k, v)
	}

	// get environment config
	viper.SetEnvPrefix(defaultEnvPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// read config file
	// note: skip handling error so that program can still run although config file is absent
	configFile, _ := cmd.Flags().GetString("config")
	viper.SetConfigFile(configFile)
	errReadConfig := viper.ReadInConfig()

	// unmarshal the config file
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("fail to unmarshal config file %q: %v", configFile, err)
	}

	if errReadConfig == nil {
		// watch for config file changes
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			// read config file
			if err := viper.ReadInConfig(); err != nil {
				log.Printf("Reload config file error: %v", err)
			}
			// unmarshal the config file
			var tmpConfig Configuration
			if err := viper.Unmarshal(&tmpConfig); err != nil {
				log.Printf("Unable to decode into struct: %v", err)
			} else {
				log.Println("Success reload config file")
				config = tmpConfig
			}
		})
	}

	return &config, nil
}
