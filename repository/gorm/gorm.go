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

var allUsers = map[string]*models.UserAccount{
	"admin": {
		ID:         1,
		FirstName:  "Administrator",
		LastName:   "",
		Department: "IT",
		UserLogin: &models.UserLogin{
			Email:        "it.developer@wahkwong.com.hk",
			LoginName:    "admin",
			PasswordHash: "$2a$10$T1OYJNv6d3iG.GCEFOUum.8smP.Ynb3UY6Qoxulz2pnPUf/wxCkIy",
			UserID:       0,
		},
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

func (r *repo) GetUserByUsername(username string) (*models.UserAccount, error) {
	//FIXME
	user, ok := allUsers[username]
	if !ok {
		return nil, fmt.Errorf("username %q not found", username)
	}
	return user, nil
}
