package models

import (
	"DevStories/api/tokens"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

func (auth *Auth) SignInUser(db *gorm.DB, email, password string) (*Auth, error) {
	user := User{}

	err := db.Debug().Model(User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return &Auth{}, err
	}

	err = VerifyHashedPassword(password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return &Auth{}, err
	}

	generatedToken, err := tokens.CreateToken(user.ID)
	if err != nil {
		return &Auth{}, err
	}

	auth.User = user
	auth.Token = generatedToken

	return auth, nil
}
