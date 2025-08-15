package main

import (
	"fmt"
	"time"
	"github.com/google/uuid"
	"strings"
	"os"
	"database/sql"
	"errors"
	"context"
	"github.com/WoutHofstra/blogGator/internal/database"
)

func handlerLogin(s *state, cmd command) error {

	if len(cmd.arguments) == 0 {
		return fmt.Errorf("A username is required")
	}

	ctx := context.Background()
	username := cmd.arguments[0]

	_, err := s.db.GetUser(ctx, username)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			os.Exit(1)
		} else {
			return fmt.Errorf("Error: Get user failed: %w", err)
		}
	}

	err = s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("Set username failed: %w", err)
	}

	fmt.Println("Username successfully set!")
	return nil

}

func handlerRegister(s *state, cmd command) error {


	if len(cmd.arguments) == 0 {
                return fmt.Errorf("A username is required")
	}

	ctx := context.Background()
	id := uuid.New()
	createdAt := time.Now()
	updatedAt := time.Now()
	username := cmd.arguments[0]

	params := database.CreateUserParams{
		ID:		id,
		CreatedAt:	createdAt,
		UpdatedAt:	updatedAt,
		Name:		username,
	}

	u, err := s.db.CreateUser(ctx, params)


	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			os.Exit(1)
		} else {
			return fmt.Errorf("Error: User creation failed: %w", err)
		}
	}

        err = s.cfg.SetUser(username)
        if err != nil {
                return fmt.Errorf("Set username failed: %w", err)
        }

	fmt.Printf("Created user: %+v\n", u)
        fmt.Printf("User created successfully!!\n")
	return nil


}

func handlerReset(s *state, cmd command) error {

        ctx := context.Background()

	err := s.db.ClearDatabase(ctx)
	if err != nil {
		os.Exit(1)
	}


	fmt.Println("Database reset!!")
	return nil
}

func handlerUsers(s *state, cmd command) error {

        ctx := context.Background()

        users, err := s.db.GetUsers(ctx)
	if err != nil {
		return fmt.Errorf("Error: Get users failed: %w", err)
	}
	for _, u := range users {
		if u == s.cfg.CurrentUserName {
			fmt.Printf("* %v (current)\n", u)
		} else {
			fmt.Printf("* %v\n", u)
		}
	}
	return nil
}
