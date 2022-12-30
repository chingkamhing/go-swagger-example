package mysql

import (
	"fmt"

	"go-swagger-example/logger"

	"go-swagger-example/gen/models"
)

type repo struct {
	log logger.Logger
	//FIXME
}

var allUsers = map[string]*models.UserInfo{
	"admin": {
		Username: "admin",
		Name:     "Administrator",
		Phone:    "1234-5678",
		Password: "$2a$10$T1OYJNv6d3iG.GCEFOUum.8smP.Ynb3UY6Qoxulz2pnPUf/wxCkIy",
	},
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
	user, ok := allUsers[username]
	if !ok {
		return nil, fmt.Errorf("username %q not found", username)
	}
	return user, nil
}
