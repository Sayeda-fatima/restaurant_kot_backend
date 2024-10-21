package validator

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/go-playground/validator"
)

type UserValidator interface {
	UserValidate(user model.User) error
}

type userValidator struct {
	validator *validator.Validate
}

func NewUserValidator() UserValidator {
	return &userValidator{
		validator: validator.New(),
	}
}

func (ur *userValidator) UserValidate(user model.User) error {
	if err := ur.validator.Struct(&user); err != nil {
		return err
	}
	return nil
}
