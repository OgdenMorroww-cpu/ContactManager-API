package controllers

import (
	"modules/modules/database"
	"modules/modules/models"

	"github.com/gofiber/fiber/v2"
)

func DeleteUsersContacts(c *fiber.Ctx) error {
	id := c.Params("id")
	users := new(models.UsersContactDetails)

	result := database.Database.Delete(&users, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.SendStatus(200)
}
