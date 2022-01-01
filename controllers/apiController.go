package controllers

import (
	"fiber-sqlite/database"
	"fiber-sqlite/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {

	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt

	result := database.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "user",
	})
}

func Logout(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "user",
	})
}

func User(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "empty",
	})
}
