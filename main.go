package main

import (
	"fmt"
	"log"

	"github.com/ahmed1bukha/Blog-Aggregator/internal/config"
)

func main (){
	cfg , err:= config.Read()
	if err !=nil {
		log.Fatalln(err)
	}
	fmt.Println(cfg.DBURL)
	// err =cfg.SetUser("bukha")
	// if err !=nil{
	// 	fmt.Println(err.Error())
	// } else{
	// 	fmt.Println("Success wrote a new user")
	// }
	
	
}
