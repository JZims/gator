package main

import (
	"context"
	"database/sql"
	"fmt"
	"gator/internal/database"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
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

		title := sql.NullString{
			String: item.Title,
			Valid:  true,
		}

		desc := sql.NullString{
			String: item.Description,
			Valid:  true,
		}

		var pubAt sql.NullTime

		pubAtTime, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err == nil {
			pubAt = sql.NullTime{
				Time:  pubAtTime,
				Valid: true,
			}
		}

		post := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       title,
			Url:         item.Link,
			Description: desc,
			PublishedAt: pubAt,
			FeedID:      markedFeed.ID,
		}

		_, err = s.db.CreatePost(context.Background(), post)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Printf("Error creating post: %v\n", err)
			continue

		}
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
