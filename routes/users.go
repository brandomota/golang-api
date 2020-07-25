package routes

import (
	"github.com/brandomota/golang-api/services"
	"github.com/gofiber/fiber"
)

// SetUserRoutes init all routes for Group **User**
func SetUserRoutes(group fiber.Router) {
	group.Get("/", services.GetAllUsers)
	group.Get("/:id", services.GeUserById)
}
