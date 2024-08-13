package main

import (
	"github.com/gofiber/fiber/v2"
	"restApi-Jwt-Authentication/data"
	"restApi-Jwt-Authentication/routes"
)

func main() {
	app := fiber.New()

	engine, err := data.CreateDBEngine()
	if err != nil {
		panic(err)
	}

	routes.SetupRoutes(app, engine)

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
