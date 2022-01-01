package main

import (
	"fiber-sqlite/database"
	"fiber-sqlite/routes"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Static("/", "./public")

	routes.RoutesSetup(app)

	app.Listen(":" + os.Getenv("PORT"))
}
