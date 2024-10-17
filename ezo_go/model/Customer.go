package model

import (
	"time"
)

type Customer struct{
	OrganizationID		uint		`json:"organization_id" gorm:"foreignKey"`
	ID 					uint		`json:"id" gorm:"primaryKey"`
	Name 				string		`json:"name"`
	PhoneNo				string		`json:"phone_no" gorm:"unique"`
	Category			string		`json:"category"`
	BillingAddress		string		`json:"billing_address"`
	BillingProvince		string		`json:"billing_province"`
	BillingPostalCode	string		`json:"billing_postal_code"`
	DeliveryAddress		string		`json:"delivery_address"`
	DeliveryProvince	string		`json:"delivery_province"`
	DeliveryPostalCode	string		`json:"delivery_postal_code"`
	GstNumber			string		`json:"gst_number"`
	BillingTerm			string		`json:"billing_term"`
	BillingType			string		`json:"billing_type"`
	DateOfBirth			string	`json:"date_of_birth"`
	WhatsappAlert		string		`json:"whatsapp_alert"`		
	CreatedAt 			time.Time	`json:"created_at"`
	UpdatedAt 			time.Time	`json:"updated_at"`
}

type CustomerResponse struct{
	ID 					uint		`json:"id" gorm:"primaryKey"`
	Name 				string		`json:"name"`
	PhoneNo				string		`json:"phone_no"`
	Category			string		`json:"category"`
	BillingAddress		string		`json:"billing_address"`
	BillingProvince		string		`json:"billing_province"`
	BillingPostalCode	string		`json:"billing_postal_code"`
	DeliveryAddress		string		`json:"delivery_address"`
	DeliveryProvince	string		`json:"delivery_province"`
	DeliveryPostalCode	string		`json:"delivery_postal_code"`
	GstNumber			string		`json:"gst_number"`
	BillingTerm			string		`json:"billing_term"`
	BillingType			string		`json:"billing_type"`
	DateOfBirth			string		`json:"date_of_birth"`
	WhatsappAlert		string		`json:"whatsapp_alert"`		
}

type CustomerUpdate struct{
	Name 				string		`json:"name"`
	PhoneNo				string		`json:"phone_no"`
	Category			string		`json:"category"`
	BillingAddress		string		`json:"billing_address"`
	BillingProvince		string		`json:"billing_province"`
	BillingPostalCode	string		`json:"billing_postal_code"`
	DeliveryAddress		string		`json:"delivery_address"`
	DeliveryProvince	string		`json:"delivery_province"`
	DeliveryPostalCode	string		`json:"delivery_postal_code"`
	GstNumber			string		`json:"gst_number"`
	BillingTerm			string		`json:"billing_term"`
	BillingType			string		`json:"billing_type"`
	DateOfBirth			time.Time	`json:"date_of_birth"`
	WhatsappAlert		string		`json:"whatsapp_alert"`	
}