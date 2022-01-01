package routes

import (
	"fiber-sqlite/controllers"

	"github.com/gofiber/fiber/v2"
)

func RoutesSetup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/logout", controllers.Logout)
	app.Get("/api/user", controllers.User)

}
