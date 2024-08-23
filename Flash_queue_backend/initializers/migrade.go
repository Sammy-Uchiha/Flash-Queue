package initializers

import "github.com/Flash_queue_backend/models"

func MirageteDB() {
	DB.AutoMigrate(&models.Customer{})
}
