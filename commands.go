package main

import (
	"errors"

	"github.com/cybergrim/gator-go/internal/config"
	"github.com/cybergrim/gator-go/internal/database"
)

type state struct {
	db  *database.Queries
	Cfg *config.Config
}

type command struct {
	Name      string
	Arguments []string
}

type commands struct {
	Cmd map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	val, ok := c.Cmd[cmd.Name]
	if ok == false {
		return errors.New("Command not found")
	}
	err := val(s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.Cmd[name] = f
}
