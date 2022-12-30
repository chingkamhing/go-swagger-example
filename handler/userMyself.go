package handler

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"go-swagger-example/gen/models"
	"go-swagger-example/gen/restapi/operations/user"
	"go-swagger-example/logger"
	"go-swagger-example/repository"
)

// userMyself handles a request for user login
type userMyself struct {
	repo repository.Interfaces
	log  logger.Logger
}

// NewUserMyself handles a request for user login
func NewUserMyself(repo repository.Interfaces, log logger.Logger) user.MyselfHandler {
	return &userMyself{
		repo: repo,
		log:  log,
	}
}

// Handle is user.Myself handler that handle user login
func (h *userMyself) Handle(params user.MyselfParams, principal interface{}) middleware.Responder {
	userInfo, ok := principal.(*models.UserAccount)
	if ok != true {
		h.log.Errorf("invalid principal")
		return user.NewMyselfDefault(http.StatusBadRequest).WithPayload(&models.Error{
			Code:    http.StatusBadRequest,
			Message: "invalid principal",
		})
	}
	return user.NewMyselfOK().WithPayload(userInfo)
}
