package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-api/src/app/models"
)

func Register(c *fiber.Ctx) error {
	user := models.User{
		Name:     "",
		Nickname: "",
		Email:    "",
		Password: "",
	}

	user.Password = ""

	return c.JSON(user)
}
