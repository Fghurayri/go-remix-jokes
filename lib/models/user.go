package models

import (
	"go-remix-jokes/lib/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string
	PasswordHash string
	Jokes        []Joke
}

func (u *User) Create(db *gorm.DB, password string) error {
	hashedPassword, err := utils.GenerateHashFromPassword(password)
	if err != nil {
		panic(err.Error())
	}

	u.PasswordHash = string(hashedPassword)
	return db.Create(u).Error
}

func (u *User) VerifyCredentials(db *gorm.DB, password string) error {
	err := db.Where("username = ?", u.Username).First(&u).Error
	if err != nil {
		return err
	}

	err = utils.CompareHashAndPassword(u.PasswordHash, password)
	if err != nil {
		return err
	}
	return nil
}
