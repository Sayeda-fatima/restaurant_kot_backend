package repository

import (
	"fmt"

	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
	"gorm.io/gorm"
)

type (
	RestaurantTableRepository interface {
		GetRestaurantTableList(restaurantTables *[]model.RestaurantTable, organizationID uint, restaurantID uint) error
		CreateRestaurantTable(restaurantTable *model.RestaurantTable) error
		UpdateRestaurantTable(restaurantTable *model.RestaurantTable, id uint, organizationID uint, restaurantID uint) error
		DeleteRestaurantTable(restaurantTable *model.RestaurantTable, id uint, organizationID uint, restaurantID uint) error
	}

	restaurantTableRepository struct{
		db *gorm.DB
	}
)

func NewRestaurantTableRepository(db *gorm.DB) RestaurantTableRepository{
	return &restaurantTableRepository{db}
}

func (rr *restaurantTableRepository) GetRestaurantTableList(restaurantTables *[]model.RestaurantTable, organizationID uint, restaurantID uint) error{

	if err := rr.db.Where("organization_id=? and restaurant_id=? and is_deleted=0", organizationID, restaurantID).Find(restaurantTables).Error; err != nil{
		return err
	}
	return nil
}

func (rr *restaurantTableRepository) CreateRestaurantTable(restaurantTable *model.RestaurantTable) error{

	if err := rr.db.Create(restaurantTable).Error; err != nil{
		return err
	}
	return nil
}

func (rr *restaurantTableRepository) UpdateRestaurantTable(restaurantTable *model.RestaurantTable, id uint, organizationID uint, restaurantID uint) error{

	result := rr.db.Model(restaurantTable).Where("id=? and organization_id=? and restaurant_id=?", id, organizationID, restaurantID).Updates(restaurantTable)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	
	return nil
}

func (rr *restaurantTableRepository) DeleteRestaurantTable(restaurantTable *model.RestaurantTable, id uint, organizationID uint, restaurantID uint) error{

	result := rr.db.Model(restaurantTable).Where("id=? and organization_id=? and restaurant_id=?", id, organizationID, restaurantID).Update("is_deleted", 1)

	if err := result.Error; err != nil{
		return err
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	
	return nil
}