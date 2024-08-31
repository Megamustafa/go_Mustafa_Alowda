package controllers

import (
	"go-wishlist-api/entities"
	"go-wishlist-api/interfaces/presenters"
	"go-wishlist-api/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WishlistController struct {
	Usecase   usecases.WishlistUsecase
	Presenter presenters.WishlistPresenter
}

func NewWishlistController(usecase usecases.WishlistUsecase, presenter presenters.WishlistPresenter) *WishlistController {
	return &WishlistController{Usecase: usecase, Presenter: presenter}
}

func (c *WishlistController) CreateWishlist(ctx echo.Context) error {
	var wishlist entities.Wishlist
	if err := ctx.Bind(&wishlist); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request")
	}

	err := c.Usecase.CreateWishlist(&wishlist)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.Presenter.Format(&wishlist, ctx)
}

func (c *WishlistController) GetWishlistByID(ctx echo.Context) error {
	id := ctx.Param("id")
	wishlist, err := c.Usecase.GetWishlistByID(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.Presenter.Format(wishlist, ctx)
}
