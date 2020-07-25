package main

import (
	"os"

	"github.com/brandomota/golang-api/routes"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

func main() {

	server := fiber.New()

	server.Use(middleware.Logger())

	// TODO: add routes
	routes.SetRoutes(server)

	port := os.Getenv("SERVER_PORT")
	if len(port) < 1 {
		port = "3000"
	}

	server.Listen(port)
}
