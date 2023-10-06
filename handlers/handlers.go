package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	msg := fmt.Sprintf("Hello %v", c.IP())
	return c.SendString(msg)
}
