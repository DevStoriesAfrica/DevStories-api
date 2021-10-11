package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
	"time"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment;unique" json:"id"`
	Username  string    `gorm:"size:255;unique;not null" json:"username"`
	Email     string    `gorm:"size:100;unique;not null" json:"email"`
	Password  string    `gorm:"size:100;not null" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (user *User) HashPassword(password string) (string, error) {
	hashedPasswordByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPasswordByte), err
}

func VerifyHashedPassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}

func (user *User) HashPasswordBeforeSave() error {
	hashedPassword, err := user.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	return nil
}

func (user *User) Prepare() {
	user.ID = 0
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.Password = html.EscapeString(strings.TrimSpace(user.Password))
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
}

func (user *User) Validate(action string) error {
	switch strings.ToLower(action) {

	case "create":
		if user.Username == "" {
			return errors.New("username required")
		}
		if user.Email == "" {
			return errors.New("email address required")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("invalid email address")
		}
		if user.Password == "" {
			return errors.New("password required")
		}

	case "login":
		if user.Email == "" {
			return errors.New("email address required")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("invalid email address")
		}
		if user.Password == "" {
			return errors.New("password required")
		}
	}

	return nil
}

func (user *User) SaveUser(db *gorm.DB) (*User, error) {
	err := user.HashPasswordBeforeSave()
	if err != nil {
		return &User{}, err
	}

	err = db.Debug().Create(&user).Error
	if err != nil {
		return &User{}, err
	}

	return user, nil
}

func (user *User) GetUser(db *gorm.DB, uid uint32) (*User, error) {
	err := db.Debug().Model(&User{}).Where("id=?", uid).Take(&user).Error
	if err != nil {
		return &User{}, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("user not found")
	}

	return user, nil
}

func (user *User) UpdateUser(db *gorm.DB, uid uint32) (*User, error) {
	err := db.Debug().Model(&User{}).Where("id=?", uid).Take(&User{}).Updates(&User{Username: user.Username, Email: user.Email, Password: user.Password, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &User{}, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("user not found")
	}

	err = db.Debug().Model(&User{}).Where("id=?", uid).Take(&user).Error
	if err != nil {
		return &User{}, err
	}

	return user, nil
}

func (user *User) DeleteUser(db *gorm.DB, uid uint32) (int64, error) {
	deleteTx := db.Debug().Model(&User{}).Where("id=?", uid).Take(&user).Delete(&user)
	if deleteTx.Error != nil {
		return 0, deleteTx.Error
	}

	if gorm.IsRecordNotFoundError(deleteTx.Error) {
		return 0, errors.New("user not found")
	}

	return deleteTx.RowsAffected, nil
}
