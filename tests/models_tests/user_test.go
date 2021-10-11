package models_test

import (
	"DevStories/api/models"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestSaveUser(t *testing.T) {
	err := dropUsersTable()
	if err != nil {
		log.Fatal(err)
		return
	}

	seededUser := models.User{
		ID:       1,
		Username: "Test User",
		Email:    "testuser@gmail.com",
		Password: "Testpassword13",
	}

	savedUser, err := seededUser.SaveUser(server.DB)
	if err != nil {
		log.Fatal(err)
		return
	}

	assert.Equal(t, seededUser.ID, savedUser.ID, "IDs should match")
	assert.Equal(t, seededUser.Username, savedUser.Username, "Usernames should match")
	assert.Equal(t, seededUser.Email, savedUser.Email, "Emails should match")
	assert.Equal(t, seededUser.Password, savedUser.Password, "Passwords should match")
}

func TestGetUser(t *testing.T) {
	seededUser, err := seedOneUser()
	if err != nil {
		log.Fatal(err)
		return
	}

	fetchedUser, err := userInstance.GetUser(server.DB, seededUser.ID)
	if err != nil {
		log.Fatal(err)
		return
	}

	assert.Equal(t, seededUser.ID, fetchedUser.ID, "IDs should match")
	assert.Equal(t, seededUser.Username, fetchedUser.Username, "Usernames should match")
	assert.Equal(t, seededUser.Email, fetchedUser.Email, "Emails should match")
	assert.Equal(t, seededUser.Password, fetchedUser.Password, "Passwords should match")
}

func TestUpdateUser(t *testing.T) {
	seededUser, err := seedOneUser()
	if err != nil {
		log.Fatal(err)
		return
	}

	seededUserUpdate := models.User{
		Username: "usernameupdate",
		Email:    "emailupdate",
		Password: "passwordupdate",
	}

	updatedUser, err := seededUserUpdate.UpdateUser(server.DB, seededUser.ID)
	if err != nil {
		log.Fatal(err)
		return
	}

	assert.Equal(t, seededUser.ID, updatedUser.ID, "IDs should match")
	assert.Equal(t, seededUserUpdate.Username, updatedUser.Username, "Usernames should match")
	assert.Equal(t, seededUserUpdate.Email, updatedUser.Email, "Emails should match")
	assert.Equal(t, seededUserUpdate.Password, updatedUser.Password, "Passwords should match")
}

func TestDeleteUser(t *testing.T) {
	seededUser, err := seedOneUser()
	if err != nil {
		log.Fatal(err)
		return
	}

	rowsAffected, err := seededUser.DeleteUser(server.DB, seededUser.ID)
	if err != nil {
		log.Fatal(err)
		return
	}

	assert.Equal(t, int64(1), rowsAffected, "Should indicate one row deleted")

	//_, err = userInstance.GetUser(server.DB, seededUser.ID)
	//assert.ErrorAs(t, err, gorm.IsRecordNotFoundError(err), "Should not find any of deleted user records")
}
