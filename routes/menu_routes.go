package routes

import (
	"os"

	"github.com/Sayeda-fatima/restaurant_kot_backend/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func MenuRoutes(e *echo.Echo, mc controller.MenuController, mi controller.MenuItemController, ma controller.MenuAllergenController){

	m := e.Group("/api/menu")
	m.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
	}))

	m.GET("", mc.GetMenuList)
	m.POST("", mc.CreateMenu)
	m.PUT("/:id", mc.UpdateMenu)
	m.DELETE("/:id", mc.DeleteMenu)

	m.GET("/:id/food-cost", mc.FoodCost)

	// menu items
	m.GET("/:menuID/item", mi.GetMenuItemList)
	m.POST("/:menuID/item", mi.CreateMenuItem)
	m.PUT("/:menuID/item/:id", mi.UpdateMenuItem)
	m.DELETE("/:menuID/item/:id", mi.DeleteMenuItem)
	m.PUT("/:menuID/item/:id/available", mi.UpdateMenuItemIsActivated)

	// menu allergen 
	m.GET("/:menuID/allergen", ma.GetMenuAllergenList)
	m.POST("/:menuID/allergen", ma.CreateMenuAllergen)
	m.PUT("/:menuID/allergen/:id", ma.UpdateMenuAllergen)
	m.DELETE("/:menuID/allergen/:id", ma.DeleteMenuAllergen)
}