package session

import (
	"context"
	"encoding/gob"
	"fmt"
	"net/http"

	"go-swagger-example/gen/models"

	"github.com/go-openapi/errors"
)

// user's session key name
const userKey = "UserSession"

func init() {
	// gob register models.User in order to support context decode
	gob.Register(&models.UserAccount{})
}

// SaveUser save user info
func (s *Session) SaveUser(user *models.UserAccount, r *http.Request) {
	s.sessionManager.Put(r.Context(), userKey, user)
}

// GetUser get user info
func (s *Session) GetUser(ctx context.Context) (user *models.UserAccount, err error) {
	any := s.sessionManager.Get(ctx, userKey)
	user, ok := any.(*models.UserAccount)
	if !ok {
		return nil, errors.Unauthenticated("invalid usser session")
	}
	return user, nil
}

// RemoveUser remove user' cookie
func (s *Session) RemoveUser(r *http.Request) {
	s.sessionManager.Remove(r.Context(), userKey)
}

// GetCookieUser get cookie's user info
func (s *Session) GetCookieUser(token string) (user *models.UserAccount, err error) {
	ctx := context.TODO()
	ctx, err = s.sessionManager.Load(ctx, token)
	if err != nil {
		return nil, errors.Unauthenticated(fmt.Sprintf("failing to load user token: %v", err))
	}
	return s.GetUser(ctx)
}
