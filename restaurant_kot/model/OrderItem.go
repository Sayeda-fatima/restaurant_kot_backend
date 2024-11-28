package model

import "time"

type OrderItem struct {
	ID             uint         `json:"id" gorm:"primaryKey"`
	OrganizationID uint         `json:"organization_id"`
	Organization   Organization `json:"-" gorm:"foreignKey:OrganizationID;references:ID" validate:"-"`
	RestaurantID   uint         `json:"restaurant_id"`
	Restaurant     Restaurant   `json:"-" gorm:"foreignKey:RestaurantID;references:ID" validate:"-"`
	MenuItemID     uint         `json:"menu_item_id" validate:"required"`
	MenuItem       MenuItem     `json:"-" gorm:"foreignKey:MenuItemID;references:ID" validate:"-"`
	ItemQuantity   int          `json:"item_quantity" validate:"required"`
	UnitItemPrice  int          `json:"unit_item_price" validate:"required"`
	TotalItemPrice int          `json:"total_item_price" validate:"required"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	IsDeleted      bool         `json:"is_deleted"`
}

type OrderItemResponse struct {
	ID             uint `json:"id" gorm:"primaryKey"`
	OrganizationID uint `json:"organization_id"`
	RestaurantID   uint `json:"restaurant_id"`
	MenuItemID     uint `json:"menu_item_id"`
	ItemQuantity   int  `json:"item_quantity"`
	UnitItemPrice  int  `json:"unit_item_price"`
	TotalItemPrice int  `json:"total_item_price"`
}
