package main

import (
	"log"
	"vault-server/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	api := app.Group("/v1")
	routes.CreateRoutes(api)

	log.Fatal(app.Listen(":8000"))
}
