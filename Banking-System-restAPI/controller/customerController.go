package controller

import (
	db "Project5-API1/config"
	Models "Project5-API1/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

func CreateCustomer(c *fiber.Ctx) error {

	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data",
		})
	}
	if data["FirstName"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Customer name is required",
		})
	}

	customer := Models.Customer{
		FirstName: data["FirstName"],
		LastName:  data["LastName"],
		Passcode:  data["Passcode"],
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	db.DB.Create(&customer)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Customer created successfully",
		"data":    customer,
	})
}

func UpdateCustomer(c *fiber.Ctx) error {

	customerId := c.Params("customerId")
	var customer Models.Customer

	db.DB.Find(&customer, "customerId = ?", customerId)

	if customer.FirstName == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Customer not found",
		})
	}

	var updatedCustomer Models.Customer
	err := c.BodyParser(&updatedCustomer)
	if err != nil {
		return err
	}

	if updatedCustomer.FirstName == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Customer name is required",
		})
	}

	customer.FirstName = updatedCustomer.FirstName
	db.DB.Save(&customer)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Customer updated successfully",
		"data":    customer,
	})
}

func DeleteCustomer(c *fiber.Ctx) error {

	customerId := c.Params("customerId")
	var customer Models.Customer

	db.DB.Where("customerId = ?", customerId).Find(&customer)
	if customer.CustomerId == 0 {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Customer not found",
		})

	}

	db.DB.Where("customerId = ?", customer.CustomerId).Delete(&Models.Customer{})
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Customer deleted successfully",
	})
}

func GetCustomerDetails(c *fiber.Ctx) error {

	customerId := c.Params("customerId")
	var customer Models.Customer

	db.DB.Select("customerId,firstName,lastName,created_at,updated_at").Where("customerId = ?", customerId).First(&customer)

	customerData := make(map[string]interface{})
	customerData["customerId"] = customer.CustomerId
	customerData["firstName"] = customer.FirstName
	customerData["lastName"] = customer.LastName
	customerData["passcode"] = customer.Passcode
	customerData["created_at"] = customer.CreatedAt
	customerData["updated_at"] = customer.UpdatedAt

	if customer.CustomerId == 0 {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Customer not found",
			"error":   map[string]interface{}{},
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Successful",
		"data":    customerData,
	})
}

func CustomerList(c *fiber.Ctx) error {
	var customers []Models.Customer
	limit, err1 := strconv.Atoi(c.Query("limit", "10"))
	skip, err2 := strconv.Atoi(c.Query("skip", "0"))
	var count int64

	if err1 != nil || err2 != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid query parameters",
		})
	}

	result := db.DB.Limit(limit).Offset(skip).Find(&customers)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Database error",
			"error":   result.Error.Error(),
		})
	}

	db.DB.Model(&Models.Customer{}).Count(&count)

	if len(customers) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "No customers found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Customer list API",
		"data":    customers,
		"count":   count,
	})
}
