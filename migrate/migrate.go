package main

import (
	"fmt"

	"github.com/Sayeda-fatima/restaurant_kot_backend/database"
	"github.com/Sayeda-fatima/restaurant_kot_backend/model"
)

func main() {
	db := database.NewDB()
	defer fmt.Println("Successfully Migrated")
	db.AutoMigrate(&model.Organization{}, &model.Restaurant{}, &model.User{}, &model.Allergen{}, &model.Recipe{}, &model.Product{}, &model.RestaurantTable{}, &model.RecipeProduct{}, &model.Menu{}, &model.MenuItem{}, &model.Cart{}, &model.CartItem{}, &model.Order{}, &model.OrderItem{}, &model.Staff{}, &model.WeeklyStaffSchedule{}, &model.Customer{}, &model.MenuAllergen{}, &model.InventoryTransaction{}, &model.CurrentInventory{})
}