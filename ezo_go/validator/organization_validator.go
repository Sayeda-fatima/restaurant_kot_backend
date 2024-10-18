package validator

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/go-playground/validator"
)

type OrganizationValidator interface{
	OrganizationValidate (organization model.Organization) error
}

type (
	Organization struct{
	Name		string		`json:"name" validate:"required"`
	AccessGiven	int64		`json:"access_given" validate:"required"`
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

func (or *organizationValidator) OrganizationValidate (organization model.Organization) error {

	if err := or.validator.Struct(&organization); err!=nil{
		return err;
	}
	return nil
}
