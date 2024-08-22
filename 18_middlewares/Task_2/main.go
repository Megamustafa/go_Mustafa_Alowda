package main

import (
	"log"
	"os"
	"post-manager/middlewares"
	"post-manager/models"
	"post-manager/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Middleware
	middlewares.LogMiddleware(e)
	middlewares.RateLimiterMiddleware(e)

	// Database connection
	db, err := gorm.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME")+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.AutoMigrate(&models.Post{}, &models.User{})

	// Routes
	routes.InitRoutes(e, db)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
