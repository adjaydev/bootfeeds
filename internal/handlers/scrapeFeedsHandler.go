package handlers

import (
	"bootfeeds/internal/config"
	"bootfeeds/internal/database"
	"bootfeeds/internal/parser"
	"bootfeeds/internal/rss"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func ScrapeFeedsHandler(s *config.State) error {
	ctx := context.Background()

	next, err := s.DB.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("Error fetching next feed.")
	}

	fmt.Printf("\n\n[ %s ]\n[ %s ]\n\n", next.Name, next.Url)

	s.DB.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{
		ID:        next.ID,
		UpdatedAt: time.Now(),
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})

	feed, err := rss.FetchFeed(ctx, next.Url)
	if err != nil {
		return fmt.Errorf("Error fetching feed content: %s", err)
	}

	items := feed.Channel.Item
	for _, item := range items {

		pubDate, err := parser.ParseTimestamp(item.PubDate)
		if err != nil {
			return fmt.Errorf("Error parsing publish date (%s): %s", item.PubDate, err)
		}

		s.DB.AddPost(ctx, database.AddPostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: pubDate,
			FeedID:      next.ID,
		})

	}

	fmt.Println()

	return nil
}
