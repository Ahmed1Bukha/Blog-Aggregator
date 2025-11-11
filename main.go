package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/ahmed1bukha/Blog-Aggregator/internal/config"
	"github.com/ahmed1bukha/Blog-Aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct{
	Cfg *config.Config
	db  *database.Queries
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

	db, err := sql.Open("postgres", cfg.DBURL)
	if err !=nil{
		log.Fatalln(err.Error())
	}
	dbQueries := database.New(db)
	
	st := state{
		Cfg: cfg,
		db: dbQueries,
	}

	
	c:= commands{
		commands:make(map[string]func(*state, command) error),
	}
	c.register("login",handlerLogin)
	c.register("register",handlerRegister)
	c.register("reset",handlerReset)
	c.register("users",handlerUsers)
	c.register("agg",handlerAgg)
	c.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	c.register("feeds",handlerFeeds)
	c.register("follow",middlewareLoggedIn(handlerFollow))
	c.register("following",middlewareLoggedIn(handlerFollowing))
	c.register("unfollow",middlewareLoggedIn(handlerUnfollow))
	c.register("browse",middlewareLoggedIn(handlerBrowse))
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

