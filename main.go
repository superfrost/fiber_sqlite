package main

import (
	"fiber-sqlite/database"
	"fiber-sqlite/routes"
	"fiber-sqlite/telegram"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.Connect()

	app := fiber.New()

	app.Static("/", "./public")

	routes.RoutesSetup(app)

	telegram.Run()

	app.Listen(":" + os.Getenv("PORT"))
}
