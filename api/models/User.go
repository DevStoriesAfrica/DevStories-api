package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID       uint32 `gorm:"primary_key;auto_increment;unique" json:"id"`
	Username string `gorm:"size:255;unique;not null" json:"username"`
	Email    string `gorm:"size:100;unique;not null" json:"email"`
	Password string `gorm:"size:100;not null" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (user *User) HashPassword(password string)(string, error){
	hashedPasswordByte,err:=bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPasswordByte),err
}
