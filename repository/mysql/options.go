package mysql

import (
	"time"
)

type options struct {
	dbname        string
	host          string
	port          int
	user          string
	password      string
	retryCount    int
	retryInterval time.Duration
}

// app default settings
var defaultOptions = options{
	dbname:        "QPAccountsDB",
	host:          "localhost",
	port:          1433,
	user:          "",
	password:      "",
	retryCount:    10,
	retryInterval: time.Duration(3 * time.Second),
}

// OptionFunc control app options behavior
type OptionFunc func(*options)

// OptionDBName set database name
func OptionDBName(dbname string) OptionFunc {
	return func(options *options) {
		options.dbname = dbname
	}
}

// OptionHost set database host address
func OptionHost(host string) OptionFunc {
	return func(options *options) {
		options.host = host
	}
}

// OptionPort set database port number
func OptionPort(port int) OptionFunc {
	return func(options *options) {
		options.port = port
	}
}

// OptionUser set database user
func OptionUser(user string) OptionFunc {
	return func(options *options) {
		options.user = user
	}
}

// OptionPassword set database password
func OptionPassword(password string) OptionFunc {
	return func(options *options) {
		options.password = password
	}
}

// OptionRetryCount set database connection retry count
func OptionRetryCount(retrycount int) OptionFunc {
	return func(options *options) {
		options.retryCount = retrycount
	}
}

// OptionRetryInterval set database connection retry interval in second
func OptionRetryInterval(retryinterval time.Duration) OptionFunc {
	return func(options *options) {
		options.retryInterval = retryinterval
	}
}
