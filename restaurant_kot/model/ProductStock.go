package model

import "time"

type ProductStock struct {
	ID             uint         `json:"id" gorm:"primaryKey"`
	OrganizationID uint         `json:"organization_id" gorm:"not null;foreignKey"`
	Organization   Organization `json:"-" gorm:"foreignKey:OrganizationID;references:ID" validate:"-"`
	RestaurantID   uint         `json:"restaurant_id" gorm:"not null;foreignKey"`
	Restaurant     Restaurant   `json:"-" gorm:"foreignKey:RestaurantID;references:ID" validate:"-"`
	ProductID      uint         `json:"product_id" gorm:"not null"`
	Product        Product      `json:"-" gorm:"foreignKey:ProductID;references:ID" validate:"-"`
	StartingStock       int64        `json:"starting_stock" gorm:"not null" validate:"required"`
	StockBeforeUpdate float64      `json:"stock_before_update" validate:"omitempty"`
	UpdateQuantity    float64      `json:"update_quantity" validate:"required"`
	UpdateType        string       `json:"update_type" validate:"required"`
	StockAfterUpdate  float64      `json:"stock_after_update" validate:"omitempty"`
	UnitCost       int64        `json:"unit_cost" gorm:"not null" validate:"required"`
	TotalCost      int64        `json:"total_cost" gorm:"not null;type:int(11)" validate:"required"`
	CreatedAt      time.Time	`json:"created_at"`
	UpdatedAt		time.Time	`json:"updated_at"`
}