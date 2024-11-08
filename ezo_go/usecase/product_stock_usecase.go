package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/NazishAhsan/easy_busy_book_go/repository"
	"github.com/NazishAhsan/easy_busy_book_go/validator"
)

type (
	ProductStockUsecase interface {
		GetProductStockList(organizationID uint) ([]model.ProductStockResponse, error)
		CreateProductStock(productStock model.ProductStock) (model.ProductStockResponse, error)
		UpdateProductStock(productStock model.ProductStock, id uint) (model.ProductStockResponse, error)
		DeleteProductStock(productStock model.ProductStock, id uint) error
	}

	productStockUsecase struct {
		pr repository.ProductStockRepository
		pv validator.ProductStockValidator
	}
)

func NewProductStockUsecase(pr repository.ProductStockRepository, pv validator.ProductStockValidator) ProductStockUsecase {
	return &productStockUsecase{pr, pv}
}

func (pu *productStockUsecase) GetProductStockList(organizationID uint) ([]model.ProductStockResponse, error) {

	productStocks := []model.ProductStock{}
	if err := pu.pr.GetProductStockList(&productStocks, organizationID); err != nil {
		return nil, err
	}

	resProductStock := []model.ProductStockResponse{}
	for _, v := range productStocks {
		res := model.ProductStockResponse{
			ID:                       v.ID,
			OrganizationID:           v.OrganizationID,
			OrderID:                  v.OrderID,
			ProductID:                v.ProductID,
			ProductName:              v.ProductName,
			ProductStockBeforeUpdate: v.ProductStockBeforeUpdate,
			ProductUpdateQuantity:    v.ProductUpdateQuantity,
			ProductUpdateType:        v.ProductUpdateType,
			ProductStockAfterUpdate:  v.ProductStockAfterUpdate,
		}
		resProductStock = append(resProductStock, res)
	}
	return resProductStock, nil
}

func (pu *productStockUsecase) CreateProductStock(productStock model.ProductStock) (model.ProductStockResponse, error) {

	if err := pu.pv.ProductStockValidate(productStock); err != nil {
		return model.ProductStockResponse{}, err
	}
	if err := pu.pr.CreateProductStock(&productStock); err != nil {
		return model.ProductStockResponse{}, err
	}

	resProductStock := model.ProductStockResponse{
		ID:                       productStock.ID,
		OrganizationID:           productStock.OrganizationID,
		OrderID:                  productStock.OrderID,
		ProductID:                productStock.ProductID,
		ProductName:              productStock.ProductName,
		ProductStockBeforeUpdate: productStock.ProductStockBeforeUpdate,
		ProductUpdateType:        productStock.ProductUpdateType,
		ProductUpdateQuantity:    productStock.ProductUpdateQuantity,
		ProductStockAfterUpdate:  productStock.ProductStockAfterUpdate,
	}
	return resProductStock, nil
}

func (pu *productStockUsecase) UpdateProductStock(productStock model.ProductStock, id uint) (model.ProductStockResponse, error) {

	if err := pu.pv.ProductStockValidate(productStock); err != nil {
		return model.ProductStockResponse{}, err
	}

	if err := pu.pr.UpdateProductStock(&productStock, id); err != nil {
		return model.ProductStockResponse{}, err
	}

	resProductStock := model.ProductStockResponse{
		ID:                       productStock.ID,
		OrganizationID:           productStock.OrganizationID,
		OrderID:                  productStock.OrderID,
		ProductID:                productStock.ProductID,
		ProductName:              productStock.ProductName,
		ProductStockBeforeUpdate: productStock.ProductStockBeforeUpdate,
		ProductUpdateType:        productStock.ProductUpdateType,
		ProductUpdateQuantity:    productStock.ProductUpdateQuantity,
		ProductStockAfterUpdate:  productStock.ProductStockAfterUpdate,
	}
	return resProductStock, nil
}

func (pu *productStockUsecase) DeleteProductStock(productStock model.ProductStock, id uint) error {

	if err := pu.pr.DeleteProductStock(&productStock, id); err != nil {
		return err
	}
	return nil
}
