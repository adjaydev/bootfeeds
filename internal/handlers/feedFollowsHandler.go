package handlers

import (
	"bootfeeds/internal/config"
	"bootfeeds/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func FeedFollowsHandler(s *config.State, cmd config.Command, u database.User) error {
	ctx := context.Background()

	feed, err := s.DB.GetFeed(ctx, cmd.Cmd[0])
	if err != nil {
		return fmt.Errorf("Feed URL does not exist.")
	}

	follow, err := s.DB.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    u.ID,
		FeedID:    feed.ID,
	})

	fmt.Println(follow)

	return nil
}
