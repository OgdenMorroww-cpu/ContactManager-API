package controllers

import (
	"modules/modules/database"
	"modules/modules/models"

	"github.com/gofiber/fiber/v2"
)

func GetUsersContacts(c *fiber.Ctx) error {
	id := c.Params("id")
	var users []models.UsersContactDetails

	result := database.Database.Find(&users, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&users)
}
