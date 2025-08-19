package main

import(
	"fmt"
        "context"
        "github.com/WoutHofstra/blogGator/internal/database"
)


func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {


	return func(s *state, cmd command) error {

	        ctx := context.Background()
	        user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	        if err != nil {
	                fmt.Println("Error:", err)
	                return err
	        }

		err = handler(s, cmd, user)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}


		return nil

	}

}
