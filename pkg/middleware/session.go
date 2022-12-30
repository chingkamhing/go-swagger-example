package middleware

import (
	"fmt"
	"strings"

	"go-swagger-example/pkg/session"
)

// CookieTokenFunc is a SessionSecurityAuth callback function that act as middleware that get user info base on the token
// note: this might double the effort to get cookie's info with sessionUser.SessionMiddleware
func CookieTokenFunc(sessionUser *session.Session) func(token string) (interface{}, error) {
	return func(token string) (interface{}, error) {
		tokens := parseTokens(token)
		myCookiePrefix := fmt.Sprintf("%s=", sessionUser.GetCookieName())
		for _, t := range tokens {
			if strings.HasPrefix(t, myCookiePrefix) {
				principal, err := sessionUser.GetCookieUser(strings.TrimPrefix(t, myCookiePrefix))
				if err != nil {
					sessionUser.Log.Debugf("Fail to get cookie user: %v", err)
					return nil, err
				}
				return principal, nil
			}
		}
		return nil, fmt.Errorf("invalid cookie name")
	}
}

// token string might contain more than one tokens that seperated by ";"; parse the raw token string to array of token string
func parseTokens(token string) []string {
	var tokens []string
	for _, t := range strings.Split(token, ";") {
		tokens = append(tokens, strings.TrimSpace(t))
	}
	return tokens
}
