package main

import "errors"

type command struct {
	name string
	args []string
}

type commands struct {
	names map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.names[name] = f
}
func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.names[cmd.name]
	if !ok {
		return errors.New("command not found")
	}
	return handler(s, cmd)
}
