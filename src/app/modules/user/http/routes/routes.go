package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-api/src/app/modules/user/http/controllers"
	"go-api/src/app/modules/user/http/middlewares"
)

func UserRoutes(app *fiber.App, handler *controllers.UsersServices) {
	app.Post("/api/register", handler.Create)
	app.Post("/api/login", handler.Login)

	app.Use(middlewares.IsAuthenticated)

	app.Get("/api/users", handler.Index)
	app.Get("/api/users/:user_id", handler.Show)
	app.Put("/api/users/:user_id", handler.Update)

	app.Post("/api/logout", controllers.Logout)
}
