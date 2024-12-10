package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/repository"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/validator"
)

type (
	CartItemUsecase interface {
		GetCartItemList(cartID uint, restaurantID uint, organizationID uint) ([]model.CartItemResponse, error)
		CreateCartItem(cartItem model.CartItem) (model.CartItemResponse, error)
		UpdateCartItem(cartItem model.CartItem, id uint, cartID uint, restaurantID uint, organizationID uint) (model.CartItemResponse, error)
		DeleteCartItem(cartItem model.CartItem, id uint, cartID uint, restaurantID uint, organizationID uint) error
	}

	cartItemUsecase struct{
		cr repository.CartItemRepository
		cv validator.CartItemValidator
	}
)

func NewCartItemUsecase(cr repository.CartItemRepository, cv validator.CartItemValidator) CartItemUsecase{
	return &cartItemUsecase{cr, cv}
}

func (cu *cartItemUsecase) GetCartItemList(cartID uint, restaurantID uint, organizationID uint) ([]model.CartItemResponse, error){

	cartItems := []model.CartItem{}
	if err := cu.cr.GetCartItemList(&cartItems, cartID, restaurantID, organizationID); err != nil{
		return nil, err
	}

	resCartItems := []model.CartItemResponse{}
	for _, v := range(cartItems){
		res := model.CartItemResponse{
			ID: v.ID,
			CartID: v.CartID,
			OrganizationID: v.OrganizationID,
			RestaurantID: v.RestaurantID,
			MenuItemID: v.MenuItemID,
			MenuItem: v.MenuItem,
			ItemQuantity: v.ItemQuantity,
		}
		resCartItems = append(resCartItems, res)
	}
	return resCartItems, nil
}

func (cu *cartItemUsecase) CreateCartItem(cartItem model.CartItem) (model.CartItemResponse, error){

	if err := cu.cv.CartItemValidate(&cartItem); err != nil{
		return model.CartItemResponse{}, err
	}

	if err := cu.cr.CreateCartitem(&cartItem); err != nil{
		return model.CartItemResponse{}, err
	}

	resCartItem := model.CartItemResponse{
		ID: cartItem.ID,
		OrganizationID: cartItem.OrganizationID,
		RestaurantID: cartItem.RestaurantID,
		CartID: cartItem.CartID,
		MenuItemID: cartItem.MenuItemID,
		MenuItem: cartItem.MenuItem,
		ItemQuantity: cartItem.ItemQuantity,
	}

	return resCartItem, nil
}

func (cu *cartItemUsecase) UpdateCartItem(cartItem model.CartItem, id uint, cartID uint, restaurantID uint, organizationID uint) (model.CartItemResponse, error){

	if err := cu.cv.CartItemValidate(&cartItem); err != nil{
		return model.CartItemResponse{}, err
	}

	if err := cu.cr.UpdateCartitem(&cartItem, id, cartID, restaurantID, organizationID); err != nil{
		return model.CartItemResponse{}, err
	}

	resCartItem := model.CartItemResponse{
		ID: cartItem.ID,
		OrganizationID: cartItem.OrganizationID,
		RestaurantID: cartItem.RestaurantID,
		CartID: cartItem.CartID,
		MenuItemID: cartItem.MenuItemID,
		MenuItem: cartItem.MenuItem,
		ItemQuantity: cartItem.ItemQuantity,
	}
	return resCartItem, nil
}

func (cu *cartItemUsecase) DeleteCartItem(cartItem model.CartItem, id uint, cartID uint, restaurantID uint, organizationID uint) error{

	if err := cu.cr.DeleteCartItem(&cartItem, id, cartID, restaurantID, organizationID); err != nil{
		return err
	}
	return nil
}