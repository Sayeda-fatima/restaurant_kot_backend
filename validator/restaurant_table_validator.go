package validator

import (
	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
	"github.com/go-playground/validator"
)

type (
	RestaurantTableValidator interface {
		RestaurantTableValidate(restaurantTable *model.RestaurantTable) error
	}

	restaurantTableValidator struct{
		validator *validator.Validate
	}
)

func NewRestaurantTableValidator() RestaurantTableValidator{
	return &restaurantTableValidator{
		validator: validator.New(),
	}
}

func (rv *restaurantTableValidator) RestaurantTableValidate(restaurantTable *model.RestaurantTable) error{

	if err := rv.validator.Struct(restaurantTable); err != nil{
		return err
	}
	return nil
}