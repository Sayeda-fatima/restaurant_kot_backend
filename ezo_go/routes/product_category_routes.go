package routes

import (
	"os"

	"github.com/NazishAhsan/easy_busy_book_go/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func ProductCategoryRoutes(e *echo.Echo, pc controller.ProductCategoryController) {

	p := e.Group("/api/product/category")
	p.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
	}))

	p.GET("", pc.GetProductCategoryList)
	p.POST("", pc.CreateProductCategory)
	p.PUT("/:id", pc.UpdateProductCategory)
	p.PUT("/:id/soft-delete", pc.DeleteProductCategory)
}
