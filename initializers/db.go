package initializers

import (
	"fmt"
	"goreact/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Unable to connect to database")
	}

}

func SyncDB() {
	DB.AutoMigrate(&models.Post{})

}
