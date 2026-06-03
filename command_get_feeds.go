package main

import (
	"context"
	"fmt"
)

func handlerGetFeeds(s *state, c command) error {
	ctx := context.Background()

	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return err
	}

	for f := range feeds {
		user, err := s.db.GetUserById(ctx, feeds[f].UserID)
		if err != nil {
			return err
		}
		fmt.Printf("Title: %v\n", feeds[f].Name)
		fmt.Printf("URL: %v\n", feeds[f].Url)
		fmt.Printf("User: %v\n", user.Name)
	}

	return nil
}
