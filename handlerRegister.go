package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/ahmed1bukha/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error{
	if len(cmd.Args) == 0{
		return fmt.Errorf("args is empty")
	}
	_,err := s.db.GetUser(context.Background(),cmd.Args[0])
	if err ==nil{
		os.Exit(1)
		return fmt.Errorf("user already exits")
	}

	userSchema := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.Args[0],
	}
	user,err:=s.db.CreateUser(context.Background(),userSchema)
	if err !=nil{
		return err
	}
	fmt.Println("user created successfully")
	s.Cfg.SetUser(user.Name)
	return nil
}