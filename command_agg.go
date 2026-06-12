package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func scrapeFeeds(s *state) {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Printf("Error getting next feed: %v\n", err)
		return
	}
	markedFeed, err := s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		log.Printf("Error marking fetched feed: %v\n", err)
		return
	}
	fetchedFeed, err := fetchFeed(context.Background(), markedFeed.Url)
	if err != nil {
		log.Printf("Error fetching feed: %v\n", err)
		return
	}
	for _, item := range fetchedFeed.Channel.Item {
		log.Printf("Feed Fetched: %v\n", item.Title)
	}
}

func handlerAgg(s *state, cmd command) error {

	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[1])
	if err != nil {
		return err
	}

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
