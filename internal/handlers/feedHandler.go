package handlers

import (
	"bootfeeds/internal/config"
	"fmt"
	"log"
	"time"
)

func FeedHandler(s *config.State, cmd config.Command) error {
	if len(cmd.Cmd[0]) < 1 {
		return fmt.Errorf("Invalid arguments. Need 1s/1m/1h.")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Cmd[0])
	if err != nil {
		return fmt.Errorf("Error parsing time_betwee_requests")
	}

	log.Printf("Collecting feeds every %s\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		err := ScrapeFeedsHandler(s)
		if err != nil {
			return fmt.Errorf("Error scraping feeds %s\n", err)
		}
	}

}
