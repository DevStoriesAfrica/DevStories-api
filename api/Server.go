package api

import (
	"DevStories/api/controllers"
	"fmt"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}
var err error

func Run() {
	err = godotenv.Load()
	if err != nil {
		fmt.Printf("Error fetching environmental variable: %s\n", err)
	} else {
		fmt.Println("Successfully fetched environmental variables")
	}

	server.InitializeServer()
	server.Run(":8081") //ToDo: Dynamically assign port by server
}
