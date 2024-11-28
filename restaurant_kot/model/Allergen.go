package model

import "time"

type Allergen struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	RecipeID     uint      `json:"recipe_id" gorm:"not null"`
	Recipe       Recipe    `json:"-" gorm:"foreignKey:RecipeID;references:ID"`
	MenuItemID   uint      `json:"menu_item_id" gorm:"not null"`
	MenuItem     MenuItem  `json:"-" gorm:"foreignKey:MenuItemID;references:ID"`
	AllergenName string    `json:"allergen_name" gorm:"not null;size:255" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	IsDeleted    bool      `json:"is_deleted" gorm:"not null;default:0"`
}

type AllergenResponse struct {
	ID           uint   `json:"id"`
	RecipeID     uint   `json:"recipe_id"`
	MenuItemID   uint   `json:"menu_item_id"`
	AllergenName string `json:"allergen_name"`
}
