package authcontrollers

import (
	"modules/modules/authdatabase"
	"modules/modules/authentication"

	"github.com/gofiber/fiber/v2"
)

func AddUsersAuth(c *fiber.Ctx) error {
	userAuth := new(authentication.UserContactAuth)

	if err := c.BodyParser(userAuth); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	authdatabase.AuthDatabase.Create(&userAuth)

	return c.Status(201).JSON(userAuth)
}
