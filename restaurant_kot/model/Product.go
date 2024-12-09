package model

import "time"

type Product struct {
	ID             uint         `json:"id" gorm:"primaryKey"`
	OrganizationID uint         `json:"organization_id" gorm:"not null"`
	Organization   Organization `json:"-" gorm:"foreignKey:OrganizationID;references:ID" validate:"-"`
	RestaurantID   uint         `json:"restaurant_id" gorm:"not null"`
	Restaurant     Restaurant   `json:"-" gorm:"foreignKey:RestaurantID;references:ID" validate:"-"`
	Name           string       `json:"name" gorm:"not null;size:255" validate:"required"`
	Description    string       `json:"description" validate:"omitempty"`
	Category       string       `json:"category" gorm:"not null;size:255" validate:"required"`
	UnitOfMeasure  string       `json:"unit_of_measure" gorm:"not null;size:255" validate:"required"`
	UnitCost       int          `json:"unit_cost" gorm:"not null;type:int(11)" validate:"required"`
	Quantity       int          `json:"quantity" gorm:"not null;type:int(11)" validate:"required"`
	InventoryValue int          `json:"inventory_value" gorm:"not null;type:int(11)"`
	//Recipes        []Recipe     `gorm:"many2many:recipe_products"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	IsDeleted      bool         `json:"is_deleted" gorm:"not null;default:0"`
}

type ProductResponse struct {
	ID             uint   `json:"id"`
	OrganizationID uint   `json:"organization_id"`
	RestaurantID   uint   `json:"restaurant_id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Category       string `json:"category"`
	UnitOfMeasure  string `json:"unit_of_measure"`
	UnitCost       int    `json:"unit_cost"`
	Quantity       int    `json:"quantity"`
	InventoryValue int    `json:"inventory_value"`
}
