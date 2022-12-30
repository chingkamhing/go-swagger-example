package repository

import (
	"errors"

	"go-swagger-example/config"
	"go-swagger-example/logger"
	"go-swagger-example/repository/mysql"
)

// ErrorDatabaseDriver is a error about "invalid database driver"
var ErrorDatabaseDriver = errors.New("invalid database driver")

// Open return repository interface
func Open(cfg *config.Configuration, log logger.Logger) (repoInterfaces Interfaces, err error) {
	// create a new sql driver
	switch cfg.Database.Driver {
	case "mysql":
		// implement for mysql driver
		repoInterfaces, err = mysql.Open(
			log,
			mysql.OptionDBName(cfg.Database.Dbname),
			mysql.OptionHost(cfg.Database.Host),
			mysql.OptionPort(cfg.Database.Port),
			mysql.OptionUser(cfg.Database.User),
			mysql.OptionPassword(cfg.Database.Password),
			mysql.OptionRetryCount(cfg.Database.RetryCount),
			mysql.OptionRetryInterval(cfg.Database.RetryInterval),
		)
		if err != nil {
			log.Errorf("repository mysql error: %v", err)
			return nil, err
		}
	default:
		// stop retry connecting to db server
		return nil, ErrorDatabaseDriver
	}

	return repoInterfaces, nil
}
