package routes

import (
	"os"

	"github.com/NazishAhsan/easy_busy_book_go/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Echo, pc controller.ProductController){

	p := e.Group("/api/product")
	p.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
	}))

	// product routes
	p.GET("", pc.GetProductList)
	p.POST("", pc.CreateProduct)
	p.PUT("/:id", pc.UpdateProduct)
	p.PUT("/:id/soft-delete", pc.DeleteProduct)
	p.GET("/search", pc.SearchProduct)
}