package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CreateRoutes(api fiber.Router) {
	api.Get("/hello", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello %v", c.IP())
		return c.SendString(msg)
	})
}
