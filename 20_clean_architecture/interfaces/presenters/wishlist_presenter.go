package presenters

import (
	"net/http"

	"go-wishlist-api/entities"

	"github.com/labstack/echo/v4"
)

// WishlistPresenter defines the interface for formatting the output
type WishlistPresenter interface {
	Format(wishlist *entities.Wishlist, ctx echo.Context) error
	FormatList(wishlists []*entities.Wishlist, ctx echo.Context) error
}

type wishlistPresenter struct{}

// NewWishlistPresenter creates a new WishlistPresenter
func NewWishlistPresenter() WishlistPresenter {
	return &wishlistPresenter{}
}

// Format formats a single wishlist into JSON and sends it to the client
func (p *wishlistPresenter) Format(wishlist *entities.Wishlist, ctx echo.Context) error {
	if wishlist == nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"message": "Wishlist not found"})
	}
	return ctx.JSON(http.StatusOK, wishlist)
}

// FormatList formats a list of wishlists into JSON and sends it to the client
func (p *wishlistPresenter) FormatList(wishlists []*entities.Wishlist, ctx echo.Context) error {
	if len(wishlists) == 0 {
		return ctx.JSON(http.StatusNotFound, map[string]string{"message": "No wishlists found"})
	}
	return ctx.JSON(http.StatusOK, wishlists)
}
