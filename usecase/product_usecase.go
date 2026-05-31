package usecase

import (
	"time"

	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
	"github.com/Sayeda-fatima/restaurant_kot_backend/repository"
	"github.com/Sayeda-fatima/restaurant_kot_backend/validator"
)

type (
	ProductUsecase interface {
		GetAllProduct(organizationID uint, restaurantID uint) ([]model.ProductResponse, error)
		CreateProduct(product model.Product) (model.ProductResponse, error)
		UpdateProduct(product model.Product, id uint, organizationID uint, restaurantID uint) (model.ProductResponse, error)
		DeleteProduct(product model.Product, id uint, organizationID uint, restaurantID uint) error
		InventoryValue(organizationID uint, restaurantID uint) (map[string]interface{}, error)
	}

	productUsecase struct{
		pr repository.ProductRepository
		pv validator.ProductValidator
		ir repository.InventoryTransactionRepository
	}
)

func NewProductUsecase(pr repository.ProductRepository, pv validator.ProductValidator, ir repository.InventoryTransactionRepository) ProductUsecase{
	return &productUsecase{pr, pv, ir}
}

func (pu *productUsecase) GetAllProduct(organizationID uint, restaurantID uint) ([]model.ProductResponse, error){

	products := []model.Product{}

	if err := pu.pr.GetAllProduct(&products, organizationID, restaurantID); err != nil{
		return nil, err
	}

	resProducts := []model.ProductResponse{}
	for _, v := range(products){
		res := model.ProductResponse{
			ID: v.ID,
			OrganizationID: v.OrganizationID,
			RestaurantID: v.RestaurantID,
			Name: v.Name,
			Description: v.Description,
			Category: v.Category,
			UnitOfMeasure: v.UnitOfMeasure,
			UnitCost: v.UnitCost,
			Quantity: v.Quantity,
			InventoryValue: v.InventoryValue,
		}
		resProducts = append(resProducts, res)
	}
	return resProducts, nil
}

func (pu *productUsecase) CreateProduct(product model.Product) (model.ProductResponse, error){

	if err := pu.pv.ProductValidate(&product); err != nil{
		return model.ProductResponse{}, err
	}

	product.InventoryValue = product.Quantity * product.UnitCost
	if err := pu.pr.CreateProduct(&product); err != nil{
		return model.ProductResponse{}, err
	}

	transaction := model.InventoryTransaction{
		OrganizationID: product.OrganizationID,
		RestaurantID: product.RestaurantID,
		ProductID: product.ID,
		TransactionType: "purchase",
		Quantity: float64(product.Quantity),
		UnitCost: float64(product.UnitCost),
		TotalCost: float64(product.UnitCost)* float64(product.Quantity),
		RecordedAt: time.Now(),
	}

	if err := pu.ir.CreateInventoryTransaction(&transaction); err != nil{
		return model.ProductResponse{}, err
	}

	resProduct := model.ProductResponse{
		ID: product.ID,
		OrganizationID: product.OrganizationID,
		RestaurantID: product.RestaurantID,
		Name: product.Name,
		Description: product.Description,
		Category: product.Category,
		UnitOfMeasure: product.UnitOfMeasure,
		UnitCost: product.UnitCost,
		Quantity: product.Quantity,
		InventoryValue: product.InventoryValue,
	}

	return resProduct, nil
}

func (pu *productUsecase) UpdateProduct(product model.Product, id uint, organizationID uint, restaurantID uint) (model.ProductResponse, error){

	if err := pu.pv.ProductValidate(&product); err != nil{
		return model.ProductResponse{}, err
	}

	product.InventoryValue = product.Quantity * product.UnitCost
	if err := pu.pr.UpdateProduct(&product, id, organizationID, restaurantID); err != nil{
		return model.ProductResponse{}, err
	}

	resProduct := model.ProductResponse{
		ID: product.ID,
		OrganizationID: product.OrganizationID,
		RestaurantID: product.RestaurantID,
		Name: product.Name,
		Description: product.Description,
		Category: product.Category,
		UnitOfMeasure: product.UnitOfMeasure,
		UnitCost: product.UnitCost,
		Quantity: product.Quantity,
		InventoryValue: product.InventoryValue,
	}

	return resProduct, nil
}

func (pu *productUsecase) DeleteProduct(product model.Product, id uint, organizationID uint, restaurantID uint) error{

	if err := pu.pr.DeleteProduct(&product, id, organizationID, restaurantID); err != nil{
		return err
	}

	return nil
}

func (pu *productUsecase) InventoryValue(organizationID uint, restaurantID uint) (map[string]interface{}, error){

	var result map[string]interface{}
	if err := pu.pr.InventoryValue(&result, organizationID, restaurantID); err != nil{
		return nil, err
	}

	return result, nil
}
