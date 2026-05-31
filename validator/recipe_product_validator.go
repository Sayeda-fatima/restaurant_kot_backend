package validator

import (
	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
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