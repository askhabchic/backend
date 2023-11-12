package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"backend/cmd/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres dbname=backend port=5432"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&models.Address{})
	database.AutoMigrate(&models.Image{})
	database.AutoMigrate(&models.Supplier{})
	database.AutoMigrate(&models.Client{})
	database.AutoMigrate(&models.Product{})

	DB = database
}
