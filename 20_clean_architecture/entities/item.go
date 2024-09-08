package entities

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name        string
	Description string
	WishlistID  uint
}