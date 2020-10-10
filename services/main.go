package services

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func ParseId(ctx *fiber.Ctx) (int64, error)  {
	id, err := strconv.ParseInt(ctx.Params("id"), 0, 36)
	if err != nil {
		log.Printf("error on parse ID: %s", err.Error())
	}

	return id, err
}
