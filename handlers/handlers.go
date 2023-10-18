package handlers

import (
	"vault-server/database"

	"github.com/gofiber/fiber/v2"
)

func GetObjects(c *fiber.Ctx) error {
	token := c.Query("token")
	if token == "" {
		return fiber.ErrBadRequest
	}
	output, err := database.ListObjects(token + "/")
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(output.Contents)
}

func GetObjectUrl(c *fiber.Ctx) error {
	key := c.Query("key")
	if key == "" {
		return fiber.ErrBadRequest
	}
	url, err := database.GetObjectUrl(key)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(fiber.Map{"url": url})
}

func DeleteObject(c *fiber.Ctx) error {
	key := c.Query("key")
	if key == "" {
		return fiber.ErrBadRequest
	}
	err := database.DeleteObject(key)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.SendStatus(200)
}
