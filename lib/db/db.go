package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"go-remix-jokes/lib/models"
)

var DB *gorm.DB

func ConnectAndMigrateDB() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// Migrate the schema
	err = db.AutoMigrate(
		&models.User{},
		&models.Joke{},
	)
	if err != nil {
		panic("failed to migrate models " + err.Error())
	}

	DB = db
}
