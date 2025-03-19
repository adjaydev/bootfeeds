package main

import (
	"bootfeeds/internal/config"
	"fmt"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	cmd  []string
}

type commands struct {
	commands map[string]func(*state, command) error
}

func (s *state) handlerLogin(cmd command) error {
	if len(cmd.cmd) == 0 {
		return fmt.Errorf("No command given.")
	}
	s.cfg.SetUser(cmd.name)
	fmt.Printf("User %s was set.", cmd.name)
	return nil
}

func main() {
	cfg := config.Read()
	cfg.SetUser("jay")
	fmt.Println(cfg.DbUrl)
}
