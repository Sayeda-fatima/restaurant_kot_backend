package model

import "time"

type Cart struct {
	OrganizationID uint            `json:"organization_id" gorm:"not null"`
	Organization   Organization    `json:"-" gorm:"foreignKey:OrganizationID;references:ID"`
	RestaurantID   uint            `json:"restaurant_id" gorm:"not null"`
	Restaurant     Restaurant      `json:"-" gorm:"foreignKey:RestaurantID;references:ID"`
	ID             uint            `json:"id" gorm:"primaryKey"`
	TableID        uint            `json:"table_id" gorm:"not null" validate:"required"`
	Table          RestaurantTable `json:"-" gorm:"foreignKey:TableID;references:ID"`
	TotalQuantity  int             `json:"total_quantity" gorm:"not null;type:int(11)"`
	CartItems      []CartItem      `json:"cart_items" gorm:"foreignKey:CartID"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}

type CartResponse struct {
	OrganizationID uint       `json:"organization_id"`
	RestaurantID   uint       `json:"restaurant_id"`
	ID             uint       `json:"id"`
	TableID        uint       `json:"table_id"`
	TotalQuantity  int        `json:"total_quantity"`
	CartItems      []CartItem `json:"cart_items"`
}
