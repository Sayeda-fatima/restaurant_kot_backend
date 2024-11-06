package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_go/common"
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/NazishAhsan/easy_busy_book_go/repository"
	"github.com/NazishAhsan/easy_busy_book_go/validator"
)

type (
	OrderUsecase interface {
		GetOrderList(organizationID uint) ([]model.OrderResponse, error)
		CreateOrder(order model.Order) (model.OrderResponse, error)
		UpdateOrder(order model.Order, id uint) (model.OrderResponse, error)
		DeleteOrder(order model.Order, id uint) error
	}

	orderUsecase struct {
		or repository.OrderRepository
		ov validator.OrderValidator
	}
)

func NewOrderUsecase(or repository.OrderRepository, ov validator.OrderValidator) OrderUsecase {
	return &orderUsecase{or, ov}
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
			ID:                     v.ID,
			OrganizationID:         v.OrganizationID,
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
		ID:                     order.ID,
		OrganizationID:         order.OrganizationID,
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
		//CartID:                 order.CartID,
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

// func (ou *orderUsecase) Checkout (cartID uint) (model.OrderResponse, error){

// 	if err := ou.or.CreateOrder(&model.Order{}); err !=nil{
// 		return model.OrderResponse{}, err
// 	}
// 	return model.OrderResponse{}, nil
// }
