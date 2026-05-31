package validator

import (
	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
	"github.com/go-playground/validator"
)

type (
	MenuValidator interface {
		MenuValidate(menu *model.Menu) error
	}

	menuValidator struct{
		validator *validator.Validate
	}
)

func NewMenuValidator() MenuValidator{
	return &menuValidator{
		validator: validator.New(),
	}
}

func (mr *menuValidator) MenuValidate(menu *model.Menu) error{

	if err := mr.validator.Struct(menu); err != nil{
		return err
	}
	return nil
}