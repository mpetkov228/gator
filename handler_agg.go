package main

import (
	"context"
	"fmt"
)

func handleAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("error fetching rss feed: %v", err)
	}

	fmt.Printf("Feed: %+v\n", feed)
	return nil
}
