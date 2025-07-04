package main

import (
	"context"
	"fmt"

	"github.com/mpetkov228/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("error getting user: %v", err)
		}

		return handler(s, cmd, user)
	}
}
