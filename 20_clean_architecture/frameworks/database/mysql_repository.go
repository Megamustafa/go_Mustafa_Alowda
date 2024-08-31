package database

import (
	"errors"
	"go-wishlist-api/entities"

	"gorm.io/gorm"
)

type GormWishlistRepository struct {
	db *gorm.DB
}

func NewGormWishlistRepository(db *gorm.DB) *GormWishlistRepository {
	return &GormWishlistRepository{db: db}
}

func (r *GormWishlistRepository) Save(wishlist *entities.Wishlist) error {
	result := r.db.Create(wishlist)
	return result.Error
}

func (r *GormWishlistRepository) FindByID(id string) (*entities.Wishlist, error) {
	var wishlist entities.Wishlist
	result := r.db.First(&wishlist, "id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &wishlist, result.Error
}

func (r *GormWishlistRepository) Update(wishlist *entities.Wishlist) error {
	result := r.db.Save(wishlist)
	return result.Error
}

func (r *GormWishlistRepository) Delete(id string) error {
	result := r.db.Delete(&entities.Wishlist{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return result.Error
}
