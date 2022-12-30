package repository

import (
	"go-swagger-example/gen/models"
)

// Interfaces defines the repository interfaces
type Interfaces interface {
	// Root APIs

	// Close close database connection
	Close() error

	// User APIs
	GetUserByUsername(username string) (*models.UserInfo, error)
	//FIXME
}
