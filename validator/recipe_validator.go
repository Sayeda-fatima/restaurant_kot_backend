package validator

import (
	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
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