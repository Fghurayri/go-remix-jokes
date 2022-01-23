package models

import "gorm.io/gorm"

type Joke struct {
	gorm.Model
	Name    string
	Content string
	UserID  uint
	User    User
}
