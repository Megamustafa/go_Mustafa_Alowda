package routes

import (
	"assignment2/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/api/v1/fruits", controllers.GetFruits)
	e.GET("/api/v1/fruits/:id", controllers.GetFruitByID)
	e.POST("/api/v1/fruits", controllers.CreateFruit)
	e.DELETE("/api/v1/fruits/:id", controllers.DeleteFruit)
}
