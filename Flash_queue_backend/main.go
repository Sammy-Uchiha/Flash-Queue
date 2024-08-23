package main

import (
	"github.com/Flash_queue_backend/controllers"
	"github.com/Flash_queue_backend/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.MirageteDB()
}

func main() {
	r := gin.Default()

	r.POST("/customers", controllers.CustomersCreate)
	// r.PUT("/customers/:id", controllers.CustomerUpdate)
	r.GET("/customers", controllers.CustomersIndex)
	r.GET("/customers/:id", controllers.CustomerShow)
	r.DELETE("/customers/:id", controllers.CustomerDelete)
	r.GET("/", controllers.CustomersStartPage)

	r.Run()
}
