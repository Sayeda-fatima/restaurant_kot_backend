package model

import "time"

type OrderItem struct {
	ID                uint         `json:"id" gorm:"primaryKey"`
	OrganizationID    uint         `json:"organization_id"`
	Organization      Organization `gorm:"foreignKey:OrganizationID;references:ID" json:"-" validate:"-"`
	OrderID           uint         `json:"order_id" validate:"required"`
	//Order             Order        `gorm:"foreignKey:OrderID;references:ID" json:"-" validate:"-"`
	ProductID         uint         `json:"product_id" validate:"required"`
	Product           Product      `gorm:"foreignKey:ProductID;references:ID" json:"-" validate:"-"`
	ProductQuantity   float64      `json:"product_quantity" validate:"required"`
	UnitProductPrice  float64      `json:"unit_product_price" validate:"required"`
	Tax               float64      `json:"tax" validate:"required"`
	TotalProductPrice float64      `json:"total_product_price" validate:"required"`
	CreatedAt         time.Time    `json:"created_at"`
	UpdatedAt         time.Time    `json:"updated_at"`
}

type OrderItemResponse struct {
	ID                uint    `json:"id"`
	OrganizationID    uint    `json:"organization_id"`
	OrderID           uint    `json:"order_id"`
	ProductID         uint    `json:"product_id"`
	ProductQuantity   float64 `json:"product_quantity"`
	UnitProductPrice  float64 `json:"unit_product_price"`
	Tax               float64 `json:"tax" validate:"required"`
	TotalProductPrice float64 `json:"total_product_price"`
}
