package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ahmed1bukha/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error{
	if len(cmd.Args) < 2{
		return fmt.Errorf("requires atleast 2 inputs")
	}
	current_user := user
	feedParams:= database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.Args[0],
		Url: cmd.Args[1],
		UserID: current_user.ID,
	}
	feed,err := s.db.CreateFeed(context.Background(),feedParams)
	if err !=nil{
		return err
	}

	feedFollowParams:= database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: current_user.ID,
		FeedID: feed.ID,
	}

	_,err= s.db.CreateFeedFollow(context.Background(),feedFollowParams)
	if err !=nil{
		return err
	}

	fmt.Println(feed)
	return nil
}