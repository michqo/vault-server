package handlers

import (
	"vault-server/database"

	"github.com/gofiber/fiber/v2"
)

func Files(c *fiber.Ctx) error {
	output, err := database.ListItems(c.Params("token") + "/")
	if err != nil {
		return fiber.ErrInternalServerError
	}
	/* 	for _, item := range output.Contents {
		fmt.Println("Name:         ", *item.Key)
		fmt.Println("Last modified:", *item.LastModified)
		fmt.Println("Size:         ", *item.Size)
		fmt.Println("Storage class:", *item.StorageClass)
		fmt.Println("")
	} */
	return c.JSON(output.Contents)
}
