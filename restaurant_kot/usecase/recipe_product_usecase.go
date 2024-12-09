package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/repository"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/validator"
)

type (
	RecipeProductUsecase interface {
		GetProductInRecipe(recipeID uint, organizationID uint, restaurantID uint) ([]model.RecipeProductResponse, error)
		CreateRecipeProduct(recipeProduct model.RecipeProduct) (model.RecipeProductResponse, error)
		UpdateRecipeProduct(recipeProduct model.RecipeProduct, id uint, organizationID uint, restaurantID uint) (model.RecipeProductResponse, error)
		DeleteRecipeProduct(recipeProduct model.RecipeProduct, id uint, organizationID uint, restaurantID uint) error
	}

	recipeProductUsecase struct{
		rr repository.RecipeProductRepository
		rv validator.RecipeProductValidator
	}
)

func NewRecipeProductUsecase (rr repository.RecipeProductRepository, rv validator.RecipeProductValidator) RecipeProductUsecase{
	return &recipeProductUsecase{rr, rv}
}

func (ru *recipeProductUsecase) GetProductInRecipe(recipeID uint, organizationID uint, restaurantID uint) ([]model.RecipeProductResponse, error){

	recipeProduct := []model.RecipeProduct{}

	if err := ru.rr.GetProductInRecipe(&recipeProduct, recipeID, organizationID, restaurantID); err != nil{
		return nil, err
	}

	resRecipeProduct := []model.RecipeProductResponse{}
	for _, v := range(recipeProduct){
		res := model.RecipeProductResponse{
			ID: v.ID,
			ProductID: v.ProductID,
			RecipeID: v.RecipeID,
			Quantity: v.Quantity,
		}
		resRecipeProduct = append(resRecipeProduct, res)
	}
	return resRecipeProduct, nil
}

func (ru *recipeProductUsecase) CreateRecipeProduct(recipeProduct model.RecipeProduct) (model.RecipeProductResponse, error){

	if err := ru.rv.RecipeProductValidate(&recipeProduct); err != nil{
		return model.RecipeProductResponse{}, err
	}

	if err := ru.rr.CreateRecipeProduct(&recipeProduct); err != nil{
		return model.RecipeProductResponse{}, err
	}

	resRecipeProduct := model.RecipeProductResponse{
		ID: recipeProduct.ID,
		RecipeID: recipeProduct.RecipeID,
		ProductID: recipeProduct.ProductID,
		Quantity: recipeProduct.Quantity,
	}
	return resRecipeProduct, nil
}

func (ru *recipeProductUsecase) UpdateRecipeProduct(recipeProduct model.RecipeProduct, id uint, organizationID uint, restaurantID uint) (model.RecipeProductResponse, error){

	if err := ru.rv.RecipeProductValidate(&recipeProduct); err != nil{
		return model.RecipeProductResponse{}, err
	}

	if err := ru.rr.UpdateRecipeProduct(&recipeProduct, id, organizationID, restaurantID); err != nil{
		return model.RecipeProductResponse{}, err
	}

	resRecipeProduct := model.RecipeProductResponse{
		ID: recipeProduct.ID,
		RecipeID: recipeProduct.RecipeID,
		ProductID: recipeProduct.ProductID,
		Quantity: recipeProduct.Quantity,
	}
	return resRecipeProduct, nil
}

func (ru *recipeProductUsecase) DeleteRecipeProduct(recipeProduct model.RecipeProduct, id uint, organizationID uint, restaurantID uint) error{

	if err := ru.rr.DeleteRecipeProduct(&recipeProduct, id, organizationID, restaurantID); err != nil{
		return err
	}
	return nil
}