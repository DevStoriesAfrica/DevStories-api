package models_test

import (
	"DevStories/api/tokens"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestSignInUser(t *testing.T) {
	seededUser, err := seedOneUser()
	if err != nil {
		log.Fatal(err)
	}

	authenticatedUser, err := authInstance.SignInUser(server.DB, seededUser.Email, seededUser.Password)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, authenticatedUser.User.ID, seededUser.ID, "IDs should match")
	assert.Equal(t, authenticatedUser.User.Username, seededUser.Username, "Usernames should match")
	assert.Equal(t, authenticatedUser.User.Email, seededUser.Email, "Emails should match")
	assert.Equal(t, authenticatedUser.User.Password, seededUser.Password, "Passwords should match")

	generatedToken, err := tokens.CreateToken(seededUser.ID)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, authenticatedUser.Token, generatedToken, "JWT Tokens should match")
}
