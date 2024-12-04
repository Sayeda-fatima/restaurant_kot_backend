package validator

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/go-playground/validator"
)

type (
	RecipeValidator interface {
		RecipeValidate(recipe *model.Recipe) error
	}

	recipeValidator struct{
		validator *validator.Validate
	}
)

func NewRecipeValidator() RecipeValidator{
	return &recipeValidator{
		validator: validator.New(),
	}
}

func (rc *recipeValidator) RecipeValidate(recipe *model.Recipe) error{

	if err := rc.validator.Struct(recipe); err != nil{
		return err
	}
	return nil
}