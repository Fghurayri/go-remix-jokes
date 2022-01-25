package models

import (
	"go-remix-jokes/lib/db"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string
	PasswordHash string
	Jokes        []Joke
}

func (u *User) CreateUser(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return db.DB.Create(&User{
		Username:     username,
		PasswordHash: string(hashedPassword),
	}).Error
}
