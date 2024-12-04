package repository

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"gorm.io/gorm"
)

type (
	AllergenRepository interface {
		GetAllergenList(allergen *[]model.Allergen, organizationID uint, restaurantID uint) error
		CreateAllergen(allergen *model.Allergen) error
		UpdateAllergen(allergen *model.Allergen, id uint) error
		DeleteAllergen(allergen *model.Allergen, id uint) error
	}

	allergenRepository struct{
		db *gorm.DB
	}
)

func NewAllergenRepository(db *gorm.DB) AllergenRepository{
	return &allergenRepository{db}
}

func (ar *allergenRepository) GetAllergenList(allergen *[]model.Allergen, organizationID uint, restaurantID uint) error{

	if err := ar.db.Where("organization_id=? and restaurant_id=? and is_deleted=0", organizationID, restaurantID).Find(allergen).Error; err != nil{
		return err
	}
	return nil
}

func (ar *allergenRepository) CreateAllergen(allergen *model.Allergen) error{

	if err := ar.db.Create(allergen).Error; err != nil{
		return err
	}
	return nil
}

func (ar *allergenRepository) UpdateAllergen(allergen *model.Allergen, id uint) error{

	result := ar.db.Model(allergen).Where("id=?",id).Updates(allergen)

	if err := result.Error; err != nil{
		return err
	}
	return nil
}

func (ar *allergenRepository) DeleteAllergen(allergen *model.Allergen, id uint) error{

	result := ar.db.Model(allergen).Where("id=?",id).Update("is_deleted",1)

	if err := result.Error; err != nil{
		return err
	}
	return nil
}