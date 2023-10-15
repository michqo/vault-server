package main

import (
	"log"
	"vault-server/database"
	"vault-server/routes"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	database.CreateDatabase()
	app := fiber.New()

	api := app.Group("/v1")
	routes.CreateRoutes(api)

	log.Fatal(app.Listen(":8000"))
}
