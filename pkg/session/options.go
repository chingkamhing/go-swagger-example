package session

import (
	"time"
)

type options struct {
	store         string        // store of: redis, file or mem
	redisNetwork  string        // if store is redis, redis network of: tcp
	redisPassword string        // if store is redis, redis password
	redisAddress  string        // if store is redis, redis ip address and port number e.g. 127.0.0.1:6379
	retryCount    int           // session retry open count
	retryInterval time.Duration // session retry open interval in second
	filePath      string        // if store is file, save session file path
	fileSize      int           // if store is file, save session file max size
	lifetime      time.Duration
	idleTimeout   time.Duration
	name          string
	httpOnly      bool
	persist       bool
	secure        bool
}

// app default settings
var defaultOptions = options{
	store:         "mem",
	redisNetwork:  "tcp",
	redisAddress:  "localhost:6379",
	retryCount:    10,
	retryInterval: time.Duration(3 * time.Second),
	filePath:      ".session",
	fileSize:      0,
	lifetime:      time.Duration(180 * time.Second),
	idleTimeout:   time.Duration(60 * time.Second),
	name:          "telesalesd",
	httpOnly:      true,
	persist:       true,
	secure:        true,
}

// OptionFunc control app options behavior
type OptionFunc func(*options)

// OptionStore set session store of: redis, file, mem
func OptionStore(store string) OptionFunc {
	return func(options *options) {
		options.store = store
	}
}

// OptionsRedisNetwork set redis network protocol
func OptionsRedisNetwork(network string) OptionFunc {
	return func(options *options) {
		options.redisNetwork = network
	}
}

// OptionsRedisAddress set redis ip and port address
func OptionsRedisAddress(address string) OptionFunc {
	return func(options *options) {
		options.redisAddress = address
	}
}

// OptionsRedisPassword set redis password
func OptionsRedisPassword(password string) OptionFunc {
	return func(options *options) {
		options.redisPassword = password
	}
}

// OptionsRetryCount set redis connection retry count
func OptionsRetryCount(retryCount int) OptionFunc {
	return func(options *options) {
		options.retryCount = retryCount
	}
}

// OptionsRetryInterval set redis connection retry interval in second
func OptionsRetryInterval(retryInterval time.Duration) OptionFunc {
	return func(options *options) {
		options.retryInterval = retryInterval
	}
}

// OptionFilePath set session store file path
func OptionFilePath(filePath string) OptionFunc {
	return func(options *options) {
		options.filePath = filePath
	}
}

// OptionFileSize set session store file size
func OptionFileSize(fileSize int) OptionFunc {
	return func(options *options) {
		options.fileSize = fileSize
	}
}

// OptionLifetime set session lifetime in minute
func OptionLifetime(lifetime time.Duration) OptionFunc {
	return func(options *options) {
		options.lifetime = lifetime
	}
}

// OptionIdleTimeout set session idle timeout time in minute
func OptionIdleTimeout(idleTimeout time.Duration) OptionFunc {
	return func(options *options) {
		options.idleTimeout = idleTimeout
	}
}

// OptionName set session name
func OptionName(name string) OptionFunc {
	return func(options *options) {
		options.name = name
	}
}

// OptionHttpOnly set session httpOnly
func OptionHttpOnly(httpOnly bool) OptionFunc {
	return func(options *options) {
		options.httpOnly = httpOnly
	}
}

// OptionPersist set session persist across browser close
func OptionPersist(persist bool) OptionFunc {
	return func(options *options) {
		options.persist = persist
	}
}

// OptionSecure set session secure (i.e. https)
func OptionSecure(secure bool) OptionFunc {
	return func(options *options) {
		options.secure = secure
	}
}
