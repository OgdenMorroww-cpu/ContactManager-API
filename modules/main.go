package main

import (
	"log"
	"modules/modules/authcontrollers"
	"modules/modules/authdatabase"
	"modules/modules/controllers"
	"modules/modules/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDataBase()
	authdatabase.ConnectAuthDB()

	//Api routes endpoints for user contacts
	app.Get("/users", controllers.GetUsers)
	app.Get("/users/:id", controllers.GetUsersContacts)
	app.Post("/users", controllers.AddUsersContacts)
	app.Patch("/users/:id", controllers.UpdateUsersContact)
	app.Delete("/users/:id", controllers.DeleteUsersContacts)

	//Api routes endpoints for user authentication endpoints
	app.Get("/userAuth", authcontrollers.GetUserAuths)
	app.Post("/userAuth", authcontrollers.AddUsersAuth)
	app.Get("/userAuth/:id", authcontrollers.GetUserAuthsID)
	app.Get("/userAuth/:id", authcontrollers.DeleteUserAuths)

	log.Fatal(app.Listen(":5000"))
}
