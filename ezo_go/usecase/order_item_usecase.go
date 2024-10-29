package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/NazishAhsan/easy_busy_book_go/repository"
	"github.com/NazishAhsan/easy_busy_book_go/validator"
)

type (
	OrderItemUsecase interface {
		GetOrderItemList(organizationID uint, orderID uint) ([]model.OrderItemResponse, error)
		CreateOrderItem(orderItem model.OrderItem) (model.OrderItemResponse, error)
		UpdateOrderItem(orderItem model.OrderItem, id uint) (model.OrderItemResponse, error)
		DeleteOrderItem(orderItem model.OrderItem, id uint) error
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
		return nil, err
	}

	resOrderItems := []model.OrderItemResponse{}
	for _, v := range(orderItems){
		res := model.OrderItemResponse{
			ID: v.ID,
			OrganizationID: v.OrganizationID,
			ProductID: v.ProductID,
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
		return model.OrderItemResponse{}, err
	}

	if err := ou.or.CreateOrderItem(&orderItem); err!=nil{
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

func (ou *orderItemUsecase) UpdateOrderItem(orderItem model.OrderItem, id uint) (model.OrderItemResponse, error){

	if err := ou.ov.OrderItemValidate(orderItem); err!=nil{
		return model.OrderItemResponse{}, err
	}

	if err := ou.or.UpdateOrderItem(&orderItem, id); err!=nil{
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

func (ou *orderItemUsecase) DeleteOrderItem(orderItem model.OrderItem, id uint) error{

	if err := ou.or.DeleteOrderItem(&orderItem, id); err!=nil{
		return err
	}
	return nil
}