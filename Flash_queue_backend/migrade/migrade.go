package main

import (
	"github.com/Flash_queue_backend/initializers"
	"github.com/Flash_queue_backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Customer{})
}
