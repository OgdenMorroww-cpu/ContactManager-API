package controllers

import (
	"modules/modules/database"
	"modules/modules/models"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var userContacts []models.UsersContactDetails
	database.Database.Find(&userContacts)
	return c.Status(200).JSON(userContacts)
}
