package main

import (
	"log"
	"os"

	"github.com/ahmed1bukha/Blog-Aggregator/internal/config"
)

type state struct{
	Cfg *config.Config
}
type command struct{
	Name string
	Args []string
}

func main(){
	cfg , err:= config.Read()
	if err !=nil {
		log.Fatalln(err)
		os.Exit(0)
	}
	st := state{
		Cfg: cfg,
	}
	
	c:= commands{
		commands:make(map[string]func(*state, command) error),
	}
	c.register("login",handlerLogin)
	args:= os.Args
	if len(args)==1{
		log.Fatal("no command has been entered")
		os.Exit(0)
	}
	
	cmd:= command{
		Name: args[1],
		Args: args[2:],
	}
	err = c.run(&st,cmd)
	if err !=nil{
		log.Fatal(err.Error())
		os.Exit(0)
	}
	
}

