package repository

import (
	"fmt"

	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"gorm.io/gorm"
)

type (
	MenuItemRepository interface {
		GetMenuItemList(menuItem *[]model.MenuItem, organizationID uint, restaurantID uint, menuID uint) error
		CreateMenuItem(menuItem *model.MenuItem) error
		UpdateMenuItem(menuItem *model.MenuItem, id uint, organizationID uint, restaurantID uint) error
		DeleteMenuItem(menuItem *model.MenuItem, id uint, organizationID uint, restaurantID uint) error
		UpdateMenuItemIsActivated(menuItem *model.MenuItem, id uint, organizationID uint, restaurantID uint, status bool) error
	}

	menuItemRepository struct{
		db *gorm.DB
	}
)

func NewMenuItemRepository(db *gorm.DB) MenuItemRepository{
	return &menuItemRepository{db}
}

func (mr *menuItemRepository) GetMenuItemList(menuItem *[]model.MenuItem, organizationID uint, restaurantID uint, menuID uint) error{

	if err := mr.db.Preload("MenuAllergens.Allergen", func (db *gorm.DB) *gorm.DB{return db.Select("id", "allergen_name")}).Where("organization_id=? and restaurant_id=? and menu_id=? and is_deleted=0", organizationID, restaurantID, menuID).Find(menuItem).Error; err != nil{
		return err
	}
	return nil
}

func (mr *menuItemRepository) CreateMenuItem(menuItem *model.MenuItem) error{

	if err := mr.db.Create(menuItem).Error; err != nil{
		return err
	}
	return nil
}

func (mr *menuItemRepository) UpdateMenuItem(menuItem *model.MenuItem, id uint, organizationID uint, restaurantID uint) error{

	result := mr.db.Model(menuItem).Where("id=? and organization_id=? and restaurant_id=?", id, organizationID, restaurantID).Updates(menuItem)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	
	return nil
}

func (mr *menuItemRepository) DeleteMenuItem(menuItem *model.MenuItem, id uint, organizationID uint, restaurantID uint) error{

	result := mr.db.Model(menuItem).Where("id=? and organization_id=? and restaurant_id=?", id, organizationID, restaurantID).Update("is_deleted", 1)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	
	return nil
}

func (mr *menuItemRepository) UpdateMenuItemIsActivated(menuItem *model.MenuItem, id uint, organizationID uint, restaurantID uint, status bool) error{

	result := mr.db.Model(menuItem).Where("id=? and restaurant_id=? and organization_id=?", id, restaurantID, organizationID).Update("is_available", status)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1{
		return fmt.Errorf("object does not exist")
	}

	return nil
}