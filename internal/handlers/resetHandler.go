package handlers

import (
	"bootfeeds/internal/config"
	"context"
	"fmt"
	"log"
)

func ResetHandler(s *config.State, cmd config.Command) error {
	ctx := context.Background()

	err := s.DB.DeleteUsers(ctx)
	if err != nil {
		return fmt.Errorf("DB Reset failed: %s", err)
	}

	log.Println("DB Reset successful")

	return nil
}
