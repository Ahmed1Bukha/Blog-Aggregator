package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ahmed1bukha/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command,user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("args should at least contain 1 url")
	}
	
	feed,err := s.db.GetFeedByUrl(context.Background(),cmd.Args[0])
	if err !=nil{
		return err
	}
	feedFollowSchema := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
		
	}
	_,err = s.db.CreateFeedFollow(context.Background(),feedFollowSchema)
	if err !=nil{
		return err
	}
	fmt.Printf("Feed follow has been created with the following user feed names: userName:%v, feedName:%v ",user.Name,feed.Name)
	return nil
}