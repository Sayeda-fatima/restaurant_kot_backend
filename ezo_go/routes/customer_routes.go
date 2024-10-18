package routes

import (
	"os"
	"github.com/NazishAhsan/easy_busy_book_go/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)


func CustomerRoutes(e *echo.Echo, cc controller.CustomerController){

	c := e.Group("/api/customer")
	c.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "header:x-auth",
	}))

	// customer routes
	c.GET("", cc.GetCustomerList)
	c.POST("", cc.CreateCustomer)
	c.PUT("/:id", cc.UpdateCustomer)
	c.PUT("/:id/soft-delete", cc.DeleteCustomer)
}