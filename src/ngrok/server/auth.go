package server

import (
	"fmt"
	"ngrok/msg"
)

type Authenticator interface {
	Auth(auth *msg.Auth) error
}

type simpleAuthenticator struct {
	authorizedTokens []string
}

func (receiver *simpleAuthenticator) Auth(auth *msg.Auth) error {
	for _, token := range receiver.authorizedTokens {
		if auth.User == token {
			return nil
		}
	}

	return fmt.Errorf("Unauthorized token %s ", auth.User)
}

func CreateAuthenticator(tokens []string) *simpleAuthenticator {
	if len(tokens) == 0 {
		return nil
	}
	return &simpleAuthenticator{
		authorizedTokens: tokens,
	}
}
