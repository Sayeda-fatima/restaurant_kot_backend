package model

import "time"

type CartItem struct {
	OrganizationID  uint         `json:"organization_id"`
	//Organization    Organization `gorm:"foreignKey:OrganizationID;references:ID" json:"-" validate:"-"`
	ID              uint         `json:"id" gorm:"primarykey"`
	CartID          uint         `json:"cart_id" validate:"required"`
	//Cart            Cart         `gorm:"foreignKey:CartID;references:ID" json:"-" validate:"-"`
	ProductID       uint         `json:"product_id" validate:"required"`
	Product         Product      `gorm:"foreignKey:ProductID;references:ID" validate:"-"`
	ProductQuantity float64      `json:"product_quantity" validate:"required"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
}

type CartItemResponse struct {
	OrganizationID  uint    `json:"organization_id"`
	ID              uint    `json:"id"`
	CartID          uint    `json:"cart_id"`
	ProductID       uint    `json:"product_id"`
	Product			Product	`json:"product"`
	ProductQuantity float64 `json:"product_quantity"`
}
