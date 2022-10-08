package database

import (
	"modules/modules/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectDataBase() error {
	var err error
	Database, err = gorm.Open(sqlite.Open("UsersContactDetails.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&models.UsersContactDetails{})

	return nil
}
