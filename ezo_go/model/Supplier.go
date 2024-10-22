package model

import (
	"time"
)

type Supplier struct {
	OrganizationID     uint         `json:"organization_id"`
	Organization       Organization `gorm:"foreignKey:OrganizationID;references:ID" json:"-" validate:"-"`
	ID                 uint         `json:"id" gorm:"primaryKey"`
	Name               string       `json:"name" validate:"required"`
	PhoneNo            string       `json:"phone_no" gorm:"unique" validate:"required"`
	Category           string       `json:"category" validate:"required"`
	BillingAddress     string       `json:"billing_address" validate:"required"`
	BillingProvince    string       `json:"billing_province" validate:"omitempty"`
	BillingPostalCode  string       `json:"billing_postal_code" validate:"omitempty"`
	DeliveryAddress    string       `json:"delivery_address" validate:"omitempty"`
	DeliveryProvince   string       `json:"delivery_province" validate:"omitempty"`
	DeliveryPostalCode string       `json:"delivery_postal_code" validate:"omitempty"`
	GstNumber          string       `json:"gst_number" validate:"omitempty"`
	BillingTerm        string       `json:"billing_term" validate:"omitempty"`
	BillingType        string       `json:"billing_type" validate:"omitempty"`
	DateOfBirth        string       `json:"date_of_birth" validate:"omitempty"`
	WhatsappAlert      string       `json:"whatsapp_alert" validate:"omitempty"`
	CreatedAt          time.Time    `json:"created_at"`
	UpdatedAt          time.Time    `json:"updated_at"`
}

type SupplierResponse struct{
	ID                 uint         `json:"id" gorm:"primaryKey"`
	Name               string       `json:"name"`
	PhoneNo            string       `json:"phone_no" gorm:"unique"`
	Category           string       `json:"category"`
	BillingAddress     string       `json:"billing_address"`
	BillingProvince    string       `json:"billing_province"`
	BillingPostalCode  string       `json:"billing_postal_code"`
	DeliveryAddress    string       `json:"delivery_address"`
	DeliveryProvince   string       `json:"delivery_province"`
	DeliveryPostalCode string       `json:"delivery_postal_code"`
	GstNumber          string       `json:"gst_number"`
	BillingTerm        string       `json:"billing_term"`
	BillingType        string       `json:"billing_type"`
	DateOfBirth        string       `json:"date_of_birth"`
	WhatsappAlert      string       `json:"whatsapp_alert"`
}