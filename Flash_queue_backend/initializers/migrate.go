package initializers

import (
	"log"

	"github.com/Flash_queue_backend/models"
)

func MirageteDB() error {
	if err := DB.AutoMigrate(&models.Customer{}); err != nil {
		log.Printf("Error migrating database: %v", err)
		// return error for handling elsewhere
		return err
	}
	// return nil if successful
	return nil
}
