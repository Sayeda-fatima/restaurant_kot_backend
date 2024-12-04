package routes

import (
	"os"

	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func MenuRoutes(e *echo.Echo, mc controller.MenuController, mi controller.MenuItemController){

	m := e.Group("/api/menu")
	m.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
	}))

	m.GET("", mc.GetMenuList)
	m.POST("", mc.CreateMenu)
	m.PUT("/:id", mc.UpdateMenu)
	m.DELETE("/:id", mc.DeleteMenu)

	// menu items
	m.GET("/:menuID/item", mi.GetMenuItemList)
	m.POST("/:menuID/item", mi.CreateMenuItem)
	m.PUT("/:menuID/item/:id", mi.UpdateMenuItem)
	m.DELETE("/:menuID/item/:id", mi.DeleteMenuItem)
}