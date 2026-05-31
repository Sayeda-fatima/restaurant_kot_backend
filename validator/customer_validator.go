package validator

import (
	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
	"github.com/go-playground/validator"
)

type (
	CustomerValidator interface {
		CustomerValidate(customer *model.Customer) error
	}

	customerValidator struct{
		validator *validator.Validate
	}
)

func NewCustomerValidator() CustomerValidator{
	return &customerValidator{
		validator: validator.New(),
	}
}

func (cr *customerValidator) CustomerValidate(customer *model.Customer) error{

	if err := cr.validator.Struct(customer); err != nil{
		return err
	}

	return nil
}