package controller

import (
	db "Project5-API1/config"
	Models "Project5-API1/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"os"
	"strconv"
	"time"
)

func Login(c *fiber.Ctx) error {
	customerId := c.Params("customerId")
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data",
		})
	}
	if data["passcode"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Passcode is required.",
			"error":   map[string]interface{}{},
		})
	}
	var customer Models.Customer
	db.DB.Where("customer_id = ?", customerId).First(&customer)
	if customer.CustomerId == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Customer not found",
			"error":   map[string]interface{}{},
		})
	}
	if customer.Passcode != data["passcode"] {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Passcode not matching",
		})

	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer":    strconv.Itoa(int(customer.CustomerId)),
		"ExpiresAt": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Error signing token",
		})

	}
	customerData := make(map[string]interface{})
	customerData["token"] = tokenString

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    customerData,
	})
}

func Logout(c *fiber.Ctx) error {
	customerId := c.Params("customerId")
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Logout successful for customer ID: " + customerId,
	})
}

func Passcode(c *fiber.Ctx) error {
	customerId := c.Params("customerId")
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data",
		})
	}
	var customer Models.Customer
	db.DB.First(&customer, customerId)
	if customer.CustomerId == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Customer not found",
		})
	}
	customer.LastName = data["passcode"]
	db.DB.Save(&customer)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Passcode updated successfully",
		"data":    customer,
	})
}
