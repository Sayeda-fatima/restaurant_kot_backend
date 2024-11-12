package usecase

import (
	"fmt"

	"github.com/NazishAhsan/easy_busy_book_go/common"
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/NazishAhsan/easy_busy_book_go/repository"
	"github.com/NazishAhsan/easy_busy_book_go/validator"
	"gorm.io/gorm"
)

type (
	OrderUsecase interface {
		GetOrderList(organizationID uint) ([]model.OrderResponse, error)
		CreateOrder(order model.Order) (model.OrderResponse, error)
		UpdateOrder(order model.Order, id uint) (model.OrderResponse, error)
		DeleteOrder(order model.Order, id uint) error
		Checkout(order model.Order, organizationID uint, cartID uint) (model.OrderResponse, error)
		InvoiceReportCustomer(organizationID uint, customerID uint, dateFrom string, dateTo string) ([]model.OrderResponse, error)
	}

	orderUsecase struct {
		or repository.OrderRepository
		ov validator.OrderValidator
		db *gorm.DB
		cr CartItemUsecase
		cu CartUsecase
	}
)

func NewOrderUsecase(or repository.OrderRepository, ov validator.OrderValidator, db *gorm.DB, cr CartItemUsecase, cu CartUsecase) OrderUsecase {
	return &orderUsecase{or, ov, db, cr, cu}
}

func (ou *orderUsecase) GetOrderList(organizationID uint) ([]model.OrderResponse, error) {

	orders := []model.Order{}

	if err := ou.or.GetOrderList(&orders, organizationID); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("GetOrderList")
		return nil, err
	}

	resOrders := []model.OrderResponse{}
	for _, v := range orders {
		res := model.OrderResponse{
			ID:             v.ID,
			OrganizationID: v.OrganizationID,
			//CartID:                 v.CartID,
			CustomerID:             v.CustomerID,
			TotalPrice:             v.TotalPrice,
			CustomerBillingAddress: v.CustomerBillingAddress,
			ModeOfPayment:          v.ModeOfPayment,
		}
		resOrders = append(resOrders, res)
	}
	return resOrders, nil
}

func (ou *orderUsecase) CreateOrder(order model.Order) (model.OrderResponse, error) {

	if err := ou.ov.OrderValidate(order); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("CreateOrder")
		return model.OrderResponse{}, err
	}

	if err := ou.or.CreateOrder(&order); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("CreateOrder")
		return model.OrderResponse{}, err
	}

	resOrder := model.OrderResponse{
		ID:             order.ID,
		OrganizationID: order.OrganizationID,
		//CartID:                 order.CartID,
		CustomerID:             order.CustomerID,
		TotalPrice:             order.TotalPrice,
		CustomerBillingAddress: order.CustomerBillingAddress,
		ModeOfPayment:          order.ModeOfPayment,
	}

	return resOrder, nil
}

func (ou *orderUsecase) UpdateOrder(order model.Order, id uint) (model.OrderResponse, error) {

	if err := ou.ov.OrderValidate(order); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("UpdateOrder")
		return model.OrderResponse{}, err
	}

	if err := ou.or.UpdateOrder(&order, id); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("UpdateOrder")
		return model.OrderResponse{}, err
	}

	resOrder := model.OrderResponse{
		ID:                     order.ID,
		OrganizationID:         order.OrganizationID,
		CustomerID:             order.CustomerID,
		TotalPrice:             order.TotalPrice,
		CustomerBillingAddress: order.CustomerBillingAddress,
		ModeOfPayment:          order.ModeOfPayment,
	}

	return resOrder, nil
}

func (ou *orderUsecase) DeleteOrder(order model.Order, id uint) error {

	if err := ou.or.DeleteOrder(&order, id); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("DeleteOrder")
		return err
	}
	return nil
}

func (ou *orderUsecase) Checkout(order model.Order, organizationID uint, cartID uint) (model.OrderResponse, error) {

	tx := ou.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	cartItems, err := ou.cr.GetCartItemList(organizationID, cartID)
	if err != nil || len(cartItems) == 0 {
		tx.Rollback()
		return model.OrderResponse{}, fmt.Errorf("cart is empty or invalid")
	}

	if err := ou.or.CreateOrder(&order); err != nil {
		tx.Rollback()
		return model.OrderResponse{}, err
	}

	orderItems := []model.OrderItem{}
	for _, cartItem := range cartItems {
		orderItem := model.OrderItem{
			OrderID:           order.ID,
			OrganizationID:    cartItem.OrganizationID,
			ProductID:         cartItem.ProductID,
			ProductQuantity:   cartItem.ProductQuantity,
			UnitProductPrice:  cartItem.Product.Mrp,
			Tax:               cartItem.Product.TaxAmount,
			TotalProductPrice: cartItem.ProductQuantity * cartItem.Product.Mrp,
		}
		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			return model.OrderResponse{}, err
		}
		orderItems = append(orderItems, orderItem)
		order.TotalPrice += orderItem.TotalProductPrice
	}

	// Update order with total price
	if err := tx.Model(&order).Update("total_price", order.TotalPrice).Error; err != nil {
		tx.Rollback()
		return model.OrderResponse{}, err
	}
	
	// delete cart and cart items
	if err := ou.cu.DeleteCart(cartID); err != nil{
		tx.Rollback()
		return model.OrderResponse{}, err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return model.OrderResponse{}, err
	}

	resOrder := model.OrderResponse{
		ID:                     order.ID,
		OrganizationID:         order.OrganizationID,
		CustomerID:             order.CustomerID,
		TotalPrice:             order.TotalPrice,
		CustomerBillingAddress: order.CustomerBillingAddress,
		ModeOfPayment:          order.ModeOfPayment,
		OrderItems:             orderItems,
	}

	return resOrder, nil
}

func (ou *orderUsecase) InvoiceReportCustomer(organizationID uint, customerID uint, dateFrom string, dateTo string) ([]model.OrderResponse, error) {

	orders := []model.Order{}

	if err := ou.or.InvoiceReportCustomer(&orders, organizationID, customerID, dateFrom, dateTo); err != nil {
		return nil, err
	}

	resOrder := []model.OrderResponse{}
	for _, v := range orders {
		res := model.OrderResponse{
			ID:                     v.ID,
			OrganizationID:         v.OrganizationID,
			CustomerID:             v.CustomerID,
			TotalPrice:             v.TotalPrice,
			CustomerBillingAddress: v.CustomerBillingAddress,
			ModeOfPayment:          v.ModeOfPayment,
		}
		resOrder = append(resOrder, res)
	}
	return resOrder, nil
}
