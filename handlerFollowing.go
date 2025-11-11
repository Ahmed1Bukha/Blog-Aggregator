package main

import (
	"context"
	"fmt"

	"github.com/ahmed1bukha/Blog-Aggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error{
	feed_follow,err := s.db.GetFeedFollowsForUser(context.Background(),user.ID)
	if err !=nil{
		return err
	}
	fmt.Printf("Feed for user %v \n",user.Name)
	for _,feed := range feed_follow{
		
		fetchedFeed,err := s.db.GetFeedByID(context.Background(),feed.FeedID)
		
		if err !=nil{
			return nil
		}
		fmt.Println(fetchedFeed.Name)
	}
	return nil
}