package logger

// options for logger
type optionLogger struct {
	format    string
	level     string
	isConsole bool
	filename  string
}

// OptionLogger control app options behavior
type OptionLogger func(*optionLogger)

// app default settings
func newOptionLogger() *optionLogger {
	return &optionLogger{
		format:    "text",
		level:     "info",
		isConsole: true,
		filename:  "",
	}
}

// WithFormat set logger output format: text, json
func WithFormat(format string) OptionLogger {
	if format == "" {
		return func(option *optionLogger) {}
	}
	return func(option *optionLogger) {
		option.format = format
	}
}

// WithLevel set logger output level: debug, info, error, warn
func WithLevel(level string) OptionLogger {
	if level == "" {
		return func(option *optionLogger) {}
	}
	return func(option *optionLogger) {
		option.level = level
	}
}

// WithIsConsole set logger output to console or not
func WithIsConsole(isConsole bool) OptionLogger {
	return func(option *optionLogger) {
		option.isConsole = isConsole
	}
}

// WithFilename set logger output to a file
func WithFilename(filename string) OptionLogger {
	if filename == "" {
		return func(option *optionLogger) {}
	}
	return func(option *optionLogger) {
		option.filename = filename
	}
}
