package main

import (
	"os"
	"time"
	"context"
        "github.com/WoutHofstra/blogGator/internal/database"
	"fmt"
        "github.com/google/uuid"
)


func handlerFeed(s *state, cmd command, user database.User) error {



        ctx := context.Background()
	now := time.Now()
        feedName := cmd.arguments[0]
        feedURL := cmd.arguments[1]
	feedUserID := uuid.NullUUID{UUID: user.ID, Valid: true}

	params:= database.CreateFeedParams {
		CreatedAt:	now,
		UpdatedAt:	now,
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

	row, err := s.db.GetFeedFromUrl(ctx, feedURL)
	if err != nil {
                fmt.Println("Error:", err)
                return err
        }



	feedID := uuid.NullUUID{UUID: row.ID, Valid: true}

	followParams := database.CreateFeedFollowParams {
                CreatedAt:      now,
                UpdatedAt:      now,
		UserID:		feedUserID,
		FeedID:		feedID,
	}


	_, err = s.db.CreateFeedFollow(ctx, followParams)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}


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


func handlerFollow(s *state, cmd command, user database.User) error {


	feedUrl := cmd.arguments[0]

	ctx := context.Background()
	feed, _ := s.db.GetFeedFromUrl(ctx, feedUrl)
	now := time.Now()
	userID := uuid.NullUUID{UUID: user.ID, Valid: true}
	feedID := uuid.NullUUID{UUID: feed.ID, Valid: true}

        params := database.CreateFeedFollowParams {
                CreatedAt:      now,
                UpdatedAt:      now,
                UserID:         userID,
                FeedID:         feedID,
        }

	s.db.CreateFeedFollow(ctx, params)

	fmt.Println(s.cfg.CurrentUserName)
	fmt.Println(feed.Name)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {

	ctx := context.Background()
        userID := uuid.NullUUID{UUID: user.ID, Valid: true}

	res, err := s.db.GetFeedFollowsForUser(ctx, userID)
        if err != nil {
                fmt.Println("Error:", err)
                return err
        }
        fmt.Println(len(res))

	for _, f := range res {
		fmt.Printf("Feed: %v, User: %v\n", f.FeedName, f.UserName)
	}
	return nil

}


func handlerUnfollow(s *state, cmd command, user database.User) error {

	ctx := context.Background()
	feedUrl := cmd.arguments[0]
	feed, err := s.db.GetFeedFromUrl(ctx, feedUrl)
	if err != nil {
                fmt.Println("Error:", err)
                return err
        }

        userID := uuid.NullUUID{UUID: user.ID, Valid: true}
	feedID := uuid.NullUUID{UUID: feed.ID, Valid: true}

	params := database.UnfollowParams {
                UserID:         userID,
                FeedID:         feedID,
	}

	err = s.db.Unfollow(ctx, params)
	if err != nil {
                fmt.Println("Error:", err)
                return err
        }
	return nil
}
