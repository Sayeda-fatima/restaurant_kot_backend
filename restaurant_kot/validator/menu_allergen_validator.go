package validator

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
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