package main

import (
	"context"
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error{
	if len(cmd.Args) == 0{
		return fmt.Errorf("args is empty")
	}
	_,err := s.db.GetUser(context.Background(),cmd.Args[0])
	if err != nil{
		os.Exit(1)
		return err
	}
	err = s.Cfg.SetUser(cmd.Args[0])
	if err !=nil{
		return err
	}
	fmt.Printf("Set user is successful")
	return nil
}
