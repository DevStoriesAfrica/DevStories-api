package models_test

import (
	"DevStories/api/controllers"
	"DevStories/api/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

var server = controllers.Server{}
var userInstance = models.User{}
var err error

func TestMain(m *testing.M) {
	err := godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting enviromental variable: %v\n", err)
	}

	Database()

	os.Exit(m.Run())
}

func Database() {
	dbDriver := os.Getenv("TEST_DB_DRIVER")
	dbHost := os.Getenv("TEST_DB_HOST")
	dbPort := os.Getenv("TEST_DB_PORT")
	dbUser := os.Getenv("TEST_DB_USER")
	dbPassword := os.Getenv("TEST_DB_PASSWORD")
	dbName := os.Getenv("TEST_DB_NAME")

	DBUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	server.DB, err = gorm.Open(dbDriver, DBUrl)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	} else {
		log.Println("Successfully connected to database")
	}
}

func dropUsersTable() error {
	err = server.DB.DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}

	err = server.DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}

	log.Println("Dropped user table")
	return nil
}

func seedOneUser() (models.User, error) {
	err:=dropUsersTable()
	if err!=nil{
		log.Fatal(err)
	}

	user := models.User{
		ID:       1,
		Username: "Test user",
		Email:    "testuser@gmail.com",
		Password: "test_password",
	}

	err = server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("Failed to seed one user: %v", err)
	}

	return user, nil
}

func seedMultipleUsers() ([]models.User, error) {
	err:=dropUsersTable()
	if err!=nil{
		log.Fatal(err)
	}

	users := []models.User{
		{ID: 1,
			Username: "Test user",
			Email:    "testuser@gmail.com",
			Password: "test_password",
		},
		{
			ID:       1,
			Username: "Test user 2",
			Email:    "testuser2@gmail.com",
			Password: "test_password2",
		},
	}

	for key := range users {
		err := server.DB.Model(&models.User{}).Create(&users[key]).Error
		if err != nil {
			log.Fatalf("Failed to seed multiple users: %v", err)
		}
	}

	return users, nil
}
