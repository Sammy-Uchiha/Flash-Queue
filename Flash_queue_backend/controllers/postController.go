package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Flash_queue_backend/initializers"
	"github.com/Flash_queue_backend/models"
	"github.com/gin-gonic/gin"
)

func CustomersCreate(c *gin.Context) {
	// Get data off a req body
	var customer struct {
		Name     string
		Number   string
		Position int
	}

	if err := c.Bind(&customer); err != nil {
		// Log the error message
		fmt.Println("Error binding request body:", err)
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	// Get the current maximum position
	var maxPosition int
	result := initializers.DB.Model(&models.Customer{}).Select("MAX(position)").Scan(&maxPosition)

	// Determine the new position based on the current maximum
	newPosition := 1
	if result.Error == nil && maxPosition > 0 {
		newPosition = maxPosition + 1
	}

	// Create a new customer with the new position
	newCustomer := models.Customer{
		Name:     customer.Name,
		Number:   customer.Number,
		Position: newPosition,
	}
	createResult := initializers.DB.Create(&newCustomer)

	if createResult.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"customer": newCustomer,
	})
}

func CustomersIndex(c *gin.Context) {
	// Get the customer
	var customer []models.Customer
	initializers.DB.Find(&customer)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": customer,
	})
}

func CustomerShow(c *gin.Context) {
	//Get id off url
	id := c.Param("id")
	// Get the customer
	var customer models.Customer
	initializers.DB.First(&customer, id)

	// Respond with them
	c.JSON(200, gin.H{
		"post": customer,
	})
}

// func CustomerUpdate(c *gin.Context) {
// 	// Get id of url
// 	id := c.Param("id")
// 	// Get the data off req body
// 	var body struct {
// 		Name   string
// 		Number string
// 	}

// 	c.Bind(&body)

// 	// Find the post were updating
// 	var customer models.Customer
// 	initializers.DB.First(&customer, id)

// 	// Update it
// 	initializers.DB.Model(&customer).Updates(models.Customer{
// 		Name:   body.Name,
// 		Number: body.Number,
// 	})
// 	// Respond with it
// 	c.JSON(200, gin.H{
// 		"post": customer,
// 	})
// }

func CustomerDelete(c *gin.Context) {
	positionStr := c.Param("position")         // Get position as a string
	position, err := strconv.Atoi(positionStr) // Convert to integer
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid position"})
		return
	}

	var customer models.Customer

	// Find the customer to be deleted by position
	result := initializers.DB.Where("position = ?", position).First(&customer)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Customer not found"})
		return
	}

	// Delete the customer
	if err := initializers.DB.Delete(&customer).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete customer"})
		return
	}

	// Update the positions of all other customers
	var customers []models.Customer
	if err := initializers.DB.Where("position > ?", position).Order("position ASC").Find(&customers).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve remaining customers"})
		return
	}

	// Update positions
	for i, cust := range customers {
		cust.Position = position + i // Correctly adjust position based on the deleted customer's position
		if err := initializers.DB.Save(&cust).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to update customer positions"})
			return
		}
	}

	// Respond with success message
	c.JSON(200, gin.H{"message": "Customer deleted successfully"})
}

func CustomersStartPage(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"post": "Healthy",
	})
}
