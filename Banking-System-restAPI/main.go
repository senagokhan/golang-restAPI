package main

import (
	db "Project5-API1/config"
	"Project5-API1/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db.Connect()

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		println("Request:", c.Method(), c.Path())
		return c.Next()
	})

	routes.Setup(app)
	err := app.Listen(":30001")
	if err != nil {
		return
	}

}
