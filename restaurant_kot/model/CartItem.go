package model

import "time"

type CartItem struct {
	OrganizationID uint         `json:"organization_id"`
	Organization   Organization `json:"-" gorm:"foreignKey:OrganizationID;references:ID"`
	RestaurantID   uint         `json:"restaurant_id"`
	Restaurant     Restaurant   `json:"-" gorm:"foreignKey:RestaurantID;references:ID"`
	ID             uint         `json:"id" gorm:"primaryKey"`
	MenuItemID     uint         `json:"menu_item_id"`
	MenuItem       MenuItem     `json:"menu_item" gorm:"foreignKey:MenuItemID;references:ID"`
	ItemQuantity   int          `json:"item_quantity"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
}

type CartItemResponse struct {
	OrganizationID uint     `json:"organization_id"`
	RestaurantID   uint     `json:"restaurant_id"`
	ID             uint     `json:"id"`
	MenuItemID     uint     `json:"menu_item_id"`
	MenuItem       MenuItem `json:"menu_item"`
	ItemQuantity   int      `json:"ItemQuantity"`
}
