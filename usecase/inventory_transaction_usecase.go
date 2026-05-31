package usecase

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Sayeda-fatima/restaurant_kot_backend/common"
	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
	"github.com/Sayeda-fatima/restaurant_kot_backend/repository"
)

type (
	InventoryTransactionUsecase interface {
		GetInventoryTransactionList(organizationID uint, restaurantID uint) ([]model.InventoryTransaction, error)
		AddStock(organizationID uint, restaurantID uint, productID uint, quantity int, unitCost int) (model.InventoryTransaction, error)
		AdjustStock(organizationID uint, restaurantID uint, productID uint, adjustmentQuantity int, reason string) (model.InventoryTransaction, error)
		RecordWaste(organizationID uint, restaurantID uint, productID uint, wasteQuantity int, reason string) (model.InventoryTransaction, error)
		CreateCurrentInventoryValue(organizationID uint, restaurantID uint) (model.CurrentInventory, error)
		GetCostOfGoodsSold(organizationID uint, restaurantID uint, dateFrom string, dateTo string) (map[string]interface{}, error)
		GetWasteDuringTimePeriod(organizationID uint, restaurantID uint, dateFrom string, dateTo string) ([]map[string]interface{}, error)
		GetDailyConsumption(organizationID uint, restaurantID uint)([]map[string]interface{}, error)
	}

	inventoryTransactionUsecase struct{
		ir repository.InventoryTransactionRepository
		pr repository.ProductRepository
		cr repository.CurrentInventoryRepository
	}
)

func NewInventoryTransactionUsecase(ir repository.InventoryTransactionRepository, pr repository.ProductRepository, cr repository.CurrentInventoryRepository) InventoryTransactionUsecase{
	return &inventoryTransactionUsecase{ir, pr, cr}
}


func (iu *inventoryTransactionUsecase) GetInventoryTransactionList(organizationID uint, restaurantID uint) ([]model.InventoryTransaction, error){

	inventoryTransactions := []model.InventoryTransaction{}
	if err := iu.ir.GetInventoryTransactionList(&inventoryTransactions, organizationID, restaurantID); err != nil{
		return nil, err
	}

	return inventoryTransactions, nil
}

func (iu *inventoryTransactionUsecase) AddStock(organizationID uint, restaurantID uint, productID uint, quantity int, unitCost int) (model.InventoryTransaction, error){

	product := model.Product{}
	if err := iu.pr.GetProduct(&product, productID, organizationID, restaurantID); err != nil{
		return model.InventoryTransaction{}, err
	}

	transaction := model.InventoryTransaction{
		OrganizationID: product.OrganizationID,
		RestaurantID: product.RestaurantID,
		ProductID: product.ID,
		StockBeforeUpdate: float64(product.Quantity),
		StockUnitPriceBeforeUpdate: float64(product.UnitCost),
		TransactionType: "purchase",
		Quantity: float64(quantity),
		UnitCost: float64(unitCost),
		TotalCost: float64(unitCost)* float64(quantity),
		RecordedAt: time.Now(),
	}

	product.Quantity += quantity
	product.UnitCost = unitCost
	product.InventoryValue = product.Quantity * product.UnitCost
	if err := iu.pr.UpdateProduct(&product, product.ID, product.OrganizationID, product.RestaurantID); err != nil{
		return model.InventoryTransaction{}, err
	}

	if err := iu.ir.CreateInventoryTransaction(&transaction); err != nil{
		return model.InventoryTransaction{}, err
	}

	return transaction, nil
}

func (iu *inventoryTransactionUsecase) AdjustStock(organizationID uint, restaurantID uint, productID uint, adjustmentQuantity int, reason string) (model.InventoryTransaction, error){

	product := model.Product{}
	if err := iu.pr.GetProduct(&product, productID, organizationID, restaurantID); err != nil{
		return model.InventoryTransaction{}, err
	}

	// calculate new stock : adjustmentQuantity +ve for add and -ve for subtract
	newQuantity := product.Quantity + adjustmentQuantity
	if newQuantity < 0{
		return model.InventoryTransaction{}, fmt.Errorf("adjustment would result in negative stock")
	}

	transaction := model.InventoryTransaction{
		OrganizationID: product.OrganizationID,
		RestaurantID: product.RestaurantID,
		ProductID: product.ID,
		StockBeforeUpdate: float64(product.Quantity),
		StockUnitPriceBeforeUpdate: float64(product.UnitCost),
		TransactionType: "adjustment",
		Quantity: float64(adjustmentQuantity),
		UnitCost: float64(product.UnitCost),
		TotalCost: float64(adjustmentQuantity)*float64(product.UnitCost),
		Reason: reason,
		RecordedAt: time.Now(),
	}

	product.Quantity = newQuantity
	product.InventoryValue = newQuantity * product.UnitCost
	if err := iu.pr.UpdateProduct(&product, product.ID, product.OrganizationID, product.RestaurantID); err != nil{
		return model.InventoryTransaction{}, err
	}

	if adjustmentQuantity < 0{
		transaction.TotalCost = -float64(adjustmentQuantity)*float64(product.UnitCost)
	}

	if err := iu.ir.CreateInventoryTransaction(&transaction); err != nil{
		return model.InventoryTransaction{}, err
	}

	return transaction, nil
}

func (iu *inventoryTransactionUsecase) RecordWaste(organizationID uint, restaurantID uint, productID uint, wasteQuantity int, reason string) (model.InventoryTransaction, error){

	product := model.Product{}
	if err := iu.pr.GetProduct(&product, productID, organizationID, restaurantID); err != nil{
		return model.InventoryTransaction{}, err
	}

	if product.Quantity < wasteQuantity{
		return model.InventoryTransaction{}, fmt.Errorf("not enough stock to record waste")
	}

	transaction := model.InventoryTransaction{
		OrganizationID: product.OrganizationID,
		RestaurantID: product.RestaurantID,
		ProductID: product.ID,
		StockBeforeUpdate: float64(product.Quantity),
		StockUnitPriceBeforeUpdate: float64(product.UnitCost),
		TransactionType: "waste",
		Quantity: -float64(wasteQuantity),
		UnitCost: float64(product.UnitCost),
		TotalCost: float64(product.UnitCost)*float64(wasteQuantity),
		Reason: reason,
		RecordedAt: time.Now(),
	}

	product.Quantity -= wasteQuantity
	product.InventoryValue = product.Quantity * product.UnitCost
	if err := iu.pr.UpdateProduct(&product, product.ID, product.OrganizationID, product.RestaurantID); err != nil{
		return model.InventoryTransaction{}, err
	}

	if err := iu.ir.CreateInventoryTransaction(&transaction); err != nil{
		return model.InventoryTransaction{}, err
	}

	return transaction, nil
}

func (iu *inventoryTransactionUsecase) CreateCurrentInventoryValue(organizationID uint, restaurantID uint) (model.CurrentInventory, error){

	var result map[string]interface{}
	if err := iu.pr.InventoryValue(&result, organizationID, restaurantID); err != nil{
		return model.CurrentInventory{}, err
	}

	common.Logger.LogInfo().Msgf("value: %s", result["total_inventory_value"])
	value, _ := strconv.ParseFloat(result["total_inventory_value"].(string), 64)
	common.Logger.LogInfo().Msgf("%f", value)

	currentInventory := model.CurrentInventory{
		OrganizationID: organizationID,
		RestaurantID: restaurantID,
		InventoryValue: float64(value),
	}

	if err := iu.cr.CreateCurrentInventory(&currentInventory); err != nil{
		return model.CurrentInventory{}, err
	}

	return currentInventory, nil
}

func (iu *inventoryTransactionUsecase) GetCostOfGoodsSold(organizationID uint, restaurantID uint, dateFrom string, dateTo string) (map[string]interface{}, error){

	var result map[string]interface{}
	if err := iu.ir.GetPurchaseDuringTimePeriod(&result, organizationID, restaurantID, dateFrom, dateTo); err != nil{
		return nil, err
	}

	if err := iu.cr.GetBeginningInventory(&result, organizationID, restaurantID, dateFrom); err != nil{
		return nil, err
	}

	currentInventory := model.CurrentInventory{}
	if err := iu.cr.GetCurrentInventory(&currentInventory, organizationID, restaurantID, dateTo); err != nil{
		return nil, err
	}

	costOfGoodsSold := result["beginning_inventory"].(float64) + result["purchased_inventory"].(float64) - currentInventory.InventoryValue

	result["cost_of_goods_sold"] = costOfGoodsSold
	result["ending_inventory"] = currentInventory.InventoryValue
	return result, nil
}

func (iu *inventoryTransactionUsecase) GetWasteDuringTimePeriod(organizationID uint, restaurantID uint, dateFrom string, dateTo string) ([]map[string]interface{}, error){

	var result []map[string]interface{}
	if err := iu.ir.GetWasteDuringTimePeriod(&result, organizationID, restaurantID, dateFrom, dateTo); err != nil{
		return nil, err
	}

	return result, nil
}

func (iu *inventoryTransactionUsecase) GetDailyConsumption(organizationID uint, restaurantID uint) ([]map[string]interface{}, error){

	var result []map[string]interface{}
	if err := iu.ir.DailyConsumption(&result, organizationID, restaurantID); err != nil{
		return nil, err
	}

	return result, nil
}