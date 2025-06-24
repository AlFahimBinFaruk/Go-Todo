package main

import (
	"go-todo/initializers"
	"go-todo/models"
	"log"

	"gorm.io/gorm"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func InitUUIDExtension(db *gorm.DB) {
	err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error
	if err != nil {
		log.Fatalf("Failed to create uuid-ossp extension: %v", err)
	}
}

func main() {
	//Ensure uuid-ossp extension is installed in db.
	InitUUIDExtension(initializers.DB)
	err := initializers.DB.AutoMigrate(&models.TODO{})
	if err != nil {
		log.Fatalf("Failed to connect to DB!!: %v", err)
	}
}
