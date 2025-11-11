package main

import (
	"context"
	"fmt"
)


func handlerReset(s *state, cmd command) error{
	err:= s.db.DeleteTableUser(context.Background())
	if err !=nil{
		return err
	}
	err= s.db.DeleteTableFeedsFollow(context.Background())
	if err !=nil{
		return err
	}
	err= s.db.DeleteTableFeeds(context.Background())
	if err !=nil{
		return err
	}
	fmt.Println("delete table successful")
	return nil
}