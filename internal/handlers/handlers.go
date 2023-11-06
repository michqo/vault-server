package handlers

import (
	"vault-server/internal/database"

	. "vault-server/cmd/config"

	"github.com/gofiber/fiber/v2"
)

type Object struct {
	Key          string `json:"key"`
	LastModified string `json:"lastModified"`
	Size         int64  `json:"size"`
}

type ObjectUrlType struct {
	Url string `json:"url"`
}

type ObjectKey struct {
	Key string `json:"key"`
}

type ObjectSlice []ObjectKey

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

func ObjectPutUrls(c *fiber.Ctx) error {
	body := ObjectSlice{}
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}
	urls := make([]ObjectUrlType, len(body))
	for i, object := range body {
		url, err := database.ObjectPutUrl(object.Key)
		if err != nil {
			return fiber.ErrInternalServerError
		}
		urls[i] = ObjectUrlType{Url: url}
	}
	return c.JSON(fiber.Map{"urls": urls})
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

func DeleteObjects(c *fiber.Ctx) error {
	body := ObjectSlice{}
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}
	for _, object := range body {
		err := database.DeleteObject(object.Key)
		if err != nil {
			return fiber.ErrInternalServerError
		}
	}
	return c.SendStatus(200)
}
