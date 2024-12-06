package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_go/common"
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/NazishAhsan/easy_busy_book_go/repository"
	"github.com/NazishAhsan/easy_busy_book_go/validator"
)

type (
	OrderItemUsecase interface {
		GetOrderItemList(organizationID uint, orderID uint) ([]model.OrderItemResponse, error)
		CreateOrderItem(orderItem model.OrderItem) (model.OrderItemResponse, error)
		UpdateOrderItem(orderItem model.OrderItem, id uint, organizationID uint) (model.OrderItemResponse, error)
		DeleteOrderItem(orderItem model.OrderItem, id uint, organizationID uint) error
		InvoiceCustomer(organizationID uint, orderID uint, dateFrom string, dateTo string)([]model.OrderItemResponse, error)
	}

	orderItemUsecase struct{
		or repository.OrderItemRepository
		ov validator.OrderItemValidator
	}
)

func NewOrderItemUsecase (or repository.OrderItemRepository, ov validator.OrderItemValidator) OrderItemUsecase{
	return &orderItemUsecase{or, ov}
}

func (ou *orderItemUsecase) GetOrderItemList(organizationID uint, orderID uint) ([]model.OrderItemResponse, error){

	orderItems := []model.OrderItem{}

	if err := ou.or.GetOrderItemList(&orderItems, organizationID, orderID); err!=nil{
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("GetOrderItemList")
		return nil, err
	}

	resOrderItems := []model.OrderItemResponse{}
	for _, v := range(orderItems){
		res := model.OrderItemResponse{
			ID: v.ID,
			OrganizationID: v.OrganizationID,
			ProductID: v.ProductID,
			Product: v.Product,
			ProductQuantity: v.ProductQuantity,
			UnitProductPrice: v.UnitProductPrice,
			Tax: v.Tax,
			TotalProductPrice: v.TotalProductPrice,
		}
		resOrderItems = append(resOrderItems, res)
	}

	return resOrderItems, nil
}

func (ou *orderItemUsecase) CreateOrderItem(orderItem model.OrderItem) (model.OrderItemResponse, error){

	if err := ou.ov.OrderItemValidate(orderItem); err!=nil{
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("CreateOrderItem")
		return model.OrderItemResponse{}, err
	}

	if err := ou.or.CreateOrderItem(&orderItem); err!=nil{
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("CreateOrderItem")
		return model.OrderItemResponse{}, err
	}

		resOrderItem := model.OrderItemResponse{
			ID: orderItem.ID,
			OrganizationID: orderItem.OrganizationID,
			OrderID: orderItem.OrderID,
			ProductID: orderItem.ProductID,
			ProductQuantity: orderItem.ProductQuantity,
			UnitProductPrice: orderItem.UnitProductPrice,
			Tax: orderItem.Tax,
			TotalProductPrice: orderItem.TotalProductPrice,
		}
	return resOrderItem, nil
}

func (ou *orderItemUsecase) UpdateOrderItem(orderItem model.OrderItem, id uint, organizationID uint) (model.OrderItemResponse, error){

	if err := ou.ov.OrderItemValidate(orderItem); err!=nil{
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("UpdateOrderItem")
		return model.OrderItemResponse{}, err
	}

	if err := ou.or.UpdateOrderItem(&orderItem, id, organizationID); err!=nil{
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("UpdateOrderItem")
		return model.OrderItemResponse{}, err
	}

	resOrderItem := model.OrderItemResponse{
		ID: orderItem.ID,
		OrganizationID: orderItem.OrganizationID,
		ProductID: orderItem.ProductID,
		ProductQuantity: orderItem.ProductQuantity,
		UnitProductPrice: orderItem.UnitProductPrice,
		Tax: orderItem.Tax,
		TotalProductPrice: orderItem.TotalProductPrice,
	}

	return resOrderItem, nil
} 

func (ou *orderItemUsecase) DeleteOrderItem(orderItem model.OrderItem, id uint, organizationID uint) error{

	if err := ou.or.DeleteOrderItem(&orderItem, id, organizationID); err!=nil{
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("DeleteOrderItem")
		return err
	}
	return nil
}

func (ou *orderItemUsecase) InvoiceCustomer(organizationID uint, orderID uint, dateFrom string, dateTo string)([]model.OrderItemResponse, error){

	orderItems := []model.OrderItem{}

	if err := ou.or.InvoiceCustomer(&orderItems, organizationID, orderID, dateFrom, dateTo); err!=nil{
		return nil, err
	}

	resOrderItem := []model.OrderItemResponse{}
	return resOrderItem, nil
}