package session

import (
	"github.com/dgraph-io/badger"

	"go-swagger-example/logger"
)

// add this badger logger
type badgerLogger struct {
	log logger.Logger
}

func newBadgerLogger(log logger.Logger) badger.Logger {
	return &badgerLogger{
		log: log,
	}
}

func (l *badgerLogger) Debugf(format string, args ...interface{}) {
	l.log.Debugf(format, args...)
}

func (l *badgerLogger) Infof(format string, args ...interface{}) {
	l.log.Infof(format, args...)
}

func (l *badgerLogger) Warningf(format string, args ...interface{}) {
	l.log.Warnf(format, args...)
}

func (l *badgerLogger) Errorf(format string, args ...interface{}) {
	l.log.Errorf(format, args...)
}
