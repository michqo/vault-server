package main

import (
	"log"
	"vault-server/cmd/config"
	"vault-server/internal/database"
	"vault-server/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadConfig()
	database.NewS3Client()

	app := fiber.New()
	api := app.Group("/v1")
	routes.CreateRoutes(api)

	log.Fatal(app.Listen(":8000"))
}
