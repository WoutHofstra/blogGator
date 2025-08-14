package main

import (
	"fmt"
)

type command struct {
	name		string
	arguments	[]string
}

type commands struct {
	cmdNames	map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {

	if _, ok := c.cmdNames[cmd.name]; ok {
		function := c.cmdNames[cmd.name]
		err := function(s, cmd)
		if err != nil {
			return fmt.Errorf("Error: failed to run command: %w", err)
		}
	} else {
		return fmt.Errorf("Command not found")
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error ) error {

	c.cmdNames[name] = f
	return nil
}
