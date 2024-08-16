package controllers

import (
	"net/http"

	"github.com/Flash_queue_backend/initializers"
	"github.com/Flash_queue_backend/models"
	"github.com/gin-gonic/gin"
)

func CustomersCreate(c *gin.Context) {
	// Get data off a req body
	var body struct {
		Name   string
		Number string
	}

	c.Bind(&body)

	// Create a post
	customer := models.Customer{Name: body.Name, Number: body.Number}
	result := initializers.DB.Create(&customer)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// Return it
	c.JSON(200, gin.H{
		"customer": customer,
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

	// Delete the post
	initializers.DB.Delete(&models.Customer{}, id)

	// Respond
	c.Status(200)
}

func CustomersStartPage(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"post": "Healthy",
	})
}
