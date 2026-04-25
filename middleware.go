package main

import (
	"context"

	"github.com/cybergrim/gator-go/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		currentUser, getError := s.db.GetUser(context.Background(), s.Cfg.CurrentUserName)
		if getError != nil {
			return getError
		}
		return handler(s, cmd, currentUser)
	}
}
