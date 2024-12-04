package repository

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"gorm.io/gorm"
)

type (
	MenuRepository interface {
		GetMenuList(menu *[]model.Menu, organizationID uint, restaurantID uint) error
		CreateMenu(menu *model.Menu) error
		UpdateMenu(menu *model.Menu, id uint) error
		DeleteMenu(menu *model.Menu, id uint) error
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

func (mr *menuRepository) CreateMenu(menu *model.Menu) error{

	if err := mr.db.Create(menu).Error; err != nil{
		return err
	}
	return nil
}

func (mr *menuRepository) UpdateMenu(menu *model.Menu, id uint) error{

	result := mr.db.Model(menu).Where("id=?", id).Updates(menu)

	if err := result.Error; err != nil{
		return err
	}
	return nil
}

func (mr *menuRepository) DeleteMenu(menu *model.Menu, id uint) error{

	result := mr.db.Model(menu).Where("id=?", id).Update("is_deleted",1)

	if err := result.Error; err != nil{
		return err
	}
	return nil
}