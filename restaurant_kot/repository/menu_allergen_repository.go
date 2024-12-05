package repository

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"gorm.io/gorm"
)

type (
	MenuAllergenRepository interface {
		GetMenuAllergenList(menuAllergen *[]model.MenuAllergen, menuID uint) error
		CreateMenuAllergen(menuAllergen *model.MenuAllergen) error
		UpdateMenuAllergen(menuAllergen *model.MenuAllergen, id uint) error
		DeleteMenuAllergen(menuAllergen *model.MenuAllergen, id uint) error
	}

	menuAllergenRepository struct{
		db *gorm.DB
	}
)

func NewMenuAllergenRepository(db *gorm.DB) MenuAllergenRepository{
	return &menuAllergenRepository{db}
}

func (mr *menuAllergenRepository) GetMenuAllergenList(menuAllergen *[]model.MenuAllergen, menuID uint) error{

	if err := mr.db.Preload("Allergen", func(db *gorm.DB) *gorm.DB{return db.Select("id", "allergen_name") }).Where("menu_item_id=?", menuID).Find(menuAllergen).Error; err != nil{
		return err
	}
	return nil
}

func (mr *menuAllergenRepository) CreateMenuAllergen(menuAllergen *model.MenuAllergen) error{

	if err := mr.db.Create(menuAllergen).Error; err != nil{
		return err
	}
	return nil
}

func (mr *menuAllergenRepository) UpdateMenuAllergen(menuAllergen *model.MenuAllergen, id uint) error{

	result := mr.db.Model(menuAllergen).Where("id=?", id).Updates(menuAllergen)

	if err := result.Error; err != nil{
		return err
	}
	return nil
}

func (mr *menuAllergenRepository) DeleteMenuAllergen(menuAllergen *model.MenuAllergen, id uint) error{

	result := mr.db.Model(menuAllergen).Where("id=?", id).Update("is_deleted", 1)

	if err := result.Error; err != nil{
		return err
	}
	return nil
}