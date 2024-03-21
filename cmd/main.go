package main

import (
	"log"
	"vault-server/cmd/config"
	"vault-server/internal/database"
	"vault-server/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.LoadConfig()
	database.NewS3Client()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	api := app.Group("/v1")
	routes.CreateRoutes(api)

	log.Fatal(app.Listen("127.0.0.1:8000"))
}
