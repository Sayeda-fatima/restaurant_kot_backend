package routes

import (
	"net/http"
	"os"

	//"github.com/NazishAhsan/easy_busy_book_go/common"
	"github.com/NazishAhsan/easy_busy_book_go/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"
)

func protected(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello umm",
	})
}

func CustomerRoutes(e *echo.Echo, cc controller.CustomerController){


	// Protected route: /protected (JWT required)
	e.GET("/protected", protected, echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
		TokenLookup: "header:x-auth",
	}))


	c := e.Group("/customer")
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