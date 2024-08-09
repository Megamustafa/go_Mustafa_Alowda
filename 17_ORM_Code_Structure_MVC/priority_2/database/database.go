package database

import (
	"assignment2/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

// Connect to the database
func Connect() {
	var err error
	DB, err = gorm.Open("sqlite3", "fruits.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB.AutoMigrate(&models.Fruit{})
}
