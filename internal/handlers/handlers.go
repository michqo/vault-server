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

type GetObjectType struct {
	Key   string `json:"key"`
	Token string `json:"token"`
}

type ObjectUrlType struct {
	Url string `json:"url"`
}

type ObjectKey struct {
	Key string `json:"key"`
}

type ObjectKeySlice []ObjectKey

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

func UploadChecks(token string) error {
	size, err1 := database.BucketSize()
	size2, err2 := database.BucketPrefixSize(token + "/")
	if err1 != nil || err2 != nil {
		return fiber.ErrInternalServerError
	} else if size >= Cfg.MaxBucketSize || size2 >= Cfg.MaxFolderSize {
		return fiber.ErrInsufficientStorage
	}
	return nil
}

func ObjectUrl(c *fiber.Ctx) error {
	key := c.Query("key")
	token := c.Query("token")
	if key == "" || token == "" {
		return fiber.ErrBadRequest
	}
	err := UploadChecks(token)
	if err != nil {
		return err
	}
	url, err := database.ObjectGetUrl(token + "/" + key)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(fiber.Map{"url": url})
}

func ObjectUrls(c *fiber.Ctx) error {
	body := ObjectKeySlice{}
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}
	urls := make([]ObjectUrlType, len(body))
	for i, object := range body {
		err := UploadChecks(object.Key)
		if err != nil {
			return err
		}
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
	body := ObjectKeySlice{}
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
