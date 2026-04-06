package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	args := cmd.Arguments
	if len(args) == 0 {
		return errors.New("No arguments found")
	}

	err := s.Cfg.SetUser(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Username set to %s\n", args[0])

	return nil
}
