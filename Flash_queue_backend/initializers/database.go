package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() (*gorm.DB, error) {
	var err error
	dsn := os.Getenv("DB_CONNECT")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
		￼￼
		
	}
	return DB, nil
}
