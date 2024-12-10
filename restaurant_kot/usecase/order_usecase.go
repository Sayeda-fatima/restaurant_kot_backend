package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/repository"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/validator"
)

type (
	OrderUsecase interface {
		GetOrderList(organizationID uint, restaurantID uint) ([]model.OrderResponse, error)
		CreateOrder(order model.Order) (model.OrderResponse, error)
		UpdateOrder(order model.Order, id uint, organizationID uint, restaurantID uint) (model.OrderResponse, error)
		DeleteOrder(order model.Order, id uint, organizationID uint, restaurantID uint) error
	}

	orderUsecase struct{
		or repository.OrderRepository
		ov validator.OrderValidator
	}
)

func NewOrderUsecase(or repository.OrderRepository, ov validator.OrderValidator) OrderUsecase{
	return &orderUsecase{or, ov}
}

func (ou *orderUsecase) GetOrderList(organizationID uint, restaurantID uint) ([]model.OrderResponse, error){

	orders := []model.Order{}
	if err := ou.or.GetOrderList(&orders, organizationID, restaurantID); err != nil{
		return nil, err
	}

	resOrders := []model.OrderResponse{}
	for _, v := range(orders){
		res := model.OrderResponse{
			ID: v.ID,
			OrganizationID: v.OrganizationID,
			RestaurantID: v.RestaurantID,
			TableID: v.TableID,
			TotalItemPrice: v.TotalItemPrice,
			Tax: v.Tax,
			ServiceCharge: v.ServiceCharge,
			Tip: v.Tip,
			TotalPrice: v.TotalPrice,
			OrderItems: v.OrderItems,
			OrderType: v.OrderType,
			OrderStatus: v.OrderStatus,
		}
		resOrders = append(resOrders, res)
	}
	return resOrders, nil
}

func (ou *orderUsecase) CreateOrder(order model.Order) (model.OrderResponse, error){

	if err := ou.ov.OrderValidate(&order); err != nil{
		return model.OrderResponse{}, err
	}

	if err := ou.or.CreateOrder(&order); err != nil{
		return model.OrderResponse{}, err
	}

	resOrder := model.OrderResponse{
		ID: order.ID,
		OrganizationID: order.OrganizationID,
		RestaurantID: order.RestaurantID,
		TableID: order.TableID,
		TotalItemPrice: order.TotalItemPrice,
		ServiceCharge: order.ServiceCharge,
		Tip: order.Tip,
		TotalPrice: order.TotalPrice,
		OrderItems: order.OrderItems,
		OrderType: order.OrderType,
		OrderStatus: order.OrderStatus,
	}
	return resOrder, nil
}

func (ou *orderUsecase) UpdateOrder(order model.Order, id uint, organizationID uint, restaurantID uint) (model.OrderResponse, error){

	if err := ou.ov.OrderValidate(&order); err != nil{
		return model.OrderResponse{}, err
	}

	if err := ou.or.UpdateOrder(&order, id, organizationID, restaurantID); err != nil{
		return model.OrderResponse{}, err
	}

	resOrder := model.OrderResponse{
		ID: order.ID,
		OrganizationID: order.OrganizationID,
		RestaurantID: order.RestaurantID,
		TableID: order.TableID,
		TotalItemPrice: order.TotalItemPrice,
		ServiceCharge: order.ServiceCharge,
		Tip: order.Tip,
		TotalPrice: order.TotalPrice,
		OrderItems: order.OrderItems,
		OrderType: order.OrderType,
		OrderStatus: order.OrderStatus,
	}

	return resOrder, nil
}

func (ou *orderUsecase) DeleteOrder(order model.Order, id uint, organizationID uint, restaurantID uint) error{

	if err := ou.or.DeleteOrder(&order, id, organizationID, restaurantID); err != nil{
		return err
	}

	return nil
}