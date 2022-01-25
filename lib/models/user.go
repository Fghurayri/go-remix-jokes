package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string
	PasswordHash string
	Jokes        []Joke
}

func (u *User) CreateUser(db *gorm.DB) error {
	return db.Create(u).Error
}

func (u *User) GetUserByUsername(db *gorm.DB, username string) error {
	err := db.Where("username = ?", username).First(&u).Error
	if err != nil {
		return err
	}
	return nil
}
