package repository

import (
	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
	"gorm.io/gorm"
)

type (
	CurrentInventoryRepository interface {
		GetCurrentInventory(currentInventory *model.CurrentInventory, organizationID uint, restaurantID uint, date string) error
		CreateCurrentInventory(currentInventory *model.CurrentInventory) error
		GetBeginningInventory(result *map[string]interface{}, organizationID uint, restaurantID uint, date string) error
	}

	currentInventoryRepository struct{
		db *gorm.DB
	}
)

func NewCurrentInventoryRepository(db *gorm.DB) CurrentInventoryRepository{
	return &currentInventoryRepository{db}
}

func (ir *currentInventoryRepository) GetCurrentInventory(currentInventory *model.CurrentInventory, organizationID uint, restaurantID uint, date string) error{

	if err := ir.db.Where("organization_id=? and restaurant_id=? and date(created_at)=?", organizationID, restaurantID, date).Find(currentInventory).Error; err != nil{
		return err
	}

	return nil
}

func (ir *currentInventoryRepository) CreateCurrentInventory(currentInventory *model.CurrentInventory) error{

	if err := ir.db.Create(currentInventory).Error; err != nil{
		return err
	}

	return nil
}

func (ir *currentInventoryRepository) GetBeginningInventory(result *map[string]interface{}, organizationID uint, restaurantID uint, date string) error{

	err := ir.db.Raw(`SELECT inventory_value as beginning_inventory from current_inventories 
							where organization_id=? and restaurant_id=? and date(created_at) = ?
						`, organizationID, restaurantID, date).Find(result).Error

	if err != nil{
		return err
	}

	return nil
}