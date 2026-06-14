package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"strconv"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := "2"
	if len(cmd.Args) > 1 {
		limit = cmd.Args[0]
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil {
		return err
	}

	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limitNum),
	}

	posts, err := s.db.GetPostsForUser(context.Background(), params)
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf("%s\n", post.Description.String)
	}
	return nil
}
