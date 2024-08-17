package main

import (
	"package-api/middleware"
	"package-api/models"
	"package-api/routes"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()

	// Middleware
	middleware.LogMiddleware(e)

	// Initialize DB
	db, _ := gorm.Open("sqlite3", "package.db")
	db.AutoMigrate(&models.Package{}, &models.User{})

	// Register routes
	routes.RegisterRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
