package route

import (
	"facetrack-gofiber/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Post("/predict", handler.PredictionHandlerCreate)
}