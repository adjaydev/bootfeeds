package handlers

import (
	"bootfeeds/internal/config"
	"fmt"
	"log"
)

func LoginHandler(s *config.State, cmd config.Command) error {
	if len(cmd.Cmd) == 0 {
		log.Fatal("Argument for username is required.")
		return fmt.Errorf("Argument for username is required.")
	}

	s.Cfg.SetUser(cmd.Cmd[0])
	return nil
}
