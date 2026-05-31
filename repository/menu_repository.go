package repository

import (
	"fmt"

	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
	"gorm.io/gorm"
)

type (
	MenuRepository interface {
		GetMenuList(menu *[]model.Menu, organizationID uint, restaurantID uint) error
		GetMenu(menu *model.Menu, id uint, organizationID uint, restaurantID uint) error
		CreateMenu(menu *model.Menu) error
		UpdateMenu(menu *model.Menu, id uint, organizationID uint, restaurantID uint) error
		DeleteMenu(menu *model.Menu, id uint, organizationID uint, restaurantID uint) error
	}

	menuRepository struct{
		db *gorm.DB
	}
)

func NewMenuRepository(db *gorm.DB) MenuRepository{
	return &menuRepository{db}
}

func (mr *menuRepository) GetMenuList(menu *[]model.Menu, organizationID uint, restaurantID uint) error{

	if err := mr.db.Preload("MenuItems", func(db *gorm.DB) *gorm.DB{return db.Select("id", "menu_id", "item_name", "description", "currency", "price") }).Where("organization_id=? and restaurant_id=? and is_deleted=0", organizationID, restaurantID).Find(menu).Error; err != nil{
		return err
	}
	return nil
}

func (mr *menuRepository) GetMenu(menu *model.Menu, id uint, organizationID uint, restaurantID uint) error{

	if err := mr.db.Preload("MenuItems.Recipe.RecipeProducts.Product").Where("id=? and restaurant_id=? and organization_id=? and is_deleted=0", id, restaurantID, organizationID).First(menu).Error; err != nil{
		return err
	}
	return nil
}

func (mr *menuRepository) CreateMenu(menu *model.Menu) error{

	if err := mr.db.Create(menu).Error; err != nil{
		return err
	}
	return nil
}

func (mr *menuRepository) UpdateMenu(menu *model.Menu, id uint, organizationID uint, restaurantID uint) error{

	result := mr.db.Model(menu).Where("id=? and organization_id=? and restaurant_id=?", id, organizationID, restaurantID).Updates(menu)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	
	return nil
}

func (mr *menuRepository) DeleteMenu(menu *model.Menu, id uint, organizationID uint, restaurantID uint) error{

	result := mr.db.Model(menu).Where("id=? and organization_id=? and restaurant_id=?", id, organizationID, restaurantID).Update("is_deleted",1)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	
	return nil
}