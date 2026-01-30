package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("no username provided")
	}
	user := cmd.args[0]
	if err := s.cfg.SetUser(user); err != nil {
		return err
	}
	fmt.Printf("Username set to %v\n", user)
	return nil
}
