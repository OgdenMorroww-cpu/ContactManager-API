package authcontrollers

import (
	"modules/modules/authdatabase"
	"modules/modules/authentication"

	"github.com/gofiber/fiber/v2"
)

func GetUserAuths(c *fiber.Ctx) error {
	var userAuths []authentication.UserContactAuth
	authdatabase.AuthDatabase.Find(&userAuths)
	return c.Status(200).JSON(userAuths)
}
