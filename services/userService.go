package services

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber"

	"github.com/brandomota/golang-api/DTOs"
	repository "github.com/brandomota/golang-api/repositories"

	"github.com/brandomota/golang-api/models"
)

// GetAllUsers return all users in database
func GetAllUsers(context *fiber.Ctx) {
	var users = repository.GetAllUsers()

	if users == nil {
		log.Printf("database is empty")
		context.Status(404).JSON(&fiber.Map{"response": "not found"})
	} else {
		context.Status(200).JSON(users)
	}
}

// GetUserById return user by **id** in database
func GetUserById(context *fiber.Ctx) {
	id, err := strconv.ParseInt(context.Params("id"), 0, 36)

	if err != nil {
		log.Printf("error on parse ID: %s", err.Error())
		context.Status(400).JSON(&fiber.Map{"error": err.Error()})
		return
	}

	var user = repository.GetUserById(id)

	if user.ID == 0 {
		log.Printf("user not found: %d", id)
		context.Status(404).JSON(&fiber.Map{"response": "not found"})
	} else {
		context.Status(200).JSON(user)
	}
}

func CreateUser(context *fiber.Ctx) {
	userDto := new(DTOs.UserDto)
	var err error
	var newUser models.User

	if err := context.BodyParser(userDto); err != nil {
		log.Printf("error on body parse: %s", err.Error())
		context.Status(400).JSON(&fiber.Map{"error": err.Error()})
		return
	}

	newUser.Age = userDto.Age
	newUser.Email = userDto.Email
	newUser.Name = userDto.Name
	newUser.UserType = userDto.UserType

	err = repository.CreateUser(&newUser)

	if err != nil {
		log.Printf("error on create new user: %s", err.Error())
		context.Status(400).JSON(&fiber.Map{"error": err.Error()})
	} else {
		context.Status(201).JSON(newUser)
	}

}

func UpdateUser(context *fiber.Ctx) {
	userData := new(models.User)

	id, err := strconv.ParseInt(context.Params("id"), 0, 36)

	if err != nil {
		log.Printf("error on parse ID: %s", err.Error())
		context.Status(400).JSON(&fiber.Map{"error": err.Error()})
		return
	}

	user := repository.GetUserById(id)

	if user.ID == 0 {
		log.Printf("user not found: %d", id)
		context.Status(404).JSON(&fiber.Map{"response": "not found"})
		return
	}

	if err := context.BodyParser(userData); err != nil {
		log.Printf("error on body parse: %s", err.Error())
		context.Status(400).JSON(&fiber.Map{"error": err.Error()})
		return
	}

	user.Name = userData.Name
	user.Age = userData.Age
	user.Email = userData.Email
	user.UserType = userData.UserType

	saveErr := repository.UpdateUser(&user)

	if saveErr != nil {
		log.Printf("error on update user: %s", saveErr.Error())
		context.Status(400).JSON(&fiber.Map{"error": saveErr.Error()})
	} else {
		context.Status(200).JSON(user)
	}

}

func DeleteUser(context *fiber.Ctx) {

	id, err := strconv.ParseInt(context.Params("id"), 0, 36)

	if err != nil {
		log.Printf("error on parse ID: %s", err.Error())
		context.Status(400).JSON(&fiber.Map{"error": err.Error()})
		return
	}

	var user = repository.GetUserById(id)

	if user.ID == 0 {
		log.Printf("user not found: %d", id)
		context.Status(404).JSON(&fiber.Map{"response": "not found"})
		return
	}

	dbErr := repository.DeleteUser(&user)

	if dbErr != nil {
		log.Printf("error on delete user: %s", dbErr.Error())
		context.Status(400).JSON(&fiber.Map{"error": dbErr.Error()})
	} else {
		context.Status(204).Send()
	}

}
