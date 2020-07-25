package routes

import (
	"github.com/gofiber/fiber"
)

// SetRoutes define router groups and init all routes
func SetRoutes(server *fiber.App) {
	SetUserRoutes(server.Group("/users"))
}
