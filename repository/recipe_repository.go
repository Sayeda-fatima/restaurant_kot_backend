package repository

import (
	"fmt"

	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
	"gorm.io/gorm"
)

type (
	RecipeRepository interface {
		GetRecipeList(recipe *[]model.Recipe, organizationID uint, restaurantID uint) error
		GetRecipe(recipe *model.Recipe, id uint, organizationID uint, restaurantID uint) error
		CreateRecipe(recipe *model.Recipe) error
		UpdateRecipe(recipe *model.Recipe, id uint, organizationID uint, restaurantID uint) error
		DeleteRecipe(recipe *model.Recipe, id uint, organizationID uint, restaurantID uint) error
	}

	recipeRepository struct{
		db *gorm.DB
	}
)

func NewRecipeRepository(db *gorm.DB) RecipeRepository{
	return &recipeRepository{db}
}

func (rr *recipeRepository) GetRecipeList(recipe *[]model.Recipe, organizationID uint, restaurantID uint) error{

	if err := rr.db.Where("organization_id=? and restaurant_id=? and is_deleted=0", organizationID, restaurantID).Find(recipe).Error; err != nil{
		return err
	}
	return nil
}

func (rr *recipeRepository) GetRecipe(recipe *model.Recipe, id uint, organizationID uint, restaurantID uint) error{

	if err := rr.db.Preload("RecipeProducts.Product").Where("id=? and organization_id=? and restaurant_id=? and is_deleted=0", id, organizationID, restaurantID).First(recipe).Error; err != nil{
		return err
	}
	return nil
}

func (rr *recipeRepository) CreateRecipe(recipe *model.Recipe) error{

	if err := rr.db.Create(recipe).Error; err != nil{
		return err
	}
	return nil
}

func (rr *recipeRepository) UpdateRecipe(recipe *model.Recipe, id uint, organizationID uint, restaurantID uint) error{

	result := rr.db.Model(recipe).Where("id=? and organization_id=? and restaurant_id=?",id, organizationID, restaurantID).Updates(recipe)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	
	return nil
}

func (rr *recipeRepository) DeleteRecipe(recipe *model.Recipe, id uint, organizationID uint, restaurantID uint) error{

	result := rr.db.Model(recipe).Where("id=? and organization_id=? and restaurant_id=?", id, organizationID, restaurantID).Update("is_deleted", 1)

	if err := result.Error; err != nil{
		return err
	}
	
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}

	return nil
}