package authdatabase

import (
	"modules/modules/authentication"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var AuthDatabase *gorm.DB

func ConnectAuthDB() error {
	var err error
	AuthDatabase, err = gorm.Open(sqlite.Open("UserAuth.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	AuthDatabase.AutoMigrate(&authentication.UserContactAuth{})

	return nil

}
