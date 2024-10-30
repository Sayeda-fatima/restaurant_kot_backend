package routes

import (
	"os"

	"github.com/NazishAhsan/easy_busy_book_go/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func CartRoutes(e *echo.Echo, cc controller.CartController, ic controller.CartItemController) {

	c := e.Group("/api/cart")
	c.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "header:x-auth",
	}))

	c.GET("", cc.GetCartList)
	c.POST("", cc.CreateCart)
	c.PUT("/:id", cc.UpdateCart)
	c.DELETE("/:id", cc.DeleteCart)

	c.GET("/:id/item", ic.GetCartItemList)
	c.POST("/:id/item", ic.CreateCartItem)
	c.PUT("/:cart/item/:id", ic.UpdateCartItem)
	c.DELETE("/:cart/item/:id", ic.DeleteCartItem)
}
