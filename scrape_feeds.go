package main

import (
	"context"
	"fmt"
)

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error returning next feed: %v", err)
	}

	feed, err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("error marking feed: %v", err)
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	fmt.Printf("Feed: %s\n", rssFeed.Channel.Title)
	for _, item := range rssFeed.Channel.Item {
		fmt.Printf(" - %s\n", item.Title)
	}

	return nil
}
