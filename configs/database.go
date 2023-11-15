package configs

import (
	"serviceemployee/models"
	"log"

  "gorm.io/gorm"
  "gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(postgres.Open("host=localhost port=5432 user=postgres dbname=serviceemployee password=postgres sslmode=disable"), &gorm.Config{})
	if err != nil {
			panic("failed to connect database")
	}

	db.AutoMigrate(&models.Employee{})

	DB = db
	log.Println("Database connected")
};