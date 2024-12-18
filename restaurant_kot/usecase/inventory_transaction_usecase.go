package usecase

import (
	"fmt"
	"time"

	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/repository"
)

type (
	InventoryTransactionUsecase interface {
		AddStock(organizationID uint, restaurantID uint, productID uint, quantity int, unitCost int) (model.InventoryTransaction, error)
		AdjustStock(organizationID uint, restaurantID uint, productID uint, adjustmentQuantity int, reason string) (model.InventoryTransaction, error)
		RecordWaste(organizationID uint, restaurantID uint, productID uint, wasteQuantity int, reason string) (model.InventoryTransaction, error)
	}

	inventoryTransactionUsecase struct{
		ir repository.InventoryTransactionRepository
		pr repository.ProductRepository
	}
)

func NewInventoryTransactionUsecase(ir repository.InventoryTransactionRepository, pr repository.ProductRepository) InventoryTransactionUsecase{
	return &inventoryTransactionUsecase{ir, pr}
}

func (iu *inventoryTransactionUsecase) AddStock(organizationID uint, restaurantID uint, productID uint, quantity int, unitCost int) (model.InventoryTransaction, error){

	product := model.Product{}
	if err := iu.pr.GetProduct(&product, productID, organizationID, restaurantID); err != nil{
		return model.InventoryTransaction{}, err
	}

	product.Quantity += quantity
	product.UnitCost = unitCost
	product.InventoryValue = product.Quantity * product.UnitCost
	if err := iu.pr.UpdateProduct(&product, product.ID, product.OrganizationID, product.RestaurantID); err != nil{
		return model.InventoryTransaction{}, err
	}

	transaction := model.InventoryTransaction{
		OrganizationID: product.OrganizationID,
		RestaurantID: product.RestaurantID,
		ProductID: product.ID,
		TransactionType: "purchase",
		Quantity: float64(quantity),
		UnitCost: float64(unitCost),
		TotalCost: float64(unitCost)* float64(quantity),
		RecordedAt: time.Now(),
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

	if err := iu.pr.UpdateProductQuantity(&product, product.ID, product.OrganizationID, product.RestaurantID, newQuantity); err != nil{
		return model.InventoryTransaction{}, err
	}

	transaction := model.InventoryTransaction{
		OrganizationID: product.OrganizationID,
		RestaurantID: product.RestaurantID,
		ProductID: product.ID,
		TransactionType: "adjustment",
		Quantity: float64(adjustmentQuantity),
		UnitCost: float64(product.UnitCost),
		TotalCost: float64(adjustmentQuantity)*float64(product.UnitCost),
		Reason: reason,
		RecordedAt: time.Now(),
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

	product.Quantity -= wasteQuantity
	if err := iu.pr.UpdateProductQuantity(&product, product.ID, product.OrganizationID, product.RestaurantID, product.Quantity); err != nil{
		return model.InventoryTransaction{}, err
	}

	transaction := model.InventoryTransaction{
		OrganizationID: product.OrganizationID,
		RestaurantID: product.RestaurantID,
		ProductID: product.ID,
		TransactionType: "waste",
		Quantity: -float64(wasteQuantity),
		UnitCost: float64(product.UnitCost),
		TotalCost: float64(product.UnitCost)*float64(wasteQuantity),
		Reason: reason,
		RecordedAt: time.Now(),
	}

	if err := iu.ir.CreateInventoryTransaction(&transaction); err != nil{
		return model.InventoryTransaction{}, err
	}

	return transaction, nil
}