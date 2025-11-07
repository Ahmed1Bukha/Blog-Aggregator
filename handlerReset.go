package main

import (
	"context"
	"fmt"
)


func handlerReset(s *state, cmd command) error{
	err:= s.db.DeleteTable(context.Background())
	if err !=nil{
		return err
	}
	fmt.Println("delete table successful")
	return nil
}