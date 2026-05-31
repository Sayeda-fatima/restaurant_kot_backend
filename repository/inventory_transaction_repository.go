package repository

import (
	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
	"gorm.io/gorm"
)

type (
	InventoryTransactionRepository interface {
		GetInventoryTransactionList(inventoryTransaction *[]model.InventoryTransaction, organizationID uint, restaurantID uint) error
		CreateInventoryTransaction(inventoryTransaction *model.InventoryTransaction) error
		UpdateInventoryTransaction(inventoryTransaction *model.InventoryTransaction, id uint, organizationID uint, restaurantID uint) error
		GetPurchaseDuringTimePeriod(result *map[string]interface{}, organizationID uint, restaurantID uint, dateFrom string, dateTo string) error
		GetWasteDuringTimePeriod(result *[]map[string]interface{}, organizationID uint, restaurantID uint, dateFrom string, dateTo string) error
		DailyConsumption(result *[]map[string]interface{}, organizationID uint, restaurantID uint) error
	}

	inventoryTransactionRepository struct{
		db *gorm.DB
	}
)

func NewInventoryTransactionRepository(db *gorm.DB) InventoryTransactionRepository{
	return &inventoryTransactionRepository{db}
}

func (ir *inventoryTransactionRepository) GetInventoryTransactionList(inventoryTransaction *[]model.InventoryTransaction, organizationID uint, restaurantID uint) error{

	if err := ir.db.Where("organization_id=? and restaurant_id=?", organizationID, restaurantID).Find(inventoryTransaction).Error; err != nil{
		return err
	}

	return nil
}

func (ir *inventoryTransactionRepository) CreateInventoryTransaction(inventoryTransaction *model.InventoryTransaction) error{

	if err := ir.db.Create(inventoryTransaction).Error; err != nil{
		return err
	}

	return nil
}

func (ir *inventoryTransactionRepository) UpdateInventoryTransaction(inventoryTransaction *model.InventoryTransaction, id uint, organizationID uint, restaurantID uint) error{

	result := ir.db.Model(inventoryTransaction).Where("id=? and organization_id=? and restaurant_id=?", id, organizationID, restaurantID).Updates(inventoryTransaction)

	if err := result.Error; err != nil{
		return err
	}

	return nil
}

func (ir *inventoryTransactionRepository) GetPurchaseDuringTimePeriod(result *map[string]interface{}, organizationID uint, restaurantID uint, dateFrom string, dateTo string) error{

	err := ir.db.Raw(`SELECT SUM(total_cost) as purchased_inventory
						from inventory_transactions 
						where transaction_type='purchase' and organization_id=? and restaurant_id=? and recorded_at between ? and ?
					`, organizationID, restaurantID, dateFrom, dateTo).Find(result).Error

	if err != nil{
		return err
	}

	return nil
}

func (ir *inventoryTransactionRepository) GetWasteDuringTimePeriod(result *[]map[string]interface{}, organizationID uint, restaurantID uint, dateFrom string, dateTo string) error{

	err := ir.db.Raw(`SELECT product_id, -1*sum(quantity) as waste_quantity, sum(total_cost) as waste_cost
						from inventory_transactions
						where transaction_type='waste' and organization_id=? and restaurant_id=? and recorded_at between ? and ?
						group by product_id
					`, organizationID, restaurantID, dateFrom, dateTo).Find(result).Error
	
	if err != nil{
		return err
	}

	return nil
}

func (ir *inventoryTransactionRepository) DailyConsumption(result *[]map[string]interface{}, organizationID uint, restaurantID uint) error{

	err := ir.db.Raw(`SELECT product_id, -1*sum(quantity) as daily_consumption
						from inventory_transactions
						where transaction_type='sale' and organization_id=? and restaurant_id=? and date(recorded_at)=CURDATE()
						group by product_id
					`, organizationID, restaurantID).Find(result).Error

	if err != nil{
		return err
	}
	
	return nil
}