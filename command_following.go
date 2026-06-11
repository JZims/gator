package main

import (
	"context"
	"fmt"
	"gator/internal/database"
)

func handlerFollowing(s *state, cmd command, u database.User) error {
	ctx := context.Background()

	feedsFollowing, err := s.db.GetFeedFollowsForUser(ctx, u.ID)
	if err != nil {
		return err
	}

	if len(feedsFollowing) == 0 {
		fmt.Printf("No follows for user: %v\n", s.cfg.CurrentUserName)
		return nil
	}

	fmt.Printf("Follows for User: %v\n", s.cfg.CurrentUserName)

	for i, feed := range feedsFollowing {
		fmt.Printf("%v: %v\n", i+1, feed.FeedName)
	}

	return nil

}
