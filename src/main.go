package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go-api/src/app/modules/user/controllers"
	"go-api/src/app/modules/user/routes"
	"go-api/src/database"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env or invalid environment")
	}
}

func main() {
	dburl := os.Getenv("DB_URL")

	services, err := database.NewRepositories(dburl)
	if err != nil {
		panic(err)
	}
	_ = services.Automigrate()

	app := fiber.New()

	app.Use(cors.New(cors.Config{AllowCredentials: true}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("root")
	})

	userHandler := controllers.UsersController(services.User)
	routes.UserRoutes(app, userHandler)

	_ = app.Listen(":3000")
}
