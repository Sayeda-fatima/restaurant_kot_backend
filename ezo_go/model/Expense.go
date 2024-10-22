package model

import "time"

type Expense struct {
	OrganizationID  uint         `json:"organization_id"`
	Organization    Organization `gorm:"foreignKey:OrganizationID;references:ID" json:"-" validate:"-"`
	ID              uint         `json:"id" gorm:"primaryKey"`
	SupplierID      uint         `json:"supplier_id"`
	Supplier        Supplier     `gorm:"foreignKey:SupplierID;references:ID" json:"-" validate:"-"`
	SupplierName    string       `json:"supplier_name" validate:"required"`
	ExpenseCategory string       `json:"expense_category" validate:"required"`
	TotalAmount     float64      `json:"total_amount" validate:"required"`
	AmountPaid      float64      `json:"amount_paid" validate:"required"`
	AmountDue       float64      `json:"amount_due" validate:"omitempty"`
	Note            string       `json:"note" validate:"omitempty"`
	ModeOfPayment   string       `json:"mode_of_payment" validate:"required"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
	IsDeleted       bool         `json:"is_deleted"`
}

type ExpenseResponse struct {
	OrganizationID  uint    `json:"organization_id"`
	ID              uint    `json:"id"`
	SupplierID      uint    `json:"supplier_id"`
	SupplierName    string  `json:"supplier_name"`
	ExpenseCategory string  `json:"expense_category"`
	TotalAmount     float64 `json:"total_amount"`
	AmountPaid      float64 `json:"amount_paid"`
	AmountDue       float64 `json:"amount_due"`
	Note            string  `json:"note"`
	ModeOfPayment   string  `json:"mode_of_payment"`
}
