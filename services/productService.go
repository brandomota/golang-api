package services

import (
	"log"
	"strconv"

	"github.com/brandomota/golang-api/DTOs"
	"github.com/brandomota/golang-api/models"
	repository "github.com/brandomota/golang-api/repositories"
	"github.com/gofiber/fiber"
	"gopkg.in/validator.v2"
)

func GetAllProducts(context *fiber.Ctx) {
	var products = repository.GetAllProducts()

	if products == nil {
		log.Print("database is empty")
		context.Status(404).JSON(&fiber.Map{"response": "not found"})
	} else {
		context.Status(200).JSON(products)
	}
}

func GetProductById(context *fiber.Ctx) {
	id, err := strconv.ParseInt(context.Params("id"), 0, 36)

	if err != nil {
		log.Print("error on parse ID: %s", err.Error())
		context.Status(400).JSON(&fiber.Map{"error": err.Error()})
		return
	}

	var product = repository.GetProductById(id)

	if product.ID == 0 {
		log.Printf("product not found : %d", id)
		context.Status(404).JSON(&fiber.Map{"response": "not found"})
	} else {
		context.Status(200).JSON(product)
	}

}

func CreateProduct(context *fiber.Ctx) {
	productDto := new(DTOs.ProductDto)
	var err error
	var newProduct models.Product

	if err := context.BodyParser(productDto); err != nil {
		log.Printf("error on body parse: %s", err.Error())
		context.Status(400).JSON(&fiber.Map{"error": err.Error()})
		return
	}

	if err := validator.Validate(productDto); err != nil {
		log.Printf("error on body validation: %s", err.Error())
		context.Status(400).JSON(&fiber.Map{"error": err.Error()})
		return
	}

	newProduct.Description = productDto.Description
	newProduct.Name = productDto.Name
	newProduct.ImageUrl = productDto.ImageUrl
	newProduct.Price = productDto.Price

	err = repository.CreateProduct(&newProduct)

	if err != nil {
		log.Printf("error on create new product: %s", err.Error())
		context.Status(400).JSON(&fiber.Map{"error": err.Error()})
	} else {
		context.Status(201).JSON(newProduct)
	}
}

func UpdateProduct(context *fiber.Ctx) {
	productDto := new(DTOs.ProductDto)
	var err error

	if err := context.BodyParser(productDto); err != nil {
		log.Printf("error on body parse: %s", err.Error())
		context.Status(400).JSON(&fiber.Map{"error": err.Error()})
		return
	}

	if err := validator.Validate(productDto); err != nil {
		log.Printf("error on body validation: %s", err.Error())
		context.Status(400).JSON(&fiber.Map{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Params("id"), 0, 36)

	if err != nil {
		log.Print("error on parse ID: %s", err.Error())
		context.Status(400).JSON(&fiber.Map{"error": err.Error()})
		return
	}

	product := repository.GetProductById(id)

	if product.ID == 0 {
		log.Printf("product not found : %d", id)
		context.Status(404).JSON(&fiber.Map{"response": "not found"})
		return
	}

	product.Price = productDto.Price
	product.ImageUrl = productDto.ImageUrl
	product.Name = productDto.Name
	product.Description = productDto.Description

	err = repository.UpdateProduct(&product)

	if err != nil {
		log.Printf("error on update product: %s", err.Error())
		context.Status(400).JSON(&fiber.Map{"error": err.Error()})
	} else {
		context.Status(200).JSON(product)
	}

}

func DeleteProduct(context *fiber.Ctx) {
	id, err := strconv.ParseInt(context.Params("id"), 0, 36)

	if err != nil {
		log.Print("error on parse ID: %s", err.Error())
		context.Status(400).JSON(&fiber.Map{"error": err.Error()})
		return
	}

	product := repository.GetProductById(id)

	if product.ID == 0 {
		log.Printf("product not found : %d", id)
		context.Status(404).JSON(&fiber.Map{"response": "not found"})
		return
	}

	err = repository.DeleteProduct(&product)

	if err != nil {
		log.Printf("error on delete product: %s", err.Error())
		context.Status(400).JSON(&fiber.Map{"error": err.Error()})
	} else {
		context.Status(204).Send()
	}

}
