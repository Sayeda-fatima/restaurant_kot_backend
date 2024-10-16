package model

import (
	"time"
)

type Customer struct{
	ID 							uint		`json:"id" gorm:"primaryKey"`
	CustomerName 				string		`json:"customer_name"`
	CustomerPhoneNo				string		`json:"customer_phone_no" gorm:"unique"`
	CustomerCategory			string		`json:"customer_category"`
	CustomerBillingAddress		string		`json:"customer_billing_address"`
	CustomerBillingProvince		string		`json:"customer_billing_province"`
	CustomerBillingPostalCode	string		`json:"customer_billing_postal_code"`
	CustomerDeliveryAddress		string		`json:"customer_delivery_address"`
	CustomerDeliveryProvince	string		`json:"customer_delivery_province"`
	CustomerDeliveryPostalCode	string		`json:"customer_delivery_postal_code"`
	CustomerGstNumber			string		`json:"customer_gst_number"`
	CustomerBillingTerm			string		`json:"customer_billing_term"`
	CustomerBillingType			string		`json:"customer_billing_type"`
	CustomerDateOfBirth			time.Time	`json:"customer_date_of_birth"`
	WhatsappAlert				string		`json:"whatsapp_alert"`		
	CreatedAt 					time.Time	`json:"created_at"`
	UpdatedAt 					time.Time	`json:"updated_at"`
}

type CustomerResponse struct{
	ID 							uint		`json:"id" gorm:"primaryKey"`
	CustomerName 				string		`json:"customer_name"`
	CustomerPhoneNo				string		`json:"customer_phone_no"`
	CustomerCategory			string		`json:"customer_category"`
	CustomerBillingAddress		string		`json:"customer_billing_address"`
	CustomerBillingProvince		string		`json:"customer_billing_province"`
	CustomerBillingPostalCode	string		`json:"customer_billing_postal_code"`
	CustomerDeliveryAddress		string		`json:"customer_delivery_address"`
	CustomerDeliveryProvince	string		`json:"customer_delivery_province"`
	CustomerDeliveryPostalCode	string		`json:"customer_delivery_postal_code"`
	CustomerGstNumber			string		`json:"customer_gst_number"`
	CustomerBillingTerm			string		`json:"customer_billing_term"`
	CustomerBillingType			string		`json:"customer_billing_type"`
	CustomerDateOfBirth			time.Time	`json:"customer_date_of_birth"`
	WhatsappAlert				string		`json:"whatsapp_alert"`		
}

type CustomerUpdate struct{
	CustomerName 				string		`json:"customer_name"`
	CustomerPhoneNo				string		`json:"customer_phone_no"`
	CustomerCategory			string		`json:"customer_category"`
	CustomerBillingAddress		string		`json:"customer_billing_address"`
	CustomerBillingProvince		string		`json:"customer_billing_province"`
	CustomerBillingPostalCode	string		`json:"customer_billing_postal_code"`
	CustomerDeliveryAddress		string		`json:"customer_delivery_address"`
	CustomerDeliveryProvince	string		`json:"customer_delivery_province"`
	CustomerDeliveryPostalCode	string		`json:"customer_delivery_postal_code"`
	CustomerGstNumber			string		`json:"customer_gst_number"`
	CustomerBillingTerm			string		`json:"customer_billing_term"`
	CustomerBillingType			string		`json:"customer_billing_type"`
	CustomerDateOfBirth			time.Time	`json:"customer_date_of_birth"`
	WhatsappAlert				string		`json:"whatsapp_alert"`	
}