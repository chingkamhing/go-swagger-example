package logger

import (
	"fmt"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New configures the logger based on options in the config.json.
func New(options ...OptionLogger) (*zap.SugaredLogger, error) {
	// populate options
	option := newOptionLogger()
	for _, function := range options {
		function(option)
	}
	// set log output format
	var format string
	{
		if option.format == "json" {
			format = "json"
		} else {
			format = "console"
		}
	}
	// set log level
	var logLevel zapcore.Level
	{
		switch option.level {
		case "debug":
			logLevel = zapcore.DebugLevel
		case "info":
			logLevel = zapcore.InfoLevel
		case "error":
			logLevel = zapcore.ErrorLevel
		case "warn":
			fallthrough
		default:
			logLevel = zapcore.WarnLevel
		}
	}
	// set log output(s)
	var writers []string
	{
		if option.isConsole {
			// enable console output
			writers = append(writers, "stdout")
		}
		if option.filename != "" {
			writers = append(writers, option.filename)
		}
	}
	logger, err := zap.Config{
		Level:            zap.NewAtomicLevelAt(logLevel),
		Development:      false,
		Encoding:         format,
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      writers,
		ErrorOutputPaths: writers,
	}.Build()
	if err != nil {
		return nil, fmt.Errorf("New NewProduction error %w", err)
	}
	return logger.Sugar(), nil
}

// Fatalln print the message and panic exit
func Fatalln(args ...interface{}) {
	// to avoid init zap error, fall back to use golang log
	log.Fatalln(args...)
}

// Fatalf print the message and panic exit
func Fatalf(format string, args ...interface{}) {
	// to avoid init zap error, fall back to use golang log
	log.Fatalf(format, args...)
}
