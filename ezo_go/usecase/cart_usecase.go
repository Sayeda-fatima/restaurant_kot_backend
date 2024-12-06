package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_go/common"
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/NazishAhsan/easy_busy_book_go/repository"
	"github.com/NazishAhsan/easy_busy_book_go/validator"
)

type (
	CartUsecase interface {
		GetCartList(organizationID uint) ([]model.CartResponse, error)
		CreateCart(cart model.Cart) (model.CartResponse, error)
		UpdateCart(cart model.Cart, id uint, organizationID uint) (model.CartResponse, error)
		DeleteCart(id uint, organizationID uint) error
	}

	cartUsecase struct {
		cr repository.CartRepository
		cv validator.CartValidator
	}
)

func NewCartUsecase(cr repository.CartRepository, cv validator.CartValidator) CartUsecase {
	return &cartUsecase{cr, cv}
}

func (cu *cartUsecase) GetCartList(organizationID uint) ([]model.CartResponse, error) {

	carts := []model.Cart{}
	if err := cu.cr.GetCartList(&carts, organizationID); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("GetCartList")
		return nil, err
	}

	resCart := []model.CartResponse{}
	for _, v := range carts {
		res := model.CartResponse{
			ID:             v.ID,
			OrganizationID: v.OrganizationID,
			CustomerID:     v.CustomerID,
			TotalQuantity:  v.TotalQuantity,
		}

		resCart = append(resCart, res)
	}

	return resCart, nil
}

func (cu *cartUsecase) CreateCart(cart model.Cart) (model.CartResponse, error) {

	if err := cu.cv.CartValidate(cart); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("CreateCart")
		return model.CartResponse{}, err
	}

	// var cartItems []CartItem
	// for _, cartItem := range cartItems{
	// 	cart.TotalQuantity += cartItem.ProductQuantity
	// }
	// Assigning organizationID to cart items
	for i := range cart.CartItems {
		cart.CartItems[i].OrganizationID = cart.OrganizationID
		// getting total quantity of product
		cart.TotalQuantity += cart.CartItems[i].ProductQuantity 
	}

	if err := cu.cr.CreateCart(&cart); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("CreateCart")
		return model.CartResponse{}, err
	}

	resCart := model.CartResponse{
		ID:             cart.ID,
		OrganizationID: cart.OrganizationID,
		CustomerID:     cart.CustomerID,
		CartItems:      cart.CartItems,
		TotalQuantity:  cart.TotalQuantity,
	}

	return resCart, nil
}

func (cu *cartUsecase) UpdateCart(cart model.Cart, id uint, organizationID uint) (model.CartResponse, error) {

	if err := cu.cv.CartValidate(cart); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("UpdateCart")
		return model.CartResponse{}, err
	}

	// Assigning organizationID to cart items
	for i := range cart.CartItems {
		cart.CartItems[i].OrganizationID = cart.OrganizationID
		// getting total quantity of product
		//cart.TotalQuantity += cart.CartItems[i].ProductQuantity 
	}

	if err := cu.cr.UpdateCart(&cart, id, organizationID); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("UpdateCart")
		return model.CartResponse{}, err
	}

	resCart := model.CartResponse{
		ID:             cart.ID,
		OrganizationID: cart.OrganizationID,
		CustomerID:     cart.CustomerID,
		CartItems:      cart.CartItems,
		TotalQuantity:  cart.TotalQuantity,
	}

	return resCart, nil
}

func (cu *cartUsecase) DeleteCart(id uint, organizationID uint) error {

	if err := cu.cr.DeleteCart(id, organizationID); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("DeleteCart")
		return err
	}
	return nil
}
