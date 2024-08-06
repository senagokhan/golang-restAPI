package routes

import (
	"Project5-API1/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/customer/:customerId/login", controller.Login)
	app.Get("/customer/:customerId/logout", controller.Logout)
	app.Post("/customer/customerId/passcode", controller.Passcode)

	app.Post("/customer", controller.CreateCustomer)
	app.Get("/customer", controller.CustomerList)
	app.Get("/customer/:customerId", controller.GetCustomerDetails)
	app.Delete("/customer/:customerId", controller.DeleteCustomer)
	app.Put("/customer/:customerId", controller.UpdateCustomer)

}
