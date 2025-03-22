package handlers

import (
	"bootfeeds/internal/config"
	"bootfeeds/internal/database"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

func RegisterHandler(s *config.State, cmd config.Command) error {
	if len(cmd.Cmd) == 0 {
		return fmt.Errorf("Invalid arguments, need USERNAME")
	}

	ctx := context.Background()

	_, err := s.DB.GetUser(ctx, cmd.Cmd[0])
	if err == nil {
		return fmt.Errorf("User already exists.")
	}

	s.DB.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Cmd[0],
	})

	s.Cfg.SetUser(cmd.Cmd[0])
	log.Printf("User %s created.", cmd.Cmd[0])
	log.Printf("User %s logged in.", cmd.Cmd[0])
	return nil
}
