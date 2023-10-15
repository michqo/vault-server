package routes

import (
	"vault-server/handlers"

	"github.com/gofiber/fiber/v2"
)

func CreateRoutes(api fiber.Router) {
	api.Get("/files/:token", handlers.Files)
}
