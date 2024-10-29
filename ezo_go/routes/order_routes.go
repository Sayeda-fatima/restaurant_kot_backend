package routes

import (
	"os"

	"github.com/NazishAhsan/easy_busy_book_go/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func OrderRoutes(e *echo.Echo, oc controller.OrderController, ic controller.OrderItemController) {

	o := e.Group("/api/order")
	o.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
		TokenLookup: "header:x-auth",
	}))

	o.GET("", oc.GetOrderList)
	o.POST("", oc.CreateOrder)
	o.PUT("/:id", oc.UpdateOrder)
	o.PUT("/:id/soft-delete", oc.DeleteOrder)

	o.GET("/:id/item", ic.GetOrderItemList)
	o.POST("/:id/item", ic.CreateOrderItem)
	o.PUT("/item/:id", ic.UpdateOrderItem)
	o.PUT("/item/:id/soft-delete", ic.DeleteOrderItem)
}