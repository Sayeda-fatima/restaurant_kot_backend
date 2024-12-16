package routes

import (
	"os"

	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func OrderRoutes(e *echo.Echo, oc controller.OrderController, oi controller.OrderItemController){

	c := e.Group("/api")
	c.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
	}))
	
	// checkout 
	c.POST("/cart/:cart/checkout", oc.Checkout)

	o := c.Group("/order")

	o.GET("", oc.GetOrderList)
	o.POST("", oc.CreateOrder)
	o.PUT("/:id", oc.UpdateOrder)
	o.DELETE("/:id", oc.DeleteOrder)

	// order items
	o.GET("/:order/item", oi.GetOrderItemList)
	o.POST("/:order/item", oi.CreateOrderItem)
	o.PUT("/:order/item/:id", oi.UpdateOrderItem)
	o.DELETE("/:order/item/:id", oi.DeleteOrderItem)
}