package web

import (
	"go-wishlist-api/interfaces/controllers"

	"github.com/labstack/echo/v4"
)

func NewEchoRouter(controller *controllers.WishlistController) *echo.Echo {
	e := echo.New()

	e.POST("/wishlists", controller.CreateWishlist)
	e.GET("/wishlists/:id", controller.GetWishlistByID)

	return e
}
