package services

import (
	repository "github.com/brandomota/golang-api/repositories"
	"github.com/gofiber/fiber/v2"
	"log"
)

func GetAllOrders(context *fiber.Ctx) error {
	var orders = repository.GetAllOrders()

	if orders == nil {
		log.Print("database is empty")
		return context.Status(404).JSON(&fiber.Map{"response": "not found"})

	} else {
		return context.Status(200).JSON(orders)
	}
}

func GetOrderById(context *fiber.Ctx) error {
	id, err := ParseId(context)

	if err != nil {
		return context.Status(400).JSON(&fiber.Map{"error": err.Error()})
	}

	var order = repository.GetOrderById(id)

	if order.ID == 0 {
		log.Printf("order not found : %d", id)
		return context.Status(404).JSON(&fiber.Map{"response": "not found"})
	} else {
		return context.Status(200).JSON(order)
	}
}

func GetByUserId(context *fiber.Ctx) error {
	id, err := ParseId(context)

	if err != nil {
		return context.Status(400).JSON(&fiber.Map{"error": err.Error()})
	}
	var orders = repository.GetOrdersByUserId(id)
	if len(orders) == 0 {
		log.Printf("no orders found")
		return context.Status(404).JSON(&fiber.Map{"response": "not found"})
	} else {
		return context.Status(200).JSON(orders)
	}
}
