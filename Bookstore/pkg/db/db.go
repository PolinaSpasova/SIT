package db

import (
	"api/pkg/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	var DBname = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "polina", "password", "bookstore")
	db, err := gorm.Open(postgres.Open(DBname), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Author{})
	return db
}
