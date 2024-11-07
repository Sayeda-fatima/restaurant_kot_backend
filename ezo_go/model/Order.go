package model

import "time"

type Order struct {
	OrganizationID         uint         `json:"organization_id"`
	Organization           Organization `gorm:"foreignKey:OrganizationID;references:ID" json:"-" validate:"-"`
	ID                     uint         `json:"id" gorm:"primaryKey"`
	CustomerID             uint         `json:"customer_id" validate:"required"`
	Customer               Customer     `gorm:"foreignKey:CustomerID;references:ID" json:"-" validate:"-"`
	TotalPrice             float64      `json:"total_price" validate:"required"`
	CustomerBillingAddress string       `json:"customer_billing_address" validate:"required"`
	ModeOfPayment          string       `json:"mode_of_payment" validate:"required"`
	OrderItems             []OrderItem  `json:"order_items" gorm:"foreignKey:OrderID"`
	CreatedBy				uint		`json:"created_by"`
	CreatedAt              time.Time    `json:"created_at"`
	UpdatedAt              time.Time    `json:"updated_at"`
	IsDeleted              bool         `json:"is_deleted"`
}

type OrderResponse struct {
	OrganizationID         uint        `json:"organization_id"`
	ID                     uint        `json:"id"`
	CustomerID             uint        `json:"customer_id"`
	TotalPrice             float64     `json:"total_price"`
	OrderItems             []OrderItem `json:"order_item"`
	CustomerBillingAddress string      `json:"customer_billing_address"`
	ModeOfPayment          string      `json:"mode_of_payment"`
}
