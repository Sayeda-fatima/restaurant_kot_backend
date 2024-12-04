package validator

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/go-playground/validator"
)

type (
	MenuItemValidator interface {
		MenuItemValidate(menuItem *model.MenuItem) error
	}

	menuItemValidator struct{
		validator *validator.Validate
	}
)

func NewMenuItemValidator() MenuItemValidator{
	return &menuItemValidator{
		validator: validator.New(),
	}
}

func (mr *menuItemValidator) MenuItemValidate(menuItem *model.MenuItem) error{

	if err := mr.validator.Struct(menuItem); err != nil{
		return err
	}
	return nil
}