package routes

import "github.com/gofiber/fiber"

func SetProductRoutes(group fiber.Router) {
	group.Get("/", services.GetAllProducts)
	group.Post("/", services.CreateProduct)
	group.Get("/:id", services.GetProductById)
	group.Post("/:id", services.UpdateProduct)
	group.Delete("/:id", services.DeleteProduct)
}
