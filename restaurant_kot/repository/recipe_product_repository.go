package repository

import (
	"fmt"

	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"gorm.io/gorm"
)

type (
	RecipeProductRepository interface {
		GetProductInRecipe(recipeProduct *[]model.RecipeProduct, recipeID uint, organizationID uint, restaurantID uint) error
		CreateRecipeProduct(recipeProduct *model.RecipeProduct) error
		UpdateRecipeProduct(recipeProduct *model.RecipeProduct, id uint, organizationID uint, restaurantID uint) error
		DeleteRecipeProduct(recipeProduct *model.RecipeProduct, id uint, organizationID uint, restaurantID uint) error
	}

	recipeProductRepository struct{
		db *gorm.DB
	}
)

func NewRecipeProductRepository(db *gorm.DB) RecipeProductRepository{
	return &recipeProductRepository{db}
}

func (rr *recipeProductRepository) GetProductInRecipe(recipeProduct *[]model.RecipeProduct, recipeID uint, organizationID uint, restaurantID uint) error{

	if err := rr.db.Preload("Product").Where("recipe_id=? and organization_id=? and restaurant_id=?", recipeID, organizationID, restaurantID).Find(recipeProduct).Error; err != nil{
		return err
	}
	return nil
}

func (rr *recipeProductRepository) CreateRecipeProduct(recipeProduct *model.RecipeProduct) error{

	if err := rr.db.Create(recipeProduct).Error; err != nil{
		return err
	}
	return nil
}

func (rr *recipeProductRepository) UpdateRecipeProduct(recipeProduct *model.RecipeProduct, id uint, organizationID uint, restaurantID uint) error{

	result := rr.db.Model(recipeProduct).Where("id=? and organization_id=? and restaurant_id=?", id, organizationID, restaurantID).Updates(recipeProduct)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1{
		return fmt.Errorf("record does not exist")
	}
	return nil
}

func (rr *recipeProductRepository) DeleteRecipeProduct(recipeProduct *model.RecipeProduct, id uint, organizationID uint, restaurantID uint) error{

	result := rr.db.Model(recipeProduct).Where("id=? and organization_id=? and restaurant_id=?", id, organizationID, restaurantID).Update("is_deleted", 1)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1{
		return fmt.Errorf("record does not exist")
	}
	return nil
}