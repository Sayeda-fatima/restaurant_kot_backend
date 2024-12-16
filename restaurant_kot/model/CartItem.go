package model

import "time"

type CartItem struct {
	OrganizationID uint         `json:"organization_id" gorm:"not null"`
	Organization   Organization `json:"-" gorm:"foreignKey:OrganizationID;references:ID" validate:"-"`
	RestaurantID   uint         `json:"restaurant_id" gorm:"not null"`
	Restaurant     Restaurant   `json:"-" gorm:"foreignKey:RestaurantID;references:ID" validate:"-"`
	ID             uint         `json:"id" gorm:"primaryKey"`
	CartID         uint         `json:"cart_id" gorm:"not null" validate:"required"`
	MenuItemID     uint         `json:"menu_item_id" gorm:"not null"`
	MenuItem       MenuItem     `json:"menu_item" gorm:"foreignKey:MenuItemID;references:ID" validate:"-"`
	ItemQuantity   int          `json:"item_quantity" gorm:"not null;type:int(11)"`
	ItemStatus     string       `json:"item_status" gorm:"not null;type:enum('pending', 'sent_to_kitchen', 'ready_to_serve', 'served', 'cancelled', 'replaced', 'complementary')"`
	Note           string       `json:"note" validate:"omitempty"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
}

type CartItemResponse struct {
	OrganizationID uint     `json:"organization_id"`
	RestaurantID   uint     `json:"restaurant_id"`
	ID             uint     `json:"id"`
	CartID         uint     `json:"cart_id"`
	MenuItemID     uint     `json:"menu_item_id"`
	MenuItem       MenuItem `json:"menu_item"`
	ItemQuantity   int      `json:"ItemQuantity"`
	ItemStatus     string   `json:"item_status"`
	Note           string   `json:"note"`
}
