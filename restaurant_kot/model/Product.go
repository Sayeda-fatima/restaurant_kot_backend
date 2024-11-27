package model

import "time"

type Product struct {
	ID             uint         `json:"id" gorm:"primaryKey"`
	OrganizationID uint         `json:"organization_id"`
	Organization   Organization `json:"-" gorm:"foreignKey:OrganizationID;references:ID" validate:"-"`
	RestaurantID   uint         `json:"restaurant_id"`
	Restaurant     Restaurant   `json:"-" gorm:"foreignKey:RestaurantID;references:ID" validate:"-"`
	Name           string       `json:"name" validate:"required"`
	Description    string       `json:"description" validate:"omitempty"`
	Category       string       `json:"category" validate:"required"`
	UnitOfMeasure  string       `json:"unit_of_measure" validate:"required"`
	UnitCost       int          `json:"unit_cost" validate:"required"`
	Quantity       int          `json:"quantity" validate:"required"`
	InventoryValue int          `json:"inventory_value"`
	Recipes		   []Recipe		`gorm:"many2many:recipe_products"`
	CreatedAt      time.Time	`json:"created_at"`
	UpdatedAt	   time.Time	`json:"updated_at"`
}

type ProductResponse struct{
	ID				uint		`json:"id"`
	OrganizationID 	uint		`json:"organization_id"`
	RestaurantID	uint		`json:"restaurant_id"`
	Name			string		`json:"name"`
	Description		string		`json:"description"`
	Category		string		`json:"category"`
	UnitOfMeasure	string		`json:"unit_of_measure"`
	UnitCost		int			`json:"unit_cost"`
	Quantity		int			`json:"quantity"`
	InventoryValue	int			`json:"inventory_value"`
}
