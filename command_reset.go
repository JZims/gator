package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, c command) error {
	ctx := context.Background()

	if err := s.db.DeleteUser(ctx); err != nil {
		return err
	}
	fmt.Println("Users successfully reset")
	return nil
}
