package authcontrollers

import (
	"modules/modules/authdatabase"
	"modules/modules/authentication"

	"github.com/gofiber/fiber/v2"
)

func DeleteUserAuths(c *fiber.Ctx) error {
	id := c.Params("id")
	var userAuth authentication.UserContactAuth

	result := authdatabase.AuthDatabase.Delete(&userAuth, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
