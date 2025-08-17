package main

import (
	"os"
	"time"
	"context"
        "github.com/WoutHofstra/blogGator/internal/database"
	"fmt"
        "github.com/google/uuid"
)


func handlerFeed(s *state, cmd command) error {



        ctx := context.Background()
	createdAt := time.Now()
	updatedAt := time.Now()
        feedName := cmd.arguments[0]
        feedURL := cmd.arguments[1]

        user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)

	feedUserID := uuid.NullUUID{UUID: user.ID, Valid: true}


	params:= database.CreateFeedParams {
		CreatedAt:	createdAt,
		UpdatedAt:	updatedAt,
		Name:		feedName,
		Url:		feedURL,
		UserID:		feedUserID,
	}

	feed, err := s.db.CreateFeed(ctx, params)
	if err != nil {
		fmt.Printf("Couldnt create feed: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Feed successfully made! %+v", feed)
	return nil
}


func handlerGetFeeds(s *state, cmd command) error {

	ctx := context.Background()

	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("Error: Get feeds failed: %w", err)
	}
	for _, f := range feeds {
		username, err := s.db.GetUserFromID(ctx, f.UserID.UUID)
		if err != nil {
			return fmt.Errorf("Error: Failed to get username from feed: %w", err)
		}

		fmt.Println(f.Name)
		fmt.Println(f.Url)
		fmt.Println(username)
	}

	return nil



}

