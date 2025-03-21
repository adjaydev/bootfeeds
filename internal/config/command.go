package config

import "errors"

type Command struct {
	Name string
	Cmd  []string
}

type Commands struct {
	Commands map[string]func(*State, Command) error
}

// Register a new handler function for a command name.
func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Commands[name] = f
}

// Run a given command with the provided state if it exists.
func (c *Commands) Run(s *State, cmd Command) error {
	f, ok := c.Commands[cmd.Name]
	if !ok {
		return errors.New("Command not available")
	}
	return f(s, cmd)
}
