package validator

import (
	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
	"github.com/go-playground/validator"
)

type (
	CartValidator interface {
		CartValidate(cart *model.Cart) error
	}

	cartValidator struct{
		validator *validator.Validate
	}
)

func NewCartValidator() CartValidator{
	return &cartValidator{
		validator: validator.New(),
	}
}

func (cr *cartValidator) CartValidate(cart *model.Cart) error{

	if err := cr.validator.Struct(cart); err != nil{
		return err
	}
	return nil
}