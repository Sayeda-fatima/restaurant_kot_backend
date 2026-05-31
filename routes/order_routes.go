package routes

import (
	"os"

	"github.com/Sayeda-fatima/restaurant_kot_backend/controller"
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
	o.GET("/popular", oi.MostOrderedItems)
	o.GET("/daily-sale", oi.DailySaleByItem)
	o.GET("/sales-report", oc.TotalSales)
	o.GET("/sales", oc.TotalSalesByOrderType)

	// order items
	o.GET("/:order/item", oi.GetOrderItemList)
	o.POST("/:order/item", oi.CreateOrderItem)
	o.PUT("/:order/item/:id", oi.UpdateOrderItem)
	o.DELETE("/:order/item/:id", oi.DeleteOrderItem)
}