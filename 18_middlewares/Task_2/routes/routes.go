package routes

import (
	"post-manager/controllers"
	"post-manager/middlewares"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	e.POST("/api/v1/register", controllers.RegisterUser(db))
	e.POST("/api/v1/login", controllers.LoginUser(db))
	e.GET("/api/v1/posts", controllers.GetPosts(db))
	e.GET("/api/v1/posts/:id", controllers.GetPost(db), middlewares.AuthMiddleware)
	e.POST("/api/v1/posts", controllers.CreatePost(db), middlewares.AuthMiddleware)
	e.PUT("/api/v1/posts/:id", controllers.UpdatePost(db), middlewares.AuthMiddleware)
	e.DELETE("/api/v1/posts/:id", controllers.DeletePost(db), middlewares.AuthMiddleware)
}
