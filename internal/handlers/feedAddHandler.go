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

func FeedAddHandler(s *config.State, cmd config.Command, u database.User) error {
	if len(cmd.Cmd) < 2 {
		return fmt.Errorf("Invalid arguments, need USERNAME NAME URL")
	}

	name := cmd.Cmd[0]
	url := cmd.Cmd[1]

	ctx := context.Background()

	f, err := s.DB.CreateFeed(ctx, database.CreateFeedParams{
		ID:        uuid.New(),
		UserID:    u.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
	})
	if err != nil {
		return fmt.Errorf("Error creating feed.")
	}

	_, err = s.DB.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    u.ID,
		FeedID:    f.ID,
	})
	if err != nil {
		return fmt.Errorf("Error creating feed follow")
	}

	log.Println(f)

	return nil
}
