package handlers

import (
	"bootfeeds/internal/config"
	"context"
	"fmt"
)

func LoginHandler(s *config.State, cmd config.Command) error {
	if len(cmd.Cmd) == 0 {
		return fmt.Errorf("Invalid arguments, need USERNAME")
	}

	ctx := context.Background()
	_, err := s.DB.GetUser(ctx, cmd.Cmd[0])
	if err != nil {
		return fmt.Errorf("Not a registered user.")
	}

	s.Cfg.SetUser(cmd.Cmd[0])
	return nil
}
