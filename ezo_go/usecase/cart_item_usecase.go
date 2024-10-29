package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_go/model"
	"github.com/NazishAhsan/easy_busy_book_go/repository"
	"github.com/NazishAhsan/easy_busy_book_go/validator"
)

type (
	CartItemUsecase interface {
		GetCartItemList(organizationID uint, cartID uint) ([]model.CartItemResponse, error)
		CreateCartItem(cartItem model.CartItem) (model.CartItemResponse, error)
		UpdateCartItem(cartItem model.CartItem, id uint) (model.CartItemResponse, error)
		DeleteCartItem(cartItem model.CartItem, id uint) error
	}

	cartItemUsecase struct {
		cr repository.CartItemRepository
		cv validator.CartItemValidator
	}
)

func NewCartItemUsecase(cr repository.CartItemRepository, cv validator.CartItemValidator) CartItemUsecase {
	return &cartItemUsecase{cr, cv}
}

func (cu *cartItemUsecase) GetCartItemList(organizationID uint, cartID uint) ([]model.CartItemResponse, error) {

	cartItems := []model.CartItem{}
	if err := cu.cr.GetCartItemList(&cartItems, organizationID, cartID); err != nil {
		return nil, err
	}

	resCartItems := []model.CartItemResponse{}
	for _, v := range cartItems {
		res := model.CartItemResponse{
			ID:              v.ID,
			OrganizationID:  v.OrganizationID,
			CartID:          v.CartID,
			ProductID:       v.ProductID,
			ProductQuantity: v.ProductQuantity,
		}
		resCartItems = append(resCartItems, res)
	}

	return resCartItems, nil
}

func (cu *cartItemUsecase) CreateCartItem(cartItem model.CartItem) (model.CartItemResponse, error) {

	if err := cu.cv.CartItemValidate(cartItem); err != nil {
		return model.CartItemResponse{}, err
	}

	if err := cu.cr.CreateCartItem(&cartItem); err != nil {
		return model.CartItemResponse{}, err
	}

	resCartitem := model.CartItemResponse{
		ID:              cartItem.ID,
		OrganizationID:  cartItem.OrganizationID,
		CartID:          cartItem.CartID,
		ProductID:       cartItem.ProductID,
		ProductQuantity: cartItem.ProductQuantity,
	}

	return resCartitem, nil
}

func (cu *cartItemUsecase) UpdateCartItem(cartItem model.CartItem, id uint) (model.CartItemResponse, error) {

	if err := cu.cv.CartItemValidate(cartItem); err != nil {
		return model.CartItemResponse{}, err
	}

	if err := cu.cr.UpdateCartItem(&cartItem, id); err != nil {
		return model.CartItemResponse{}, err
	}

	resCartItem := model.CartItemResponse{
		ID:              cartItem.ID,
		OrganizationID:  cartItem.OrganizationID,
		CartID:          cartItem.CartID,
		ProductID:       cartItem.ProductID,
		ProductQuantity: cartItem.ProductQuantity,
	}

	return resCartItem, nil
}

func (cu *cartItemUsecase) DeleteCartItem(cartItem model.CartItem, id uint) error {

	if err := cu.cr.DeleteCartItem(&cartItem, id); err != nil {
		return err
	}

	return nil
}
