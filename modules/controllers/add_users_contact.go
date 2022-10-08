package controllers

import (
	"modules/modules/database"
	"modules/modules/models"

	"github.com/gofiber/fiber/v2"
)

func AddUsersContacts(c *fiber.Ctx) error {
	users := new(models.UsersContactDetails)

	if err := c.BodyParser(users); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.Database.Create(&users)
	return c.Status(201).JSON(users)
}
