package handler

import (
	"net/http"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/crypto/bcrypt"

	"go-swagger-example/gen/models"
	"go-swagger-example/gen/restapi/operations/auth"
	"go-swagger-example/logger"
	"go-swagger-example/pkg/session"
	"go-swagger-example/repository"
)

// authLogin handles a request for user login
type authLogin struct {
	repo        repository.Interfaces
	log         logger.Logger
	sessionUser *session.Session
}

// NewAuthLogin handles a request for user login
func NewAuthLogin(repo repository.Interfaces, log logger.Logger, sessionUser *session.Session) auth.LoginHandler {
	return &authLogin{
		repo:        repo,
		log:         log,
		sessionUser: sessionUser,
	}
}

// Handle is auth.Login handler that handle user login
func (h *authLogin) Handle(params auth.LoginParams) middleware.Responder {
	// get the user from database by the request's username
	user, err := h.repo.GetUserByUsername(params.Body.Username)
	if err != nil {
		return loginDefaultError(h, http.StatusBadRequest, "user not found", err)
	}
	// compare the request's password to the one in database
	isPasswordValid := comparePassword(params.Body.Password, user.UserLogin.PasswordHash)
	if !isPasswordValid {
		return loginDefaultError(h, http.StatusUnauthorized, "invalid password", err)
	}
	// valid user, save the user info into user session and response OK
	h.log.Infof("User %q login at %v", params.Body.Username, time.Now())
	// remove the password for security reason
	user.UserLogin.PasswordHash = ""
	// renew session's token after login for security reason
	err = h.sessionUser.RenewToken(params.HTTPRequest.Context())
	if err != nil {
		return loginDefaultError(h, http.StatusInternalServerError, "fail to renew session", err)
	}
	h.sessionUser.SaveUser(user, params.HTTPRequest)
	return auth.NewLoginOK().WithPayload(user)
}

// comparePassword compare input password against the password in database
func comparePassword(password string, against string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(against), []byte(password))
	return err == nil
}

// loginDefaultError is a helper function to return default error
func loginDefaultError(h *authLogin, code int, message string, err error) *auth.LoginDefault {
	h.log.Infof("%s: %v", message, err)
	return auth.NewLoginDefault(code).WithPayload(&models.Error{
		Code:    int32(code),
		Message: message,
	})
}
