package validator

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/go-playground/validator"
)

type (
	OrderValidator interface {
		OrderValidate(order *model.Order) error
	}

	orderValidator struct{
		validator *validator.Validate
	}
)

func NewOrderValidator() OrderValidator{
	return &orderValidator{
		validator: validator.New(),
	}
}

func (or *orderValidator) OrderValidate(order *model.Order) error{

	if err := or.validator.Struct(order); err != nil{
		return err
	}
	return nil
}