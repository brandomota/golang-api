package main

import (
	"github.com/brandomota/golang-api/repositories"
	"github.com/brandomota/golang-api/routes"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
)

func main() {

	godotenv.Load()

	repositories.InitDatabase()

	server := fiber.New()

	server.Use(logger.New())

	routes.SetRoutes(server)

	//port := os.Getenv("PORT")

	server.Listen(":3000")
}
