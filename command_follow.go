package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, u database.User) error {
	ctx := context.Background()
	url := cmd.Args[0]
	now := time.Now()

	matchedFeed, err := s.db.GetFeedByUrl(ctx, url)
	if err != nil {
		return err
	}

	newFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    u.ID,
		FeedID:    matchedFeed.ID,
	}

	follow, err := s.db.CreateFeedFollow(ctx, newFollow)
	if err != nil {
		return err
	}

	fmt.Printf("User: %v\n", follow.UserName)
	fmt.Printf("Feed Followed: %v\n", follow.FeedName)

	return nil
}
