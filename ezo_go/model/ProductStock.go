package model

import "time"

type ProductStock struct {
	ID                       uint         `json:"id" gorm:"primaryKey"`
	OrganizationID           uint         `json:"organization_id"`
	Organization             Organization `json:"-" gorm:"foreignKey:OrganizationID;references:ID"`
	OrderID                  uint         `json:"order_id" validate:"required"`
	ProductID                uint         `json:"product_id" validate:"required"`
	Product                  Product      `json:"-" gorm:"foreignKey:ProductID;references:ID"`
	ProductName              string       `json:"product_name" validate:"required"`
	ProductStockBeforeUpdate float64      `json:"product_stock_before_update" validate:"required"`
	ProductUpdateQuantity    float64      `json:"product_update_quantity" validate:"required"`
	ProductUpdateType        string       `json:"product_update_type" validate:"required"`
	ProductStockAfterUpdate  float64      `json:"product_stock_after_update" validate:"required"`
	IsDeleted                bool         `json:"is_deleted" validate:"-"`
	CreatedAt                time.Time    `json:"created_at"`
	UpdatedAt                time.Time    `json:"updated_at"`
}

type ProductStockResponse struct {
	ID                       uint    `json:"id" gorm:"primaryKey"`
	OrganizationID           uint    `json:"organization_id"`
	OrderID                  uint    `json:"order_id"`
	ProductID                uint    `json:"product_id"`
	ProductName              string  `json:"product_name"`
	ProductStockBeforeUpdate float64 `json:"product_stock_before_update"`
	ProductUpdateQuantity    float64 `json:"product_update_quantity"`
	ProductUpdateType        string  `json:"product_update_type"`
	ProductStockAfterUpdate  float64 `json:"product_stock_after_update"`
}
