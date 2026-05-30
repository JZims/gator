package main

import (
	"fmt"
	config "gator/internal"
	"gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	Name string
	Args []string
}

type commands struct {
	options map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	fn, ok := c.options[cmd.Name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}
	return fn(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) error {
	if _, exists := c.options[name]; exists {
		return fmt.Errorf("command %s already registered", name)
	}
	c.options[name] = f
	return nil
}
