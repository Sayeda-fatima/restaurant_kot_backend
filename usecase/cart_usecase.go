package usecase

import (
	"math"
	"time"

	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
	"github.com/Sayeda-fatima/restaurant_kot_backend/repository"
	"github.com/Sayeda-fatima/restaurant_kot_backend/validator"
	"gorm.io/gorm"
)

type (
	CartUsecase interface {
		GetCartList(organizationID uint, restaurantID uint) ([]model.CartResponse, error)
		GetCart(organizationID uint, restaurantID uint, cartID uint) (model.Cart, error)
		CreateCart(cart model.Cart) (model.CartResponse, error)
		UpdateCart(cart model.Cart, id uint, organizationID uint, restaurantID uint) (model.CartResponse, error)
		UpdateCartStatus(cart model.Cart, id uint, organizationID uint, restaurantID uint, status string) (error)
		DeleteCart(id uint, organizationID uint, restaurantID uint) error
		SendCartToKitchen(id uint, organizationID uint, restaurantID uint) (model.CartResponse, error)
		CheckCartActive(organizationID uint, restaurantID uint, tableID uint) (model.CartResponse, error)
	}

	cartUsecase struct {
		cr repository.CartRepository
		cv validator.CartValidator
		db *gorm.DB
	}
)

func NewCartUsecase(cr repository.CartRepository, cv validator.CartValidator, db *gorm.DB) CartUsecase {
	return &cartUsecase{cr, cv, db}
}

func (cu *cartUsecase) GetCartList(organizationID uint, restaurantID uint) ([]model.CartResponse, error) {

	carts := []model.Cart{}

	if err := cu.cr.GetCartList(&carts, organizationID, restaurantID); err != nil {
		return nil, err
	}

	resCarts := []model.CartResponse{}
	for _, v := range carts {
		res := model.CartResponse{
			ID:             v.ID,
			OrganizationID: v.OrganizationID,
			RestaurantID:   v.RestaurantID,
			TableID:        v.TableID,
			TotalQuantity:  v.TotalQuantity,
			CartType:       v.CartType,
			CartStatus:     v.CartStatus,
			CartItems:      v.CartItems,
		}
		resCarts = append(resCarts, res)
	}
	return resCarts, nil
}

func (cu *cartUsecase) GetCart(organizationID uint, restaurantID uint, cartID uint) (model.Cart, error){

	cart := model.Cart{}
	if err := cu.cr.GetCart(&cart, cartID, organizationID, restaurantID); err != nil{
		return model.Cart{}, err
	}

	return cart, nil
}

func (cu *cartUsecase) CreateCart(cart model.Cart) (model.CartResponse, error) {

	if err := cu.cv.CartValidate(&cart); err != nil {
		return model.CartResponse{}, err
	}

	// Assigning organization_id and restaurant_id to cart_items
	for i := range cart.CartItems {
		cart.CartItems[i].OrganizationID = cart.OrganizationID
		cart.CartItems[i].RestaurantID = cart.RestaurantID
		// getting total quantity of product
		cart.TotalQuantity += cart.CartItems[i].ItemQuantity
	}

	if err := cu.cr.CreateCart(&cart); err != nil {
		return model.CartResponse{}, err
	}

	resCart := model.CartResponse{
		ID:             cart.ID,
		OrganizationID: cart.OrganizationID,
		RestaurantID:   cart.RestaurantID,
		TableID:        cart.TableID,
		TotalQuantity:  cart.TotalQuantity,
		CartType:       cart.CartType,
		CartStatus:     cart.CartStatus,
		CartItems:      cart.CartItems,
	}

	return resCart, nil
}

func (cu *cartUsecase) UpdateCart(cart model.Cart, id uint, organizationID uint, restaurantID uint) (model.CartResponse, error) {

	if err := cu.cv.CartValidate(&cart); err != nil {
		return model.CartResponse{}, err
	}

	// Assigning organization_id and restaurant_id to cart_items
	for i := range cart.CartItems {
		cart.CartItems[i].OrganizationID = cart.OrganizationID
		cart.CartItems[i].RestaurantID = cart.RestaurantID
		// getting total quantity of product
		cart.TotalQuantity += cart.CartItems[i].ItemQuantity
	}

	if err := cu.cr.UpdateCart(&cart, id, organizationID, restaurantID); err != nil {
		return model.CartResponse{}, err
	}

	resCart := model.CartResponse{
		ID:             cart.ID,
		OrganizationID: cart.OrganizationID,
		RestaurantID:   cart.RestaurantID,
		TableID:        cart.TableID,
		TotalQuantity:  cart.TotalQuantity,
		CartType:       cart.CartType,
		CartStatus:     cart.CartStatus,
		CartItems:      cart.CartItems,
	}
	return resCart, nil
}

func (cu *cartUsecase) UpdateCartStatus(cart model.Cart, id uint, organizationID uint, restaurantID uint, status string) (error) {

	if err := cu.cr.UpdateCartStatus(&cart, id, organizationID, restaurantID, status); err != nil {
		return err
	}

	return nil
}

func (cu *cartUsecase) DeleteCart(id uint, organizationID uint, restaurantID uint) error {

	if err := cu.cr.DeleteCart(id, organizationID, restaurantID); err != nil {
		return err
	}
	return nil
}

func (cu *cartUsecase) SendCartToKitchen(id uint, organizationID uint, restaurantID uint) (model.CartResponse, error) {

	cart := model.Cart{}
	if err := cu.cr.GetPendingCart(&cart, id, organizationID, restaurantID); err != nil {
		return model.CartResponse{}, err
	}

	// calculate updated product quantities
	productUpdates := map[uint]int{}
	for i, cartItem := range cart.CartItems {
		// return cart items with updated status
		cart.CartItems[i].ItemStatus = "sent_to_kitchen" // update status
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

	// change cart items status to sent to kitchen
	for _, cartItem := range cart.CartItems {

		if err := tx.Model(&cartItem).Where("id=? AND cart_id=? AND restaurant_id=? AND organization_id=?", cartItem.ID, cartItem.CartID, restaurantID, organizationID).Update("item_status", cartItem.ItemStatus).Error; err != nil {
			tx.Rollback()
			return model.CartResponse{}, err
		}

	}

	for productID, quantityChange := range productUpdates {

		product := model.Product{}
		if err := tx.Where("id=?", productID).First(&product).Error; err != nil{
			return model.CartResponse{}, err
		}
		inventoryTransaction := model.InventoryTransaction{
			OrganizationID: organizationID,
			RestaurantID: restaurantID,
			ProductID: productID,
			StockBeforeUpdate: float64(product.Quantity),
			StockUnitPriceBeforeUpdate: float64(product.UnitCost),
			Quantity: float64(quantityChange),
			UnitCost: float64(product.UnitCost),
			TotalCost: -float64(product.UnitCost) * float64(quantityChange),
			TransactionType: "sale",
			RecordedAt: time.Now(),
		}

		// update product quantity in product table in bulk
		if err := tx.Model(&model.Product{}).Where("id = ?", productID).
			UpdateColumn("quantity", gorm.Expr("quantity + ?", quantityChange)).
			UpdateColumn("inventory_value", gorm.Expr("quantity * unit_cost")).
			Error; err != nil {
			tx.Rollback()
			return model.CartResponse{}, err
		}

		// log consumption of products in inventory_consumption table
		if err := tx.Create(&inventoryTransaction).Error; err != nil{
			tx.Rollback()
			return model.CartResponse{}, err
		}
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return model.CartResponse{}, err
	}

	resCart := model.CartResponse{
		ID:             cart.ID,
		OrganizationID: cart.OrganizationID,
		RestaurantID:   cart.RestaurantID,
		TableID:        cart.TableID,
		TotalQuantity:  cart.TotalQuantity,
		CartType:       cart.CartType,
		CartStatus:     cart.CartStatus,
		CartItems:      cart.CartItems,
	}

	return resCart, nil
}

func (cu *cartUsecase) CheckCartActive(organizationID uint, restaurantID uint, tableID uint) (model.CartResponse, error) {

	cart := model.Cart{}

	if err := cu.cr.CheckCartActive(&cart, organizationID, restaurantID, tableID); err != nil {
		return model.CartResponse{}, err
	}

	resCart := model.CartResponse{
		ID:             cart.ID,
		OrganizationID: cart.OrganizationID,
		RestaurantID:   cart.RestaurantID,
		TableID:        cart.TableID,
		TotalQuantity:  cart.TotalQuantity,
		CartType:       cart.CartType,
		CartStatus:     cart.CartStatus,
		CartItems:      cart.CartItems,
	}

	return resCart, nil
}
