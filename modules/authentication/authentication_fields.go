package authentication

import "gorm.io/gorm"

type UserContactAuth struct {
	*gorm.Model
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}
