package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error{
	users,err:= s.db.GetUsers(context.Background())
	if err !=nil{
		return err
	}
	for _,name:= range users{
		if name.Name == s.Cfg.CurrentUserName {
			fmt.Printf("* %s (current)",name.Name)
		} else{
			fmt.Printf("* %s",name.Name)
		}
	}
	return nil
}