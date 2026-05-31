package validator

import (
	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
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