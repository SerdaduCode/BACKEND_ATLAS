package main

import (
	"backend-atlas/config"
	"backend-atlas/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.ConnectDB()

	routes.ProvinceRoutes(app)

	app.Listen(":3000")
}
