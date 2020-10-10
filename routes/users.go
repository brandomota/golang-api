package routes

import (
	"github.com/brandomota/golang-api/services"
	"github.com/gofiber/fiber/v2"
)

// SetUserRoutes init all routes for Group **User**
func SetUserRoutes(group fiber.Router) {
	group.Get("/", services.GetAllUsers)
	group.Post("/", services.CreateUser)
	group.Get("/:id", services.GetUserById)
	group.Post("/:id", services.UpdateUser)
	group.Delete("/:id", services.DeleteUser)
}
