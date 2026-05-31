package validator

import (
	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
	"github.com/go-playground/validator"
)

type (
	StaffValidator interface {
		StaffValidate(staff *model.Staff) error
	}

	staffValidator struct{
		validator *validator.Validate
	}
)

func NewStaffValidator () StaffValidator {
	return &staffValidator{
		validator: validator.New(),
	}
}

func (sr *staffValidator) StaffValidate(staff *model.Staff) error{

	if err := sr.validator.Struct(staff); err != nil{
		return err
	}
	return nil
}