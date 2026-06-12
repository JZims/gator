package main

import (
	"context"
	"database/sql"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, u database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("addfeed takes two arguments: addfeed <name_of_feed> <feed_url>")
	}

	ctx := context.Background()

	url := cmd.Args[1]

	now := time.Now()
	nullableTime := sql.NullTime{
		Time:  now,
		Valid: true,
	}

	newFeed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: nullableTime,
		UpdatedAt: nullableTime,
		Name:      cmd.Args[0],
		Url:       url,
		UserID:    u.ID,
	}

	feed, err := s.db.CreateFeed(ctx, newFeed)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", feed)

	newFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    u.ID,
		FeedID:    feed.ID,
	}

	_, err = s.db.CreateFeedFollow(ctx, newFollow)
	if err != nil {
		return err
	}

	return nil
}
