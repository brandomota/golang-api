package services

import (
	"log"
	"strconv"

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

func GetOrderById(context *fiber.Ctx) {
	id, err := strconv.ParseInt(context.Params("id"), 0, 36)
	if err != nil {
		log.Printf("error on parse ID: %s", err.Error())
		context.Status(400).JSON(&fiber.Map{"error": err.Error()})
		return
	}

	var order = repository.GetOrderById(id)

	if order.ID == 0 {
		log.Printf("order not found : %d", id)
		context.Status(404).JSON(&fiber.Map{"response": "not found"})
	} else {
		context.Status(200).JSON(order)
	}
}
