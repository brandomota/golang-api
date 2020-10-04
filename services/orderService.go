package services

import (
	"log"

	repository "github.com/brandomota/golang-api/repositories"
	"github.com/gofiber/fiber"
)

func GetAllOrders(context *fiber.Ctx) {
	var orders = repository.GetAllOrders()

	if orders == nil {
		log.Print("database is empty")
		context.Status(404).JSON(&fiber.Map{"response": "not found"})

	} else {
		context.Status(200).JSON(orders)
	}
}
