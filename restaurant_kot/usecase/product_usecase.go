package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/repository"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/validator"
)

type (
	ProductUsecase interface {
		GetAllProduct(organizationID uint, restaurantID uint) ([]model.ProductResponse, error)
		CreateProduct(product model.Product) (model.ProductResponse, error)
		UpdateProduct(product model.Product, id uint) (model.ProductResponse, error)
		DeleteProduct(product model.Product, id uint) error
	}

	productUsecase struct{
		pr repository.ProductRepository
		pv validator.ProductValidator
	}
)

func NewProductUsecase(pr repository.ProductRepository, pv validator.ProductValidator) ProductUsecase{
	return &productUsecase{pr, pv}
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

	if err := pu.pr.CreateProduct(&product); err != nil{
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

func (pu *productUsecase) UpdateProduct(product model.Product, id uint) (model.ProductResponse, error){

	if err := pu.pv.ProductValidate(&product); err != nil{
		return model.ProductResponse{}, err
	}

	if err := pu.pr.UpdateProduct(&product, id); err != nil{
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

func (pu *productUsecase) DeleteProduct(product model.Product, id uint) error{

	if err := pu.pr.DeleteProduct(&product, id); err != nil{
		return err
	}

	return nil
}
