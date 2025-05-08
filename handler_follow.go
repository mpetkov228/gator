package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mpetkov228/gator/internal/database"
)

func handleFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error getting feed: %v", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feedFollow: %v", err)
	}

	fmt.Printf("User: %s | Feed: %s\n", feedFollow.UserName, feedFollow.FeedName)

	return nil
}

func handleFollowing(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting user follows: %v", err)
	}

	fmt.Printf("Feeds followed by %s:\n", user.Name)
	for _, followed := range feedFollows {
		fmt.Printf(" - %s\n", followed.FeedName)
	}

	return nil
}
