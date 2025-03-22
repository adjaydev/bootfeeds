package handlers

import (
	"bootfeeds/internal/config"
	"context"
	"fmt"
)

func FeedsHandler(s *config.State, cmd config.Command) error {
	ctx := context.Background()

	feeds, err := s.DB.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("Error getting feeds: %s", err)
	}

	for _, f := range feeds {
		fmt.Println(f.Name, f.Url, f.Username)
	}

	return nil
}
