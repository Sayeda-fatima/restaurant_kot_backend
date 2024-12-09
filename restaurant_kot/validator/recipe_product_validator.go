package validator

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/go-playground/validator"
)

type (
	RecipeProductValidator interface {
		RecipeProductValidate(recipeProduct *model.RecipeProduct) error
	}
	recipeProductValidator struct{
		validator *validator.Validate
	}
)

func NewRecipeProductValidator() RecipeProductValidator{
	return &recipeProductValidator{
		validator: validator.New(),
	}
}

func (rp *recipeProductValidator) RecipeProductValidate(recipeProduct *model.RecipeProduct) error{

	if err := rp.validator.Struct(recipeProduct); err != nil{
		return err
	}
	return nil
}