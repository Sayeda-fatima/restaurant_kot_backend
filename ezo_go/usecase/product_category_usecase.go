package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_go/common"
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/NazishAhsan/easy_busy_book_go/repository"
	"github.com/NazishAhsan/easy_busy_book_go/validator"
)

type(
	ProductCategoryUsecase interface{
		GetProductCategoryList (organizationID uint) ([]model.ProductCategoryResponse, error)
		CreateProductCategory (productCategory model.ProductCategory) (model.ProductCategoryResponse, error)
		UpdateProductCategory (productCategory model.ProductCategory, id uint) (model.ProductCategoryResponse, error)
		DeleteProductCategory (productCategory model.ProductCategory, id uint) error
	}

	productCategoryUsecase struct{
		pr repository.ProductCategoryRepository
		pv validator.ProductCategoryValidator
	}
)

func NewProductCategoryUsecase (pr repository.ProductCategoryRepository, pv validator.ProductCategoryValidator) ProductCategoryUsecase{
	return &productCategoryUsecase{pr,pv}
}

func (pu *productCategoryUsecase) GetProductCategoryList (organizationID uint) ([]model.ProductCategoryResponse, error){

	productCategory := []model.ProductCategory{}

	if err := pu.pr.GetProductCategoryList(&productCategory, organizationID); err!=nil{
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("GetProductCategoryList")
		return nil, err
	}

	resProductCategory := []model.ProductCategoryResponse{}

	for _,v :=range(productCategory){
		res := model.ProductCategoryResponse{
			ID: v.ID,
			Category: v.Category,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}

		resProductCategory = append(resProductCategory, res)
	}

	return resProductCategory, nil
}

func (pu *productCategoryUsecase) CreateProductCategory (productCategory model.ProductCategory) (model.ProductCategoryResponse, error){

	if err := pu.pv.ProductCategoryValidate(productCategory); err!=nil{
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("CreateProductCategoryList")
		return model.ProductCategoryResponse{}, err
	}

	if err := pu.pr.CreateProductCategory(&productCategory); err!=nil{
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("CreateProductCategoryList")
		return model.ProductCategoryResponse{}, err
	}

	resProductCategory := model.ProductCategoryResponse{
		ID: productCategory.ID,
		Category: productCategory.Category,
		CreatedAt: productCategory.CreatedAt,
		UpdatedAt: productCategory.UpdatedAt,
	}

	return resProductCategory, nil
}

func (pu *productCategoryUsecase) UpdateProductCategory (productCategory model.ProductCategory, id uint) (model.ProductCategoryResponse, error){

	if err := pu.pv.ProductCategoryValidate(productCategory); err!=nil{
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("UpdateProductCategoryList")
		return model.ProductCategoryResponse{}, err
	}

	if err := pu.pr.UpdateProductCategory(&productCategory, id); err!=nil{
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("UpdateProductCategoryList")
		return model.ProductCategoryResponse{}, err
	}

	resProductCategory := model.ProductCategoryResponse{
		ID: productCategory.ID,
		Category: productCategory.Category,
		CreatedAt: productCategory.CreatedAt,
		UpdatedAt: productCategory.UpdatedAt,
	}

	return resProductCategory, nil
}

func (pu *productCategoryUsecase) DeleteProductCategory (productCategory model.ProductCategory, id uint) error{

	if err := pu.pr.DeleteProductCategory(&productCategory, id); err!=nil{
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("DeleteProductCategoryList")
		return err
	}
	return nil
}