package routes

import (
	"os"

	"github.com/Sayeda-fatima/restaurant_kot_backend/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RestaurantTableRoutes(e *echo.Echo, rc controller.RestaurantTableController){

	rt := e.Group("/api/restaurant/table")
	rt.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
	}))

	rt.GET("", rc.GetRestaurantTableList)
	rt.POST("", rc.CreateRestaurantTable)
	rt.PUT("/:id", rc.UpdateRestaurantTable)
	rt.DELETE("/:id", rc.DeleteRestaurantTable)
}