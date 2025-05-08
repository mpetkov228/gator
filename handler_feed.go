package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mpetkov228/gator/internal/database"
)

func handleAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed: %v", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed follow: %v", err)
	}

	fmt.Println("Feed created successfully!")
	printFeed(feed)
	fmt.Println()
	fmt.Println("============================")

	return nil
}

func handleFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error getting feeds: %v", err)
	}

	for i, feed := range feeds {
		fmt.Println()
		fmt.Printf("%d. %s - %s\n", i+1, feed.Name, feed.Url)
		fmt.Printf("by %s\n", feed.Username)
	}

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:             %s\n", feed.ID)
	fmt.Printf("* Created:        %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:        %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:           %s\n", feed.Name)
	fmt.Printf("* Url:            %s\n", feed.Url)
	fmt.Printf("* User ID:        %s\n", feed.UserID)
}
