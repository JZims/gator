package main

import (
	"context"
	"database/sql"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("you must provide a name to be registered")
	}
	userName := cmd.Args[0]
	ctx := context.Background()
	now := time.Now()
	nullableTime := sql.NullTime{
		Time:  now,
		Valid: true,
	}

	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: nullableTime,
		UpdatedAt: nullableTime,
		Name:      userName,
	}
	_, err := s.db.CreateUser(ctx, newUser)
	if err != nil {
		return err
	}

	if err := s.cfg.SetUser(userName); err != nil {
		return err
	}

	fmt.Printf("New User: %v created\n", userName)

	return nil
}
