package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		err := errors.New("Two arguments required")
		return err
	}
	if len(cmd.Args) > 1 {
		err := errors.New("Too many arguments given (expected 1)")
		return err
	}

	err := s.Config.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Printf("User set to: %v\n", cmd.Args[0])

	return nil
}
