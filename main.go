package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

func main() {

	server := fiber.New()

	server.Use(middleware.Logger())

	// TODO: add routes

	server.Listen()
}
