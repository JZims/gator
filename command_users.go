package main

import (
	"context"
	"fmt"
)

func handlerGetUsers(s *state, c command) error {
	ctx := context.Background()

	users, err := s.db.GetUsers(ctx)
	if err != nil {
		return err
	}

	for i := range users {
		if users[i].Name == s.cfg.CurrentUserName {
			fmt.Printf("*%v (current) \n", users[i].Name)
		} else {
			fmt.Printf("*%v\n", users[i].Name)
		}
	}

	return nil

}
