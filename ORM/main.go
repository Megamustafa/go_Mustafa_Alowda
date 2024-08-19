package main

import (
	"orm/controllers"
	"orm/database"

	"github.com/labstack/echo/v4"
)

func main() {

	database.InitDB()

	database.MigrateDB()
	packageController := controllers.InitPackageController()
	e := echo.New()
	e.GET("/api/v1/packages", packageController.GetAll)
	e.GET("/api/v1/packages/:id", packageController.GetByID)
	e.POST("/api/v1/packages", packageController.Create)
	e.PUT("/api/v1/packages/:id", packageController.Update)
	e.DELETE("/api/v1/packages/:id", packageController.Delete)

	e.Logger.Fatal(e.Start(":1323"))

}
