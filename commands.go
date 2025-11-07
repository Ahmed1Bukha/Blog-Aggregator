package main

import "fmt"

type commands struct{
	commands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error{
	handler,ok := c.commands[cmd.Name]
	if !ok{
		return fmt.Errorf("this command doesn't exit in the map")
	}
	err := handler(s,cmd)
	if err !=nil{
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error){
	 c.commands[name] = f
}