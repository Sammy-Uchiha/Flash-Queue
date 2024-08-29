package controllers

import (
	"fmt"
	"net/http"

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
	// Get the id off the url
	id := c.Param("id")

	// Find the customer to be deleted
	var customer models.Customer
	result := initializers.DB.First(&customer, id)
	if result.Error != nil {
		c.Status(404)
		return
	}

	// Delete the customer
	initializers.DB.Delete(&customer)

	// Update the positions of all other customers
	var customers []models.Customer
	initializers.DB.Where("id <> ?", id).Order("position ASC").Find(&customers)

	for i, c := range customers {
		c.Position = i + 1
		initializers.DB.Save(&c)
	}

	// Respond
	c.Status(200)
}

func CustomersStartPage(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"post": "Healthy",
	})
}
