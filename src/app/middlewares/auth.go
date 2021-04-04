package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"go-api/src/app/utils"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("go-api")

	if _, err := utils.ParseJwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return c.Next()
}
