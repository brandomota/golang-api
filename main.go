package main

import (
	"os"

	"github.com/brandomota/golang-api/repositories"
	"github.com/brandomota/golang-api/routes"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

func main() {

	godotenv.Load()

	repositories.InitDatabase()

	server := fiber.New()

	server.Use(middleware.Logger())

	routes.SetRoutes(server)

	port := os.Getenv("SERVER_PORT")

	server.Listen(port)
}
