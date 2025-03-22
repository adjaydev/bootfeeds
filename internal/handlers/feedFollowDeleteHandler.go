package handlers

import (
	"bootfeeds/internal/config"
	"bootfeeds/internal/database"
	"context"
	"fmt"
)

func FeedFollowDeleteHandler(s *config.State, cmd config.Command, u database.User) error {
	ctx := context.Background()

	err := s.DB.DeleteFeedFollow(ctx, database.DeleteFeedFollowParams{
		UserID: u.ID,
		Url:    cmd.Cmd[0],
	})
	if err != nil {
		return fmt.Errorf("Error deleting feed follow: %s", err)
	}

	fmt.Printf("%s unfollowed.\n", cmd.Cmd[0])

	return nil
}
