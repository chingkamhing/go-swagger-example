package session

import (
	"context"
	"net/http"
	"os"

	"github.com/alexedwards/scs/badgerstore"
	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/alexedwards/scs/v2/memstore"
	"github.com/dgraph-io/badger"
	"github.com/gomodule/redigo/redis"

	"go-swagger-example/logger"
	"go-swagger-example/pkg/retry"
)

// badgerMinFileSize define the min badger value log file size
// note: not sure the min size, arbitrary start with 32MB
const badgerMinFileSize = 32 * 1024 * 1024

// Session hold session manager info
type Session struct {
	sessionManager *scs.SessionManager
	store          string
	pool           *redis.Pool
	dbBadger       *badger.DB
	Log            logger.Logger
}

var currentOptions options

// NewSession init session base on config.Config settings
func NewSession(log logger.Logger, optionFunctions ...OptionFunc) *Session {
	// populate currOptions
	currentOptions = defaultOptions
	for _, function := range optionFunctions {
		function(&currentOptions)
	}
	s := &Session{
		sessionManager: scs.New(),
		store:          currentOptions.store,
		Log:            log,
	}
	// init session
	s.sessionManager.Lifetime = currentOptions.lifetime
	s.sessionManager.IdleTimeout = currentOptions.idleTimeout
	s.sessionManager.Cookie.Name = currentOptions.name
	s.sessionManager.Cookie.HttpOnly = currentOptions.httpOnly
	s.sessionManager.Cookie.Persist = currentOptions.persist
	s.sessionManager.Cookie.Secure = currentOptions.secure

	var err error
	switch s.store {
	case "redis":
		// establish a redigo connection pool
		redisOptions := redis.DialPassword(currentOptions.redisPassword)
		s.pool = &redis.Pool{
			MaxIdle: 10,
			Dial: func() (redis.Conn, error) {
				return redis.Dial(currentOptions.redisNetwork, currentOptions.redisAddress, redisOptions)
			},
		}
		if err != nil {
			log.Fatalf("session redis error: %s", err)
		}
		// try connecting the redis server
		var redisConn redis.Conn
		err = retry.Do(currentOptions.retryCount, currentOptions.retryInterval, func(i int) error {
			redisConn, err = s.pool.Dial()
			if err != nil {
				log.Errorf("session connect redis failed %v times: %v", i, err)
				return err
			}
			return nil
		})
		if err != nil {
			log.Fatalf("session redis error: %s", err)
		}
		defer redisConn.Close()
		// initialize a new session manager and configure it to use redisstore as the session store
		s.sessionManager.Store = redisstore.New(s.pool)
		log.Infof("session connected to redis at %s", currentOptions.redisAddress)
	case "file":
		// open badger database
		// badgerFileSize define the badger value log file size
		// note: the default vlog file is huge (2GB) in Windows if WithValueLogFileSize() is not called
		// note: after setting WithValueLogFileSize(), somehow, the actual size is doubled in Windows; while the file size in Linux is growing from 0 byte
		// note: as the vlog file size is fine in Linux (i.e. grow from 0 byte instead 2GB in Windows), remove calling WithValueLogFileSize()
		// note: if the server is killed without waiting proper shutdown, start the server next time will always have badger.ErrTruncateNeeded error, need to set "Truncate = true" to get around this problem
		badgerOption := badger.DefaultOptions(currentOptions.filePath)
		if currentOptions.fileSize > 0 {
			var badgerFileSize = currentOptions.fileSize * 1024 * 1024
			if badgerFileSize < badgerMinFileSize {
				badgerFileSize = badgerMinFileSize
			}
			badgerOption = badgerOption.WithValueLogFileSize(int64(badgerFileSize))
		}
		badgerOption.Truncate = true
		badgerOption.Logger = newBadgerLogger(log)
		err = retry.Do(currentOptions.retryCount, currentOptions.retryInterval, func(i int) error {
			s.dbBadger, err = badger.Open(badgerOption)
			if err == badger.ErrTruncateNeeded {
				os.RemoveAll(currentOptions.filePath)
				log.Errorf("session file corrupted, clear it and re-try %d...", i)
			} else if err != nil {
				log.Errorf("session file error: %v", err)
				return err
			}
			return nil
		})
		if err != nil {
			log.Fatalf("NewSession() file error: %s", err)
		}
		s.sessionManager.Store = badgerstore.New(s.dbBadger)
	case "mem":
		fallthrough
	default:
		s.sessionManager.Store = memstore.New()
	}

	return s
}

// close save login user of models.User to session manager
func (s *Session) close() {
	if s.dbBadger != nil {
		// xlose badger db
		s.dbBadger.Close()
	}
	if s.pool != nil {
		// close redis pool
		s.pool.Close()
	}
}

// RenewToken get session's cookie name
func (s *Session) RenewToken(ctx context.Context) error {
	return s.sessionManager.RenewToken(ctx)
}

// Remove remove session's cookie key
func (s *Session) Remove(ctx context.Context, key string) {
	s.sessionManager.Remove(ctx, key)
}

// GetCookieName get session's cookie name
func (s *Session) GetCookieName() string {
	return s.sessionManager.Cookie.Name
}

// SessionMiddleware get session middleware
func (s *Session) SessionMiddleware(next http.Handler) http.Handler {
	return s.sessionManager.LoadAndSave(next)
}
