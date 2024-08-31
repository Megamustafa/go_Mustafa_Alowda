package repositories

import "go-wishlist-api/entities"

type WishlistRepository interface {
	Save(wishlist *entities.Wishlist) error
	FindByID(id string) (*entities.Wishlist, error)
	Update(wishlist *entities.Wishlist) error
	Delete(id string) error
}
