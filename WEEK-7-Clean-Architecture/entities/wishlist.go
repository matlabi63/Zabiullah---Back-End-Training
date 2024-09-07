package entities

import "gorm.io/gorm"

type Wishlist struct {
	gorm.Model
	Name        string
	Description string
	Items       []Item
}