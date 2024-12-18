package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/repository"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/validator"
)

type (
	OrderItemUsecase interface {
		GetOrderItemList(organizationID uint, restaurantID uint, orderID uint) ([]model.OrderItemResponse, error)
		CreateOrderItem(orderItem model.OrderItem) (model.OrderItemResponse, error)
		UpdateOrderItem(orderItem model.OrderItem, id uint, orderID uint, organizationID uint, restaurantID uint) (model.OrderItemResponse, error)
		DeleteOrderItem(orderItem model.OrderItem, id uint, orderID uint, organizationID uint, restaurantID uint) error
		MostOrderedItems(organizationID uint, restaurantID uint, dateFrom string, dateTo string) ([]map[string]interface{}, error)
		DailySaleByItem(organizationID uint, restaurantID uint, dateFrom string, dateTo string, page int) ([]map[string]interface{}, error)
	}

	orderItemUsecase struct{
		or repository.OrderItemRepository
		ov validator.OrderItemValidator
	}
)

func NewOrderItemUsecase(or repository.OrderItemRepository, ov validator.OrderItemValidator) OrderItemUsecase{
	return &orderItemUsecase{or, ov}
}

func (ou *orderItemUsecase) GetOrderItemList(organizationID uint, restaurantID uint, orderID uint) ([]model.OrderItemResponse, error){

	orderItems := []model.OrderItem{}

	if err := ou.or.GetOrderItemList(&orderItems, organizationID, restaurantID, orderID); err != nil{
		return nil, err
	}

	resOrderItems := []model.OrderItemResponse{}
	for _, v := range(orderItems){
		res := model.OrderItemResponse{
			ID: v.ID,
			OrganizationID: v.OrganizationID,
			RestaurantID: v.RestaurantID,
			OrderID: v.OrderID,
			MenuItemID: v.MenuItemID,
			ItemQuantity: v.ItemQuantity,
			UnitItemPrice: v.UnitItemPrice,
			TotalItemPrice: v.TotalItemPrice,
		}
		resOrderItems = append(resOrderItems, res)
	}

	return resOrderItems, nil
}

func (ou *orderItemUsecase) CreateOrderItem(orderItem model.OrderItem) (model.OrderItemResponse, error){

	if err := ou.ov.OrderItemValidate(&orderItem); err != nil{
		return model.OrderItemResponse{}, err
	}

	if err := ou.or.CreateOrderItem(&orderItem); err != nil{
		return model.OrderItemResponse{}, err
	}

	resOrderItem := model.OrderItemResponse{
		ID: orderItem.ID,
		OrganizationID: orderItem.OrganizationID,
		RestaurantID: orderItem.RestaurantID,
		OrderID: orderItem.OrderID,
		MenuItemID: orderItem.MenuItemID,
		ItemQuantity: orderItem.ItemQuantity,
		UnitItemPrice: orderItem.UnitItemPrice,
		TotalItemPrice: orderItem.TotalItemPrice,
	}
	return resOrderItem, nil
}

func (ou *orderItemUsecase) UpdateOrderItem(orderItem model.OrderItem, id uint, orderID uint, organizationID uint, restaurantID uint) (model.OrderItemResponse, error){

	if err := ou.ov.OrderItemValidate(&orderItem); err != nil{
		return model.OrderItemResponse{}, err
	}

	if err := ou.or.UpdateOrderItem(&orderItem, id, orderID, organizationID, restaurantID); err != nil{
		return model.OrderItemResponse{}, err
	}

	resOrderItem := model.OrderItemResponse{
		ID: orderItem.ID,
		OrganizationID: orderItem.OrganizationID,
		RestaurantID: orderItem.RestaurantID,
		OrderID: orderItem.OrderID,
		MenuItemID: orderItem.MenuItemID,
		ItemQuantity: orderItem.ItemQuantity,
		UnitItemPrice: orderItem.UnitItemPrice,
		TotalItemPrice: orderItem.TotalItemPrice,
	}
	return resOrderItem, nil
}

func (ou *orderItemUsecase) DeleteOrderItem(orderItem model.OrderItem, id uint, orderID uint, organizationID uint, restaurantID uint) error{

	if err := ou.or.DeleteOrderItem(&orderItem, id, orderID, organizationID, restaurantID); err != nil{
		return err
	}

	return nil
}

func (ou *orderItemUsecase) MostOrderedItems(organizationID uint, restaurantID uint, dateFrom string, dateTo string) ([]map[string]interface{}, error){

	var result []map[string]interface{}

	if err := ou.or.MostOrderedItems(&result, organizationID, restaurantID, dateFrom, dateTo); err != nil{
		return nil, err
	}

	return result, nil
}

func (ou *orderItemUsecase) DailySaleByItem(organizationID uint, restaurantID uint, dateFrom string, dateTo string, page int) ([]map[string]interface{}, error){

	var result []map[string]interface{}

	if err := ou.or.DailySaleByItem(&result, organizationID, restaurantID, dateFrom, dateTo, page); err != nil{
		return nil, err
	}

	return result, nil
}