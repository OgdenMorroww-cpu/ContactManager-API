package models

import "gorm.io/gorm"

type UsersContactDetails struct {
	*gorm.Model
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Age           uint   `json:"age"`
	PhoneNumber   uint   `json:"phone_number"`
	NationalID    uint   `json:"national_id"`
	BVN           uint   `json:"bvn"`
	AccountNumber uint   `json:"account_number"`
}
