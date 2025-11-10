package main

import (
	"context"
	"fmt"
)


func handlerFeeds(s *state, cmd command) error{
	feeds,err := s.db.GetFeeds(context.Background())
	if err !=nil{
		return err
	}
	for index,feed := range feeds{
		user,err:= s.db.GetUserByID(context.Background(),feed.UserID)
		if err !=nil{
			return err
			
		}
		fmt.Printf("%v: Feed Name:%v, URL:%v, user: %v\n",index,feed.Name,feed.Url,user.Name)
	}
	return nil
}