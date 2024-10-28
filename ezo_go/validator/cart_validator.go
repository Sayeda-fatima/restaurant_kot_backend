package validator

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/go-playground/validator"
)

type (
	CartValidator interface {
		CartValidate(cart model.Cart) error
	}

	cartValidator struct {
		validator *validator.Validate
	}
)

func NewCartValidator() CartValidator {
	return &cartValidator{
		validator: validator.New(),
	}
}

func (cr *cartValidator) CartValidate(cart model.Cart) error {

	if err := cr.validator.Struct(&cart); err != nil {
		return err
	}
	return nil
}
