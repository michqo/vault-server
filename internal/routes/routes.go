package routes

import (
	"vault-server/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func CreateRoutes(api fiber.Router) {
	api.Get("/objects", handlers.GetObjects)
	api.Get("/object-url", handlers.ObjectUrl)
	api.Post("/object-urls", handlers.ObjectPutUrls)
	api.Delete("/object", handlers.DeleteObject)
	api.Delete("/objects", handlers.DeleteObjects)
}
