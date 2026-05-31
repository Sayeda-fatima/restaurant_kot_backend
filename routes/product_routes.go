package routes

import (
	"os"

	"github.com/Sayeda-fatima/restaurant_kot_backend/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Echo, pc controller.ProductController){

	p := e.Group("/api/product")
	p.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
	}))

	p.GET("", pc.GetAllProduct)
	p.POST("", pc.CreateProduct)
	p.PUT("/:id", pc.UpdateProduct)
	p.DELETE("/:id", pc.DeleteProduct)
	p.GET("/inventory", pc.InventoryValue)
}