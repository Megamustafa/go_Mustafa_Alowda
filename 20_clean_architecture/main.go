package main

import (
	"log"

	"go-wishlist-api/config"
	"go-wishlist-api/entities"
	"go-wishlist-api/frameworks/database"
	"go-wishlist-api/frameworks/web"
	"go-wishlist-api/interfaces/controllers"
	"go-wishlist-api/interfaces/presenters"
	"go-wishlist-api/middleware"
	"go-wishlist-api/usecases"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/labstack/echo/v4"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize MySQL database connection
	dsn := cfg.DBUser + ":" + cfg.DBPassword + "@tcp(" + cfg.DBHost + ":" + cfg.DBPort + ")/" + cfg.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	db.AutoMigrate(&entities.Wishlist{}, &entities.Item{})

	// Initialize repository
	wishlistRepo := database.NewGormWishlistRepository(db)

	// Initialize use case
	wishlistUsecase := usecases.NewWishlistUsecase(wishlistRepo)

	// Initialize presenter
	wishlistPresenter := presenters.NewWishlistPresenter()

	// Initialize controller with use case and presenter
	wishlistController := controllers.NewWishlistController(wishlistUsecase, wishlistPresenter)

	// Initialize Echo router
	e := web.NewEchoRouter(wishlistController)

	protectedRoutes := e.Group("")
	protectedRoutes.Use(middleware.JWTMiddleware())

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
