package usecase

import (
	"fmt"

	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/repository"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/validator"
	"gorm.io/gorm"
)

type (
	OrderUsecase interface {
		GetOrderList(organizationID uint, restaurantID uint) ([]model.OrderResponse, error)
		CreateOrder(order model.Order) (model.OrderResponse, error)
		UpdateOrder(order model.Order, id uint, organizationID uint, restaurantID uint) (model.OrderResponse, error)
		DeleteOrder(order model.Order, id uint, organizationID uint, restaurantID uint) error
		Checkout(order model.Order, organizationID uint, restaurantID uint, cartID uint) (model.OrderResponse, error)
	}

	orderUsecase struct{
		or repository.OrderRepository
		ov validator.OrderValidator
		db *gorm.DB
		cr CartUsecase
	}
)

func NewOrderUsecase(or repository.OrderRepository, ov validator.OrderValidator, db *gorm.DB, cr CartUsecase) OrderUsecase{
	return &orderUsecase{or, ov, db, cr}
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

func (ou *orderUsecase) Checkout(order model.Order, organizationID uint, restaurantID uint, cartID uint) (model.OrderResponse, error){

	cart, err := ou.cr.GetCart(organizationID, restaurantID, cartID)

	if err != nil || len(cart.CartItems) == 0{
		return model.OrderResponse{}, fmt.Errorf("cart is empty or invalid")
	}

	// begin transaction
	tx := ou.db.Begin()
	defer func(){
		if r := recover(); r != nil{
			tx.Rollback()
		}
	}()

	order.OrganizationID = cart.OrganizationID
	order.RestaurantID = cart.RestaurantID
	order.OrderType = cart.CartType
	order.TableID = cart.TableID
	order.TotalPrice = order.TotalItemPrice + order.Tax + order.ServiceCharge + order.Tip
	if err := ou.or.CreateOrder(&order); err != nil{
		tx.Rollback()
		return model.OrderResponse{}, err
	}

	orderItems := []model.OrderItem{}
	for _, cartItem := range(cart.CartItems){

		orderItem := model.OrderItem{
			OrganizationID: organizationID,
			RestaurantID: restaurantID,
			OrderID: order.ID,
			MenuItemID: cartItem.MenuItemID,
			ItemQuantity: cartItem.ItemQuantity,
			UnitItemPrice: cartItem.MenuItem.Price,
			TotalItemPrice: cartItem.MenuItem.Price * cartItem.ItemQuantity,
			ItemStatus: cartItem.ItemStatus,
			Note: cartItem.Note,
		}

		if err := tx.Create(&orderItem).Error; err != nil{
			tx.Rollback()
			return model.OrderResponse{}, err
		}

		orderItems = append(orderItems, orderItem)
		order.TotalItemPrice += orderItem.TotalItemPrice
		order.TotalPrice += orderItem.TotalItemPrice
	}

	// update order with total price
	if err := tx.Model(&order).UpdateColumn("total_item_price", order.TotalItemPrice).UpdateColumn("total_price", order.TotalPrice).Error; err != nil{
		tx.Rollback()
		return model.OrderResponse{}, err
	}

	// delete cart and cart items
	if err := tx.Model(&cart).Delete(&cart).Error; err != nil{
		tx.Rollback()
		return model.OrderResponse{}, err
	}

	// commit transaction
	if err := tx.Commit().Error; err != nil {
		return model.OrderResponse{}, err
	}

	resOrder := model.OrderResponse{
		ID: order.ID,
		OrganizationID: order.OrganizationID,
		RestaurantID: order.RestaurantID,
		TableID: order.TableID,
		TotalItemPrice: order.TotalItemPrice,
		Tax: order.Tax,
		ServiceCharge: order.ServiceCharge,
		Tip: order.Tip,
		TotalPrice: order.TotalPrice,
		OrderType: order.OrderType,
		OrderStatus: order.OrderStatus,
		OrderItems: orderItems,
	}

	return resOrder, nil
}