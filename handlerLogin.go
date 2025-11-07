package main

import "fmt"

func handlerLogin(s *state, cmd command) error{
	if len(cmd.Args) == 0{
		return fmt.Errorf("args is empty")
	}
	err := s.Cfg.SetUser(cmd.Args[0])
	if err !=nil{
		return err
	}
	fmt.Printf("Set user is successful")
	return nil
}
