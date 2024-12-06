package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/repository"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/validator"
)

type (
	RecipeUsecase interface {
		GetRecipeList(organizationID uint, restaurantID uint) ([]model.RecipeResponse, error)
		CreateRecipe(recipe model.Recipe) (model.RecipeResponse, error)
		UpdateRecipe(recipe model.Recipe, id uint, organizationID uint, restaurantID uint) (model.RecipeResponse, error)
		DeleteRecipe(recipe model.Recipe, id uint, organizationID uint, restaurantID uint) error
	}

	recipeUsecase struct{
		rr repository.RecipeRepository
		rv validator.RecipeValidator
	}
)

func NewRecipeUsecase(rr repository.RecipeRepository, rv validator.RecipeValidator) RecipeUsecase{
	return &recipeUsecase{rr,rv}
}

func (ru *recipeUsecase) GetRecipeList(organizationID uint, restaurantID uint) ([]model.RecipeResponse, error){

	recipe := []model.Recipe{}

	if err := ru.rr.GetRecipeList(&recipe, organizationID, restaurantID); err != nil{
		return nil, err
	}

	resRecipe := []model.RecipeResponse{}
	for _, v := range(recipe){
		res := model.RecipeResponse{
			ID: v.ID,
			Name: v.Name,
			Instruction: v.Instruction,
			CookingTime: v.CookingTime,
			Serving: v.Serving,
		}
		resRecipe = append(resRecipe, res)
	}
	return resRecipe, nil
}

func (ru *recipeUsecase) CreateRecipe(recipe model.Recipe) (model.RecipeResponse, error){

	if err := ru.rv.RecipeValidate(&recipe); err != nil{
		return model.RecipeResponse{}, err
	}

	if err := ru.rr.CreateRecipe(&recipe); err != nil{
		return model.RecipeResponse{}, err
	}

	resRecipe := model.RecipeResponse{
		ID: recipe.ID,
		Name: recipe.Name,
		CookingTime: recipe.CookingTime,
		Instruction: recipe.Instruction,
		Serving: recipe.Serving,
	}
	return resRecipe, nil
}

func (ru *recipeUsecase) UpdateRecipe(recipe model.Recipe, id uint, organizationID uint, restaurantID uint) (model.RecipeResponse, error){

	if err := ru.rv.RecipeValidate(&recipe); err != nil{
		return model.RecipeResponse{}, err
	}

	if err := ru.rr.UpdateRecipe(&recipe, id, organizationID, restaurantID); err != nil{
		return model.RecipeResponse{}, err
	}

	resRecipe := model.RecipeResponse{
		ID: recipe.ID,
		Name: recipe.Name,
		CookingTime: recipe.CookingTime,
		Instruction: recipe.Instruction,
		Serving: recipe.Serving,
	}
	return resRecipe, nil
}

func (ru *recipeUsecase) DeleteRecipe(recipe model.Recipe, id uint, organizationID uint, restaurantID uint) error{

	if err := ru.rr.DeleteRecipe(&recipe, id, organizationID, restaurantID); err != nil{
		return err
	}
	return nil
}