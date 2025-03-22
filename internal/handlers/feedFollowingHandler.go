package handlers

import (
	"bootfeeds/internal/config"
	"bootfeeds/internal/database"
	"context"
	"fmt"
)

func FeedFollowingHandler(s *config.State, cmd config.Command, u database.User) error {
	ctx := context.Background()

	feeds, err := s.DB.GetFeedFollowsForUser(ctx, u.ID)
	if err != nil {
		return fmt.Errorf("User does not follow any feeds.")
	}

	for _, f := range feeds {
		fmt.Printf("> %s %s\n", f.FeedName, f.UserName)
	}

	return nil
}
