package model

import "time"

type OrderItem struct {
	ID             uint         `json:"id" gorm:"primaryKey"`
	OrganizationID uint         `json:"organization_id" gorm:"not null"`
	Organization   Organization `json:"-" gorm:"foreignKey:OrganizationID;references:ID" validate:"-"`
	RestaurantID   uint         `json:"restaurant_id" gorm:"not null"`
	Restaurant     Restaurant   `json:"-" gorm:"foreignKey:RestaurantID;references:ID" validate:"-"`
	OrderID        uint         `json:"order_id" gorm:"not null" validate:"required"`
	MenuItemID     uint         `json:"menu_item_id" gorm:"not null" validate:"required"`
	MenuItem       MenuItem     `json:"-" gorm:"foreignKey:MenuItemID;references:ID" validate:"-"`
	ItemQuantity   int          `json:"item_quantity" gorm:"not null;type:int(11)" validate:"required"`
	UnitItemPrice  int          `json:"unit_item_price" gorm:"not null;type:int(11)" validate:"required"`
	TotalItemPrice int          `json:"total_item_price" gorm:"not null;type:int(11)" validate:"required"`
	ItemStatus		string		`json:"item_status" gorm:"not null;type:enum('normal', 'replaced', 'complementary')"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	IsDeleted      bool         `json:"is_deleted" gorm:"not null;default:0"`
}

type OrderItemResponse struct {
	ID             uint `json:"id" gorm:"primaryKey"`
	OrganizationID uint `json:"organization_id"`
	RestaurantID   uint `json:"restaurant_id"`
	OrderID        uint `json:"order_id"`
	MenuItemID     uint `json:"menu_item_id"`
	ItemQuantity   int  `json:"item_quantity"`
	UnitItemPrice  int  `json:"unit_item_price"`
	TotalItemPrice int  `json:"total_item_price"`
	ItemStatus		string	`json:"item_status"`
}
