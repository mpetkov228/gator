package main

import (
	"fmt"
	"time"
)

func handleAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.Name)
	}

	duration, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error parsing time arg: %v", err)
	}

	ticker := time.NewTicker(duration)

	fmt.Printf("Collecting feeds every %vm%vs\n", duration.Minutes(), duration.Seconds())

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
