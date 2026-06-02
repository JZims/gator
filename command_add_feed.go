package main

import (
	"context"
	"database/sql"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("addfeed takes two arguments: addfeed <name_of_feed> <feed_url>")
	}

	ctx := context.Background()
	user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	url := cmd.Args[1]
	nullableUrl := sql.NullString{
		String: url,
		Valid:  true,
	}

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
		Url:       nullableUrl,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(ctx, newFeed)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", feed)

	return nil
}
