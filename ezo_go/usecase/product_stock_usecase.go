package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/NazishAhsan/easy_busy_book_go/repository"
	"github.com/NazishAhsan/easy_busy_book_go/validator"
	"gorm.io/gorm"
)

type (
	ProductStockUsecase interface {
		GetProductStockList(organizationID uint) ([]model.ProductStockResponse, error)
		CreateProductStock(organizationID uint, productID uint, quantity int) (model.ProductStockResponse, error)
		UpdateProductStock(productStock model.ProductStock, id uint, organizationID uint) (model.ProductStockResponse, error)
		DeleteProductStock(productStock model.ProductStock, id uint, organizationID uint) error
		GetProductStockListByUpdateType(organizationID uint, term string)([]model.ProductStockResponse, error)
	}

	productStockUsecase struct {
		pr repository.ProductStockRepository
		pv validator.ProductStockValidator
		ps ProductUsecase
		db *gorm.DB
	}
)

func NewProductStockUsecase(pr repository.ProductStockRepository, pv validator.ProductStockValidator, ps ProductUsecase, db *gorm.DB) ProductStockUsecase {
	return &productStockUsecase{pr, pv, ps, db}
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

func (pu *productStockUsecase) CreateProductStock(organizationID uint, productID uint, quantity int) (model.ProductStockResponse, error) {

	product, _ := pu.ps.GetProduct(organizationID, productID) 
	productStock := model.ProductStock{
		OrganizationID: product.OrganizationID,
		ProductID: productID,
		ProductName: product.Name,
		ProductStockBeforeUpdate: float64(product.Quantity),
		ProductUpdateType: "add",
		ProductUpdateQuantity: float64(quantity),
		ProductStockAfterUpdate: float64(product.Quantity) + float64(quantity),
	}
	// start transaction
	tx := pu.db.Begin()
	defer func(){
		if r := recover(); r != nil{
			tx.Rollback()
		}
	}()

	if err := pu.pv.ProductStockValidate(productStock); err != nil {
		tx.Rollback()
		return model.ProductStockResponse{}, err
	}
	if err := pu.pr.CreateProductStock(&productStock); err != nil {
		tx.Rollback()
		return model.ProductStockResponse{}, err
	}

	// update quantity in product table
	if _, err := pu.ps.UpdateStockOfProduct(productStock.Product, productStock.ProductID, int(productStock.ProductStockAfterUpdate)); err != nil{
		tx.Rollback()
		return model.ProductStockResponse{}, err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
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

func (pu *productStockUsecase) UpdateProductStock(productStock model.ProductStock, id uint, organizationID uint) (model.ProductStockResponse, error) {

	if err := pu.pv.ProductStockValidate(productStock); err != nil {
		return model.ProductStockResponse{}, err
	}

	if err := pu.pr.UpdateProductStock(&productStock, id, organizationID); err != nil {
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

func (pu *productStockUsecase) DeleteProductStock(productStock model.ProductStock, id uint, organizationID uint) error {

	if err := pu.pr.DeleteProductStock(&productStock, id, organizationID); err != nil {
		return err
	}
	return nil
}

func (pu *productStockUsecase) GetProductStockListByUpdateType(organizationID uint, term string) ([]model.ProductStockResponse, error){

	productStocks := []model.ProductStock{}

	if err := pu.pr.GetProductStockListByUpdateType(&productStocks, organizationID, term); err != nil{
		return nil, err
	}

	resProductStock := []model.ProductStockResponse{}
	for _, v := range(productStocks){
		res := model.ProductStockResponse{
			ID: v.ID,
			OrganizationID: v.OrganizationID,
			OrderID: v.OrderID,
			ProductID: v.ProductID,
			ProductName: v.ProductName,
			ProductStockBeforeUpdate: v.ProductStockBeforeUpdate,
			ProductUpdateQuantity: v.ProductUpdateQuantity,
			ProductUpdateType: v.ProductUpdateType,
			ProductStockAfterUpdate: v.ProductStockAfterUpdate,
		}
		resProductStock = append(resProductStock, res)
	}

	return resProductStock, nil
}
