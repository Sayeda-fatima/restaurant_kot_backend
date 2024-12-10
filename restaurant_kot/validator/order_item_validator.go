package validator

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/go-playground/validator"
)

type (
	OrderItemValidator interface {
		OrderItemValidate(orderItem *model.OrderItem) error
	}

	orderItemValidator struct{
		validator *validator.Validate
	}
)

func NewOrderItemValidator() OrderItemValidator{
	return &orderItemValidator{
		validator: validator.New(),
	}
}

func (or *orderItemValidator) OrderItemValidate(orderItem *model.OrderItem) error{

	if err := or.validator.Struct(orderItem); err != nil{
		return err
	}
	return nil
}