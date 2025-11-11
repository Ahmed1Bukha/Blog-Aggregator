package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ahmed1bukha/Blog-Aggregator/internal/database"
)



func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		
		username := s.Cfg.CurrentUserName
		if username == "" {
			return fmt.Errorf("not logged in")
		}

		
		user, err := s.db.GetUser(context.Background(), username)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("user %q not found", username)
			}
			return err
		}
		return handler(s, cmd, user)
	}
}