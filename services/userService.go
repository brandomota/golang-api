package services

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/validator.v2"
	"log"

	"github.com/brandomota/golang-api/DTOs"
	repository "github.com/brandomota/golang-api/repositories"

	"github.com/brandomota/golang-api/models"
)

// GetAllUsers return all users in database
func GetAllUsers(context *fiber.Ctx) error {
	var users = repository.GetAllUsers()

	if users == nil {
		log.Printf("database is empty")
		return context.Status(404).JSON(&fiber.Map{"response": "not found"})
	} else {
		return context.Status(200).JSON(users)
	}
}

// GetUserById return user by **id** in database
func GetUserById(context *fiber.Ctx) error {
	id, err := ParseId(context)

	if err != nil {
		return context.Status(400).JSON(&fiber.Map{"error": err.Error()})
	}

	var user = repository.GetUserById(id)

	if user.ID == 0 {
		log.Printf("user not found: %d", id)
		return context.Status(404).JSON(&fiber.Map{"response": "not found"})
	} else {
		return context.Status(200).JSON(user)
	}
}

func CreateUser(context *fiber.Ctx) error {
	userDto := new(DTOs.UserDto)
	var err error
	var newUser models.User

	if err = context.BodyParser(userDto); err != nil {
		log.Printf("error on body parse: %s", err.Error())
		return context.Status(400).JSON(&fiber.Map{"error": err.Error()})
	}

	if err = validator.Validate(userDto); err != nil {
		log.Printf("error on body validation: %s", err.Error())
		return context.Status(400).JSON(&fiber.Map{"error": err.Error()})
	}

	newUser.Age = userDto.Age
	newUser.Email = userDto.Email
	newUser.Name = userDto.Name
	newUser.UserType = userDto.UserType

	err = repository.CreateUser(&newUser)

	if err != nil {
		log.Printf("error on create new user: %s", err.Error())
		return context.Status(400).JSON(&fiber.Map{"error": err.Error()})
	} else {
		return context.Status(201).JSON(newUser)
	}

}

func UpdateUser(context *fiber.Ctx) error {
	userData := new(models.User)

	id, err := ParseId(context)

	if err != nil {
		log.Printf("error on parse ID: %s", err.Error())
		return context.Status(400).JSON(&fiber.Map{"error": err.Error()})
	}

	user := repository.GetUserById(id)

	if user.ID == 0 {
		log.Printf("user not found: %d", id)
		return context.Status(404).JSON(&fiber.Map{"response": "not found"})
	}

	if err := context.BodyParser(userData); err != nil {
		log.Printf("error on body parse: %s", err.Error())
		return context.Status(400).JSON(&fiber.Map{"error": err.Error()})
	}

	if err = validator.Validate(userData); err != nil {
		log.Printf("error on body validation: %s", err.Error())
		return context.Status(400).JSON(&fiber.Map{"error": err.Error()})
	}

	user.Name = userData.Name
	user.Age = userData.Age
	user.Email = userData.Email
	user.UserType = userData.UserType

	saveErr := repository.UpdateUser(&user)

	if saveErr != nil {
		log.Printf("error on update user: %s", saveErr.Error())
		return context.Status(400).JSON(&fiber.Map{"error": saveErr.Error()})
	} else {
		return context.Status(200).JSON(user)
	}

}

func DeleteUser(context *fiber.Ctx) error {

	id, err := ParseId(context)

	if err != nil {
		return context.Status(400).JSON(&fiber.Map{"error": err.Error()})
	}

	var user = repository.GetUserById(id)

	if user.ID == 0 {
		log.Printf("user not found: %d", id)
		return context.Status(404).JSON(&fiber.Map{"response": "not found"})
	}

	dbErr := repository.DeleteUser(&user)

	if dbErr != nil {
		log.Printf("error on delete user: %s", dbErr.Error())
		return context.Status(400).JSON(&fiber.Map{"error": dbErr.Error()})
	} else {
		return context.SendStatus(204)
	}

}
