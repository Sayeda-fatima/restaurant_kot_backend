package routes

import (
	"os"

	"github.com/NazishAhsan/easy_busy_book_laravel/restaurant_kot/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RestaurantTableRoutes(e *echo.Echo, rc controller.RestaurantTableController){

	m := e.Group("/api/restaurant/table")
	m.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
	}))

	m.GET("", rc.GetRestaurantTableList)
	m.POST("", rc.CreateRestaurantTable)
	m.PUT("/:id", rc.UpdateRestaurantTable)
	m.DELETE("/:id", rc.DeleteRestaurantTable)
}