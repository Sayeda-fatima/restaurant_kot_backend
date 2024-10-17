package validator

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/go-playground/validator"
)

type CustomerValidator interface{
	CustomerValidate(customer model.Customer) error
}

type (
	Customer struct{
		CustomerName 				string		`json:"customer_name" validate:"required"`
		CustomerPhoneNo				string		`json:"customer_phone_no" validate:"required"`
		CustomerCategory			string		`json:"customer_category" validate:"required"`
		CustomerBillingAddress		string		`json:"customer_billing_address" validate:"required"`
		CustomerBillingProvince		string		`json:"customer_billing_province" validate:"omitempty"`
		CustomerBillingPostalCode	string		`json:"customer_billing_postal_code" validate:"omitempty"`
		CustomerDeliveryAddress		string		`json:"customer_delivery_address" validator:"omitempty"`
		CustomerDeliveryProvince	string		`json:"customer_delivery_province" validator:"omitempty"`
		CustomerDeliveryPostalCode	string		`json:"customer_delivery_postal_code" validator:"omitempty"`
		CustomerGstNumber			string		`json:"customer_gst_number" validator:"omitempty"`
		CustomerBillingTerm			string		`json:"customer_billing_term" validator:"omitempty"`
		CustomerBillingType			string		`json:"customer_billing_type" validator:"omitempty"`
		CustomerDateOfBirth			string	`json:"customer_date_of_birth" validator:"omitempty"`
		WhatsappAlert				string		`json:"whatsapp_alert" validator:"omitempty"`	
	}

	customerValidator struct{
		validator *validator.Validate
	}
)

func NewCustomerValidator () CustomerValidator{
	return &customerValidator{
		validator: validator.New(),
	}
}

func (cr *customerValidator) CustomerValidate (customer model.Customer) error {
	if err := cr.validator.Struct(&customer); err!=nil{
		return err
	}
	return nil
}

