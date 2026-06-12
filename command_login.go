package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		err := fmt.Errorf("Login name required")
		return err
	}

	ctx := context.Background()

	if _, err := s.db.GetUser(ctx, cmd.Args[0]); err != nil {
		return err
	}

	err := s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("User set to: %v\n", cmd.Args[0])

	return nil
}
