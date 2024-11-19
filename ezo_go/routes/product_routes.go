package routes

import (
	"os"

	"github.com/NazishAhsan/easy_busy_book_go/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Echo, pc controller.ProductController, ps controller.ProductStockController, pr controller.ProductImageController){

	p := e.Group("/api/product")
	p.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
	}))

	// product routes
	p.GET("", pc.GetProductList)
	p.POST("", pc.CreateProduct)
	p.GET("/:id", pc.GetProduct)
	p.PUT("/:id", pc.UpdateProduct)
	p.PUT("/:id/soft-delete", pc.DeleteProduct)
	p.GET("/search", pc.SearchProduct)

	// product stock
	p.GET("/stock", ps.GetProductStockList)
	p.POST("/:id/stock", ps.CreateProductStock)
	p.PUT("/stock/:id", ps.UpdateProductStock)
	p.PUT("/stock/:id/soft-delete", ps.DeleteProductStock)
	p.GET("/stock/type", ps.GetProductStockListByUpdateType)

	// product image
	p.GET("/:id/image", pr.GetProductImageList)
	p.POST("/:id/image", pr.AddProductImage)
	p.DELETE("/image/:id", pr.DeleteProductImage)
}