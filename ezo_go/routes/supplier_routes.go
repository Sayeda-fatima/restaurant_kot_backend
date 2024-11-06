package routes

import (
	"os"

	"github.com/NazishAhsan/easy_busy_book_go/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func SupplierRoutes( e *echo.Echo, sc controller.SupplierController){

	s := e.Group("/api/supplier")
	s.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
	}))

	s.GET("", sc.GetSupplierList)
	s.POST("", sc.CreateSupplier)
	s.PUT("/:id", sc.UpdateSupplier)
	s.PUT("/:id/soft-delete", sc.DeleteSupplier)
}