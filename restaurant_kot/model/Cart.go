package model

import "time"

type Cart struct {
	OrganizationID uint            `json:"organization_id" gorm:"not null"`
	Organization   Organization    `json:"-" gorm:"foreignKey:OrganizationID;references:ID" validate:"-"`
	RestaurantID   uint            `json:"restaurant_id" gorm:"not null"`
	Restaurant     Restaurant      `json:"-" gorm:"foreignKey:RestaurantID;references:ID" validate:"-"`
	ID             uint            `json:"id" gorm:"primaryKey"`
	TableID        uint            `json:"table_id" gorm:"not null" validate:"required"`
	Table          RestaurantTable `json:"-" gorm:"foreignKey:TableID;references:ID" validate:"-"`
	TotalQuantity  int             `json:"total_quantity" gorm:"not null;type:int(11)"`
	CartType       string          `json:"cart_type" gorm:"not null;type:enum('dine_in', 'takeaway', 'delivery')"`
	CartStatus     string          `json:"cart_status" gorm:"not null;type:enum('active', 'ready_for_checkout', 'checked_out')"`
	CartItems      []CartItem      `json:"cart_items" gorm:"foreignKey:CartID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}

type CartResponse struct {
	OrganizationID uint       `json:"organization_id"`
	RestaurantID   uint       `json:"restaurant_id"`
	ID             uint       `json:"id"`
	TableID        uint       `json:"table_id"`
	TotalQuantity  int        `json:"total_quantity"`
	CartType       string     `json:"cart_type"`
	CartStatus     string     `json:"cart_status"`
	CartItems      []CartItem `json:"cart_items"`
}
