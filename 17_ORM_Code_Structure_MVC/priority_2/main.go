package main

import (
	"assignment2/database"
	"assignment2/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize Echo
	e := echo.New()

	// Connect to the database
	database.Connect()

	// Set up routes
	routes.SetupRoutes(e)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
