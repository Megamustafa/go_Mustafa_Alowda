package main

import (
	"assignment1/config"
	"assignment1/models"
	"assignment1/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.Connect()
	models.Migrate(config.DB)
	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
