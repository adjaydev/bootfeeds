package handlers

import (
	"bootfeeds/internal/config"
	"bootfeeds/internal/rss"
	"context"
	"fmt"
)

func FeedHandler(s *config.State, cmd config.Command) error {
	url := "https://www.wagslane.dev/index.xml"
	ctx := context.Background()
	feed, err := rss.FetchFeed(ctx, url)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	fmt.Print(feed)

	return nil
}
