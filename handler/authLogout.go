package handler

import (
	"github.com/go-openapi/runtime/middleware"

	"go-swagger-example/gen/restapi/operations/auth"
	"go-swagger-example/logger"
	"go-swagger-example/pkg/session"
)

// authLogout handles a request for user logout
type authLogout struct {
	log         logger.Logger
	sessionUser *session.Session
}

// NewAuthLogout handles a request for user logout
func NewAuthLogout(log logger.Logger, sessionUser *session.Session) auth.LogoutHandler {
	return &authLogout{
		log:         log,
		sessionUser: sessionUser,
	}
}

// Handle is auth.Logout handler that handle user logout
func (h *authLogout) Handle(params auth.LogoutParams, principal interface{}) middleware.Responder {
	// as the cookie is HttpOnly, need to clear user's session and cookie on the server side
	h.sessionUser.RemoveUser(params.HTTPRequest)
	return auth.NewLogoutOK()
}
