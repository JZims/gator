package main

import (
	"context"
	"gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, c command, user database.User) error) func(*state, command) error {

	return func(s *state, c command) error {
		ctx := context.Background()
		user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, c, user)
	}
}
