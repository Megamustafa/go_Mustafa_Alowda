package usecases

import (
	"go-wishlist-api/entities"
	"go-wishlist-api/interfaces/repositories"
)

type WishlistUsecase interface {
	CreateWishlist(wishlist *entities.Wishlist) error
	GetWishlistByID(id string) (*entities.Wishlist, error)
	UpdateWishlist(wishlist *entities.Wishlist) error
	DeleteWishlist(id string) error
}

type wishlistUsecase struct {
	wishlistRepo repositories.WishlistRepository
}

func NewWishlistUsecase(repo repositories.WishlistRepository) WishlistUsecase {
	return &wishlistUsecase{wishlistRepo: repo}
}

func (u *wishlistUsecase) CreateWishlist(wishlist *entities.Wishlist) error {
	return u.wishlistRepo.Save(wishlist)
}

func (u *wishlistUsecase) GetWishlistByID(id string) (*entities.Wishlist, error) {
	return u.wishlistRepo.FindByID(id)
}

func (u *wishlistUsecase) UpdateWishlist(wishlist *entities.Wishlist) error {
	return u.wishlistRepo.Update(wishlist)
}

func (u *wishlistUsecase) DeleteWishlist(id string) error {
	return u.wishlistRepo.Delete(id)
}
