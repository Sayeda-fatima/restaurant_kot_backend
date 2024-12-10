package usecase

import (
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/repository"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/validator"
)

type (
	CartUsecase interface {
		GetCartList(organizationID uint, restaurantID uint) ([]model.CartResponse, error)
		CreateCart(cart model.Cart) (model.CartResponse, error)
		UpdateCart(cart model.Cart, id uint, organizationID uint, restaurantID uint) (model.CartResponse, error)
		UpdateCartStatus(cart model.Cart, id uint, organizationID uint, restaurantID uint, status string) (model.CartResponse, error)
		DeleteCart(id uint, organizationID uint, restaurantID uint) error
	}

	cartUsecase struct{
		cr repository.CartRepository
		cv validator.CartValidator
	}
)

func NewCartUsecase(cr repository.CartRepository, cv validator.CartValidator) CartUsecase{
	return &cartUsecase{cr, cv}
}

func (cu *cartUsecase) GetCartList(organizationID uint, restaurantID uint) ([]model.CartResponse, error){

	carts := []model.Cart{}

	if err := cu.cr.GetCartList(&carts, organizationID, restaurantID); err != nil{
		return nil, err
	}

	resCarts := []model.CartResponse{}
	for _, v := range(carts){
		res := model.CartResponse{
			ID: v.ID,
			OrganizationID: v.OrganizationID,
			RestaurantID: v.RestaurantID,
			TableID: v.TableID,
			TotalQuantity: v.TotalQuantity,
			CartType: v.CartType,
			CartStatus: v.CartStatus,
			CartItems: v.CartItems,
		}
		resCarts = append(resCarts, res)
	}
	return resCarts, nil
}

func (cu *cartUsecase) CreateCart(cart model.Cart) (model.CartResponse, error){

	if err := cu.cv.CartValidate(&cart); err != nil{
		return model.CartResponse{}, err
	}

	// Assigning organization_id and restaurant_id to cart_items
	for i := range(cart.CartItems){
		cart.CartItems[i].OrganizationID = cart.OrganizationID
		cart.CartItems[i].RestaurantID = cart.RestaurantID
		// getting total quantity of product
		cart.TotalQuantity += cart.CartItems[i].ItemQuantity 
	}

	if err := cu.cr.CreateCart(&cart); err != nil{
		return model.CartResponse{}, err
	}

	resCart := model.CartResponse{
		ID: cart.ID,
		OrganizationID: cart.OrganizationID,
		RestaurantID: cart.RestaurantID,
		TableID: cart.TableID,
		TotalQuantity: cart.TotalQuantity,
		CartType: cart.CartType,
		CartStatus: cart.CartStatus,
		CartItems: cart.CartItems,
	}

	return resCart, nil
}

func (cu *cartUsecase) UpdateCart(cart model.Cart, id uint, organizationID uint, restaurantID uint) (model.CartResponse, error){

	if err := cu.cv.CartValidate(&cart); err != nil{
		return model.CartResponse{}, err
	}

	// Assigning organization_id and restaurant_id to cart_items
	for i := range(cart.CartItems){
		cart.CartItems[i].OrganizationID = cart.OrganizationID
		cart.CartItems[i].RestaurantID = cart.RestaurantID
		// getting total quantity of product
		cart.TotalQuantity += cart.CartItems[i].ItemQuantity 
	}
	
	if err := cu.cr.UpdateCart(&cart, id, organizationID, restaurantID); err != nil{
		return model.CartResponse{}, err
	}

	resCart := model.CartResponse{
		ID: cart.ID,
		OrganizationID: cart.OrganizationID,
		RestaurantID: cart.RestaurantID,
		TableID: cart.TableID,
		TotalQuantity: cart.TotalQuantity,
		CartType: cart.CartType,
		CartStatus: cart.CartStatus,
		CartItems: cart.CartItems,
	}
	return resCart, nil
}

func (cu *cartUsecase) UpdateCartStatus(cart model.Cart, id uint, organizationID uint, restaurantID uint, status string) (model.CartResponse, error){

	if err := cu.cr.UpdateCartStatus(&cart, id, organizationID, restaurantID, status); err != nil{
		return model.CartResponse{}, err
	}

	resCart := model.CartResponse{
		ID: cart.ID,
		OrganizationID: cart.OrganizationID,
		RestaurantID: cart.RestaurantID,
		TableID: cart.TableID,
		TotalQuantity: cart.TotalQuantity,
		CartType: cart.CartType,
		CartStatus: cart.CartStatus,
		CartItems: cart.CartItems,
	}

	return resCart, nil
}

func (cu *cartUsecase) DeleteCart(id uint, organizationID uint, restaurantID uint) error{

	if err := cu.cr.DeleteCart(id, organizationID, restaurantID); err != nil{
		return err
	}
	return nil
}