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

func handlerAddFeed(s *state, cmd command) error {
	args := cmd.Arguments
	if len(args) != 2 {
		return errors.New("Incorrect number of arguments")
	}
	name := args[0]
	url := args[1]
	currentUser, usrError := s.db.GetUser(context.Background(), s.Cfg.CurrentUserName)
	if usrError != nil {
		return usrError
	}

	feed, feedErr := s.db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      name,
			Url:       url,
			UserID:    currentUser.ID,
		},
	)
	if feedErr != nil {
		return feedErr
	}

	_, newFeedFollowErr := s.db.CreateFeedFollow(context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    currentUser.ID,
			FeedID:    feed.ID,
		},
	)
	if newFeedFollowErr != nil {
		return newFeedFollowErr
	}

	fmt.Println(feed)

	return nil
}

func handlerGetFeeds(s *state, cmd command) error {
	args := cmd.Arguments
	if len(args) != 0 {
		return errors.New("This command requires no arguments")
	}

	feeds, getError := s.db.GetFeeds(context.Background())
	if getError != nil {
		return getError
	}

	for _, feed := range feeds {
		feedUser, getUserErr := s.db.GetUserByID(context.Background(), feed.UserID)
		if getUserErr != nil {
			return getUserErr
		}
		fmt.Printf("Feed Name: %s\n", feed.Name)
		fmt.Printf("Feed URL: %s\n", feed.Url)
		fmt.Printf("Feed User: %s\n", feedUser.Name)
	}

	return nil
}

func handlerFollowFeeds(s *state, cmd command) error {
	args := cmd.Arguments
	if len(args) != 1 {
		return errors.New("This command requires a single argument (URL)")
	}

	currentUser, usrError := s.db.GetUser(context.Background(), s.Cfg.CurrentUserName)
	if usrError != nil {
		return usrError
	}

	currentFeed, feedError := s.db.GetFeedsByURL(context.Background(), args[0])
	if feedError != nil {
		return feedError
	}

	feeds, getError := s.db.CreateFeedFollow(context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    currentUser.ID,
			FeedID:    currentFeed.ID,
		},
	)
	if getError != nil {
		return getError
	}

	fmt.Printf("Feed: %s\nUser: %s\n", feeds.FeedName, currentUser.Name)

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	args := cmd.Arguments
	if len(args) != 0 {
		return errors.New("This command takes no arguments")
	}

	currentUser, usrError := s.db.GetUser(context.Background(), s.Cfg.CurrentUserName)
	if usrError != nil {
		return usrError
	}

	userFeeds, getError := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if getError != nil {
		return getError
	}

	for i := range userFeeds {
		fmt.Printf("Feed: %s\n", userFeeds[i].FeedName)
	}

	return nil
}
