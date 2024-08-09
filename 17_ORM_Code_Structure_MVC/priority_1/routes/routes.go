package routes

import (
	"assignment1/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/packages", controllers.GetPackages)
	e.GET("/api/v1/packages/:id", controllers.GetPackageByID)
	e.POST("/api/v1/packages", controllers.CreatePackage)
	e.PUT("/api/v1/packages/:id", controllers.UpdatePackage)
	e.DELETE("/api/v1/packages/:id", controllers.DeletePackage)
}
