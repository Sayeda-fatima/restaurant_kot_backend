package routes

import (
	"os"

	"github.com/NazishAhsan/easy_busy_book_go/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func OrderRoutes(e *echo.Echo, oc controller.OrderController, ic controller.OrderItemController) {

	c := e.Group("/api")
	c.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
	}))
	// checkout 
	c.POST("/:customer/cart/:id/checkout", oc.Checkout)

	o := c.Group("/order")

	o.GET("", oc.GetOrderList)
	o.POST("", oc.CreateOrder)
	o.GET("/:id", oc.GetInvoice)
	o.PUT("/:id", oc.UpdateOrder)
	o.PUT("/:id/soft-delete", oc.DeleteOrder)
	o.GET("/customer/:id/invoice", oc.InvoiceReportCustomer)
	o.GET("/report", oc.SaleReport)

	o.GET("/:id/item", ic.GetOrderItemList)
	o.POST("/:id/item", ic.CreateOrderItem)
	o.PUT("/item/:id", ic.UpdateOrderItem)
	o.PUT("/item/:id/soft-delete", ic.DeleteOrderItem)
	o.GET("/customer/:customer/invoice/:id", ic.InvoiceCustomer)
}