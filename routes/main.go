package routes

import (
	"github.com/gofiber/fiber/v2"
)

// SetRoutes define router groups and init all routes
func SetRoutes(server *fiber.App) {
	SetUserRoutes(server.Group("/users"))
	SetProductRoutes(server.Group("/products"))
	SetOrderRoutes(server.Group("/orders"))
}
