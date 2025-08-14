package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {

	if len(cmd.arguments) == 0 {
		return fmt.Errorf("A username is required")
	}

	username := cmd.arguments[0]
	err := s.config.SetUser(username)
	if err != nil {
		return fmt.Errorf("Set username failed: %w", err)
	}

	fmt.Println("Username successfully set!")
	return nil

}
