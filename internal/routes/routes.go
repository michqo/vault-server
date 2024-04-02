package routes

import (
	"vault-server/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func CreateRoutes(api fiber.Router) {
	api.Get("/objects", handlers.GetObjects)
	api.Get("/object-url", handlers.ObjectUrl)
	api.Get("/random-token", handlers.GetToken)
	api.Post("/object-urls", handlers.ObjectUrls)
	api.Delete("/object", handlers.DeleteObject)
	api.Delete("/objects", handlers.DeleteObjects)
}
