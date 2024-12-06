package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_go/common"
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/NazishAhsan/easy_busy_book_go/repository"
	"github.com/NazishAhsan/easy_busy_book_go/validator"
)

type (
	ProductUsecase interface{
		GetProductList(organizationID uint) ([]model.ProductListResponse, error)
		CreateProduct(product model.Product) (model.ProductResponse, error)
		GetProduct(organizationID uint, id uint) (model.ProductResponse, error)
		UpdateProduct(product model.Product, id uint, organizationID uint) (model.ProductResponse, error)
		DeleteProduct(product model.Product, id uint, organizationID uint) error
		SearchProduct(organizationID uint, term string) ([]model.ProductResponse, error)
		GetProductByBarcode(organizationID uint, barcode string)(model.ProductResponse, error)
		UpdateStockOfProduct(product model.Product, id uint, quantity int) (model.ProductResponse, error)
	}

	productUsecase struct{
		pr repository.ProductRepository
		pv validator.ProductValidator
	}
)

func NewProductUsecase (pr repository.ProductRepository, pv validator.ProductValidator) ProductUsecase{
	return &productUsecase{pr,pv}
}

func (pu *productUsecase) GetProductList (organizationID uint) ([]model.ProductListResponse, error){

	products := []model.Product{}

	if err := pu.pr.GetProductList(&products, organizationID); err!=nil{
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("GetProductList")
		return nil, err
	}

	resProducts := []model.ProductListResponse{}

	for _, v := range(products){
		res := model.ProductListResponse{
			OrganizationID: v.OrganizationID,
			ID: v.ID,
			Name: v.Name,
			Image: v.Image,
			SellPrice: v.SellPrice,
			Quantity: v.Quantity,
			CategoryID: v.CategoryID,
		}

		resProducts = append(resProducts, res)
	}
	return resProducts, nil
}

func (pu *productUsecase) CreateProduct (product model.Product) (model.ProductResponse, error){

	if err := pu.pv.ProductValidate(product); err!=nil{
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("CreateProduct")
		return model.ProductResponse{}, err
	}

	if err := pu.pr.CreateProduct(&product); err!=nil{
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("CreateProduct")
		return model.ProductResponse{}, err
	}

	resProduct := model.ProductResponse{
		ID: product.ID,
		OrganizationID: product.OrganizationID,
		Name: product.Name,
		Image: product.Image,
		SellPrice: product.SellPrice,
		MeasuringUnit: product.MeasuringUnit,
		CategoryID: product.CategoryID,
		Quantity: product.Quantity,
		Mrp: product.Mrp,
		PurchasePrice: product.PurchasePrice,
		AcSalePrice: product.AcSalePrice,
		NonAcSalePrice: product.NonAcSalePrice,
		OnlineDeliverySellPrice: product.OnlineDeliverySellPrice,
		OnlineSellPrice: product.OnlineSellPrice,
		Tax: product.Tax,
		PriceWithTax: product.PriceWithTax,
		Cess: product.Cess,
		HsnCode: product.HsnCode,
		Description: product.Description,
		MinimumStockRequired: product.MinimumStockRequired,
		StorageLocation: product.StorageLocation,
		BulkPurchaseUnit: product.BulkPurchaseUnit,
		RetailSaleUnitPerBulkPurchase: product.RetailSaleUnitPerBulkPurchase,
		BulkPurchaseUnitPerRetailSale: product.BulkPurchaseUnitPerRetailSale,
		ExpiryDate: product.ExpiryDate,
		ShowProductOnlineStore: product.ShowProductOnlineStore,
	}

	return resProduct, nil
}

func (pu *productUsecase) GetProduct(organizationID uint, id uint) (model.ProductResponse, error){

	product := model.Product{}
	if err := pu.pr.GetProduct(&product, organizationID, id); err != nil{
		return model.ProductResponse{}, err
	}

	resProduct := model.ProductResponse{
		ID: product.ID,
		OrganizationID: product.OrganizationID,
		Name: product.Name,
		Image: product.Image,
		SellPrice: product.SellPrice,
		MeasuringUnit: product.MeasuringUnit,
		CategoryID: product.CategoryID,
		Quantity: product.Quantity,
		Mrp: product.Mrp,
		PurchasePrice: product.PurchasePrice,
		AcSalePrice: product.AcSalePrice,
		NonAcSalePrice: product.NonAcSalePrice,
		OnlineDeliverySellPrice: product.OnlineDeliverySellPrice,
		OnlineSellPrice: product.OnlineSellPrice,
		Tax: product.Tax,
		PriceWithTax: product.PriceWithTax,
		Cess: product.Cess,
		HsnCode: product.HsnCode,
		Description: product.Description,
		MinimumStockRequired: product.MinimumStockRequired,
		StorageLocation: product.StorageLocation,
		BulkPurchaseUnit: product.BulkPurchaseUnit,
		RetailSaleUnitPerBulkPurchase: product.RetailSaleUnitPerBulkPurchase,
		BulkPurchaseUnitPerRetailSale: product.BulkPurchaseUnitPerRetailSale,
		ExpiryDate: product.ExpiryDate,
		ShowProductOnlineStore: product.ShowProductOnlineStore,
		ProductImages: product.ProductImages,
	}

	return resProduct, nil
}

func (pu *productUsecase) UpdateProduct (product model.Product, id uint, organizationID uint) (model.ProductResponse, error){

	if err := pu.pv.ProductValidate(product); err!=nil{
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("UpdateProduct")
		return model.ProductResponse{}, err
	}

	if err := pu.pr.UpdateProduct(&product, id, organizationID); err!=nil{
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("UpdateProduct")
		return model.ProductResponse{}, err
	}

	resProduct := model.ProductResponse{
		ID: product.ID,
		OrganizationID: product.OrganizationID,
		Name: product.Name,
		Image: product.Image,
		SellPrice: product.SellPrice,
		MeasuringUnit: product.MeasuringUnit,
		CategoryID: product.CategoryID,
		Quantity: product.Quantity,
		Mrp: product.Mrp,
		PurchasePrice: product.PurchasePrice,
		AcSalePrice: product.AcSalePrice,
		NonAcSalePrice: product.NonAcSalePrice,
		OnlineDeliverySellPrice: product.OnlineDeliverySellPrice,
		OnlineSellPrice: product.OnlineSellPrice,
		Tax: product.Tax,
		PriceWithTax: product.PriceWithTax,
		Cess: product.Cess,
		HsnCode: product.HsnCode,
		Description: product.Description,
		MinimumStockRequired: product.MinimumStockRequired,
		StorageLocation: product.StorageLocation,
		BulkPurchaseUnit: product.BulkPurchaseUnit,
		RetailSaleUnitPerBulkPurchase: product.RetailSaleUnitPerBulkPurchase,
		BulkPurchaseUnitPerRetailSale: product.BulkPurchaseUnitPerRetailSale,
		ExpiryDate: product.ExpiryDate,
		ShowProductOnlineStore: product.ShowProductOnlineStore,
	}
	return resProduct, nil
}

func (pu *productUsecase) DeleteProduct (product model.Product, id uint, organizationID uint) error{

	if err := pu.pr.DeleteProduct(&product, id, organizationID); err!=nil{
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("DeleteProduct")
		return err
	}
	return nil
}

func (pu *productUsecase) SearchProduct(organizationID uint, term string) ([]model.ProductResponse, error){

	products := []model.Product{}
	if err := pu.pr.SearchProduct(&products, organizationID, term); err!=nil{
		return nil, err
	}

	resProducts := []model.ProductResponse{}
	for _, v := range products{
		res := model.ProductResponse{
			OrganizationID: v.OrganizationID,
			ID: v.ID,
			Name: v.Name,
			Image: v.Image,
			SellPrice: v.SellPrice,
			Quantity: v.Quantity,
			CategoryID: v.CategoryID,
		}
		resProducts = append(resProducts, res)
	}

	return resProducts, nil
}

func (pu *productUsecase) GetProductByBarcode(organizationID uint, barcode string) (model.ProductResponse, error){

	product := model.Product{}
	if err := pu.pr.GetProductByBarcode(&product, organizationID, barcode); err != nil{
		return model.ProductResponse{}, nil
	}

	resProduct := model.ProductResponse{
		ID: product.ID,
		OrganizationID: product.OrganizationID,
		Name: product.Name,
		Image: product.Image,
		SellPrice: product.SellPrice,
		MeasuringUnit: product.MeasuringUnit,
		CategoryID: product.CategoryID,
		Quantity: product.Quantity,
		Mrp: product.Mrp,
		PurchasePrice: product.PurchasePrice,
		AcSalePrice: product.AcSalePrice,
		NonAcSalePrice: product.NonAcSalePrice,
		OnlineDeliverySellPrice: product.OnlineDeliverySellPrice,
		OnlineSellPrice: product.OnlineSellPrice,
		Tax: product.Tax,
		PriceWithTax: product.PriceWithTax,
		Cess: product.Cess,
		HsnCode: product.HsnCode,
		Description: product.Description,
		MinimumStockRequired: product.MinimumStockRequired,
		StorageLocation: product.StorageLocation,
		BulkPurchaseUnit: product.BulkPurchaseUnit,
		RetailSaleUnitPerBulkPurchase: product.RetailSaleUnitPerBulkPurchase,
		BulkPurchaseUnitPerRetailSale: product.BulkPurchaseUnitPerRetailSale,
		ExpiryDate: product.ExpiryDate,
		ShowProductOnlineStore: product.ShowProductOnlineStore,
		Barcode: product.Barcode,
	}
	return resProduct, nil
}

func (pu *productUsecase) UpdateStockOfProduct(product model.Product, id uint, quantity int) (model.ProductResponse, error){

	if err := pu.pr.UpdateStockOfProduct(&product, id, quantity); err != nil{
		return model.ProductResponse{}, err
	}

	resProduct := model.ProductResponse{
		ID: product.ID,
		OrganizationID: product.OrganizationID,
		Name: product.Name,
		Image: product.Image,
		SellPrice: product.SellPrice,
		MeasuringUnit: product.MeasuringUnit,
		CategoryID: product.CategoryID,
		Quantity: product.Quantity,
		//Mrp: product.Mrp,
	}
	return resProduct, nil
}