package routes

import (
	"github.com/brandomota/golang-api/services"
	"github.com/gofiber/fiber/v2"
)

func SetOrderRoutes(group fiber.Router) {
	group.Get("/", services.GetAllOrders)
	group.Get("/:id", services.GetOrderById)
	// group.Get("/byUser/:id", services.GetByUserId) Added in next branch
	//group.Post("/", services.CreateOrder)
	//group.Post("/:id/status", services.UpdateStatus)
}
