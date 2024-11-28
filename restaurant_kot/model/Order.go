package model

import "time"

type Order struct {
	ID             uint            `json:"id" gorm:"primaryKey"`
	OrganizationID uint            `json:"organization_id" gorm:"not null"`
	Organization   Organization    `json:"-" gorm:"foreignKey:OrganizationID;references:ID" validate:"-"`
	RestaurantID   uint            `json:"restaurant_id" gorm:"not null"`
	Restaurant     Restaurant      `json:"-" gorm:"foreignKey:RestaurantID;references:ID" validate:"-"`
	TableID        uint            `json:"table_id" gorm:"not null" validate:"required"`
	Table          RestaurantTable `json:"-" gorm:"foreignKey:TableID;references:ID" validate:"-"`
	TotalItemPrice int             `json:"total_item_price" gorm:"not null;type:int(11)" validate:"required"`
	Tax            int             `json:"tax" gorm:"not null;type:int(11)" validate:"required"`
	ServiceCharge  int             `json:"service_charge" gorm:"not null;type:int(11)" validate:"required"`
	Tip            int             `json:"tip" gorm:"not null;type:int(11)" validate:"required"`
	TotalPrice     int             `json:"total_price" gorm:"not null;type:int(11)" validate:"required"`
	OrderItems     []OrderItem     `json:"order_items" gorm:"foreignKey:OrderID"`
	OrderStatus    string          `json:"order_status" gorm:"not null;type:enum('placed','in_progress','ready_to_serve','served','paid')" validate:"required"` // set: placed, in_progress, ready_to_serve, served, paid
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	IsDeleted      bool            `json:"is_deleted" gorm:"not null;default:0"`
}

type OrderResponse struct {
	ID             uint        `json:"id"`
	OrganizationID uint        `json:"organization_id"`
	RestaurantID   uint        `json:"restaurant_id"`
	TableID        uint        `json:"table_id"`
	TotalItemPrice int         `json:"total_item_price"`
	Tax            int         `json:"tax"`
	ServiceCharge  int         `json:"service_charge"`
	Tip            int         `json:"tip"`
	TotalPrice     int         `json:"total_price"`
	OrderItems     []OrderItem `json:"order_items"`
	OrderStatus    string      `json:"order_status"`
}
