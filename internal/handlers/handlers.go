package handlers

import (
	"vault-server/internal/database"

	. "vault-server/cmd/config"

	"github.com/gofiber/fiber/v2"
)

type Object struct {
	Key          string
	LastModified string
	Size         int64
}

func GetObjects(c *fiber.Ctx) error {
	token := c.Query("token")
	if token == "" {
		return fiber.ErrBadRequest
	}
	output, err := database.ListObjects(token + "/")
	if err != nil {
		return fiber.ErrInternalServerError
	}
	objects := make([]Object, len(output.Contents))
	for i, object := range output.Contents {
		objects[i] = Object{Key: *object.Key, LastModified: object.LastModified.String(), Size: *object.Size}
	}
	return c.JSON(objects)
}

func ObjectUrl(c *fiber.Ctx) error {
	urlType := c.Query("type")
	key := c.Query("key")
	token := c.Query("token")
	if urlType == "" || key == "" || token == "" {
		return fiber.ErrBadRequest
	}
	switch urlType {
	case "GET":
		url, err := database.ObjectGetUrl(token + "/" + key)
		if err != nil {
			return fiber.ErrInternalServerError
		}
		return c.JSON(fiber.Map{"url": url})
	case "PUT":
		size, err1 := database.BucketSize()
		size2, err2 := database.BucketPrefixSize(token + "/")
		if err1 != nil || err2 != nil {
			return fiber.ErrInternalServerError
		}
		if size >= Cfg.MaxBucketSize || size2 >= Cfg.MaxFolderSize {
			return fiber.ErrInsufficientStorage
		}
		url, err := database.ObjectPutUrl(token + "/" + key)
		if err != nil {
			return fiber.ErrInternalServerError
		}
		return c.JSON(fiber.Map{"url": url})
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Type not allowed"})
	}
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
