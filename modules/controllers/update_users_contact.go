package controllers

import (
	"modules/modules/database"
	"modules/modules/models"

	"github.com/gofiber/fiber/v2"
)

func UpdateUsersContact(c *fiber.Ctx) error {
	users := new(models.UsersContactDetails)
	id := c.Params("id")

	if err := c.BodyParser(users); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.Database.Where("id = ?", id).Updates(&users)
	return c.Status(200).JSON(users)
}
