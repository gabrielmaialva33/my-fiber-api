package main

import (
	"github.com/gofiber/fiber/v2"
	"go-api/src/database"
	"go-api/src/routes"
)

func main() {
	app := fiber.New()

	database.Connect()

	routes.Setup(app)

	_ = app.Listen(":3000")
}
