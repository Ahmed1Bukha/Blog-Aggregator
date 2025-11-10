package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ahmed1bukha/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error{
	if len(cmd.Args) < 2{
		return fmt.Errorf("requires atleast 2 inputs")
	}
	current_user,err := s.db.GetUser(context.Background(),s.Cfg.CurrentUserName)
	if err !=nil{
		return err
	}
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
	fmt.Println(feed)
	return nil
}