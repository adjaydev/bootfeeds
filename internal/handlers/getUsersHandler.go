package handlers

import (
	"bootfeeds/internal/config"
	"context"
	"fmt"
)

func GetUsersHandler(s *config.State, cmd config.Command) error {
	ctx := context.Background()

	users, err := s.DB.GetUsers(ctx)
	if err != nil {
		return fmt.Errorf("No users found.")
	}

	for _, u := range users {
		current := ""
		if s.Cfg.CurrentUserName == u.Name {
			current = " (current)"
		}
		fmt.Printf("* %s%s\n", u.Name, current)
	}

	return nil
}
