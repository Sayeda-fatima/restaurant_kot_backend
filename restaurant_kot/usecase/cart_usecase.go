package usecase

import (
	"math"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/common"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/model"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/repository"
	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/validator"
	"gorm.io/gorm"
)

type (
	CartUsecase interface {
		GetCartList(organizationID uint, restaurantID uint) ([]model.CartResponse, error)
		CreateCart(cart model.Cart) (model.CartResponse, error)
		UpdateCart(cart model.Cart, id uint, organizationID uint, restaurantID uint) (model.CartResponse, error)
		UpdateCartStatus(cart model.Cart, id uint, organizationID uint, restaurantID uint, status string) (model.CartResponse, error)
		DeleteCart(id uint, organizationID uint, restaurantID uint) error
		SendCartToKitchen(id uint, organizationID uint, restaurantID uint) (model.CartResponse, error)
		CheckCartActive(organizationID uint, restaurantID uint, tableID uint) (model.CartResponse, error)
	}

	cartUsecase struct{
		cr repository.CartRepository
		cv validator.CartValidator
		db *gorm.DB
		ci CartItemUsecase
	}
)

func NewCartUsecase(cr repository.CartRepository, cv validator.CartValidator, db *gorm.DB, ci CartItemUsecase) CartUsecase{
	return &cartUsecase{cr, cv, db, ci}
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

func (cu *cartUsecase) SendCartToKitchen(id uint, organizationID uint, restaurantID uint) (model.CartResponse, error){

	// begin transaction
	tx := cu.db.Begin()
	defer func(){
		if r := recover(); r != nil{
			tx.Rollback()
		}
	}()

	cart := model.Cart{}
	if err := cu.cr.GetCart(&cart, id, organizationID, restaurantID); err != nil{
		tx.Rollback()
		return model.CartResponse{}, err
	}

	// change cart items status to sent to kitchen
	for i, cartItem := range(cart.CartItems){
		_, err := cu.ci.UpdateCartItemStatus(cartItem, cartItem.ID, cartItem.CartID, restaurantID, organizationID, "sent_to_kitchen")
		// return cart items with updated status
		cart.CartItems[i].ItemStatus = "sent_to_kitchen"
		if err != nil{
			tx.Rollback()
			return model.CartResponse{}, err
		}

		// update product quantity in product table		
		for _, v := range(cartItem.MenuItem.Recipe.RecipeProducts){

			quantityAfterUpdate := v.Product.Quantity - int(math.Ceil((float64(v.Quantity)/float64(cartItem.MenuItem.Recipe.Serving))*float64(cartItem.MenuItem.Serving)*float64(cartItem.ItemQuantity)))
			common.Logger.LogInfo().Msgf("quantity: %d", quantityAfterUpdate)
			
			result := tx.Model(&model.Product{}).Select("quantity", "inventory_value").Where("id=?", v.ProductID).Updates(map[string]interface{}{"quantity": quantityAfterUpdate, "inventory_value": quantityAfterUpdate*v.Product.UnitCost}) 
			if err := result.Error; err != nil{
				tx.Rollback()
				return model.CartResponse{}, err
			}
		}
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
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

func (cu *cartUsecase) CheckCartActive(organizationID uint, restaurantID uint, tableID uint) (model.CartResponse, error){

	cart := model.Cart{}

	if err := cu.cr.CheckCartActive(&cart, organizationID, restaurantID, tableID); err != nil{
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