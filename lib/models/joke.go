package models

import (
	"gorm.io/gorm"
)

type Joke struct {
	gorm.Model
	Name    string
	Content string
	UserID  uint
	User    User
}

func (j *Joke) Create(db *gorm.DB) error {
	return db.Create(j).Error
}

func (j *Joke) Delete(db *gorm.DB) error {
	return db.Delete(j).Error
}

func (j *Joke) GetById(db *gorm.DB) error {
	result := db.First(&j)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (j *Joke) GetAll(db *gorm.DB, dst *[]Joke) error {
	result := db.Find(&dst)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
