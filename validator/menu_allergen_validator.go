package validator

import (
	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
	"github.com/go-playground/validator"
)

type (
	MenuAllergenValidator interface {
		MenuAllergenValidate(menuAllergen *model.MenuAllergen) error
	}

	menuAllergenValidator struct{
		validator *validator.Validate
	}
)

func NewMenuAllergenValidator() MenuAllergenValidator{
	return &menuAllergenValidator{
		validator: validator.New(),
	}
}

func (mr *menuAllergenValidator) MenuAllergenValidate(menuAllergen *model.MenuAllergen) error{

	if err := mr.validator.Struct(menuAllergen); err != nil{
		return err
	}
	return nil
}