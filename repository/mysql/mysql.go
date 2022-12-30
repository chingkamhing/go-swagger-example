package mysql

import (
	"go-swagger-example/logger"

	"go-swagger-example/gen/models"
)

type repo struct {
	log logger.Logger
	//FIXME
}

func Open(log logger.Logger, options ...OptionFunc) (*repo, error) {
	//FIXME
	return &repo{
		log: log,
	}, nil
}

func (r *repo) Close() error {
	//FIXME
	return nil
}

func (r *repo) GetUserByUsername(username string) (*models.UserInfo, error) {
	//FIXME
	return nil, nil
}
