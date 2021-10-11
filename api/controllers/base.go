package controllers

import (
	"DevStories/api/models"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"os"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) InitializeServer() {
	var err error

	dbDriver := os.Getenv("DB_DRIVER")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	DBUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUsername, dbPassword, dbName)

	server.DB, err = gorm.Open(dbDriver, DBUrl)
	if err != nil {
		log.Fatal("Database connection error: ", err)
	} else {
		fmt.Println("Successfully connected to database")
	}

	server.DB.Debug().AutoMigrate(&models.User{})

	server.Router = mux.NewRouter()

	//server.InitializeRoutes()
}

func (server *Server) Run(address string) {
	fmt.Printf("Listening and serving on port: %s", address)
	log.Fatal(http.ListenAndServe(address, server.Router))
}
