package validator

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/go-playground/validator"
)

type (
	CartItemValidator interface {
		CartItemValidate(cartItem *model.CartItem) error
	}

	cartItemValidator struct{
		validator *validator.Validate
	}
)

func NewCartItemValidator() CartItemValidator{
	return &cartItemValidator{
		validator: validator.New(),
	}
}

func (cr *cartItemValidator) CartItemValidate(cartItem *model.CartItem) error{

	if err := cr.validator.Struct(cartItem); err != nil{
		return err
	}
	return nil
}