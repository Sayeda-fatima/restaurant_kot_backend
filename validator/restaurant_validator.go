package validator

import (
	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
	"github.com/go-playground/validator"
)

type (
	RestaurantValidator interface {
		RestaurantValidate(restaurant model.Restaurant) error
	}
	restaurantValidator struct{
		validator *validator.Validate
	}
)

func NewRestaurantValidator() RestaurantValidator{
	return &restaurantValidator{
		validator: validator.New(),
	}
}

func (rv *restaurantValidator) RestaurantValidate(restaurant model.Restaurant) error{

	if err := rv.validator.Struct(restaurant); err != nil{
		return err
	}
	return nil
}