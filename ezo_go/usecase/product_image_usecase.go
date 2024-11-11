package usecase

import (
	"mime/multipart"

	"github.com/NazishAhsan/easy_busy_book_go/common"
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/NazishAhsan/easy_busy_book_go/repository"
	"github.com/NazishAhsan/easy_busy_book_go/validator"
)

type (
	ProductImageUsecase interface {
		GetProductImageList(organizationID uint, productID uint) ([]model.ProductImageResponse, error)
		AddProductImage(file *multipart.FileHeader, organizationID uint, productID uint) (model.ProductImageResponse, error)
		DeleteProductImage(productImage model.ProductImage, id uint) error
	}

	productImageUsecase struct {
		pr repository.ProductImageRepository
		pv validator.ProductImageValidator
		ic common.ImageUpload
	}
)

func NewProductImageUsecase(pr repository.ProductImageRepository, pv validator.ProductImageValidator, ic common.ImageUpload) ProductImageUsecase {
	return &productImageUsecase{pr, pv, ic}
}

func (pu *productImageUsecase) GetProductImageList(organizationID uint, productID uint) ([]model.ProductImageResponse, error) {

	productImages := []model.ProductImage{}
	if err := pu.pr.GetProductImageList(&productImages, organizationID, productID); err != nil {
		return nil, err
	}

	resProductImage := []model.ProductImageResponse{}
	for _, v := range productImages {
		res := model.ProductImageResponse{
			ID:             v.ID,
			OrganizationID: v.OrganizationID,
			ProductID:      v.ProductID,
			Url:            v.Url,
		}
		resProductImage = append(resProductImage, res)
	}
	return resProductImage, nil
}

func (pu *productImageUsecase) AddProductImage(file *multipart.FileHeader, organizationID uint, productID uint) (model.ProductImageResponse, error) {

	imagePath, err := pu.ic.UploadImage(file, "public/product_images/")
	if err != nil {
		return model.ProductImageResponse{}, err
	}
	productImage := model.ProductImage{
		OrganizationID: organizationID,
		ProductID: productID,
		Url:       imagePath,
	}

	if err := pu.pv.ProductImageValidate(productImage); err != nil {
		return model.ProductImageResponse{}, err
	}

	if err := pu.pr.AddProductImage(&productImage); err != nil {
		return model.ProductImageResponse{}, err
	}

	resProductImage := model.ProductImageResponse{
		ID:             productImage.ID,
		OrganizationID: productImage.OrganizationID,
		ProductID:      productImage.ProductID,
		Url:            productImage.Url,
	}

	return resProductImage, nil
}

func (pu *productImageUsecase) DeleteProductImage(productImage model.ProductImage, id uint) error{

	if err := pu.pr.DeleteProductImage(&productImage, id); err!=nil{
		return err
	}
	return nil
}
