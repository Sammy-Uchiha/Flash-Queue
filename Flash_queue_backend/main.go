package main

import (
	"log"

	"github.com/Flash_queue_backend/controllers"
	"github.com/Flash_queue_backend/initializers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.MirageteDB()
}

func main() {
	r := gin.Default()

	// Configure the CORS middleware

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.POST("/customers", controllers.CustomersCreate)
	// r.PUT("/customers/:id", controllers.CustomerUpdate)
	r.GET("/customers", controllers.CustomersIndex)
	r.GET("/customers/:id", controllers.CustomerShow)
	r.DELETE("/customers/:id", controllers.CustomerDelete)
	r.GET("/", controllers.CustomersStartPage)

	err := r.Run()
	if err != nil {
		log.Fatalf("Error running the server: %v", err)
	}
}
