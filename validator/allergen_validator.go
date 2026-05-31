package validator

import (
	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
	"github.com/go-playground/validator"
)

type (
	AllergenValidator interface {
		AllergenValidate(allergen *model.Allergen) error
	}

	allergenValidator struct{
		validator *validator.Validate
	}
)

func NewAllergenValidator() AllergenValidator{
	return &allergenValidator{
		validator: validator.New(),
	}
}

func (ar *allergenValidator) AllergenValidate(allergen *model.Allergen) error{

	if err := ar.validator.Struct(allergen); err != nil{
		return err
	}
	return nil
}