package validator

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/go-playground/validator"
)

type (
	OrganizationValidator interface {
		OrganizationValidate(organization model.Organization) error
	}
	organizationValidator struct{
		validator *validator.Validate
	}
)

func NewOrganizationValidator() OrganizationValidator{
	return &organizationValidator{
		validator: validator.New(),
	}
}

func (or *organizationValidator) OrganizationValidate(organization model.Organization) error{

	if err := or.validator.Struct(&organization); err != nil{
		return err
	}
	return nil
}