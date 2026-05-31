package routes

import (
	"os"

	"github.com/Sayeda-fatima/restaurant_kot_backend/controller"
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
	c.GET("/:table/active", cc.CheckCartActive)
	c.PUT("/:id", cc.UpdateCart)
	c.PUT("/:id/status", cc.UpdateCartStatus)
	c.DELETE("/:id", cc.DeleteCart)

	// send to kitchen
	c.POST("/:id/send-to-kitchen", cc.SendCartToKitchen)

	// cart items
	c.GET("/:cart/item", ci.GetCartItemList)
	c.POST("/:cart/item", ci.CreateCartItem)
	c.PUT("/:cart/item/:id", ci.UpdateCartItem)
	c.PUT("/:cart/item/:id/status", ci.UpdateCartItemStatus)
	c.DELETE("/:cart/item/:id", ci.DeleteCartItem)

	// send to kitchen
	c.POST("/:cart/item/send-to-kitchen", ci.SendCartItemToKitchen)
}