package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-api/src/app/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/api/register", controllers.Register)
}
