package main

import (
	"context"
	"fmt"
)

func handleGetUsers(s *state, c command) error {
	ctx := context.Background()

	users, err := s.db.GetUsers(ctx)

	if err != nil {
		return err
	}

	for i := 0; i < len(users); i++ {
		if users[i].Name == s.cfg.CurrentUserName {
			fmt.Printf("*%v (current) \n", users[i].Name)
		} else {
			fmt.Printf("*%v\n", users[i].Name)
		}
	}

	return nil

}
