package routes

import (
	"package-api/controllers"
	"package-api/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.POST("/api/v1/register", controllers.Register)
	e.POST("/api/v1/login", controllers.Login)

	r := e.Group("/api/v1/packages")
	r.Use(middleware.JWTMiddleware())

	r.GET("", controllers.GetPackages)
	r.GET("/:id", controllers.GetPackageByID)
	r.POST("", controllers.CreatePackage)
	r.PUT("/:id", controllers.UpdatePackage)
	r.DELETE("/:id", controllers.DeletePackage)
}
