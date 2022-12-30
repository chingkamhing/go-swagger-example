package session

import (
	"context"
	"encoding/gob"
	"fmt"
	"net/http"

	"go-swagger-example/gen/models"
)

// user's session key name
const userKey = "user"

func init() {
	// gob register models.User in order to support context decode
	gob.Register(&models.UserInfo{})
}

// SaveUser save user info
func (s *Session) SaveUser(user *models.UserInfo, r *http.Request) {
	s.sessionManager.Put(r.Context(), userKey, user)
}

// GetUser get user info
func (s *Session) GetUser(r *http.Request) (user *models.UserInfo, err error) {
	any := s.sessionManager.Get(r.Context(), userKey)
	user, ok := any.(*models.UserInfo)
	if !ok {
		return nil, fmt.Errorf("invalid user session")
	}
	return user, nil
}

// RemoveUser remove user' cookie
func (s *Session) RemoveUser(r *http.Request) {
	s.sessionManager.Remove(r.Context(), userKey)
}

// GetCookieUser get cookie's user info
func (s *Session) GetCookieUser(token string) (user *models.UserInfo, err error) {
	ctx := context.TODO()
	ctx, err = s.sessionManager.Load(ctx, token)
	if err != nil {
		return nil, err
	}
	any := s.sessionManager.Get(ctx, userKey)
	user, ok := any.(*models.UserInfo)
	if !ok {
		return nil, fmt.Errorf("invalid user session")
	}
	return user, nil
}
