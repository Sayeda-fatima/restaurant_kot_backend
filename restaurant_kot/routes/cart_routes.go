package routes

import (
	"os"

	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func CartRoutes(e *echo.Echo, cc controller.CartController, ci controller.CartItemController){

	c := e.Group("/api/cart")
	c.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
	}))

	c.GET("", cc.GetCartList)
	c.POST("", cc.CreateCart)
	c.PUT("/:id", cc.UpdateCart)
	c.PUT("/:id/status", cc.UpdateCartStatus)
	c.DELETE("/:id", cc.DeleteCart)

	// cart items
	c.GET("/:cart/item", ci.GetCartItemList)
	c.POST("/:cart/item", ci.CreateCartItem)
	c.PUT("/:cart/item/:id", ci.UpdateCartItem)
	c.DELETE("/:cart/item/:id", ci.DeleteCartItem)
}