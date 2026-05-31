package model

import "time"

type MenuAllergen struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	MenuItemID uint      `json:"menu_item_id" gorm:"not null"`
	MenuItem   MenuItem  `json:"-" gorm:"foreignKey:MenuItemID;references:ID" validate:"-"`
	AllergenID uint      `json:"allergen_id" gorm:"not null"`
	Allergen   Allergen  `json:"allergen" gorm:"foreignKey:AllergenID;references:ID" validate:"-"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type MenuAllergenResponse struct {
	ID         uint     `json:"id"`
	MenuItemID uint     `json:"menu_item_id"`
	AllergenID uint     `json:"allergen_id"`
	Allergen   Allergen `json:"allergen"`
}
