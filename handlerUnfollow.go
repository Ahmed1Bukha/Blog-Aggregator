package main

import (
	"context"
	"fmt"

	"github.com/ahmed1bukha/Blog-Aggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error{
	if len(cmd.Args) <1 {
		return fmt.Errorf("should at least have 1 argument for url")
	}

	feed,err:= s.db.GetFeedByUrl(context.Background(),cmd.Args[0])
	if err !=nil{
		return err
	}
	unFollowParams:= database.UnFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	err= s.db.UnFollow(context.Background(),unFollowParams)
	if err !=nil{
		return err
	}
	fmt.Printf("Unfollowed is successful")
	return nil
}