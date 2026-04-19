package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cybergrim/gator-go/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	args := cmd.Arguments
	if len(args) == 0 {
		return errors.New("No arguments found")
	}

	_, usrError := s.db.GetUser(context.Background(), args[0])
	if usrError != nil {
		return usrError
	}

	err := s.Cfg.SetUser(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Username set to %s\n", args[0])

	return nil
}

func handlerRegister(s *state, cmd command) error {
	args := cmd.Arguments
	if len(args) == 0 {
		return errors.New("No arguments found")
	}

	usr, creationError := s.db.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      args[0],
		},
	)

	if creationError != nil {
		return creationError
	}

	err := s.Cfg.SetUser(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("User %s created.\nUser Data:%s\n", args[0], usr)

	return nil
}

func handlerReset(s *state, cmd command) error {
	resetError := s.db.Reset(context.Background())
	if resetError != nil {
		return resetError
	}

	fmt.Printf("The %s database has been reset\n", s.Cfg.DBUrl)

	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	allUsers, getUserError := s.db.GetUsers(context.Background())
	if getUserError != nil {
		return getUserError
	}

	for i := range allUsers {
		if allUsers[i].Name == s.Cfg.CurrentUserName {
			fmt.Printf("%s (current)", allUsers[i].Name)
		} else {
			fmt.Println(allUsers[i].Name)
		}

	}

	return nil
}

func handlerAgg(s *state, cmd command) error {
	hardcodedURLFeed := "https://www.wagslane.dev/index.xml"
	feed, feedErr := fetchFeed(context.Background(), hardcodedURLFeed)
	if feedErr != nil {
		return feedErr
	}
	fmt.Printf("%+v\n", feed)

	return nil
}
