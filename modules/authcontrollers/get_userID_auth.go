package authcontrollers

import (
	"modules/modules/authdatabase"
	"modules/modules/authentication"

	"github.com/gofiber/fiber/v2"
)

func GetUserAuthsID(c *fiber.Ctx) error {
	id := c.Params("id")
	var usersAuth authentication.UserContactAuth
	result := authdatabase.AuthDatabase.Find(&usersAuth, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&usersAuth)
}
