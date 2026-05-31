package usecase

import (
	"math"

	"github.com/Sayeda-fatima/restaurant_kot_backend/common"
	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
	"github.com/Sayeda-fatima/restaurant_kot_backend/repository"
	"github.com/Sayeda-fatima/restaurant_kot_backend/validator"
	"gorm.io/gorm"
)

type (
	CartItemUsecase interface {
		GetCartItemList(cartID uint, restaurantID uint, organizationID uint) ([]model.CartItemResponse, error)
		GetPendingCartItemList(cartID uint, restaurantID uint, organizationID uint) ([]model.CartItem, error)
		CreateCartItem(cartItem model.CartItem) (model.CartItemResponse, error)
		UpdateCartItem(cartItem model.CartItem, id uint, cartID uint, restaurantID uint, organizationID uint) (model.CartItemResponse, error)
		UpdateCartItemStatus(cartItem model.CartItem, id uint, cartID uint, restaurantID uint, organizationID uint, status string) (error)
		DeleteCartItem(cartItem model.CartItem, id uint, cartID uint, restaurantID uint, organizationID uint) error
		SendCartItemToKitchen(cartID uint, organizationID uint, restaurantID uint) ([]model.CartItemResponse, error)
	}

	cartItemUsecase struct {
		cr repository.CartItemRepository
		cv validator.CartItemValidator
		db *gorm.DB
	}
)

func NewCartItemUsecase(cr repository.CartItemRepository, cv validator.CartItemValidator, db *gorm.DB) CartItemUsecase {
	return &cartItemUsecase{cr, cv, db}
}

func (cu *cartItemUsecase) GetCartItemList(cartID uint, restaurantID uint, organizationID uint) ([]model.CartItemResponse, error) {

	cartItems := []model.CartItem{}
	if err := cu.cr.GetCartItemList(&cartItems, cartID, restaurantID, organizationID); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("GetCartItemList")
		return nil, err
	}

	resCartItems := []model.CartItemResponse{}
	for _, v := range cartItems {
		res := model.CartItemResponse{
			ID:             v.ID,
			CartID:         v.CartID,
			OrganizationID: v.OrganizationID,
			RestaurantID:   v.RestaurantID,
			MenuItemID:     v.MenuItemID,
			MenuItem:       v.MenuItem,
			ItemQuantity:   v.ItemQuantity,
			ItemStatus:     v.ItemStatus,
		}
		resCartItems = append(resCartItems, res)
	}
	return resCartItems, nil
}

func (cu *cartItemUsecase) GetPendingCartItemList(cartID uint, restaurantID uint, organizationID uint) ([]model.CartItem, error) {

	cartItems := []model.CartItem{}
	if err := cu.cr.GetPendingCartItemList(&cartItems, cartID, restaurantID, organizationID); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("GetPendingCartItemList")
		return nil, err
	}

	return cartItems, nil
}

func (cu *cartItemUsecase) CreateCartItem(cartItem model.CartItem) (model.CartItemResponse, error) {

	if err := cu.cv.CartItemValidate(&cartItem); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("CreateCartItem")
		return model.CartItemResponse{}, err
	}

	if err := cu.cr.CreateCartitem(&cartItem); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("CreateCartItem")
		return model.CartItemResponse{}, err
	}

	resCartItem := model.CartItemResponse{
		ID:             cartItem.ID,
		OrganizationID: cartItem.OrganizationID,
		RestaurantID:   cartItem.RestaurantID,
		CartID:         cartItem.CartID,
		MenuItemID:     cartItem.MenuItemID,
		MenuItem:       cartItem.MenuItem,
		ItemQuantity:   cartItem.ItemQuantity,
		ItemStatus:     cartItem.ItemStatus,
	}

	return resCartItem, nil
}

func (cu *cartItemUsecase) UpdateCartItem(cartItem model.CartItem, id uint, cartID uint, restaurantID uint, organizationID uint) (model.CartItemResponse, error) {

	if err := cu.cv.CartItemValidate(&cartItem); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("UpdateCartItem")
		return model.CartItemResponse{}, err
	}

	if err := cu.cr.UpdateCartitem(&cartItem, id, cartID, restaurantID, organizationID); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("UpdateCartItem")
		return model.CartItemResponse{}, err
	}

	resCartItem := model.CartItemResponse{
		ID:             cartItem.ID,
		OrganizationID: cartItem.OrganizationID,
		RestaurantID:   cartItem.RestaurantID,
		CartID:         cartItem.CartID,
		MenuItemID:     cartItem.MenuItemID,
		MenuItem:       cartItem.MenuItem,
		ItemQuantity:   cartItem.ItemQuantity,
		ItemStatus:     cartItem.ItemStatus,
	}
	return resCartItem, nil
}

func (cu *cartItemUsecase) UpdateCartItemStatus(cartItem model.CartItem, id uint, cartID uint, restaurantID uint, organizationID uint, status string) (error) {

	if err := cu.cr.UpdateCartItemStatus(&cartItem, id, cartID, restaurantID, organizationID, status); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("UpdateCartItemStatus")
		return err
	}

	return nil
}

func (cu *cartItemUsecase) DeleteCartItem(cartItem model.CartItem, id uint, cartID uint, restaurantID uint, organizationID uint) error {

	if err := cu.cr.DeleteCartItem(&cartItem, id, cartID, restaurantID, organizationID); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("DeleteCartItem")
		return err
	}
	return nil
}

// func (cu *cartItemUsecase) SendCartItemToKitchen(cartID uint, organizationID uint, restaurantID uint) ([]model.CartItemResponse, error) {

// 	cartItems := []model.CartItem{}
// 	if err := cu.cr.GetPendingCartItemList(&cartItems, cartID, restaurantID, organizationID); err != nil {
// 		return nil, err
// 	}

// 	// begin transaction
// 	tx := cu.db.Begin()
// 	defer func() {
// 		if r := recover(); r != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	// change cart items status to sent to kitchen
// 	for i, cartItem := range cartItems {
// 		if err := tx.Model(&cartItem).Where("id=? and cart_id=? and restaurant_id=? and organization_id=?", cartItem.ID, cartItem.CartID, restaurantID, organizationID).Update("item_status", "sent_to_kitchen").Error; err != nil {
// 			tx.Rollback()
// 			return nil, err
// 		}

// 		// return cart items with updated status
// 		cartItems[i].ItemStatus = "sent_to_kitchen"

// 		// update product quantity in product table
// 		for j, v := range cartItem.MenuItem.Recipe.RecipeProducts {

// 			common.Logger.LogInfo().Msgf("initial_quantity: %d", cartItems[i].MenuItem.Recipe.RecipeProducts[j].Product.Quantity)

// 			cartItems[i].MenuItem.Recipe.RecipeProducts[j].Product.Quantity = cartItems[i].MenuItem.Recipe.RecipeProducts[j].Product.Quantity - int(math.Ceil((float64(v.Quantity)/float64(cartItem.MenuItem.Recipe.Serving))*float64(cartItem.MenuItem.Serving)*float64(cartItem.ItemQuantity)))

// 			common.Logger.LogInfo().Msgf("quantity_left: %d", cartItems[i].MenuItem.Recipe.RecipeProducts[j].Product.Quantity)
// 			common.Logger.LogInfo().Msgf("quantity_subtracted: %d", int(math.Ceil((float64(v.Quantity)/float64(cartItem.MenuItem.Recipe.Serving))*float64(cartItem.MenuItem.Serving)*float64(cartItem.ItemQuantity))))

// 			for k:=i; k< len(cartItems); k++ {

// 				for l := range cartItems[k].MenuItem.Recipe.RecipeProducts {
// 					if cartItems[i].MenuItem.Recipe.RecipeProducts[j].Product.ID == cartItems[k].MenuItem.Recipe.RecipeProducts[l].Product.ID {

// 						cartItems[k].MenuItem.Recipe.RecipeProducts[l].Product.Quantity = cartItems[i].MenuItem.Recipe.RecipeProducts[j].Product.Quantity
// 					}

// 				}
// 			}
// 			if err := tx.Model(&model.Product{}).Select("quantity", "inventory_value").Where("id=?", v.ProductID).Updates(map[string]interface{}{"quantity": cartItems[i].MenuItem.Recipe.RecipeProducts[j].Product.Quantity, "inventory_value": cartItems[i].MenuItem.Recipe.RecipeProducts[j].Product.Quantity * v.Product.UnitCost}).Error; err != nil {
// 				tx.Rollback()
// 				return nil, err
// 			}
// 		}
// 	}

// 	// Commit transaction
// 	if err := tx.Commit().Error; err != nil {
// 		return nil, err
// 	}

// 	resCartItems := []model.CartItemResponse{}
// 	for _, v := range cartItems {
// 		res := model.CartItemResponse{
// 			ID:             v.ID,
// 			CartID:         v.CartID,
// 			OrganizationID: v.OrganizationID,
// 			RestaurantID:   v.RestaurantID,
// 			MenuItemID:     v.MenuItemID,
// 			MenuItem:       v.MenuItem,
// 			ItemQuantity:   v.ItemQuantity,
// 			ItemStatus:     v.ItemStatus,
// 		}
// 		resCartItems = append(resCartItems, res)
// 	}
// 	return resCartItems, nil
// }

func (cu *cartItemUsecase) SendCartItemToKitchen(cartID uint, organizationID uint, restaurantID uint) ([]model.CartItemResponse, error) {

	cartItems := []model.CartItem{}
	if err := cu.cr.GetPendingCartItemList(&cartItems, cartID, restaurantID, organizationID); err != nil {
		common.Logger.LogError().Fields(map[string]interface{}{"error": err.Error()}).Msg("SendCartItemToKitchen")
		return nil, err
	}

	// calculate updated product quantities
	productUpdates := map[uint]int{}
	for i, cartItem := range cartItems {
		cartItems[i].ItemStatus = "sent_to_kitchen" // update status
		for _, recipeProduct := range cartItem.MenuItem.Recipe.RecipeProducts {
			// calculate product quantity to subtract
			subtractedQty := int(math.Ceil((float64(recipeProduct.Quantity) / float64(cartItem.MenuItem.Recipe.Serving)) *
				float64(cartItem.MenuItem.Serving) * float64(cartItem.ItemQuantity)))

			// update the product quantity in the map
			productUpdates[recipeProduct.ProductID] -= subtractedQty
		}
	}

	// begin transaction
	tx := cu.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// update cart items status
	for _, cartItem := range cartItems {
		if err := tx.Model(&cartItem).Where("id=? AND cart_id=? AND restaurant_id=? AND organization_id=?", cartItem.ID, cartItem.CartID, restaurantID, organizationID).Update("item_status", cartItem.ItemStatus).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// update product quantity in product table in bulk
	for productID, quantityChange := range productUpdates {
		if err := tx.Model(&model.Product{}).Where("id = ?", productID).
			UpdateColumn("quantity", gorm.Expr("quantity + ?", quantityChange)).
			UpdateColumn("inventory_value", gorm.Expr("quantity * unit_cost")).
			Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// commit transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	resCartItems := make([]model.CartItemResponse, len(cartItems))
	for i, cartItem := range cartItems {
		resCartItems[i] = model.CartItemResponse{
			ID:             cartItem.ID,
			CartID:         cartItem.CartID,
			OrganizationID: cartItem.OrganizationID,
			RestaurantID:   cartItem.RestaurantID,
			MenuItemID:     cartItem.MenuItemID,
			MenuItem:       cartItem.MenuItem,
			ItemQuantity:   cartItem.ItemQuantity,
			ItemStatus:     cartItem.ItemStatus,
		}
	}

	return resCartItems, nil
}
