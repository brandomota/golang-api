package services

import (
	repository "github.com/brandomota/golang-api/repositories"
	"github.com/gofiber/fiber"
)

// GetAllUsers return all users in database
func GetAllUsers(context *fiber.Ctx) {
	var users = repository.GetAllUsers()

	if users == nil {
		context.Status(404).JSON(&fiber.Map{"response": "not found"})
	} else {
		context.Status(200).JSON(users)
	}
}

// GetUserById return user by **id** in database
func GetUserById(context *fiber.Ctx) {
	panic("not implemented yet!")
}

func CreateUser(context *fiber.Ctx) {
	panic("not implemented yet!")
}

func UpdateUser(context *fiber.Ctx) {
	panic("not implemented yet!")
}

func DeleteUser(context *fiber.Ctx) {
	panic("not implemented yet")
}
