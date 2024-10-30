package model

import "time"

type Cart struct {
	OrganizationID uint         `json:"organization_id"`
	Organization   Organization `gorm:"foreignKey:OrganizationID;references:ID" json:"-" validate:"-"`
	ID             uint         `json:"id" gorm:"primarykey"`
	CustomerID     uint         `json:"customer_id" validate:"required"`
	Customer       Customer     `gorm:"foreignKey:CustomerID;references:ID" json:"-" validate:"-"`
	TotalQuantity  float64      `json:"total_quantity" validate:"-"`
	CartItems      []CartItem   `json:"cart_items" gorm:"foreignKey:CartID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	IsDeleted      bool         `json:"is_deleted"`
}

type CartResponse struct {
	OrganizationID uint       `json:"organization_id"`
	ID             uint       `json:"id"`
	CustomerID     uint       `json:"customer_id"`
	CartItems      []CartItem `json:"cart_items"`
	TotalQuantity  float64    `json:"total_quantity"`
}
