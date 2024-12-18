package usecase

import (
	"fmt"
	"strconv"

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
		TotalSales(organizationID uint, restaurantID uint, dateFrom string, dateTo string) (map[string]interface{}, error)
		TotalSalesByOrderType(organizationID uint, restaurantID uint, dateFrom string, dateTo string) ([]map[string]interface{}, error)
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
			TotalItemPrice: strconv.FormatFloat(float64(v.TotalItemPrice)/100, 'f', -1, 64),
			Tax: strconv.FormatFloat(float64(v.Tax)/100, 'f', -1, 64),
			ServiceCharge: strconv.FormatFloat(float64(v.ServiceCharge)/100, 'f', -1, 64),
			Tip: strconv.FormatFloat(float64(v.Tip)/100, 'f', -1, 64),
			TotalPrice: strconv.FormatFloat(float64(v.TotalPrice)/100, 'f', -1, 64),
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
		TotalItemPrice: strconv.FormatFloat(float64(order.TotalItemPrice)/100, 'f', -1, 64),
		Tax: strconv.FormatFloat(float64(order.Tax)/100, 'f', -1, 64),
		ServiceCharge: strconv.FormatFloat(float64(order.ServiceCharge)/100, 'f', -1, 64),
		Tip: strconv.FormatFloat(float64(order.Tip)/100, 'f', -1, 64),
		TotalPrice: strconv.FormatFloat(float64(order.TotalPrice)/100, 'f', -1, 64),
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
		TotalItemPrice: strconv.FormatFloat(float64(order.TotalItemPrice)/100, 'f', -1, 64),
		Tax: strconv.FormatFloat(float64(order.Tax)/100, 'f', -1, 64),
		ServiceCharge: strconv.FormatFloat(float64(order.ServiceCharge)/100, 'f', -1, 64),
		Tip: strconv.FormatFloat(float64(order.Tip)/100, 'f', -1, 64),
		TotalPrice: strconv.FormatFloat(float64(order.TotalPrice)/100, 'f', -1, 64),
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

	if err != nil || len(cart.CartItems) == 0 || cart.CartStatus != "ready_for_checkout"{
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
		TotalItemPrice: strconv.FormatFloat(float64(order.TotalItemPrice)/100, 'f', -1, 64),
		Tax: strconv.FormatFloat(float64(order.Tax)/100, 'f', -1, 64),
		ServiceCharge: strconv.FormatFloat(float64(order.ServiceCharge)/100, 'f', -1, 64),
		Tip: strconv.FormatFloat(float64(order.Tip)/100, 'f', -1, 64),
		TotalPrice: strconv.FormatFloat(float64(order.TotalPrice)/100, 'f', -1, 64),
		OrderType: order.OrderType,
		OrderStatus: order.OrderStatus,
		OrderItems: orderItems,
	}

	return resOrder, nil
}

func (ou *orderUsecase) TotalSales(organizationID uint, restaurantID uint, dateFrom string, dateTo string) (map[string]interface{}, error){

	var result map[string]interface{}
	if err := ou.or.TotalSales(&result, organizationID, restaurantID, dateFrom, dateTo); err != nil{
		return nil, err
	}
	return result, nil
}

func (ou *orderUsecase) TotalSalesByOrderType(organizationID uint, restaurantID uint, dateFrom string, dateTo string) ([]map[string]interface{}, error){

	var result []map[string]interface{}
	if err := ou.or.TotalSalesByOrderType(&result, organizationID, restaurantID, dateFrom, dateTo); err != nil{
		return nil, err
	}

	return result, nil
}