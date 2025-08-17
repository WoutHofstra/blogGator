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
	"sort"
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

	files, err := os.ReadDir("sql/schema")
	if err != nil {
		return fmt.Errorf("Error: read directory failed: %w", err)
	}

	var filenames []string
	for _, f := range files {
		if !f.IsDir() {
			filenames = append(filenames, f.Name())
		}
	}
	sort.Strings(filenames)

	for _, fname := range filenames {
		contents, err := os.ReadFile("sql/schema/" + fname)
		if err != nil {
			return fmt.Errorf("Error: Couldnt read file: %w, %v", err, fname)
		}

		queries := strings.Split(string(contents), ";")
		for _, query := range queries {
			if query != "" {
				_, err := s.dbConn.ExecContext(ctx, query)
				if err != nil {
					return fmt.Errorf("Error running query: %w\n%v", err, query)
				}
			}
		}
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
