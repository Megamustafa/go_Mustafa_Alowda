package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("packages.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}
}
