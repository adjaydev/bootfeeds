package handlers

import (
	"bootfeeds/internal/config"
	"bootfeeds/internal/database"
	"context"
	"fmt"
	"strconv"
)

func BrowseHandler(s *config.State, cmd config.Command, u database.User) error {
	var limit int
	limit = 2

	if len(cmd.Cmd) == 1 {
		var err error
		limit, err = strconv.Atoi(cmd.Cmd[0])
		if err != nil {
			return fmt.Errorf("Invalid interger [%s]: %s", cmd.Cmd[0], err)
		}
	}

	ctx := context.Background()
	posts, err := s.DB.GetPostsForUser(ctx, database.GetPostsForUserParams{
		Name:  u.Name,
		Limit: int32(limit),
	})
	if err != nil {
		return fmt.Errorf("Error fetching posts for user: %s", err)
	}

	for _, p := range posts {
		fmt.Printf("Title: %s\n", p.Title)
		fmt.Printf("Description: %s\n", p.Description)
		fmt.Printf("URL: %s\n", p.Url)
		fmt.Printf("Published at: %s\n\n", p.PublishedAt)
	}

	return nil
}
