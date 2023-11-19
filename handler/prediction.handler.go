package handler

import (
	"github.com/gofiber/fiber/v2"
)

func PredictionHandlerCreate(ctx *fiber.Ctx) error {

	return ctx.JSON(fiber.Map{
		"message": "terprediksi",
	}) 
}