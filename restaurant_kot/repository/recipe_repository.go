package repository

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"gorm.io/gorm"
)

type (
	RecipeRepository interface {
		GetRecipeList(recipe *[]model.Recipe, organizationID uint, restaurantID uint) error
		CreateRecipe(recipe *model.Recipe) error
		UpdateRecipe(recipe *model.Recipe, id uint) error
		DeleteRecipe(recipe *model.Recipe, id uint) error
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

func (rr *recipeRepository) CreateRecipe(recipe *model.Recipe) error{

	if err := rr.db.Create(recipe).Error; err != nil{
		return err
	}
	return nil
}

func (rr *recipeRepository) UpdateRecipe(recipe *model.Recipe, id uint) error{

	result := rr.db.Model(recipe).Where("id=?",id).Updates(recipe)

	if err := result.Error; err != nil{
		return err
	}
	return nil
}

func (rr *recipeRepository) DeleteRecipe(recipe *model.Recipe, id uint) error{

	result := rr.db.Model(recipe).Where("id=?", id).Update("is_deleted", 1)

	if err := result.Error; err != nil{
		return err
	}
	return nil
}