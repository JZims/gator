package main

import (
	"context"
	"fmt"
	"gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, u database.User) error {
	ctx := context.Background()

	url := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(ctx, url)
	if err != nil {
		return err
	}

	params := database.DeleteFeedFollowParams{
		UserID: u.ID,
		FeedID: feed.ID,
	}

	err = s.db.DeleteFeedFollow(ctx, params)
	if err != nil {
		return err
	}

	fmt.Printf("User: %v Unfollowed Feed: %v (%v)", u.Name, feed.Name, feed.Url)

	return nil
}
