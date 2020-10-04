package routes

import (
	"github.com/gofiber/fiber"
)

func SetOrderRoutes(group fiber.Router) {
	group.Get("/", service.GetAllOrders)
	group.Get("/:id", service.GetOrder)
	group.Get("/byUser/:id", service.GetByUserId)
	group.Post("/", service.CreateOrder)
}
