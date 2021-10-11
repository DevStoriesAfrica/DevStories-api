package controllers_tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignIn(t *testing.T) {
	err := dropUsersTable()
	if err != nil {
		log.Fatal(err)
	}

	seededUser, err := seedOneUser()
	if err != nil {
		log.Fatal(err)
	}

	samples := []struct {
		email        string
		password     string
		errorMessage string
	}{
		{
			email:        seededUser.Email,
			password:     "password",
			errorMessage: "",
		},
		{
			email:        seededUser.Email,
			password:     "Wrong password",
			errorMessage: "crypto/bcrypt: hashedPassword is not the hash of the given password",
		},
		{
			email:        "Wrong email",
			password:     "password",
			errorMessage: "record not found",
		},
	}

	for _, value := range samples {
		token, err := authInstance.SignInUser(server.DB, value.email, value.password)
		if err != nil {
			assert.Equal(t, err, errors.New(value.errorMessage))
		} else {
			assert.NotEqual(t, token, "")
		}
	}
}

func TestLogin(t *testing.T) {
	err := dropUsersTable()
	if err != nil {
		log.Fatal(err)
	}

	_, err = seedOneUser()
	if err != nil {
		fmt.Printf("This is the error %v\n", err)
	}

	samples := []struct {
		inputJSON    string
		statusCode   int
		email        string
		password     string
		errorMessage string
	}{
		{
			inputJSON:    `{"email": "testuser@gmail.com", "password": "test_password"}`,
			statusCode:   200,
			errorMessage: "",
		},
		{
			inputJSON:    `{"email": "testuser@gmail.com", "password": "wrong password"}`,
			statusCode:   422,
			errorMessage: "Incorrect Password",
		},
		{
			inputJSON:    `{"email": "frank@gmail.com", "password": "test_password"}`,
			statusCode:   422,
			errorMessage: "Incorrect Details",
		},
		{
			inputJSON:    `{"email": "testuser.com", "password": "password"}`,
			statusCode:   422,
			errorMessage: "Invalid Email",
		},
		{
			inputJSON:    `{"email": "", "password": "password"}`,
			statusCode:   422,
			errorMessage: "Required Email",
		},
		{
			inputJSON:    `{"email": "kan@gmail.com", "password": ""}`,
			statusCode:   422,
			errorMessage: "Required Password",
		},
		{
			inputJSON:    `{"email": "", "password": "password"}`,
			statusCode:   422,
			errorMessage: "Required Email",
		},
	}

	for _, value := range samples {
		req, err := http.NewRequest(http.MethodPost, "/user/login", bytes.NewBufferString(value.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v", err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.LoginUser)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, value.statusCode)
		if value.statusCode == http.StatusOK {
			assert.NotEqual(t, rr.Body.String(), "")
		}

		if value.statusCode == http.StatusInternalServerError && value.errorMessage != "" {
			responseMap := make(map[string]interface{})
			err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
			if err != nil {
				t.Errorf("Cannot convert to json: %v", err)
			}
			assert.Equal(t, responseMap["error"], value.errorMessage)
		}
	}
}
