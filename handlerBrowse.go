package main

import (
	"context"
	"fmt"

	"github.com/ahmed1bukha/Blog-Aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error{
	browsedFeeds,err := s.db.GetFeedFollowsForUser(context.Background(),user.ID)
	if err !=nil{
		return err
	}
	for _,feed:= range browsedFeeds{
		fmt.Println(feed)
	}
	return nil
}