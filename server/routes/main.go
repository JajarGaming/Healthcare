package routes

import (
	loginhandlers "api/handlers/login"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	//Login
	app.Get("/login", loginhandlers.Login)
	app.Get("/logout", loginhandlers.Logout)
	app.Get("/sessionStorage", loginhandlers.GetUsers)

}
