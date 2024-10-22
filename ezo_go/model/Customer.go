package model

import (
	"time"
)

type Customer struct {
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
	IsDeleted          bool         `json:"is_deleted"`
}

type CustomerResponse struct {
	ID                 uint   `json:"id" gorm:"primaryKey"`
	Name               string `json:"name" gorm:"column:name"`
	PhoneNo            string `json:"phone_no" gorm:"column:phone_no"`
	Category           string `json:"category" gorm:"column:category"`
	BillingAddress     string `json:"billing_address" gorm:"column:billing_address"`
	BillingProvince    string `json:"billing_province" gorm:"column:billing_province"`
	BillingPostalCode  string `json:"billing_postal_code" gorm:"column:billing_postal_code"`
	DeliveryAddress    string `json:"delivery_address" gorm:"column:delivery_address"`
	DeliveryProvince   string `json:"delivery_province" gorm:"column:delivery_province"`
	DeliveryPostalCode string `json:"delivery_postal_code" gorm:"column:delivery_postal_code"`
	GstNumber          string `json:"gst_number" gorm:"column:gst_number"`
	BillingTerm        string `json:"billing_term" gorm:"column:billing_term"`
	BillingType        string `json:"billing_type" gorm:"column:billing_type"`
	DateOfBirth        string `json:"date_of_birth" gorm:"column:date_of_birth"`
	WhatsappAlert      string `json:"whatsapp_alert" gorm:"column:whatsapp_alert"`
}

type CustomerUpdate struct {
	OrganizationID    uint   `json:"organization_id" gorm:"foreignKey"`
	ID                uint   `json:"id" gorm:"primaryKey"`
	Name              string `json:"name"`
	PhoneNo           string `json:"phone_no"`
	Category          string `json:"category"`
	BillingAddress    string `json:"billing_address"`
	BillingProvince   string `json:"billing_province"`
	BillingPostalCode string `json:"billing_postal_code"`
	// DeliveryAddress		string		`json:"delivery_address"`
	// DeliveryProvince	string		`json:"delivery_province"`
	// DeliveryPostalCode	string		`json:"delivery_postal_code"`
	// GstNumber			string		`json:"gst_number"`
	// BillingTerm			string		`json:"billing_term"`
	// BillingType			string		`json:"billing_type"`
	// DateOfBirth			time.Time	`json:"date_of_birth"`
	// WhatsappAlert		string		`json:"whatsapp_alert"`
}
