package services

import (
	"github.com/brandomota/golang-api/DTOs"
	"github.com/brandomota/golang-api/models"
	repository "github.com/brandomota/golang-api/repositories"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/validator.v2"
	"log"
)

func GetAllProducts(context *fiber.Ctx) error {
	var products = repository.GetAllProducts()

	if products == nil {
		log.Print("database is empty")
		return context.Status(404).JSON(&fiber.Map{"response": "not found"})
	} else {
		return context.Status(200).JSON(products)
	}
}

func GetProductById(context *fiber.Ctx) error {
	id, err := ParseId(context)

	if err != nil {
		return context.Status(400).JSON(&fiber.Map{"error": err.Error()})
	}

	var product = repository.GetProductById(id)

	if product.ID == 0 {
		log.Printf("product not found : %d", id)
		return context.Status(404).JSON(&fiber.Map{"response": "not found"})
	} else {
		return context.Status(200).JSON(product)
	}

}

func CreateProduct(context *fiber.Ctx) error {
	productDto := new(DTOs.ProductDto)
	var err error
	var newProduct models.Product

	if err := context.BodyParser(productDto); err != nil {
		log.Printf("error on body parse: %s", err.Error())
		return context.Status(400).JSON(&fiber.Map{"error": err.Error()})
	}

	if err := validator.Validate(productDto); err != nil {
		log.Printf("error on body validation: %s", err.Error())
		return context.Status(400).JSON(&fiber.Map{"error": err.Error()})
	}

	newProduct.Description = productDto.Description
	newProduct.Name = productDto.Name
	newProduct.ImageUrl = productDto.ImageUrl
	newProduct.Price = productDto.Price

	err = repository.CreateProduct(&newProduct)

	if err != nil {
		log.Printf("error on create new product: %s", err.Error())
		return context.Status(400).JSON(&fiber.Map{"error": err.Error()})
	} else {
		return context.Status(201).JSON(newProduct)
	}
}

func UpdateProduct(context *fiber.Ctx) error {
	productDto := new(DTOs.ProductDto)
	var err error

	if err := context.BodyParser(productDto); err != nil {
		log.Printf("error on body parse: %s", err.Error())
		return context.Status(400).JSON(&fiber.Map{"error": err.Error()})
	}

	if err := validator.Validate(productDto); err != nil {
		log.Printf("error on body validation: %s", err.Error())
		return context.Status(400).JSON(&fiber.Map{"error": err.Error()})
	}

	id, err := ParseId(context)

	if err != nil {
		return context.Status(400).JSON(&fiber.Map{"error": err.Error()})
	}

	product := repository.GetProductById(id)

	if product.ID == 0 {
		log.Printf("product not found : %d", id)
		return context.Status(404).JSON(&fiber.Map{"response": "not found"})
	}

	product.Price = productDto.Price
	product.ImageUrl = productDto.ImageUrl
	product.Name = productDto.Name
	product.Description = productDto.Description

	err = repository.UpdateProduct(&product)

	if err != nil {
		log.Printf("error on update product: %s", err.Error())
		return context.Status(400).JSON(&fiber.Map{"error": err.Error()})
	} else {
		return context.Status(200).JSON(product)
	}

}

func DeleteProduct(context *fiber.Ctx) error {
	id, err := ParseId(context)

	if err != nil {
		return context.Status(400).JSON(&fiber.Map{"error": err.Error()})
	}

	product := repository.GetProductById(id)

	if product.ID == 0 {
		log.Printf("product not found : %d", id)
		return context.Status(404).JSON(&fiber.Map{"response": "not found"})
	}

	err = repository.DeleteProduct(&product)

	if err != nil {
		log.Printf("error on delete product: %s", err.Error())
		return context.Status(400).JSON(&fiber.Map{"error": err.Error()})
	} else {
		return context.SendStatus(204)
	}

}
